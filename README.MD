# MQTT (Message Queuing Telemetry Transport)

MQTT adalah protokol messaging yang dibentuk dengan TCP/IP berdasarkan model messaging publish-subscribe. Publisher mengirim pesan, subscriber menerima pesan yang mereka sukai, dan broker akan menyampaikan pesan dari pengirim ke penerima. Publisher dan subscriber adalah klien MQTT yang hanya berkomunikasi dengan broker MQTT. Klien MQTT dapat berupa perangkat atau aplikasi apapun (dari mikro kontroler seperti Arduino sampai dengan aplikasi penuh yang di host di Cloud) yang menjalankan MQTT library dan mengkoneksikan ke broker MQTT melalui sebuah jaringan. Broker MQTT mengelola penerimaan pesan dari publisher dan pengiriman pesan ke subscriber (dan juga mengelola daftar topik yang disukai subscriber).

# Cara Run MQTT

1. `go run main.go` (subscriber)
2. `go run publisher.go` (publisher)
3. send message on `publisher.go`

# Author

I Z E A H .

# References

-   https://gist.github.com/ustropo/58d9de16570682d8151e14138a5e0c95
-   https://medium.com/pujanggateknologi/berkenalan-dengan-teknologi-
-   https://www.emqx.com/id/blog/introduction-to-mqtt-qos
-   https://www.emqx.com/id/blog/mqtt-is-the-leading-protocol-for-connecting-iot-devices
