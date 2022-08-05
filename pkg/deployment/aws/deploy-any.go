package aws

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/progress"
)

func deployAny() error {
	m := progress.NewLinear()
	m.Title = "Provisioning EKS..."

	if err := tea.NewProgram(m).Start(); err != nil {
		return err
	}

	m = progress.NewLinear()
	m.Title = "Deploying PMM..."

	if err := tea.NewProgram(m).Start(); err != nil {
		return err
	}

	m = progress.NewLinear()
	m.Title = "Waiting for PMM to be ready..."

	if err := tea.NewProgram(m).Start(); err != nil {
		return err
	}

	return nil
}
