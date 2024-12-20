package cmdTcp

import (
	"client-server/src/tcp"
	"github.com/spf13/cobra"
	"log"
	"time"
)

var TcpClientCmd = &cobra.Command{
	Use:   "client",
	Short: "Simple tcp client",
	Run: func(cmd *cobra.Command, args []string) {

		// get hostname from flag
		host, err := cmd.Flags().GetString("host")
		if err != nil {
			log.Fatal(err)
		}

		// get port from flag
		port, err := cmd.Flags().GetString("port")
		if err != nil {
			log.Fatal(err)
		}

		// get timeout from flag
		timeout, err := cmd.Flags().GetString("timeout")
		if err != nil {
			log.Fatal(err)
		}
		timeDuration, err := time.ParseDuration(timeout)
		if err != nil {
			log.Fatal(err)
		}

		tcp.TcpClient(host, port, timeDuration)
	},
}

func init() {
	TcpClientCmd.Flags().String("host", "127.0.0.1", "tcp server hostname")
	TcpClientCmd.Flags().String("port", "9090", "tcp port for listening")
	TcpClientCmd.Flags().String("timeout", "3s", "timeout before tcp requests")
}
