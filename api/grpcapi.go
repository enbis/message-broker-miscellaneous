package api

import (
	"context"
	"errors"
	"log"

	models "github.com/enbis/message-broker-miscellaneous/models/src"
)

type RedirectServer struct {
	NatsHandler *NatsTransport
}

func NewRedirectServer(nats *NatsTransport) *RedirectServer {
	return &RedirectServer{
		NatsHandler: nats,
	}
}

func (redirect *RedirectServer) Send(ctx context.Context, in *models.PingMessage) (*models.Empty, error) {
	log.Printf("Received message %s for topic %s\n", string(in.Payload), in.Topic)

	if redirect.NatsHandler.Conn == nil || redirect.NatsHandler.Conn.IsClosed() {
		err := redirect.NatsHandler.Connect()
		if err != nil {
			return &models.Empty{}, errors.New("Unable to connect to Nats")
		}
	}

	redirect.NatsHandler.Publish(in.Topic, in.Payload)

	return &models.Empty{}, nil
}
