package ydb

import (
	"context"
	"path"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func CreateTable(userMap map[string]string) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

	db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
		  	return s.CreateTable(ctx, path.Join(db.Name(), userMap["nameDb"]),
				options.WithColumn("user_id", types.TypeUint64),  // not null column
				options.WithColumn("fName", types.Optional(types.TypeUTF8)),
				options.WithColumn("lName", types.Optional(types.TypeUTF8)),
				options.WithColumn("nName", types.Optional(types.TypeUTF8)),
				options.WithColumn("email", types.Optional(types.TypeUTF8)),
				options.WithColumn("password", types.Optional(types.TypeUTF8)),
				options.WithPrimaryKeyColumn("user_id"),
			)
		},
	)
	db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
		  	return s.CreateTable(ctx, path.Join(db.Name(), nameDb),
				options.WithColumn("user_id", types.TypeUint64),  // not null column
				options.WithColumn("fName", types.Optional(types.TypeUTF8)),
				options.WithColumn("lName", types.Optional(types.TypeUTF8)),
				options.WithColumn("nName", types.Optional(types.TypeUTF8)),
				options.WithColumn("email", types.Optional(types.TypeUTF8)),
				options.WithColumn("password", types.Optional(types.TypeUTF8)),
				options.WithPrimaryKeyColumn("user_id"),
			)
		},
	)
}