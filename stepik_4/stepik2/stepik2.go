package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func Server() {

	// Устанавливаем прослушивание порта
	ln, err := net.Listen("tcp", "127.0.0.1:8081")

	if err != nil {
		log.Println(err)
	}
	// Открываем порт
	conn, err := ln.Accept()

	if err != nil {
		log.Println(err)
	}

	defer conn.Close()

	conn.Write([]byte("message"))
	time.Sleep(10)
	conn.Write([]byte("MesSaGe"))
	time.Sleep(10)
	conn.Write([]byte("MESSAGE"))

}

func main() {
	// Подключение к серверу по протоколу TCP
	conn, err := net.Dial("tcp", "127.0.0.1:8081")

	if err != nil {
		fmt.Println("Ошибка подключения:", err)
		return
	}
	defer conn.Close()

	buffer := make([]byte, 1024)

	for i := 0; i < 3; i++ {
		// Чтение сообщения в буфер
		bytesRead, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Ошибка чтения:", err)
			return
		}
		// Преобразование прочитанных байт в строку и перевод в верхний регистр
		message := strings.ToUpper(string(buffer[:bytesRead]))
		fmt.Print(message)

	}

}
