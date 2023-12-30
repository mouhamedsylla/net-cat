package utils

import (
	"netcat/Internal/App/client"
	"fmt"
	"sync"
)


// The function `BroadcastMessage` sends a message to multiple clients concurrently.
func BroadcastMessage(toSend []client.Client, message []byte) {
	var wg sync.WaitGroup
	for _, user := range toSend {
		wg.Add(1)
		go func(user client.Client) {
			user.Connect.Write(message)
			remis := fmt.Sprintf("[%s]:", user.Name)
			user.Connect.Write([]byte(remis))
			wg.Done()
		}(user)
		wg.Wait()
	}
}

// The function `NotifGroup` sends a notification to a group of clients.
func NotifGroup(notif string, toSend []client.Client) {
	for _, client := range toSend {
		data := fmt.Sprint("\n", notif)
		client.Connect.Write([]byte(data))
		remis := fmt.Sprintf("[%s]:", client.Name)
		client.Connect.Write([]byte(remis))
	}
}
