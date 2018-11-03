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

// fmCmd represents the fm command
var fmCmd = &cobra.Command{
	Use:   "fm",
	Short: "find meetings that are bewteen starttime and endtime",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Query Meeting called")
		starttime, _ := cmd.Flags().GetString("starttime")
		endtime, _ := cmd.Flags().GetString("endtime")
		if starttime == "" || endtime == "" {
			fmt.Println("Please input start time and end time both")
			return
		}
		if user, flag := service.GetCurUser(); flag != true {
			fmt.Println("Error: please login")
		} else {
			if ml, flag := service.QueryMeeting(user.Name, starttime, endtime); flag != true {
				fmt.Println("Error. Please input the date as yyyy-mm-dd/hh:mm and make sure that starttime is before endtime")
			} else {
				for _, m := range ml {
					fmt.Println("----------------")
					fmt.Println("Title:", m.Title)
					fmt.Println("Start Time:", m.StartTime.ToString())
					fmt.Println("End Time:", m.EndTime.ToString())
					fmt.Printf("Participator(s): ")
					for _, p := range m.Participators {
						fmt.Printf("%s ", p)
					}
					fmt.Printf("\n")
					fmt.Println("----------------")
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(fmCmd)
	fmCmd.Flags().StringP("starttime", "s", "", "the start time of the meeting")
	fmCmd.Flags().StringP("endtime", "e", "", "the end time of the meeting")
}
