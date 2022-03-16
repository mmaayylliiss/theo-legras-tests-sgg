// Package tags provides and implements the tags plugin.
package tags

import (
	"path/filepath"
	"strings"
	"sync"

	"github.com/verless/verless/model"
	"github.com/verless/verless/tree"
)

const (
	// tagsDir is the target directory for all tag directories.
	tagsDir string = "/tags"
)

// New creates a new tags plugin that uses templates from the given
// build path and outputs the tag directories to outputDir.
func New() *tags {
	t := tags{
		tags: make(map[string]*model.ListPage),
	}

	return &t
}

// tags is the actual tags plugin that maintains a map with all
// tags from all processed pages.
type tags struct {
	tags      map[string]*model.ListPage
	tagsMutex sync.Mutex
}

// ProcessPage creates a new map entry for each tag in the processed
// page and adds the page to the entry's list page.
func (t *tags) ProcessPage(page *model.Page) error {
	for _, tag := range page.Tags {
		// Sanitizing the tags like "Making Coffee" to "making-coffee".
		tag.Name = strings.Replace(tag.Name, " ", "-", -1)
		tag.Name = strings.ToLower(tag.Name)

		t.tagsMutex.Lock()
		_, tagExists := t.tags[tag.Name]

		if !tagExists {
			t.createListPage(tag.Name)
		}

		t.tags[tag.Name].Pages = append(t.tags[tag.Name].Pages, page)
		t.tagsMutex.Unlock()
	}

	return nil
}

// PreWrite registers each list page in the site model. Those list
// pages will be rendered by the writer.
func (t *tags) PreWrite(site *model.Site) error {
	node := model.NewNode()
	node.ListPage.Route = tagsDir

	if err := tree.CreateNode(tagsDir, site.Root, node); err != nil {
		return err
	}

	for tag, listPage := range t.tags {
		path := filepath.ToSlash(filepath.Join(tagsDir, tag))

		node := model.NewNode()
		node.ListPage = *listPage

		if err := tree.CreateNode(path, site.Root, node); err != nil {
			return err
		}
	}

	return nil
}

// PostWrite isn't needed by the tags plugin.
func (t *tags) PostWrite() error {
	return nil
}

// createListPage initializes a new list page for a given key.
func (t *tags) createListPage(key string) {
	t.tags[key] = &model.ListPage{
		Pages: make([]*model.Page, 0),
		Page: model.Page{
			Route: tagsDir + "/" + key,
		},
	}
}
