package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var DefaultRotateOption = RotateOption{
	MaxAge: func(d time.Duration) *time.Duration { return &d }(time.Hour * 24 * 7),
	Size:   func(i int64) *int64 { return &i }(1 << 30),
}

// 配置日志输出：
// 1.　Error|Fatal|Panic日志输出到dir/app.err*.log文件和Stderr
// 2. 其余级别日志输出到app*.log和Stdout
// 3. 文件日志启用轮转
func Setup(dir string, app string, level logrus.Level, opt RotateOption) {
	logrus.SetLevel(level)

	writers := []io.Writer{os.Stderr}
	levelSplitWriter := NewLevelSplitWriter()
	levelSplitWriter.AddWriter(io.MultiWriter(writers...), logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel)

	writers = []io.Writer{os.Stdout}
	if dir != "" {
		option := opt
		option.LinkName = filepath.Join(dir, fmt.Sprintf("%s.log", app))
		if out, err := NewRotateLogger(option); err != nil {
			fmt.Fprintf(os.Stderr, "NewRotateLogger %s failed:%v\n", option.LinkName, err)
		} else {
			writers = append(writers, out)
		}
	}
	levelSplitWriter.AddWriter(io.MultiWriter(writers...), logrus.WarnLevel, logrus.InfoLevel, logrus.DebugLevel)

	logrus.StandardLogger().SetReportCaller(true)
	logrus.StandardLogger().AddHook(levelSplitWriter)
	//formatter := NewSeraphFormatter(Config{AppName: app}, &logrus.JSONFormatter{DisableHTMLEscape: true})
	//logrus.StandardLogger().SetFormatter(formatter)
}
