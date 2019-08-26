package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"

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

	info()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "amount, a",
			Usage: "Number of UUID result",
			Value: 1,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "gen",
			Aliases: []string{"g"},
			Usage:   "generate uuid and copy to clipboard",
			Action: func(c *cli.Context) error {
				fmt.Println(c.Args())
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		no := c.GlobalInt("a")
		rs := []string{}
		for idx := 0; idx < no; idx++ {
			uid := uuid.NewV4()
			rs = append(rs, uid.String())
		}
		osStr := runtime.GOOS
		cb := clipboard.NewClipboard(osStr)
		cb.Write(strings.Join(rs, "\n"))
		fmt.Println("copied to clipboard")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
