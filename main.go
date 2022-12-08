package main

import (
	"ciunikine/server"
	"io"
	"log"
	"os"
	"syscall"
	"unsafe"

	"github.com/gliderlabs/ssh"
)

func setWinsize(f *os.File, w, h int) {
	syscall.Syscall(syscall.SYS_IOCTL, f.Fd(), uintptr(syscall.TIOCSWINSZ),
		uintptr(unsafe.Pointer(&struct{ h, w, x, y uint16 }{uint16(h), uint16(w), 0, 0})))
}

func main() {
	r := server.NewRouter()
	go func() {
		log.Println("starting ssh server on port 2222...")

		forwardHandler := &ssh.ForwardedTCPHandler{}

		server := ssh.Server{
			LocalPortForwardingCallback: ssh.LocalPortForwardingCallback(func(ctx ssh.Context, dhost string, dport uint32) bool {
				log.Println("Accepted forward", dhost, dport)
				return true
			}),
			Addr: ":2222",
			Handler: ssh.Handler(func(s ssh.Session) {
				io.WriteString(s, "Remote forwarding available...\n")
				select {}
			}),
			ReversePortForwardingCallback: ssh.ReversePortForwardingCallback(func(ctx ssh.Context, host string, port uint32) bool {
				log.Println("attempt to bind", host, port, "granted")
				return true
			}),
			RequestHandlers: map[string]ssh.RequestHandler{
				"tcpip-forward":        forwardHandler.HandleSSHRequest,
				"cancel-tcpip-forward": forwardHandler.HandleSSHRequest,
			},
		}

		log.Fatal(server.ListenAndServe())
	}()

	r.Run(":8080")
}
