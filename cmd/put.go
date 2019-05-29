// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

// putCmd represents the put command
var putCmd = &cobra.Command{
	Use:   "put <PATH> <REMOTE_PATH>",
	Short: "Put an Object",
	Long: `Put an Object`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {

		bucket := viper.GetString("cos_bucket")
		if bucket == "" {
			panic(errors.New("bucket invalid"))
		}

		region := viper.GetString("cos_region")
		if region == "" {
			panic(errors.New("region invalid"))
		}

		secretId := viper.GetString("cos_secret_id")
		if secretId == "" {
			panic(errors.New("secretId invalid"))
		}

		secretKey := viper.GetString("cos_secret_key")
		if secretKey == "" {
			panic(errors.New("secretKey invalid"))
		}

		path := args[0]
		remotePath := args[1]

		reader, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer reader.Close()

		bucketUrl, err := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucket, region))
		if err != nil {
			panic(err)
		}

		client := cos.NewClient(&cos.BaseURL{
			BucketURL: bucketUrl,
		}, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID: secretId,
				SecretKey: secretKey,
			},
		})

		_, err = client.Object.Put(context.Background(), remotePath, reader, nil)
		if err != nil {
			panic(err)
		}

		log.Println(fmt.Sprintf("%s ---> %s", path, remotePath))
	},
}

func init() {
	rootCmd.AddCommand(putCmd)
}
