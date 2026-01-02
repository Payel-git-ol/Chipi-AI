package logerr

import "log"

type Logerr struct{}

func (l Logerr) LogerrStatusOk(method string, path string, status string) {
	log.Printf(
		"%s %s %s",
		"      "+"| "+method+" |",
		"      "+"| "+path+" |",
		"      "+"| "+status+" |",
	)
}

func (l Logerr) LogerrStatusBad(method string, path string, status string, err error) {
	log.Printf(
		"%s %s %s",
		"      "+"| "+method+" |",
		"      "+"| "+path+" |",
		"      "+"| "+status+" |",
		"      "+"| "+err.Error()+" |",
	)
}
