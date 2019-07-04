package error

import "fmt"

func Report(did string, err error) error {
	return fmt.Errorf("failed to %s: %s", did, err)
}
