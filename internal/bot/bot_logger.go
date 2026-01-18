package bot

import (
	"github.com/go-logr/logr"
)

type BotLogger struct {
	Logger     logr.LogSink
	BotManager *BotManager
}

func (b *BotLogger) Init(info logr.RuntimeInfo) {
	b.Logger.Init(info)
}

func (b *BotLogger) Enabled(level int) bool {
	return b.Logger.Enabled(level)
}

func (b *BotLogger) Info(level int, msg string, keysAndValues ...interface{}) {
	b.BotManager.LogInfo(msg, keysAndValues...)

	b.Logger.Info(level, msg, keysAndValues...)
}

func (b *BotLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	if err != nil {
		b.BotManager.LogError(err, msg, keysAndValues...)

		b.Logger.Error(err, msg, keysAndValues...)
	}
}

func (b *BotLogger) WithValues(keysAndValues ...interface{}) logr.LogSink {
	return &BotLogger{
		Logger:     b.Logger.WithValues(keysAndValues...),
		BotManager: b.BotManager,
	}
}

func (b *BotLogger) WithName(name string) logr.LogSink {
	return &BotLogger{
		Logger:     b.Logger.WithName(name),
		BotManager: b.BotManager,
	}
}
