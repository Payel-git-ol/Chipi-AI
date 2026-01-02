package logerr

type LoggerRepository interface {
	LogerrStatusOk(method string, path string, status string)
	LogerrStatusBad(method string, path string, status string, err error)
}
