package commands

import (
	"fmt"
	"github.com/andygrunwald/watson/client"
	"github.com/andygrunwald/watson/storage"
	"github.com/andygrunwald/watson/storage/identity"
	"github.com/codegangsta/cli"
	"log"
	"os"
	"sync"
)

// Crawl crawls a Gerrit instance
func Crawl(c *cli.Context) {
	watson, err := client.NewGerritClient(c.GlobalString("instance"), 60)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	watson.Authentication(c.String("auth-mode"), c.String("username"), c.String("password"))

	var wg sync.WaitGroup
	storageChan := storage.GetStorage(c.String("storage"), &wg)
	identityChan := identity.GetStorage(c.String("identity-storage"), &wg)

	crawl := client.NewCrawler(watson)
	crawl.ChangeSetQueryLimit = watson.GetQueryLimit()
	crawl.Storage = storageChan
	crawl.IdentityStorage = identityChan

	log.Println("Start crawling ...")

	projects, err := crawl.Projects()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	i := 0
	for name := range *projects {
		i++
		if i > 100 {
			continue
		}
		wg.Add(1)
		log.Printf("Crawling project %s ...", name)

		go func(crawl *client.Crawler, name string) {
			defer wg.Done()
			crawl.Changesets(name)

			// * proceedChangeSetsDependsOnRelation
			// * proceedChangeSetsNeededByRelation
		}(crawl, name)
	}

	wg.Wait()
}
