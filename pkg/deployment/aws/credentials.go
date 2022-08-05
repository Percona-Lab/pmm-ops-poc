package aws

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/textinput"
)

func getCredentials() error {
	m := textinput.New()
	m.Label = "AWS Access Key:"

	if err := tea.NewProgram(m).Start(); err != nil {
		return err
	}

	m = textinput.NewPassword()
	m.Label = "AWS Secret Key:"

	if err := tea.NewProgram(m).Start(); err != nil {
		return err
	}

	fmt.Print("\n\n")
	return nil
}
