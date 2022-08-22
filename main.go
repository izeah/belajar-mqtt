package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
}

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
	options.SetClientID("izeah01")
	options.SetDefaultPublishHandler(messagePubHandler)
	options.OnConnect = connectHandler
	options.OnConnectionLost = connectionLostHandler

	client := mqtt.NewClient(options)
	token := client.Connect()
	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "volta/stb/chz123"
	token = client.Subscribe(topic, 0, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic \"%s\"\n", topic)

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	signal.Notify(stop, syscall.SIGTERM)

	//Recieve shutdown signals.
	<-stop

	client.Disconnect(100)
}
