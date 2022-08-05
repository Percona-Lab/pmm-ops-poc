package aws

import (
	"fmt"

	"github.com/percona-lab/pmm-ops-poc/pkg/bubbles/styles"
)

func Wizard() error {
	resource, err := selectResourceType()
	if err != nil {
		return err
	}

	err = getCredentials()
	if err != nil {
		return err
	}

	err = deployAny()
	if err != nil {
		return err
	}

	successMsg := fmt.Sprintf(`🚀 Great job! PMM is now available in %s.

Visit https://pmm.aws.eks to sign in.`, resource)

	fmt.Println(
		styles.SuccessBox.Inherit(styles.SuccessTextStyle).Render(successMsg),
	)

	return nil
}
