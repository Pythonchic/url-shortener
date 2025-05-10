package main

import (
	"fmt"
	"url-shortener/internal/config"
)

// "log/slog"
func main () {
	config := config.MustLoad()
	fmt.Println(config )
}
