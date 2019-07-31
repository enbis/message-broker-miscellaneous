package api

import (
	"context"
	"log"

	models "github.com/enbis/message-broker-miscellaneous/models/src"
)

type RedirectServer struct {
}

func NewRedirectServer() *RedirectServer {
	return &RedirectServer{}
}

func (mb *RedirectServer) Send(ctx context.Context, in *models.PingMessage) (*models.Empty, error) {
	log.Printf("Received %s for topic %s\n", string(in.Payload), in.Topic)
	return &models.Empty{}, nil
}
