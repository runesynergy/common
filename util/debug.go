package util

import (
	"runtime"

	"github.com/minio/highwayhash"
)

var key = []byte("MEDIEVALSOFTWAREMEDIEVALSOFTWARE")

func UID() (uid uint32) {
	var rpc [1]uintptr
	runtime.Callers(3, rpc[:])
	frame, _ := runtime.CallersFrames(rpc[:]).Next()
	sum := highwayhash.Sum64([]byte(frame.File), key)
	sum += highwayhash.Sum64([]byte(frame.Function), key)
	uid += uint32(sum >> 32)
	uid += uint32(sum) + uint32(frame.Line)
	return
}
