package cmd

import (
	cmd "client-server/cmd/cmdTcp"

	"github.com/spf13/cobra"
)

var tcpCmd = &cobra.Command{
	Use:   "tcp",
	Short: "Tcp client/server",
}

func init() {

	tcpCmd.AddCommand(cmd.TcpServerCmd)
	tcpCmd.AddCommand(cmd.TcpClientCmd)
	rootCmd.AddCommand(tcpCmd)
}
