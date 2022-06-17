package snap

import (
	"runtime"
)

type Status struct {
	host       string
	canRunSnap bool
}

var status Status

func init() {
	host := runtime.GOOS
	status = Status{
		host:       host,
		canRunSnap: host == "linux",
	}
}

func CanRunSnap() (string, bool) {
	return status.host, status.canRunSnap
}

func HasSystemd() bool {
	if !status.canRunSnap {
		// doesn't matter any way
		return false
	}
	return true
}
