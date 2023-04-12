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
		cmd := os.Args[1]
		fmt.Printf("cmd %s\n", cmd)
		f, err := excelize.OpenFile("Book1.xlsx")
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
		cell, err := f.GetCellValue("Sheet1", "B2")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(cell)

		//}
		//-------------------------------------------------------------
	default:

		fmt.Printf("\n Syntax : %s <file> \n\n", os.Args[0])

	}
}
