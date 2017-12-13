// Hermetix
// augustzf@gmail.com

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type handler func(string, string) error

func main() {
	user := os.Getenv("USER")
	if user == "" {
		panic("Missing env: USER")
	}
	host := os.Getenv("HOST")
	if host == "" {
		panic("Missing env: HOST")
	}
	handler := func(rec, msg string) error {
		return hostExec(user, host, sendMsg(rec, msg))
	}
	serve(handler)
}

// process http requests on port 8009
func serve(fn handler) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rec := r.URL.Query().Get("rec")
		msg := r.URL.Query().Get("msg")
		if rec == "" || msg == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fn(rec, msg)
		w.WriteHeader(http.StatusOK)
	})
	fmt.Println("Listening on port 8009")
	log.Fatal(http.ListenAndServe(":8009", nil))
}

// run command on Docker host
func hostExec(username, host, hostCmd string) error {
	// this assumes the user's public key has been added to ~/.ssh/authorized_keys
	cmd := exec.Command("ssh", "-o", "StrictHostKeyChecking=no", "-i", "./ssh/id_rsa",
		fmt.Sprintf("%v@%v", username, host), hostCmd)
	return cmd.Run()
}

// applescript to send message via the Messages app
func sendMsg(rec, msg string) string {
	f := `osascript <<'END'
		tell application "Messages"
			set targetService to 1st service whose service type = iMessage
			set rec to buddy %q of targetService
			send %q to rec
		end tell%vEND`
	return fmt.Sprintf(f, rec, msg, "\n")
}
