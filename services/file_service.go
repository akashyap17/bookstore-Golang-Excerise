package services

import (
	"fmt"
	"os"
)

func SaveToFile(filename, content string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

}
