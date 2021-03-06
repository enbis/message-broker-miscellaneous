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

	"github.com/spf13/cobra"
	"github.com/streadway/amqp"
)

// rmqreceivingCmd represents the rmqreceiving command
var rmqreceivingCmd = &cobra.Command{
	Use:   "rmqreceiving",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rmqreceiving called")

		conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
		if err != nil {
			log.Fatal("Failed to connect to RabbitMQ ", err)
		}
		defer conn.Close()

		ch, err := conn.Channel()
		if err != nil {
			log.Fatal("Failed to open a channel ", err)
		}
		defer ch.Close()

		q, err := ch.QueueDeclare(
			"hello1234",
			false,
			false,
			false,
			false,
			nil,
		)
		if err != nil {
			log.Fatal("Failed to declare a Queue ", err)
		}

		msgs, err := ch.Consume(
			q.Name, // queue
			"",     // consumer
			true,   // auto-ack
			false,  // exclusive
			false,  // no-local
			false,  // no-wait
			nil,    // args
		)
		if err != nil {
			log.Fatal("Failed to declare a consumer ", err)
		}

		//Reading from empty channel allow block program until there's nothing to receive.
		//Like the infinite loop, but not using the 100% of the CPU
		forever := make(chan bool)

		go func() {
			for d := range msgs {
				log.Printf("RabbitMQ received message: %s\n", d.Body)
			}
		}()

		<-forever

	},
}

func init() {
	rootCmd.AddCommand(rmqreceivingCmd)
}
