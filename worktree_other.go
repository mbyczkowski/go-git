// +build !linux,!darwin,!freebsd,!netbsd,!openbsd

package git

import (
	"syscall"

	"srcd.works/go-git.v4/plumbing/format/index"
)

func fillSystemInfo(e *index.Entry, os *syscall.Stat_t) {

}
