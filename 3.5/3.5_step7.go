package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func createTestFile(autoClose bool) *os.File {
	f, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	_, _ = f.WriteString("123456\nabc\nFOOBAR")

	if autoClose {
		defer f.Close()
	}

	return f
}

func main() {
	// bufio.Reader - позволяет читать данные по байтам, рунам, строкам и пр., указывать символ,
	// на котором необходимо прекратить чтение. Когда данные будут прочитаны до конца, метод вернет
	// ошибку io.EOF.
	{
		f := createTestFile(true)

		f, err := os.Open(f.Name())
		if err != nil {
			panic(err)
		}
		defer func() {
			f.Close()
			os.Remove(f.Name())
		}()

		rd := bufio.NewReader(f)

		buf := make([]byte, 5)
		n, err := rd.Read(buf) // читаем в buf 10 байт из ранее открытого файла
		if err != nil && err != io.EOF {
			// io.EOF не совсем ошибка - это состояние, указывающее, что файл прочитан до конца
			panic(err)
		}
		fmt.Printf("Прочитано %d байт: %q\n", n, buf)
		// Прочитано 5 байт: "12345"

		s, err := rd.ReadString('\n') // читаем данные до разрыва абзаца ('\n')
		fmt.Printf("Text (%v): %q EOF=%v\n", len(s), s, err == io.EOF)
		// Text (2): "6\n" EOF=false

		s, err = rd.ReadString('\n')
		fmt.Printf("Text (%v): %q EOF=%v\n", len(s), s, err == io.EOF)
		// Text (4): "abc\n" EOF=false

		s, err = rd.ReadString('\n')
		fmt.Printf("Text (%v): %q EOF=%v\n", len(s), s, err == io.EOF)
		// Text (6): "FOOBAR" EOF=true

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// bufio.Writer - создан для записи в объекты, удовлетворяющие интерфейсу io.Writer,
	// но предоставляет ряд более высокоуровневых методов, в частности метод WriteString(s string)
	{
		f, err := os.Create("test.txt")
		if err != nil {
			panic(err)
		}
		defer func() {
			f.Close()
			os.Remove(f.Name())
		}()

		w := bufio.NewWriter(f)
		n, err := w.WriteString("Запишем строку")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Записано %d байт\n", n)
		// Записано 27 байт

		data, _ := ioutil.ReadFile(f.Name())
		text := string(data)
		fmt.Printf("Text (%v): %q\n", len(text), text)
		// Text (0): ""

		// bufio.Writer имеет собственный буфер, чтобы быть уверенным, что данные точно записаны,
		// вызываем метод Flush()
		w.Flush()

		data, _ = ioutil.ReadFile(f.Name())
		text = string(data)
		fmt.Printf("Text (%v): %q\n", len(text), text)
		// Text (27): "Запишем строку"

		fmt.Println("\n" + strings.Repeat("-", 100) + "\n")
	}

	// bufio.Scanner - создан для построчного чтения данных.
	{
		f := createTestFile(false)

		f, err := os.Open(f.Name())
		if err != nil {
			panic(err)
		}
		defer func() {
			f.Close()
			os.Remove(f.Name())
		}()

		s := bufio.NewScanner(f)

		// Возвращает true, пока файл не будет прочитан до конца
		for s.Scan() {
			// s.Text() содержит данные, считанные на данной итерации
			line := s.Text()
			fmt.Printf("Line (%v) %q\n", len(line), line)
		}
		// Line (6) "123456"
		// Line (3) "abc"
		// Line (6) "FOOBAR"
	}
}
