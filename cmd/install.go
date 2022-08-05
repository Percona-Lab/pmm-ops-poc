package cmd

import (
	"fmt"

	bubblesList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/list"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/styles"
	"github.com/percona-lab/pmm-ops-poc/pkg/deployment/aws"
	"github.com/percona-lab/pmm-ops-poc/pkg/deployment/docker"
	"github.com/spf13/cobra"
)

var (
	installCmd = &cobra.Command{
		Use:   "install",
		Short: "Run installation wizard",
		RunE:  runInstallCmd,
	}

	targetDocker = list.Item("Docker")
	targetAWS    = list.Item("AWS deployment")
	targetGCP    = list.Item("Google Cloud deployment")
)

func init() {
	rootCmd.AddCommand(installCmd)
}

func renderInstallList(m list.Model) string {
	if m.Choice != "" {
		return styles.QuitTextStyle.Render(fmt.Sprintf("%s? Sounds great 👍", m.Choice))
	}
	if m.Quitting {
		return styles.QuitTextStyle.Render("See you later maybe?")
	}
	return "\n" + m.List.View()
}

func runInstallCmd(cmd *cobra.Command, args []string) error {
	items := []bubblesList.Item{
		targetDocker,
		targetAWS,
		targetGCP,
	}

	m := list.New(items, renderInstallList)
	m.List.Title = "Select deployment method:"

	pm, err := tea.NewProgram(m).StartReturningModel()
	if err != nil {
		return err
	}

	m = pm.(list.Model)
	selected, ok := m.List.SelectedItem().(list.Item)
	if !ok {
		return fmt.Errorf("cannot find selected choice")
	}

	switch selected {
	case targetDocker:
		return docker.Install()
	case targetAWS:
		return aws.Wizard()
	default:
		msg := "🚀 Great job!\n\nTry AWS option"
		fmt.Println(
			styles.SuccessBox.Inherit(styles.SuccessTextStyle).Render(msg),
		)
	}

	return nil
}
