package main

import (
	"fmt"
	"telegram_bot_ozon_notification/config"
)

func main() {
	token := config.Token()

	fmt.Println("Programm is started!!!")
	fmt.Println("Token = ", token)
}
