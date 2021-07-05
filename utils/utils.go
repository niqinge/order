package utils

import (
	"fmt"
	"go/build"
	"time"
)

func FilePath(fileName string) string {
	return fmt.Sprintf("%s/%d_%s", build.Default.GOPATH, time.Now().UnixNano(), fileName)
}
