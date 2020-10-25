package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err // Если по какой-то причине мы получили ошибку, проигнорируем эту итерацию
	}

	// Проигнорируем директории
	if info.IsDir() {
		// Не спускаемся в директории .git и .idea
		if info.Name() == ".git" || info.Name() == ".idea" {
			return filepath.SkipDir
		}

		return nil
	}

	fmt.Printf("Name: %s\tSize: %d byte\tPath: %s\n", info.Name(), info.Size(), path)
	return nil
}

func main() {
	const root = "."

	if err := filepath.Walk(root, walkFunc); err != nil {
		fmt.Printf("Какая-то ошибка: %v\n", err)
	}
	// Name: .gitignore	Size: 269 byte	Path: .gitignore
	// Name: 1.10_step10.go	Size: 909 byte	Path: 1.10\1.10_step10.go
	// Name: 1.10_step2.go	Size: 96 byte	Path: 1.10\1.10_step2.go
	// Name: 1.10_step3.go	Size: 508 byte	Path: 1.10\1.10_step3.go
	// Name: 1.10_step4.go	Size: 736 byte	Path: 1.10\1.10_step4.go
	// ...
}
