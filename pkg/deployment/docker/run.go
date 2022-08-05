package docker

import (
	"fmt"

	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/styles"
)

func Install() error {
	msg := "🚀 Great job!\n\nTry AWS option"
	fmt.Println(
		styles.SuccessBox.Inherit(styles.SuccessTextStyle).Render(msg),
	)
	return nil
}
