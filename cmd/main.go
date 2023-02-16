package main

import (
	"fmt"
	"telegram_bot_ozon_notification/config"
	"telegram_bot_ozon_notification/pkg"
)

func main() {
	token := config.Token()
	pkg.TakeAPIToTelegram(token)

	fmt.Println("Programm is started!!!")
}
