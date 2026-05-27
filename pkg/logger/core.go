package logger

import (
	"go.uber.org/zap/zapcore"
	"regexp"
)

type customCore struct {
	zapcore.LevelEnabler
	enc    zapcore.Encoder
	out    zapcore.WriteSyncer
	filter []*regexp.Regexp
}

func newCore(enc zapcore.Encoder, ws zapcore.WriteSyncer, enab zapcore.LevelEnabler, filter []*regexp.Regexp) zapcore.Core {
	return &customCore{
		LevelEnabler: enab,
		enc:          enc,
		out:          ws,
		filter:       filter,
	}
}

func (c *customCore) With(fields []zapcore.Field) zapcore.Core {
	clone := c.clone()
	addFields(clone.enc, fields)
	return clone
}

func (c *customCore) Check(ent zapcore.Entry, ce *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	// if !c.Enabled(ent.Level) {
	// 	return ce
	// }
	// if c.filter == nil || len(c.filter) == 0 {
	// 	return ce
	// }
	// _, filePath, _, ok := runtime.Caller(c.option.CallerSkip)
	// flag := false

	// for _, exp := range c.filter {
	// 	if ok && exp.MatchString(filePath) {
	// 		flag = true
	// 		break
	// 	}
	// }

	// if !flag {
	// 	return ce.AddCore(ent, c)
	// }
	// return ce
	if c.Enabled(ent.Level) {
		return ce.AddCore(ent, c)
	}
	return ce
}

func (c *customCore) Write(ent zapcore.Entry, fields []zapcore.Field) error {
	isFilter := false
	if ent.Caller.File != "" {
		for _, exp := range c.filter {
			if exp.MatchString(ent.Caller.File) {
				isFilter = true
				break
			}
		}
	}
	if isFilter {
		return nil
	}
	buf, err := c.enc.EncodeEntry(ent, fields)
	if err != nil {
		return err
	}
	_, err = c.out.Write(buf.Bytes())
	buf.Free()
	if err != nil {
		return err
	}
	if ent.Level > zapcore.ErrorLevel {
		_ = c.Sync()
	}
	return nil
}

func (c *customCore) Sync() error {
	return c.out.Sync()
}

func (c *customCore) clone() *customCore {
	return &customCore{
		LevelEnabler: c.LevelEnabler,
		enc:          c.enc.Clone(),
		out:          c.out,
		filter:       c.filter,
	}
}

func addFields(enc zapcore.ObjectEncoder, fields []zapcore.Field) {
	for i := range fields {
		fields[i].AddTo(enc)
	}
}
