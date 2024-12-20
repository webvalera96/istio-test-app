package cmdTcp

import (
	"client-server/src/tcp"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var TcpServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Simple tcp server",
	Run: func(cmd *cobra.Command, args []string) {

		// get hostname from flag
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}

		// get port from flag
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		tcp.TcpServer(host, port)
	},
}

func init() {
	TcpServerCmd.Flags().String("host", "127.0.0.1", "tcp server hostname")
	TcpServerCmd.Flags().String("port", "9090", "tcp port for listening")
}
