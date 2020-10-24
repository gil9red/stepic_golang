package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func removeFile(f *os.File) (err error) {
	_ = f.Close()

	err = os.Remove(f.Name())
	if err != nil {
		return err
	}

	return nil
}

func main() {
	// os.Create | os.Remove -- with Close
	{
		f, err := os.Create("fileName.txt")
		if err != nil {
			fmt.Println(err)
		}
		f.Close()

		err = os.Remove(f.Name())
		fmt.Println(err)
		// <nil>

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// os.Create | os.Remove -- without Close
	{
		f, err := os.Create("fileName.txt")
		// f.Close()

		err = os.Remove(f.Name())
		fmt.Println(err)
		// remove fileName.txt: The process cannot access the file because it is being used by another process.

		_ = removeFile(f)

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// os.Create | os.Rename | os.Remove
	{
		// Создание файла
		f, _ := os.Create("text.txt")
		f.Close()

		// Переименование файла
		os.Rename("text.txt", "new_text.txt")

		// Удаление файла
		os.Remove("new_text.txt")

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// os.Create | WriteString | ioutil.ReadFile | strconv.Quote
	{
		f, _ := os.Create("text.txt")
		f.WriteString("1 строка\n")
		f.WriteString("2 строка\n")
		f.Close()

		data, _ := ioutil.ReadFile(f.Name())
		text := string(data)
		fmt.Printf("text (%d): %s\n", len(text), strconv.Quote(text))
		// data (30): "1 строка\n2 строка\n"

		removeFile(f)

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// os.Create | Stat
	{
		f, _ := os.Create("text.txt")
		f.WriteString("1234")
		info, _ := f.Stat()

		fmt.Printf(
			"name: %v, size: %v, mode: %v, modTime: %v, sys: %v\n",
			info.Name(), info.Size(), info.Mode(), info.ModTime(), info.Sys(),
		)
		// name: text.txt, size: 4, mode: -rw-rw-rw-, modTime: 2020-10-24 18:35:04.5977649 +0500 +05, sys: &{32 {2065564442 30845450} {2065574449 30845450} {2065574449 30845450} 0 4}

		removeFile(f)

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// os.Exit
	{
		const fileName = "fileName.txt"
		os.Remove(fileName)

		f, _ := os.Create(fileName)
		defer removeFile(f)

		fmt.Println("Before os.Exit")
		os.Exit(0)
		fmt.Println("After os.Exit")
	}
}
