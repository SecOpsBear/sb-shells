/*
Copyright © 2023 SecOpsBear bear@secopsbear.com

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

	"github.com/secopsbear/sb-shells/colors"
	"github.com/secopsbear/sb-shells/validators"
	"github.com/spf13/cobra"
)

var (
	outputFile     string
	ipAddressLocal string
	portLocal      string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sb-shells",
	Short: "Generate shell codes for various environments Powershell, bash, php..",
	Long:  `Generate shell codes for various environments Powershell, bash, php..`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&ipAddressLocal, "ipAddrLocal", "i", "", "Local IP address")
	rootCmd.MarkPersistentFlagRequired("ipAddrLocal")
	rootCmd.PersistentFlags().StringVarP(&portLocal, "portLocal", "p", "", "Local Port")
	rootCmd.MarkPersistentFlagRequired("portLocal")
}

// Check IP address and ports are valid format.
func validateIpAndPorts(cmd *cobra.Command, args []string) {
	notValidInputs := false
	if !validators.IsIPaddressValid(ipAddressLocal) {
		fmt.Println(colors.RedColor("Error: ") + colors.YellowColor(ipAddressLocal) + colors.RedColor(" is not a  valid IP Address"))
		notValidInputs = true
	}
	if !validators.IsIPaddressPortValid(portLocal) {
		fmt.Println(colors.RedColor("Error: ") + colors.YellowColor(portLocal) + colors.RedColor(" is not a  valid port"))
		notValidInputs = true
	}

	if notValidInputs {
		cmd.Help()
		os.Exit(1)
	}

}
