// Copyright 2022 The Gitea Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package v1_16 //nolint

import (
	"code.gitea.io/gitea/models/migrations/base"

	"xorm.io/xorm"
	"xorm.io/xorm/schemas"
)

func MigrateUserPasswordSalt(x *xorm.Engine) error {
	dbType := x.Dialect().URI().DBType
	// For SQLITE, the max length doesn't matter.
	if dbType == schemas.SQLITE {
		return nil
	}

	if err := base.ModifyColumn(x, "user", &schemas.Column{
		Name: "rands",
		SQLType: schemas.SQLType{
			Name: "VARCHAR",
		},
		Length: 32,
		// MySQL will like us again.
		Nullable:       true,
		DefaultIsEmpty: true,
	}); err != nil {
		return err
	}

	return base.ModifyColumn(x, "user", &schemas.Column{
		Name: "salt",
		SQLType: schemas.SQLType{
			Name: "VARCHAR",
		},
		Length:         32,
		Nullable:       true,
		DefaultIsEmpty: true,
	})
}
