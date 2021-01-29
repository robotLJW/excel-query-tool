package sheet

import (
	"excel-query-tool/pkg/rule"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Sheet struct {
	Name   string
	Cols   [][]string
	RowNum int
	ColNum int
}

func NewSheet(sheetName string, cols [][]string, rowNum int, colNum int) *Sheet {
	sheet := &Sheet{
		Name:   sheetName,
		Cols:   cols,
		RowNum: rowNum,
		ColNum: colNum,
	}
	return sheet
}

func (s *Sheet) FilterRules(f *excelize.File, fules ...*rule.Rule) []int {
	deleteData := make([]int, 0)
	fulesIndex := findRulesColNameIndex(s.Cols, fules...)
	style, _ := f.NewStyle(`{"fill":{"type":"pattern","pattern":1,"color":["#AFEEEE"]}}`)
	for i := 1; i < s.RowNum; i++ {
		flag := false
		for j := 0; j < len(fulesIndex); j++ {
			content := fules[j].Segmentation(s.Cols[fulesIndex[j]][i])
			tempFlag := fules[j].HasPrefix(content)
			if tempFlag {
				f.SetCellStyle(s.Name, "D7", "D7", style)
				flag = true
			}
		}
		if !flag {
			deleteData = append(deleteData, i)
		}
	}
	return deleteData
}

func findRulesColNameIndex(cols [][]string, fules ...*rule.Rule) []int {
	inds := make([]int, len(fules))
	for i := 0; i < len(fules); i++ {
		for j := 0; j < len(cols); j++ {
			if fules[i].Name == cols[j][0] {
				inds[i] = j
				break
			}
		}
	}
	return inds
}
