package api

import (
	"errors"
	"fmt"
	"log"
	"time"

	paho "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

type MqttTransport struct {
	Client paho.Client
}

func NewMqttTransport() *MqttTransport {
	t := new(MqttTransport)

	opts := paho.NewClientOptions().AddBroker(viper.GetString("mqtt_url"))
	opts.AutoReconnect = true
	t.Client = paho.NewClient(opts)

	return t
}

func (t *MqttTransport) Connect() error {
	if token := t.Client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return errors.New("Connect mqtt error")
	}
	return nil
}

func (t *MqttTransport) Disconnect() error {
	t.Client.Disconnect(500)
	time.Sleep(time.Millisecond * 1000)
	if t.Client.IsConnected() {
		return errors.New("Errore disconnessione")
	}
	return nil
}

func (t *MqttTransport) Subscribe(topic string) (chan []byte, error) {
	if t.Client == nil || !t.Client.IsConnected() {
		return nil, errors.New("Client not connected")
	}

	c := make(chan []byte)

	if token := t.Client.Subscribe(topic, 0, func(client paho.Client, msg paho.Message) {
		c <- msg.Payload()
	}); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		return nil, errors.New("Subscribe error")
	}

	log.Printf("MQTT subscribed to topic %s\n", topic)

	return c, nil
}

func (t *MqttTransport) Unsubscribe(topic string) error {
	if token := t.Client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return errors.New("MQTT Unsubscribe error")
	}
	return nil
}

func (t *MqttTransport) Publish(topic string, data []byte) error {
	if t.Client == nil {
		return errors.New("Client MQTT not connected")
	}

	if !t.Client.IsConnected() {
		err := t.Client.Connect()
		if err != nil {
			return errors.New(fmt.Sprintf("Error publishing on MQTT ", err))
		}
	}

	if token := t.Client.Publish(topic, 0, false, data); token.Wait() && token.Error() != nil {
		return errors.New(fmt.Sprintf("MQTT publish error %s \n", token.Error().Error()))
	}

	return nil
}
