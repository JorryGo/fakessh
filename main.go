package main

import (
	"fmt"
	"github.com/gliderlabs/ssh"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	logFile := "connections.log"
	file, err := os.OpenFile(logFile,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	ssh.Handle(func(s ssh.Session) {
		io.WriteString(s, "Welcome to Ubuntu 18.04.4 LTS (GNU/Linux 4.15.0-101-generic x86_64)\n")
		io.WriteString(s, s.User()+"@yourMom:~# fuck you\n")

		file.WriteString(time.Now().Format("2006-01-02 15:04:05") + ": " + s.User())
		file.WriteString(" - " + s.RemoteAddr().String() + "\n")
		file.Sync()

		select {
		case <-s.Context().Done():
			fmt.Println("close")
			return
		}
	})

	log.Fatal(ssh.ListenAndServe(":22", nil))
}
