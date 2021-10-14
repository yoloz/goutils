package excelUtil

import (
	"fmt"
	"log"
	"time"

	"github.com/xuri/excelize/v2"
	"github.com/yoloz/gos/calendarUtil"
)

func h1Style(file *excelize.File) (int, error) {

	return file.NewStyle(`{
		"border": [
		{
			"type": "left",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "top",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "bottom",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "right",
			"color": "#000000",
			"style": 2
		}],
		"alignment":
		{
			"horizontal": "center",
			"vertical": "center"
		},
		"font":
		{
			"bold": true,
			"family": "Times New Roman",
			"size": 16,
			"color": "#777777"
		}
	}`)
}

func h2Style(file *excelize.File) (int, error) {

	return file.NewStyle(`{
		"border": [
		{
			"type": "left",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "top",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "bottom",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "right",
			"color": "#000000",
			"style": 2
		}],
		"alignment":
		{
			"horizontal": "center",
			"vertical": "center"
		},
		"font":
		{
			"bold": true,
			"family": "Times New Roman",
			"size": 12,
			"color": "#777777"
		}
	}`)
}

func hbStyle(file *excelize.File) (int, error) {
	return file.NewStyle(`{
		"border": [
		{
			"type": "left",
			"color": "#000000",
			"style": 2
		},
		{
			"type": "right",
			"color": "#000000",
			"style": 2
		}],
		"alignment":
		{
			"horizontal": "center",
			"vertical": "center"
		},
		"font":
		{
			"bold": false,
			"family": "Times New Roman",
			"size": 10,
			"color": "#777777"
		}
	}`)
}

func leftBorderStyle(file *excelize.File) (int, error) {
	return file.NewStyle(`{
		"border": [
		{
			"type": "left",
			"color": "#000000",
			"style": 2
		}]
	}`)
}

func rightBoderStyle(file *excelize.File) (int, error) {
	return file.NewStyle(`{
		"border": [
		{
			"type": "right",
			"color": "#000000",
			"style": 2
		}]
	}`)
}

func fillWeekendStyle(file *excelize.File) (int, error) {
	return file.NewStyle(`{
		"fill":
		{
			"type":"pattern",
			"color":["#E0E0E0"],
			"pattern":1
		}
	}`)
}

func GeneratePlan(year int, months []int) {
	var (
		days      int          //当月天数
		err       error        //错误
		startWeek time.Weekday //当月第一天星期几
		weekIndex int          //当月第几周
		colIndex  int          //列索引
		colName   string       //列名称
		mstartcol string       //月份起始列
		mendcol   string       //月份结束列
		wstartcol string       //周起始列
		wendcol   string       //周结束列

	)

	f := excelize.NewFile()
	sheetName := "Sheet1"
	f.SetSheetName("Sheet1", sheetName)

	h1, err := h1Style(f)
	if err != nil {
		fmt.Printf("h1 style error %v", err)
	}

	f.SetCellValue(sheetName, "B2", "任务")
	f.SetCellStyle(sheetName, "B2", "B5", h1)
	f.MergeCell(sheetName, "B2", "B5")

	headBody, err := hbStyle(f)
	if err != nil {
		fmt.Printf("head body style error %v", err)
	}

	f.SetCellStyle(sheetName, "B6", "B30", headBody)

	if err = f.SetColWidth(sheetName, "B", "B", 40); err != nil {
		log.Fatal(err)
	}

	h2, err := h2Style(f)
	if err != nil {
		fmt.Printf("h2 style error %v", err)
	}

	leftBorder, err := leftBorderStyle(f)
	if err != nil {
		fmt.Printf("left border style error %v", err)
	}

	rightBorder, err := rightBoderStyle(f)
	if err != nil {
		fmt.Printf("right border style error %v", err)
	}

	weekendBack, err := fillWeekendStyle(f)
	if err != nil {
		fmt.Printf("weekend backgroud style error %v", err)
	}

	colIndex = 3 //从C列开始是数据

	for _, month := range months {

		if days, err = calendarUtil.GetMonthDay(year, month); err != nil {
			log.Fatal(err)
		}

		if startWeek, err = calendarUtil.GetWeekday(year, month, 1); err != nil {
			log.Fatal(err)
		}

		if mstartcol, err = excelize.ColumnNumberToName(colIndex); err != nil {
			log.Fatal(err)
		}

		weekIndex = 1
		for day := 1; day <= days; day++ {
			week := (int(startWeek) + day - 1) % 7

			if week == 0 || day == days {
				if wendcol, err = excelize.ColumnNumberToName(colIndex); err != nil {
					log.Fatal(err)
				}
				f.SetCellValue(sheetName, wstartcol+"3", fmt.Sprintf("第%d周", weekIndex))
				f.SetCellStyle(sheetName, wstartcol+"3", wendcol+"3", h2)
				f.MergeCell(sheetName, wstartcol+"3", wendcol+"3")

				f.SetCellStyle(sheetName, wendcol+"6", wendcol+"30", rightBorder)

				if week == 0 {
					saturday, err := excelize.ColumnNumberToName(colIndex - 1)
					if err != nil {
						fmt.Printf("get staurday column name error %v", err)
					}
					f.SetCellStyle(sheetName, saturday+"6", saturday+"30", weekendBack)
					f.SetCellStyle(sheetName, wendcol+"6", wendcol+"30", weekendBack)
				}

				weekIndex++
			}

			if week == 1 || day == 1 {
				if wstartcol, err = excelize.ColumnNumberToName(colIndex); err != nil {
					log.Fatal(err)
				}

				f.SetCellStyle(sheetName, wstartcol+"6", wstartcol+"30", leftBorder)
			}

			weekday := string([]rune(calendarUtil.Weekday_zh(time.Weekday(week)))[2:])
			if colName, err = excelize.ColumnNumberToName(colIndex); err != nil {
				log.Fatal(err)
			}
			f.SetCellValue(sheetName, colName+"4", weekday)
			f.SetCellStyle(sheetName, colName+"4", colName+"4", h2)
			f.SetCellValue(sheetName, colName+"5", day)
			f.SetCellStyle(sheetName, colName+"5", colName+"5", h2)
			colIndex++
		}

		//不包含最后一列，否则最后三个月合并成一个月
		if mendcol, err = excelize.ColumnNumberToName(colIndex - 1); err != nil {
			log.Fatal(err)
		}
		f.SetCellValue(sheetName, mstartcol+"2", fmt.Sprintf("%d月", month))
		f.SetCellStyle(sheetName, mstartcol+"2", mendcol+"2", h2)
		f.MergeCell(sheetName, mstartcol+"2", mendcol+"2")
	}

	if err = f.SaveAs("plan.xlsx"); err != nil {
		log.Fatal(err)
	}
}
