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

// loCmd represents the lo command
var loCmd = &cobra.Command{
	Use:   "lo",
	Short: "logout",
	Run: func(cmd *cobra.Command, args []string) {
		errLog.Println("Logout called")
		if err := service.UserLogout(); err != true {
			fmt.Println("Some error happened when log out, please read error.log for details")
		} else {
			fmt.Println("Logout Successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(loCmd)
}
