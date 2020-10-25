package main

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
Поиск файла в заданном формате и его обработка
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath, хотя для решения
может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется
именно в этом формате).

В тестовом архиве, который мы можете скачать из нашего репозитория на github.com, содержится
набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV, прочие же
файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица
10х10, разделителем является запятая), а в качестве ответа необходимо указать число, находящееся
на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func download(fileName, url string) *os.File {
	f, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rs, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer rs.Body.Close()

	_, err = io.Copy(f, rs.Body)
	if err != nil {
		panic(err)
	}

	return f
}

func main() {
	const URL = "https://github.com/semyon-dev/stepik-go/raw/master/work_with_files/task_csv_1/task.zip"
	const FileName = "task_step13.zip"

	ok := Exists(FileName)
	if !ok {
		fmt.Printf("Download from %v\n", URL)
		download(FileName, URL)
	}

	zf, err := zip.OpenReader(FileName)
	if err != nil {
		panic(err)
	}
	defer zf.Close()

	for _, file := range zf.File {
		f, err := file.Open()
		if err != nil {
			continue
		}
		defer f.Close()

		lines, _ := csv.NewReader(f).ReadAll()
		if len(lines) > 1 {
			fmt.Println(lines[4][2])
		}
	}
}
