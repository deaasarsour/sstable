package util

func RunInLoop(exec func()) {
	for {
		exec()
	}
}
