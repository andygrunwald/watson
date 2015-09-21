package main

import (
	"fmt"
	"github.com/andygrunwald/watson/commands"
	"github.com/codegangsta/cli"
	"os"
)

const (
	// Name is the name of Watson :)
	Name = "Watson"
	// Version is the current version of Watson
	Version = "0.0.1"
	// Usage is a small and catchy sentence
	Usage = "Crawl your Gerrit!"
	// Author is the name of the app author
	Author = "Andy Grunwald"
	// Email is the email of the app author
	Email = "andygrunwald@gmail.com"
)

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Author = Author
	app.Email = Email
	app.Usage = Usage
	app.Action = func(c *cli.Context) {
		fmt.Printf("Hi, i am %s. Nice to meet you.\n", Name)
		fmt.Println("Use the -help parameter to get more information how to use me!")
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "instance, i",
			Value:  "",
			Usage:  "URL for the Gerrit instance",
			EnvVar: "WATSON_INSTANCE",
		},
	}

	app.Commands = commands.Commands()

	app.Run(os.Args)
}
