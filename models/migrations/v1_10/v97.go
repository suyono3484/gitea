// Copyright 2019 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_10 //nolint

import "xorm.io/xorm"

func AddRepoAdminChangeTeamAccessColumnForUser(x *xorm.Engine) error {
	type User struct {
		RepoAdminChangeTeamAccess bool `xorm:"NOT NULL DEFAULT false"`
	}

	return x.Sync2(new(User))
}
