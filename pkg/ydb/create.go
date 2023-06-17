package ydb

import (
	"context"
	"path"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/options"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func CreateTable(nameDb string) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

	db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
		  return s.CreateTable(ctx, path.Join(db.Name(), nameDb),
			options.WithColumn("series_id", types.TypeUint64),  // not null column
			options.WithColumn("title", types.Optional(types.TypeUTF8)),
			options.WithColumn("series_info", types.Optional(types.TypeUTF8)),
			options.WithColumn("release_date", types.Optional(types.TypeDate)),
			options.WithColumn("comment", types.Optional(types.TypeUTF8)),
			options.WithPrimaryKeyColumn("series_id"),
		  )
		},
	  )
}