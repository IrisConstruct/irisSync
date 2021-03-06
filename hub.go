package iris

import (
	"fmt"
	"net/http"

	"github.com/jsimnz/wsHub"
	"github.com/likexian/simplejson"
)

type SyncBoard struct {
	hub   wsHub.WsHub
	admin *wsHub.Client
	kill  chan bool
}

// Construct
func NewSyncBoard() SyncBoard {
	s := SyncBoard{
		hub:  wsHub.NewHub(),
		kill: make(chan bool),
	}
	return s
}

// Run loop
func (s SyncBoard) Run() {

}

// Stop
func (s SyncBoard) Kill() {
	s.kill <- true
}

// Client screen connection
func (s SyncBoard) IrisClient(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got new screen client")
	client, err := wsHub.NewClient(w, r)
	if err != nil {
		fmt.Println("Error getting websocket connection:", err)
	}

	s.hub.RegisterClient(client)
	go client.Start()
	defer func() {
		fmt.Println("Screen disconnected")
		s.hub.UnregisterClient(client)
	}()

	for {
		cmd, err := client.Read()
		if err != nil {
			fmt.Println("Couldnt read command from screen:", err)
			break
		}

		cmdjson, err := simplejson.Loads(string(cmd))
		if cmdjson.Exists("cmd") {
			cmdString, err := cmdjson.Get("cmd").String()
			if err != nil {
				fmt.Println("Invalid command from screen")
			} else {
				switch cmdString {
				case "PAUSE":
					// pause playback
				case "PLAY":
					// start/resume playback
				case "PLAY_AT":
					// start/resume playback @
				case "MAKE_ADMIN":
					// upgrade connection to admin
				}
			}
		}

	}
}

// Admin connection
func (s SyncBoard) IrisAdmin(w http.ResponseWriter, r *http.Request) {

}

// Handlers
func (s SyncBoard) UpgradeToAdmin(client *wsHub.Client) {
	s.admin = admin

	defer func() {
		fmt.Println("Upgrading connection to admin")
		s.hub.UnregisterClient(s.admin)
	}()

	for {
		cmd, err := s.admin.ReadString()
		if err != nil {
			break
		}

		switch cmd {

		}
	}

}

// Utils
