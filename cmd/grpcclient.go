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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	models "github.com/enbis/message-broker-miscellaneous/models/src"
)

// grpcclientCmd represents the grpcclient command
var grpcclientCmd = &cobra.Command{
	Use:   "grpcclient",
	Short: "Launch gRPC Client",
	Long: `The gRPC Client interact with the gRPC Server. 
	It sends the protocol buffer which contains all the necessary information for the Nats communication`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("grpcclient called")

		var conn *grpc.ClientConn

		dialPort := fmt.Sprintf(":%d", viper.GetInt("grpc_port"))

		conn, err := grpc.Dial(dialPort, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("did not connect: %s", err)
		}
		defer conn.Close()
		c := models.NewRedirectClient(conn)

		_, err = c.Send(context.Background(), &models.PingMessage{Topic: viper.GetString("topic"), Payload: []byte(viper.GetString("payload"))})
		if err != nil {
			log.Fatalf("Error sending the data to gRPC Server: %s", err)
		}

		log.Printf("Information %s %s sent to the gRPC Server \n", viper.GetString("topic"), viper.GetString("payload"))

	},
}

func init() {
	rootCmd.AddCommand(grpcclientCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// grpcclientCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// grpcclientCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
