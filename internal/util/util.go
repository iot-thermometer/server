package util

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetTokenFromContext(c echo.Context) string {
	return strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
}

func ParseParamID(param string) uint {
	id, err := strconv.Atoi(param)
	if err != nil {
		return 0
	}
	return uint(id)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Uint64(n uint64) *uint64 {
	return &n
}

func Int64(n int64) *int64 {
	return &n
}

func Int(n int) *int {
	return &n
}

func Float32(n float32) *float32 {
	return &n
}
