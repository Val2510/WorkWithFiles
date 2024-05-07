package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	file, err := os.OpenFile("messages.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return
	}

	if fileInfo.Size() == 0 {
		fmt.Println("Файл messages.txt пустой")
	} else {
		fmt.Println("Размер файла messages.txt:", fileInfo.Size(), "байт")
	}

	writer := bufio.NewWriter(file)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите строки. Для завершения введите 'exit':")

	for scanner.Scan() {
		line := scanner.Text()
		if line == "exit" {
			break
		}

		currentTime := time.Now().Format("2006-01-02 15:04:05")

		message := fmt.Sprintf("%s %s\n", currentTime, line)
		_, err := writer.WriteString(message)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}

		writer.Flush()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Ошибка чтения ввода:", err)
		return
	}

	fmt.Println("Программа завершена.")
}
