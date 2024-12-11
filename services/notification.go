package services

import (
	"fmt"
)

func SendNotification(message string, ch chan string) {
	ch <- fmt.Sprintf("Notification sent: %s ", message)
}

func NotifyAllCustomers() {
	ch := make(chan string)
	go SendNotification("New Books available", ch)
	fmt.Println(<-ch)
}
