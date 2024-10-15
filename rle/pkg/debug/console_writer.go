package debug

type ConsoleWriter struct{}

func (c *ConsoleWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}
