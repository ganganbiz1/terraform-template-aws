package logger

import (
	"log/slog"
	"os"
)

func Setting() {
	// JSON形式で出力
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func Info(msg string) {
	slog.Info(msg)
}

func InfoWithParams(msg string, params map[string]interface{}) {
	slog.Info(msg, makeParams(params)...)
}

func Warn(msg string) {
	slog.Warn(msg)
}

func WarnWithParams(msg string, params map[string]interface{}) {
	slog.Warn(msg, makeParams(params)...)
}

func Error(msg string, err error) {
	slog.Error(msg, "Error:", err.Error())
}

func ErrorWithParams(msg string, params map[string]interface{}) {
	slog.Error(msg, makeParams(params)...)
}

func makeParams(m map[string]interface{}) []interface{} {
	params := make([]interface{}, 0, len(m))
	for key, v := range m {
		params = append(params, key)
		params = append(params, v)
	}
	return params
}
