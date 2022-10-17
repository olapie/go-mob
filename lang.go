package mobile

import (
	"os"
	"strings"

	"code.olapie.com/log"
)

type langT int

const (
	english = iota
	simplifiedChinese
	traditionalChinese
)

var lang langT = english

func SetLang(l string) {
	if err := os.Setenv("LANG", l); err != nil {
		log.S().Errorf("Cannot set env:lang=%s, %v", lang, err)
	}
	l = strings.ToLower(l)
	strings.Replace(l, "-", "_", -1)
	switch {
	case strings.Contains(l, "hans"):
		lang = simplifiedChinese
	case strings.Contains(l, "zh_cn"):
		lang = simplifiedChinese
	case strings.Contains(l, "zh_sg"):
		lang = simplifiedChinese
	}
}

func IsSimplifiedChinese() bool {
	return lang == simplifiedChinese
}

func IsTraditionalChinese() bool {
	return lang == traditionalChinese
}
