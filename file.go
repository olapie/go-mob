package mobile

import (
	"os"

	"code.olapie.com/log"
)

func MustMkdir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.S().Fatalf("Make dir: %s, %v", dir, err)
	}
}
