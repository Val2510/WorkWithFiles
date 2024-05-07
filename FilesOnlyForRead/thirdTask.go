package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile("readonly.txt", os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Эта запись только для чтения.")
	if err != nil {
		fmt.Println("Ошибка записи в файл:", err)
		return
	}

	fmt.Println("Данные успешно записаны в файл.")
}
