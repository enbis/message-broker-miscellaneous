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

// natspublisherCmd represents the natspublisher command
var natspublisherCmd = &cobra.Command{
	Use:   "natspublisher",
	Short: "This command allows to publish a message to a specific Nats topic",
	Long: `This command is useful for testing the NATS messaging system standalone.
	First of all launch the natssubscriber selecting preferred topic, 
	then launch the natspublisher specifying the message being sure to publish the message
	on the correct topic. There are two parameters, topic and payload.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("natspublisher called")

		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			topic = viper.GetString("topic")
		}

		payload, err := cmd.Flags().GetString("payload")
		if err != nil {
			payload = viper.GetString("payload")
		}

		nats := api.NewNatsTransport()

		err = nats.Publish(topic, []byte(payload))
		if err != nil {
			log.Fatal("Error publishing on Nats ", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(natspublisherCmd)
	natspublisherCmd.Flags().StringP("topic", "t", "", "Preferred Nats topic on which to publish the message")
	natspublisherCmd.Flags().StringP("payload", "p", "", "Preferred Nats payload to publish")
}
