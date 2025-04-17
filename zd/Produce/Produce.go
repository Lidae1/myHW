package Produce

import (
	"fmt"
	"io"
	"os"
)

type RealProduce struct {
}

func (r *RealProduce) Produce(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	buf := make([]byte, 100) // Буфер для чтения
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break // Достигнут конец файла
		}
		if err != nil {
			fmt.Println("Ошибка при чтении:", err)
		}
		buf = buf[:n]
	}
	return string(buf)
}
