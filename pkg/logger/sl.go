package sl

import "golang.org/x/exp/slog"

func Error(err error) slog.Attr {
	return slog.String("error", err.Error())
}
