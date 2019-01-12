// Copyright © 2019 nephe
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

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use: "get",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("get called with all flag: %v and %s output\n", all, output)
	},
}
var all bool
var output string

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVar(&all, "all", false, "Help message for all")
	getCmd.PersistentFlags().StringVar(&output, "output", "default", "Type of output")
}
