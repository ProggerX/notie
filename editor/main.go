package editor

import (
	"net"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/bubbletea"
	"github.com/charmbracelet/wish/logging"
)

func isFileExist(path string) bool {
	if _, err := os.Stat(".notie/notes/" + path); err == nil {
		return true
	}
	return false
}

func StartEditor(host string, port string) {
	srv, _ := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),

		wish.WithHostKeyPath(".notie/.ssh/editor/id_ed25519"),

		wish.WithMiddleware(
			func(ssh.Handler) ssh.Handler {
				return func(s ssh.Session) {
					wish.Println(s, "Good luck! To view saved file use viewer port and note like command")
				}
			},
			bubbletea.Middleware(teaHandler),
			activeterm.Middleware(),
			logging.Middleware(),
		),
	)

	srv.ListenAndServe()
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	pty, _, _ := s.Pty()

	ni := textinput.New()
	ni.Focus()

	ta := textarea.New()
	ta.Blur()
	ta.CharLimit = 0
	ta.SetWidth(pty.Window.Width - 6)
	ta.SetHeight(pty.Window.Height - 6)

	m := model{
		Term:      pty.Term,
		Width:     pty.Window.Width,
		Height:    pty.Window.Height,
		NameInput: ni,
		TextArea:  ta,
	}

	return m, []tea.ProgramOption{tea.WithAltScreen()}
}
