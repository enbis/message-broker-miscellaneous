/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/enbis/message-broker-miscellaneous/api"
	models "github.com/enbis/message-broker-miscellaneous/models/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// grpcserverCmd represents the grpcserver command
var grpcserverCmd = &cobra.Command{
	Use:   "grpcserver",
	Short: "Launch gRPC Server",
	Long:  `Launch gRPC Server listening for gRPC Client request, and foreward it as pub/sub on nats`,
	Run: func(cmd *cobra.Command, args []string) {

		log.Printf("gRPC Server")

		grpcPort := fmt.Sprintf(":%d", viper.GetInt("grpc_port"))
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		nats := api.NewNatsTransport()

		s := api.RedirectServer{NatsHandler: nats}

		grpcServer := grpc.NewServer()

		models.RegisterRedirectServer(grpcServer, &s)

		// start the server
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %s", err)
		}

	},
}

func init() {
	rootCmd.AddCommand(grpcserverCmd)
}
