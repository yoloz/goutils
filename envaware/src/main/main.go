package main

import (
	"encoding/json"
	"envaware/src/psutil"
	"errors"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	m := make(map[string]interface{}, 5)
	go send(m)
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		close(idleConnsClosed)
	}()

	<-idleConnsClosed
}

func send(m map[string]interface{}) {
	for {
		m["hostinfo"], _ = psutil.HostInfo()
		m["cpuinfo"], _ = psutil.CpuInfo()
		m["diskinfo"], _ = psutil.PartitionInfo()
		m["nclist"], _ = psutil.NCList()
		m["procinfo"], _ = psutil.ProcessList()
		dataType, _ := json.Marshal(m)
		dataString := string(dataType)
		fp, _ := getCurrentPath()
		openFile, e := os.OpenFile(fp+"envaware.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
		if e != nil {
			log.Fatal(e)
		}
		openFile.WriteString(dataString)
		openFile.Close()

		time.Sleep(time.Second * 5)
	}
}

func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\"`)
	}
	return string(path[0 : i+1]), nil
}
