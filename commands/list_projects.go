package commands

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
	"github.com/codegangsta/cli"
	"os"
)

// ListProjects will list all projects of a given Gerrit instance
func ListProjects(c *cli.Context) {
	instance := c.GlobalString("instance")

	client, err := gerrit.NewClient(instance, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	opt := &gerrit.ProjectOptions{
		Description: true,
	}
	projects, _, err := client.Projects.ListProjects(opt)
	for name, p := range *projects {
		if len(p.Description) > 0 {
			fmt.Printf("%s - %s\n", name, p.Description)
		} else {
			fmt.Println(name)
		}
	}
}
