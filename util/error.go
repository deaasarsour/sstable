package util

func TryRunAll(execs ...func() error) error {
	for _, exec := range execs {
		if err := exec(); err != nil {
			return err
		}
	}
	return nil
}
