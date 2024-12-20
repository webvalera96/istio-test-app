package tcp

import (
	"client-server/src/logger"
	"net"
	"time"
)

func TcpClient(host, port string, timeout time.Duration) {
	logger.Log.Sugar().Infof("Current timeout: %v", timeout)
	tcpAddr, err := net.ResolveTCPAddr("tcp", host+":"+port)
	if err != nil {
		logger.Log.Sugar().Info(err)
	}

	// TODO: написать генератор нагрузки, пока "Hello"
	for {
		conn, err := net.DialTCP("tcp", nil, tcpAddr)
		if err != nil {
			logger.Log.Sugar().Info(err)
			logger.Log.Sugar().Info("retrying ... ")
			time.Sleep(timeout)
			continue
		}

		sendBytes, err := conn.Write([]byte("Hello"))
		if err != nil {
			logger.Log.Sugar().Info(err)
		}
		logger.Log.Sugar().Infof("Sent: %d; Msg: Hello", sendBytes)

		// получение результата
		reply := make([]byte, 1024)

		receivedBytes, err := conn.Read(reply)
		if err != nil {
			logger.Log.Sugar().Info(err)
		}

		logger.Log.Sugar().Infof("Received: %d; Msg: %s", receivedBytes, reply)

		conn.Close()

		// timeout
		time.Sleep(timeout)
	}

}
