package main

import (
	"excel-query-tool/pkg/rule"
	"excel-query-tool/pkg/sheet"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("data.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	ruleOne := rule.NewRule("Addresses", "]", "Zhejiang Univ Technol", 1)
	ruleTwo := rule.NewRule("Reprint Addresses", "),", "Zhejiang Univ Technol", 1)
	list := f.GetSheetList()
	// style, _ := f.NewStyle(`{"fill":{"type":"pattern","pattern":1,"color":["#AFEEEE"]}}`)
	for i := 0; i < len(list); i++ {
		sheetName := list[i]
		cols, _ := f.GetCols(sheetName)
		s := sheet.NewSheet(sheetName, cols, len(cols[0]), len(cols))
		deleteData := s.FilterRules(f, ruleOne, ruleTwo)
		fmt.Println(deleteData)
	}

	fmt.Println(list)
}
