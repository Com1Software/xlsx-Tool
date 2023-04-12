package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/xuri/excelize/v2"
)

func main() {
	fmt.Printf("\n xlsx Tool - Operating System : %s\n", runtime.GOOS)
	//	args := os.Args
	switch {
	//-------------------------------------------------------------
	case len(os.Args) == 2:
		file := os.Args[1]
		fmt.Printf("File : %s\n", file)
		f, err := excelize.OpenFile(file)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			// Close the spreadsheet.
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		// Get value from cell by given worksheet name and cell reference.
		sheetlist := f.GetSheetList()

		fmt.Printf("Sheet List %s\n", sheetlist)

		//}
		//-------------------------------------------------------------
	default:

		fmt.Printf("\n Syntax : %s <file> \n\n", os.Args[0])

	}
}
