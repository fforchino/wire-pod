package main

import (
	"github.com/digital-dream-labs/vector-cloud/internal/clad/cloud"
	"github.com/digital-dream-labs/vector-cloud/internal/ipc"
	"github.com/digital-dream-labs/vector-cloud/internal/voice"
	"log"
	"time"
)

func TriggerWakeWord() {
	var cloudSock ipc.Conn

	log.Println("Creating Test Client Socket to send messages to vic-cloud")
	cloudSock = getSocketWithRetry(ipc.GetSocketPath("cp_test"), "cp_test_client")
	defer cloudSock.Close()
	log.Println("Socket created")

	location_currentzone, _ := time.LoadLocation("Europe/Paris")
	log.Println("Triggering hotword: " + location_currentzone.String())
	hw := cloud.Hotword{Mode: cloud.StreamType_Normal, Locale: "en-US", Timezone: location_currentzone.String(), NoLogging: true}
	message := cloud.NewMessageWithHotword(&hw)

	log.Println("Creating sender")
	testSender := voice.IPCMsgSender{Conn: cloudSock}
	log.Println("Sending message")
	testSender.Send(message)
	log.Println("DONE")
}

func main() {
	log.Println("Now triggering wake word")
	TriggerWakeWord()
}

func getSocketWithRetry(name string, client string) ipc.Conn {
	for {
		sock, err := ipc.NewUnixgramClient(name, client)
		if err != nil {
			log.Println("Couldn't create socket", name, "- retrying:", err)
			time.Sleep(5 * time.Second)
		} else {
			return sock
		}
	}
}
