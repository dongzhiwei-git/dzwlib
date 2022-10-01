package log

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"time"
)

type RotateOption struct {
	LinkName string
	Count    *uint
	MaxAge   *time.Duration
	Size     *int64
}

func NewRotateLogger(option RotateOption) (*rotatelogs.RotateLogs, error) {

	var opts []rotatelogs.Option

	opts = append(opts, rotatelogs.WithLinkName(option.LinkName))
	if option.Count != nil {
		opts = append(opts, rotatelogs.WithRotationCount(*option.Count))
	}
	if option.MaxAge != nil {
		opts = append(opts, rotatelogs.WithMaxAge(*option.MaxAge))
	}
	if option.Size != nil {
		opts = append(opts, rotatelogs.WithRotationSize(*option.Size))
	}
	return rotatelogs.New(option.LinkName+".%Y%m%d%H%M", opts...)
}
