package main

import (
	"bufio"
	"fmt"
	"os"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected")
}

var connectionLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connection Lost: %s\n", err.Error())
}

func main() {
	var broker = "tcp://test.mosquitto.org:1883"
	options := mqtt.NewClientOptions()
	options.AddBroker(broker)
	options.SetClientID("IOT Traccar")
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	buf := bufio.NewReader(os.Stdin)

	var input string
	for {
		fmt.Print("Publish a message : ")
		input, _ = buf.ReadString('\n')
		token = client.Publish("volta/stb/chz123", 0, false, input)
		token.Wait()
	}
}
