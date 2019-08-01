package api

import (
	"errors"
	"log"

	nats "github.com/nats-io/go-nats"
	"github.com/spf13/viper"
)

var subscriptions = make(map[string]chan *nats.Msg)

type NatsTransport struct {
	URL  string
	Conn *nats.Conn
}

func NewNatsTransport() *NatsTransport {
	transport := new(NatsTransport)
	transport.URL = viper.GetString("nats_url")
	return transport
}

func (transport *NatsTransport) Connect() error {
	nc, err := nats.Connect(transport.URL)
	if err != nil {
		return err
	}
	transport.Conn = nc
	return nil
}

func (transport *NatsTransport) Disconnect() error {
	if transport.Conn == nil {
		return nil
	}
	transport.Conn.Close()
	transport.Conn = nil
	return nil
}

func (transport *NatsTransport) Subscribe(subj string) (chan []byte, error) {
	if transport.Conn == nil || transport.Conn.IsClosed() {
		return nil, errors.New("Not connected")
	}

	c := make(chan []byte)
	var ch chan *nats.Msg

	ch = make(chan *nats.Msg, 64)
	transport.Conn.ChanSubscribe(subj, ch)
	log.Printf("Subscribed to %s\n", subj)

	go func() {
		for {
			select {
			case msg := <-ch:
				if msg == nil {
					return
				}
				c <- msg.Data
				break
			}
		}
	}()
	subscriptions[subj] = ch
	return c, nil
}

func (transport *NatsTransport) Unsubscribe(subj string) error {
	if ch, ok := subscriptions[subj]; ok {
		ch <- nil
		delete(subscriptions, subj)
	}
	return nil
}

func (transport *NatsTransport) Publish(subj string, data []byte) error {
	if transport.Conn == nil || transport.Conn.IsClosed() {
		return errors.New("Not connected")
	}
	log.Printf("Published %s on topic %s\n", data, subj)
	return transport.Conn.Publish(subj, data)
}

func (transport *NatsTransport) GetSubscriptions() []string {
	subs := make([]string, len(subscriptions))
	i := 0
	for s := range subscriptions {
		subs[i] = s
		i++
	}
	return subs
}
