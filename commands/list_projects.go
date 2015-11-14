package commands

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
	"github.com/codegangsta/cli"
	"os"
	"text/template"
	"bytes"
)

func ListProjectsCommandDefinition() cli.Command {
	return cli.Command{
		Name:    "list-projects",
		Aliases: []string{"lp"},
		Usage:   "Lists all projects of a Gerrit instance",
		Action:  ListProjectsCommand,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "template, t",
				Value:  "{{ .IDEscaped }}",
				Usage:  "Template to apply at a project",
				EnvVar: "WATSON_LISTPROJECT_TEMPLATE",
			},
			cli.StringFlag{
				Name:   "filter, f",
				Value:  "",
				Usage:  "Regular expression to filter projects",
				EnvVar: "WATSON_LISTPROJECT_FILTER",
			},
		},
	}
}

// ListProjects will list all projects of a given Gerrit instance
func ListProjectsCommand(c *cli.Context) {
	instance := c.GlobalString("instance")
	client, err := gerrit.NewClient(instance, nil)
	if err != nil {

	}

	type WrappedProjectInfo struct {
		gerrit.ProjectInfo
		IDEscaped string
	}

	// Pre render template
	tmpl, err := template.New("template").Parse(c.String("template"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	opt := &gerrit.ProjectOptions{
		Description: true,
		Regex: c.String("filter"),
	}
	projects, _, err := client.Projects.ListProjects(opt)
	for name, p := range *projects {
		project := WrappedProjectInfo{p, name,}

		// Render template
		b, err := renderTemplate(tmpl, project)
		if err != nil {
			continue
		}

		fmt.Println(string(b))
	}
}

func renderTemplate(tmpl *template.Template, data interface{}) ([]byte, error) {
	var buf []byte
	buffer := bytes.NewBuffer(buf)
	err := tmpl.Execute(buffer, data)
	return buffer.Bytes(), err
}