// Copyright 2020 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package repo

import (
	"github.com/gogs/git-module"
	"gogs.io/gogs/internal/context"
)

const (
	SearchPage = "repo/search"
)

// SearchFile renders file search page
func SearchFile(c *context.Context) {
	var rev = c.Params("revision")

	tree, err := c.Repo.GitRepo.LsTree(rev, git.LsTreeOptions{Recursive: true})

	if err != nil {
		c.NotFound()
		return
	}

	c.Data["PageIsViewFiles"] = true
	c.Data["PageSearchFiles"] = true
	c.Data["FilesTree"] = tree.Files()
	c.Data["Revision"] = rev

	c.Success(SearchPage)
}
