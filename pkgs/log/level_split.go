package log

import (
	"github.com/sirupsen/logrus"
	"io"
	"sync"
)

var (
	levelMap = map[logrus.Level]string{
		logrus.PanicLevel: "PANIC",
		logrus.FatalLevel: "FATAL",
		logrus.ErrorLevel: "ERROR",
		logrus.WarnLevel:  "WARN",
		logrus.InfoLevel:  "INFO",
		logrus.DebugLevel: "DEBUG",
	}
)

type LevelSplitWriter struct {
	writers map[logrus.Level]io.Writer
	levels  []logrus.Level
	locker  sync.RWMutex
}

func NewLevelSplitWriter() *LevelSplitWriter  {
	return &LevelSplitWriter{
		writers: map[logrus.Level]io.Writer{},
	}
}

func (h *LevelSplitWriter) Levels() []logrus.Level {
	return h.levels
}

func (h *LevelSplitWriter) AddWriter(w io.Writer, levels ...logrus.Level) {
	h.locker.Lock()
	defer h.locker.Unlock()

	for _, l := range levels {
		if _, ok := h.writers[l]; ok {
			continue
		}
		h.writers[l] = w
		h.levels = append(h.levels, l)
	}
}

func (h *LevelSplitWriter) Fire(entry *logrus.Entry) error {

	h.locker.RLock()
	defer h.locker.RUnlock()

	if l, ok := h.writers[entry.Level]; ok {
		entry.Logger.Out = l
	}
	return nil
}
