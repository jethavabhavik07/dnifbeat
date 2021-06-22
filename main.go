package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/jethavabhavik07/dnifbeat/beater"
)

func main() {
	err := beat.Run("dnifbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
