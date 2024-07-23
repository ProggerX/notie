package viewer

import (
	"net"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"
)

func isFileExist(path string) bool {
	if _, err := os.Stat(".notie/notes/" + path); err == nil {
		return true
	}
	return false
}

func StartViewer(host string, port string) {
	srv, _ := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),

		wish.WithHostKeyPath(".notie/.ssh/viewer/id_ed25519"),

		wish.WithMiddleware(
			func(ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					if len(s.Command()) != 1 {
						wish.Println(s, "I expected only one argument - name of note")
					} else if isFileExist(s.Command()[0]) && !strings.Contains(s.Command()[0], "/") {
						bts, _ := os.ReadFile(".notie/notes/" + s.Command()[0])
						out, _ := glamour.Render(string(bts), "dark")
						wish.Println(s, out)
					} else {
						wish.Printf(s, "No such note: %s\n\n", s.Command()[0])
					}
				}
			},
			logging.Middleware(),
		),
	)

	srv.ListenAndServe()
}
