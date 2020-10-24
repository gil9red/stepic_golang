package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	// ReadFile | WriteFile
	{
		const fileName = "test.txt"
		var err error
		var dataFromFile []byte

		_ = os.Remove(fileName)

		dataFromFile, err = ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("dataFromFile (%d): %s\n", len(dataFromFile), dataFromFile)
		// open test.txt: The system cannot find the file specified.
		// dataFromFile: []

		fmt.Println()

		dataForFile := []byte("Тестовая строка, предназначенная для записи в файл")

		// Создаем новый файл и записываем в него данные dataForFile
		if err := ioutil.WriteFile(fileName, dataForFile, 0600); err != nil {
			fmt.Println(err)
		}
		dataFromFile, err = ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("dataFromFile (%d): %s\n", len(dataFromFile), dataFromFile)
		// dataFromFile: Тестовая строка, предназначенная для записи в файл

		// Сравниваем исходные данные с записанными в файл и прочитанными из него
		fmt.Printf("dataForFile == dataFromFile: %v\n", bytes.Equal(dataFromFile, dataForFile))

		_ = os.Remove(fileName)

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// ReadDir
	{
		var filesFromDir []os.FileInfo
		var err error

		filesFromDir, err = ioutil.ReadDir(".")
		if err != nil {
			fmt.Println(err)
		}

		// Print directories
		for _, file := range filesFromDir {
			if file.IsDir() {
				fmt.Printf("Directory: %s\n", file.Name())
			}
		}

		fmt.Println()

		filesFromDir, err = ioutil.ReadDir("1.10")
		if err != nil {
			fmt.Println(err)
		}

		// Print files
		for _, file := range filesFromDir {
			if !file.IsDir() {
				fmt.Printf("name: %s, size: %d\n", file.Name(), file.Size())
			}
		}
	}
}
