/*
Author:  Md Kawsar Islam Yeasin <mdkawsarislam2002@gmail.com>
*/

package main

import (
	"embed"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yeasin2002/better-next-app/cmd"
)

//go:embed templates
var templatesFS embed.FS

func main() {
	// handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	err := cmd.Execute(templatesFS)
	if err != nil {
		fmt.Println("Something Went Wrong!!")
		os.Exit(1)
	}

}
