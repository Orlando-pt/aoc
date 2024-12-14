package utils

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

func StrToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func IntAbs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func SetTestLogger() {
	var programLevel = new(slog.LevelVar)
	h := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel})

	slog.SetDefault(slog.New(h))
	programLevel.Set(slog.LevelDebug)
}

func Debug(something ...any) {
	slog.Debug(fmt.Sprint(something...))
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ArrStrContains(arr []string, value string) bool {
	for _, v := range arr {
		if v == value {
			return true
		}
	}
	return false
}
