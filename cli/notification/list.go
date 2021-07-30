/*
Copyright © 2021 Andy Zhou

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
package notification

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
	u "github.com/z-t-y/flogo/utils"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your notifications",
	Long: `List all your notifications to the output.

Example Output:

Message: Lorem Ipsum
Time:    Wed, 28 Jul 2021 08:24:23 UTC

Note that the time is NOT displayed in your local time.
`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken, err := u.GetLocalAccessToken()
		cobra.CheckErr(err)
		notifications, err := GetNotificationList(accessToken)
		cobra.CheckErr(err)
		fmt.Println(len(notifications))
		for _, notification := range notifications {
			fmt.Println(u.Segmenter)
			fmt.Println("Message:", notification.Message)
			fmt.Println("Time:   ", notification.Time.Format(time.RFC1123Z))
		}
	},
}

func GetNotificationList(accessToken string) (notifications []u.Notification, err error) {
	client := http.Client{}
	req, err := http.NewRequest("GET", u.URLFor("/api/v3/notification/all"), nil)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	err = u.CheckStatusCode(resp, 200)
	if err != nil {
		return
	}
	json.NewDecoder(resp.Body).Decode(&notifications)
	return
}

func init() {
	notificationCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
