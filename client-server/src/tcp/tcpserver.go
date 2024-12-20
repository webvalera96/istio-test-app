package tcp

import (
	"client-server/src/logger"
	"fmt"
	"log"
	"net"
	"os"
)

func TcpServer(host, port string) {
	listen, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// close listener
	defer listen.Close()
	logger.Log.Info(fmt.Sprintf("Listening on %s...", host+":"+port))
	for {
		conn, err := listen.Accept()
		logger.Log.Sugar().Infof("Connection from: %s to %s", conn.RemoteAddr(), conn.LocalAddr())
		if err != nil {
			logger.Log.Sugar().Error(err)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	// incoming request
	buffer := make([]byte, 1024)
	receivedBytes, err := conn.Read(buffer)
	logger.Log.Sugar().Infof("Received: %d bytes; Data: %s", receivedBytes, buffer)
	if err != nil {
		logger.Log.Sugar().Error(err)
		return
	}

	// write data to response
	responseStr := fmt.Sprintf("Data: %v", string(buffer[:]))
	sendBytes, err := conn.Write([]byte(responseStr))
	if err != nil {
		logger.Log.Sugar().Error(err)
		return
	}
	logger.Log.Sugar().Infof("Send: %d bytes; %s", sendBytes, responseStr)
}
