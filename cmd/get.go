/*
Copyright © 2020  Mike Herbert <mikeherbert.dev@gmail.com>

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
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Does an HTTP get and returns the body",
	Long:  `cli-go get will attempt an HTTP request and return the body to stdout`,
	Run: func(cmd *cobra.Command, args []string) {
		verbose, _ := rootCmd.Flags().GetBool("verbose")
		if verbose {
			fmt.Println("Running get command")
		}
		url, err := cmd.Flags().GetString("url")
		if err != nil {
			// handle error
		}
		getURL(url)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("url", "u", "", "URL to be requested")
}

func getURL(url string) {
	if len(url) == 0 {
		fmt.Println("You must pass in a url if you are enabling the flag --url (-u)")
		os.Exit(1)
	}
	println(url)
}
