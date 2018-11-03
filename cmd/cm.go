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

// cmCmd represents the cm command
var cmCmd = &cobra.Command{
	Use:   "cm",
	Short: "create a meeting",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Create Meeting Called")
		title, _ := cmd.Flags().GetString("title")
		participator, _ := cmd.Flags().GetStringSlice("user")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		if title == "" || len(participator) == 0 || starttime == "" || endtime == "" {
			fmt.Println("Please input title, starttime[yyyy-mm-dd/hh:mm], endtime, user(s)(input like \"name1, name2\")")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login")
			return
		} else {
			if flag := service.CreateMeeting(user.Name, title, starttime, endtime, participator); flag != true {
				fmt.Println("Error: create Failed. Please check error.log for more detail")
				return
			} else {
				fmt.Println("Create meeting successfully!")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(cmCmd)
	cmCmd.Flags().StringP("title", "t", "", "the title of meeting")
	cmCmd.Flags().StringSliceP("user", "u", nil, "the user(s) of the meeting, input like \"name1, name2\"")
	cmCmd.Flags().StringP("starttime","s","","the startTime of the meeting")
	cmCmd.Flags().StringP("endtime", "e", "", "the endTime of the meeting")
}
