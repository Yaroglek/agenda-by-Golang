// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"agenda/service"
	"github.com/spf13/cobra"
)

// dmCmd represents the dm command
var dmCmd = &cobra.Command{
	Use:   "dm",
	Short: "delete a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Delete Meeting called")
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			fmt.Println("Error: Please input meeting title")
			return
		}
		if user,flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login")
		} else {
			if c, f := service.DeleteMeeting(user.Name, title); f == false {
				return
			} else if c == 0 {
				fmt.Println("Error: Meeting not exist or you're not a Sponsor of it")
			} else {
				fmt.Println("Delete Successfully")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dmCmd)
	dmCmd.Flags().StringP("title", "t", "", "the title of meeting")
}
