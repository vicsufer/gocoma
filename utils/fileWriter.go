package utils

import (
	"fmt"
	"os"
)

func WriteToFile(filename, content string) {
	f ,err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Printf("Can't create %sfile:\n%s", filename, err)
		os.Exit(1)
	}
	f.WriteString(content)
}
