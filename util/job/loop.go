package jobutil

func RunInLoop(exec func()) {
	for {
		exec()
	}
}
