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

// mqttpublisherCmd represents the mqttpublisher command
var mqttpublisherCmd = &cobra.Command{
	Use:   "mqttpublisher",
	Short: "This command allows to publish a message to a specific MQTT topic",
	Long: `This command is useful for testing the MQTT messaging system standalone.
	First of all launch the mqttsubscriber command selecting preferred topic, 
	then launch the mqttpublisher command specifying the message being sure to publish the message
	on the correct topic. There are two parameters, topic and payload.`,
	Run: func(cmd *cobra.Command, args []string) {

		topic, err := cmd.Flags().GetString("topic")
		if err != nil {
			topic = viper.GetString("topic")
		}

		payload, err := cmd.Flags().GetString("payload")
		if err != nil {
			payload = viper.GetString("payload")
		}

		mqtt := api.NewMqttTransport()
		err = mqtt.Connect()
		if err != nil {
			log.Fatal("Error publishing on MQTT ", err)
		}

		err = mqtt.Publish(topic, []byte(payload))
		if err != nil {
			log.Fatal("Error publishing on MQTT ", err)
		}
		fmt.Println(fmt.Sprintf("MQTT - Published %s to topic %s", payload, topic))

	},
}

func init() {
	rootCmd.AddCommand(mqttpublisherCmd)
	mqttpublisherCmd.Flags().StringP("topic", "t", "", "Preferred MQTT topic on which to publish the message")
	mqttpublisherCmd.Flags().StringP("payload", "p", "", "Preferred MQTT payload to publish")
}
