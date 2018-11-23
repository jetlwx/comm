package ftp

import (
	"time"

	"github.com/jlaffaye/ftp"
)

func Conn(addr string, timemout time.Duration) (*ftp.ServerConn, error) {
	return ftp.DialTimeout(addr, timemout)
}
