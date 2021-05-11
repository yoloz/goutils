//遍历文件生成docsify中的sidebar,readme
//全局执行顺序，先执行全局变量，在执行init(),在执行main方法
package main

import (
	"container/list"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var subList *list.List

func init() {
	subList = list.New()
}

func main() {

	var docsDir string
	fmt.Printf("Please enter docsify directory: ")
	fmt.Scanln(&docsDir)

	fd, err := os.ReadDir(docsDir)
	if err != nil {
		log.Fatal(err)
	}

	flist := list.New()
	templ := list.New()
	var str strings.Builder

	for _, catlog := range fd { //遍历目录
		if !catlog.IsDir() || strings.HasPrefix(catlog.Name(), "_") || strings.HasPrefix(catlog.Name(), ".") {
			continue
		}
		flist.Init()
		templ.Init()
		hasDir := false

		catalogPath := filepath.Join(docsDir, catlog.Name())
		td, err := os.ReadDir(catalogPath)
		if err != nil {
			log.Fatal(err)
		}
		flist.PushBack("- 文件\n")

		for _, file := range td { //遍历文件
			str.Reset()
			if file.Type().IsRegular() && !strings.HasPrefix(file.Name(), "_") &&
				strings.Compare(file.Name(), "README.md") != 0 {
				str.WriteString("  - [")
				str.WriteString(strings.TrimSuffix(file.Name(), ".md"))
				str.WriteString("](/")
				str.WriteString(catlog.Name())
				str.WriteString("/")
				str.WriteString(file.Name())
				str.WriteString(")\n")
				flist.PushBack(str.String())
			} else if file.Type().IsDir() {
				go scanSubDir(filepath.Join(catalogPath, file.Name()))
				str.WriteString("  - [")
				str.WriteString(strings.TrimSuffix(file.Name(), ".md"))
				str.WriteString("](/")
				str.WriteString(catlog.Name())
				str.WriteString("/")
				str.WriteString(file.Name())
				str.WriteString("/")
				str.WriteString("README.md)\n")
				templ.PushBack(str.String())
			}
		}

		if templ.Len() > 0 {
			for e := templ.Front(); e != nil; e = e.Next() {
				flist.PushFront(e.Value.(string))
			}
			hasDir = true
		}

		if hasDir {
			flist.PushFront("- 文件夹\n")
		}

		//生成文件
		siderbar, err := os.Create(filepath.Join(catalogPath, "_sidebar.md"))
		if err != nil {
			log.Fatal(err)
		}
		defer siderbar.Close()

		readme, err := os.Create(filepath.Join(catalogPath, "README.md"))
		if err != nil {
			log.Fatal(err)
		}
		defer readme.Close()

		for e := flist.Front(); e != nil; e = e.Next() {
			line := e.Value.(string)
			siderbar.WriteString(line)
			readme.WriteString(line)
		}

	}
}

//遍历二级目录
func scanSubDir(dirPath string) {

	files, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	subList.Init()
	subList.PushBack("- 文件\n")
	var str strings.Builder

	for _, file := range files {
		if !file.Type().IsRegular() || strings.HasPrefix(file.Name(), "_") || strings.Compare(file.Name(), "README.md") == 0 {
			continue
		}
		str.Reset()
		str.WriteString("  - [")
		str.WriteString(strings.TrimSuffix(file.Name(), ".md"))
		str.WriteString("](/")
		str.WriteString(filepath.Base(filepath.Dir(dirPath)))
		str.WriteString("/")
		str.WriteString(filepath.Base(dirPath))
		str.WriteString("/")
		str.WriteString(file.Name())
		str.WriteString(")\n")
		subList.PushBack(str.String())
	}

	//生成文件
	siderbar, err := os.Create(filepath.Join(dirPath, "_sidebar.md"))
	if err != nil {
		log.Fatal(err)
	}
	defer siderbar.Close()

	readme, err := os.Create(filepath.Join(dirPath, "README.md"))
	if err != nil {
		log.Fatal(err)
	}
	defer readme.Close()

	for e := subList.Front(); e != nil; e = e.Next() {
		line := e.Value.(string)
		siderbar.WriteString(line)
		readme.WriteString(line)
	}
}
