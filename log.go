package main // import "github.com/dayvillefire/iardisplay"

import (
	"io"
	"sync"
)

type LockedWriter struct {
	w io.Writer
	*sync.Mutex
}

func (l LockedWriter) Write(b []byte) (int, error) {
	l.Lock()
	defer l.Unlock()
	return l.w.Write(b)
}

// usage
// l := LockedWriter{&lumberjack.Logger{ /*config*/ } }
// use your locked writer
