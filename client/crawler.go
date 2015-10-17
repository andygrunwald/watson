package client

import (
	"fmt"
	"github.com/andygrunwald/go-gerrit"
	"github.com/andygrunwald/watson/storage"
	"github.com/andygrunwald/watson/storage/identity"
	"log"
)

type Crawler struct {
	Client              *Client
	ChangeSetQueryLimit int
	Storage             chan *storage.ChangeSet
	IdentityStorage     chan *identity.Identity
}

func NewCrawler(c *Client) *Crawler {
	crawler := &Crawler{
		Client: c,
	}

	return crawler
}

func (c *Crawler) Projects() (*map[string]storage.Project, error) {
	/*
		opt := &gerrit.ProjectOptions{
			Description: true,
			Tree:        true,
			Type:        "ALL",
		}

		projects, _, err := c.Client.Gerrit.Projects.ListProjects(opt)
	*/
	projects, _, err := c.Client.Gerrit.Projects.ListProjects(nil)

	res := make(map[string]storage.Project, len(*projects))

	for name, p := range *projects {
		// TODO Store name, p
		// type Project *gerrit.ProjectInfo
		res[name] = storage.Project(p)
	}

	return &res, err
}

func (c *Crawler) Changesets(project string) {
	for startNum := 0; ; {
		endNum := startNum + c.ChangeSetQueryLimit
		log.Printf("Querying for changes %d...%d for project %s", startNum, endNum, project)

		opt := &gerrit.QueryChangeOptions{
			Start: startNum,
		}
		opt.Query = []string{fmt.Sprintf("project:%s", project)}
		opt.Limit = c.ChangeSetQueryLimit
		opt.AdditionalFields = []string{
			"DETAILED_ACCOUNTS",
			/*
				"LABELS",
				"WEB_LINKS",
				"ALL_FILES",
				"MESSAGES",
				"CHANGE_ACTIONS",
				"REVIEWED",
				"WEB_LINKS",
				"COMMIT_FOOTERS",
				"ALL_REVISIONS",
				"DOWNLOAD_COMMANDS",
				"CURRENT_COMMIT",
				"ALL_COMMITS",
				"CURRENT_FILES",
				"CURRENT_REVISION",
				"DETAILED_LABELS",
			*/
		}

		changes, resp, err := c.Client.Gerrit.Changes.QueryChanges(opt)
		if err != nil {
			log.Printf("ERROR ... %+v", err)
			continue
		}

		if changes == nil {
			log.Printf("changes is nil ... %+v", resp)
			continue
		}

		numOfChangesets := len(*changes)
		log.Printf(">>>> Received %d changes to process for project %s", numOfChangesets, project)
		startNum += numOfChangesets

		for _, change := range *changes {
			cs := &storage.ChangeSet{
				Change: &change,
			}
			c.Storage <- cs
			c.IdentityStorage <- identity.AccountInfo(change.Owner).Identify()
			// TODO Add all identities to IdentityStorage
			// GitPersonInfo, AccountInfo, EmailInfo

			// Import this changeset :)
			//$this->proceedChangeset($changeSet, $project);
		}

		// Last changeset have a key: _more_changes (set to true)
		// TODO
		if numOfChangesets == 0 || (numOfChangesets > 0 && numOfChangesets < c.ChangeSetQueryLimit) {
			break
		}
	}
}
