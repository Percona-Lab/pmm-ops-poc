package aws

import (
	"fmt"

	bubblesList "github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/list"
	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/styles"
)

var (
	resourceEKS   = list.Item("AWS EKS")
	resourceRocky = list.Item("AWS EC2 - Rocky Linux")
	resourceAMI   = list.Item("AWS EC2 - PMM AMI")
)

func renderResourcesList(m list.Model) string {
	if m.Choice != "" {
		return styles.QuitTextStyle.Render(fmt.Sprintf("%s? Great choice.", m.Choice))
	}
	if m.Quitting {
		return styles.QuitTextStyle.Render("Not hungry? That's cool.")
	}
	return "\n" + m.List.View()
}

func selectResourceType() (list.Item, error) {
	items := []bubblesList.Item{
		resourceEKS,
		resourceRocky,
		resourceAMI,
	}

	m := list.New(items, renderResourcesList)
	pm, err := tea.NewProgram(m).StartReturningModel()
	if err != nil {
		return list.Item(""), err
	}

	m = pm.(list.Model)
	selected, ok := m.List.SelectedItem().(list.Item)
	if !ok {
		return list.Item(""), fmt.Errorf("cannot find selected choice")
	}

	return selected, nil
}
