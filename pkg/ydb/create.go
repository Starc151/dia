package ydb

import (
	"context"
	"path"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func CreateTable(nName string) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

	db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
		  	return s.CreateTable(ctx, path.Join(db.Name(), nName),
				options.WithColumn("res_id", types.TypeUint64),  // not null column
				options.WithColumn("glucose", types.Optional(types.TypeFloat)),
				options.WithColumn("lName", types.Optional(types.TypeUTF8)),
				options.WithColumn("nName", types.Optional(types.TypeUTF8)),
				options.WithColumn("email", types.Optional(types.TypeUTF8)),
				options.WithColumn("password", types.Optional(types.TypeUTF8)),
				options.WithPrimaryKeyColumn("user_id"),
			)
		},
	)
}