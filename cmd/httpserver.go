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
	"context"
	"fmt"
	"log"
	"net/http"

	models "github.com/enbis/message-broker-miscellaneous/models/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// httpserverCmd represents the httpserver command
var httpserverCmd = &cobra.Command{
	Use:   "httpserver",
	Short: "Launch HTTP Server",
	Long:  `Launch HTTP Server listening for HTTP Client request, and foreward it to the gRPC Server`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("httpserver listening on %s\n", viper.GetString("http_port"))

		http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {

			var topic string
			var payload string

			topicParam, ok := r.URL.Query()["topic"]
			if !ok || len(topicParam[0]) < 1 {
				log.Println("Url Param topic is missing")
				topic = viper.GetString("topic")
			} else {
				topic = topicParam[0]
			}

			payloadParam, ok := r.URL.Query()["payload"]
			if !ok || len(payloadParam[0]) < 1 {
				log.Println("Url Param payload is missing")
				payload = viper.GetString("topic")
			} else {
				payload = payloadParam[0]
			}

			sendDataTogRPCServer(topic, payload)

		})

		http_port := fmt.Sprintf(":%s", viper.GetString("http_port"))
		http.ListenAndServe(http_port, nil)

		select {}
	},
}

func sendDataTogRPCServer(topic, payload string) {
	var conn *grpc.ClientConn

	dialPort := fmt.Sprintf(":%d", viper.GetInt("grpc_port"))

	conn, err := grpc.Dial(dialPort, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := models.NewRedirectClient(conn)
	_, err = c.Send(context.Background(), &models.PingMessage{Topic: topic, Payload: []byte(payload)})

	if err != nil {
		log.Fatalf("Error sending the data to gRPC Server: %s", err)
	}

	log.Printf("Information %s %s sent to the gRPC Server \n", topic, payload)

}

func init() {
	rootCmd.AddCommand(httpserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// httpserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// httpserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
