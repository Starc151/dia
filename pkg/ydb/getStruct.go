package ydb

import (
	"context"
	"fmt"
	"path"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
)

func GetTableStruct(tableName string) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

	err := db.Table().Do(ctx,
		func(ctx context.Context, s table.Session) (err error) {
			desc, err := s.DescribeTable(ctx, path.Join(db.Name(), tableName))
			if err != nil {
				return
			}
			fmt.Printf("> describe table: %s\n", tableName)
			for _, c := range desc.Columns {
				fmt.Printf("  > column, name: %s, %s\n", c.Type, c.Name)
			}
			return
		},
	)
	if err != nil {
		fmt.Println("Не удалось выполнить запрос")
	}
}