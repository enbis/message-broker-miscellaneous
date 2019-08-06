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

	"github.com/enbis/message-broker-miscellaneous/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// natssubscriberCmd represents the natssubscriber command
var natssubscriberCmd = &cobra.Command{
	Use:   "natssubscriber",
	Short: "This command allows to subscribe to a specific Nats topic",
	Long: `This command allows to subscribe to a specific Nats topic.
	There is a parameter to specify the Nats topic to subscribe, be careful that match with the topic selected from the http request
	otherwise you won't see the message coming.
	If no topic is provided it reads it from the config.yml.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("natssubscriber called")

		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			topic = viper.GetString("topic")
		}

		nats := api.NewNatsTransport()
		if nats.Conn == nil || nats.Conn.IsClosed() {
			err := nats.Connect()
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
		}

		ch, err := nats.Subscribe(topic)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		for {
			select {
			case msg := <-ch:
				log.Println("Received message: ", string(msg))
			}
		}

	},
}

func init() {
	rootCmd.AddCommand(natssubscriberCmd)
	natssubscriberCmd.Flags().StringP("topic", "t", "", "Preferred Nats topic to subscribe")
}
