// Package related provides and implements the related plugin.
package related

import (
	"sync"

	"github.com/verless/verless/model"
	"github.com/verless/verless/tree"
)

type related struct {
	pages      map[string]*model.Page
	pagesMutex sync.RWMutex
}

// New initializes and returns a related plugin instance.
func New() *related {
	return &related{
		pages: make(map[string]*model.Page),
	}
}

// ProcessPage adds a given pointer to a Page instance to the plugin's page
// map. This prevents that each page has to be resolved from the tre later.
func (r *related) ProcessPage(page *model.Page) error {
	r.pagesMutex.Lock()
	r.pages[page.Href] = page
	r.pagesMutex.Unlock()
	return nil
}

// PreWrite walks the site's route tree, iterates over all pages of the
// current node and attempts to resolve the provided related page URIs.
//
// If the page URI has been stored in the page map by ProcessPage before,
// the particular pointer will be assigned to the page's Related slice.
//
// ToDo: Log a warning if the page URI cannot be resolved.
func (r *related) PreWrite(site *model.Site) error {
	resolver := func(path string, node tree.Node) error {
		pages := node.(*model.Node).Pages

		for i, _ := range pages {
			for _, related := range pages[i].ProvidedRelated() {
				if p, ok := r.pages[related]; ok {
					pages[i].Related = append(pages[i].Related, p)
				}
			}
		}
		return nil
	}

	return tree.Walk(site.Root, resolver, -1)
}

// PostWrite isn't needed by the related plugin.
func (r *related) PostWrite() error {
	return nil
}
