package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
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

		log.Printf("workMode = %s, address = %s, msgLen = %v, delayBetweenMsgsMs = %v", workMode, address, msgLen, delayBetweenMsgsMs)
		startClient(address, msgLen, delayBetweenMsgsMs)
	} else {

		// Server

		listenPort := os.Getenv("LISTEN_PORT")
		if len(listenPort) == 0 {
			listenPort = "80"
		}
		address := "127.0.0.1" + ":" + listenPort
		log.Printf("workMode = %s, address = %s", workMode, address)
		startServer(address)
	}
}
