package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	workMode := os.Getenv("WORK_MODE")

	if strings.EqualFold(workMode, "client") {

		// Client

		targetPort := os.Getenv("TARGET_PORT")
		if len(targetPort) == 0 {
			targetPort = "80"
		}

		targetHost := os.Getenv("TARGET_HOST")
		if len(targetHost) == 0 {
			targetHost = "localhost"
		}

		address := targetHost + ":" + targetPort

		msgLenStr := os.Getenv("MSG_LEN")
		msgLen, err := strconv.Atoi(msgLenStr)
		if err != nil {
			msgLen = 32000
		}

		delayBetweenMsgsStr := os.Getenv("DELAY_BETWEEN_MSGS_MS")
		delayBetweenMsgsMs, err := strconv.Atoi(delayBetweenMsgsStr)
		if err != nil {
			delayBetweenMsgsMs = 2000
		}

		//writeBufferSizeStr := os.Getenv("WR_BUF")
		//writeBufferSize, err := strconv.Atoi(writeBufferSizeStr)
		//if err != nil {
		//	writeBufferSize = 64 * 1024 * 1024
		//}
		//
		//initialWindowSizeStr := os.Getenv("WIN_SIZE")
		//initialWindowSize, err := strconv.Atoi(initialWindowSizeStr)
		//if err != nil {
		//	initialWindowSize = 128 * 1024 * 1024
		//}
		//
		//initialConnWindowSizeStr := os.Getenv("CONN_WIN_SIZE")
		//initialConnWindowSize, err := strconv.Atoi(initialConnWindowSizeStr)
		//if err != nil {
		//	initialConnWindowSize = 128 * 1024 * 1024
		//}

		//log.Printf("workMode = %s, address = %s, msgLen = %v, delayBetweenMsgsMs = %v, writeBufferSize = %v, initialWindowSize = %v, initialConnWindowSize = %v", workMode, address, msgLen, delayBetweenMsgsMs, writeBufferSize, initialWindowSize, initialConnWindowSize)
		//startClient(address, msgLen, delayBetweenMsgsMs, writeBufferSize, int32(initialWindowSize), int32(initialConnWindowSize))
		startClient(address, msgLen, delayBetweenMsgsMs)
	} else {

		// Server

		listenPort := os.Getenv("LISTEN_PORT")
		if len(listenPort) == 0 {
			listenPort = "80"
		}
		address := "0.0.0.0" + ":" + listenPort
		log.Printf("workMode = %s, address = %s", workMode, address)
		startServer(address)
	}
}
