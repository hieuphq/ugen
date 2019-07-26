package main

import (
	"runtime"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli"

	"github.com/hieuphq/ugen/src/clipboard"
)

var app = cli.NewApp()

func info() {
	app.Name = "Generate UUID CLI"
	app.Usage = "Helper to generate a UUID and copy to clipboard"
	app.Author = "Hieuphq"
	app.Version = "0.0.1"
}

func main() {
	// info()

	// err := app.Run(os.Args)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	os := runtime.GOOS
	uid := uuid.NewV4()
	c := clipboard.NewClipboard(os)
	c.Write(uid.String())
	fmt.Println(uid.String()+"\ncopied to clipboard")
}
