package log_tool

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

func NewLogTool() *logrus.Logger {
	log := logrus.New()
	log.SetReportCaller(true)
	formatter := &logrus.TextFormatter{
		TimestampFormat:        "02-01-2006 15:04:05", // the "time" field configuratiom
		FullTimestamp:          true,
		DisableLevelTruncation: true, // log level field configuration
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			return "", fmt.Sprintf(" %s:%d ", formatFilePath(f.File), f.Line)
		},
	}
	log.SetFormatter(formatter)
	return log
}

func formatFilePath(path string) string {
	arr := strings.Split(path, "/")
	return arr[len(arr)-1]
}
