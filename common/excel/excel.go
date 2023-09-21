package excel

import (
	"fmt"
	"io"
	"reflect"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// ExportExcelByStruct excel导出(数据源为Struct)
func ExportByStruct(c io.Writer, titleList []string, data []interface{}, sheetName string) error {
	f := excelize.NewFile()
	f.SetSheetName("Sheet1", sheetName)
	header := make([]string, 0)
	for _, v := range titleList {
		header = append(header, v)
	}
	rowStyleID, _ := f.NewStyle(`{"font":{"color":"#666666","size":13,"family":"arial"},"alignment":{"vertical":"center","horizontal":"center"}}`)
	_ = f.SetSheetRow(sheetName, "A1", &header)
	_ = f.SetRowHeight("Sheet1", 1, 30)
	length := len(titleList)
	headStyle := letter(length)
	var lastRow string
	var widthRow string
	for k, v := range headStyle {
		if k == length-1 {
			lastRow = fmt.Sprintf("%s1", v)
			widthRow = v
		}
	}
	f.SetColWidth(sheetName, "A", widthRow, 30)
	rowNum := 1
	for _, v := range data {
		t := reflect.TypeOf(v)
		value := reflect.ValueOf(v)
		row := make([]interface{}, 0)
		for l := 0; l < t.NumField(); l++ {
			val := value.Field(l).Interface()
			row = append(row, val)
		}
		rowNum++
		err := f.SetSheetRow(sheetName, "A"+strconv.Itoa(rowNum), &row)
		_ = f.SetCellStyle(sheetName, fmt.Sprintf("A%d", rowNum), fmt.Sprintf("%s", lastRow), rowStyleID)
		if err != nil {
			return err
		}
	}
	return f.Write(c)
}

// Letter 遍历a-z
func letter(length int) []string {
	var str []string
	for i := 0; i < length; i++ {
		str = append(str, string(rune('A'+i)))
	}
	return str
}
