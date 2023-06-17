package ydb

import (
	"context"
	"fmt"
	// "time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
)

func Select(query string) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

    //столбцы таблицы
    var user struct{
        idT2    uint64
        Col1T2  string
        Col2T2  string
    }

    db.Table().Do(ctx, func(ctx context.Context, s table.Session) (err error) {
    _, res, _ := s.Execute(ctx, table.DefaultTxControl(), query, nil)

    defer res.Close()

    if err = res.NextResultSetErr(ctx); err != nil {
        return err
    }

    for res.NextRow() {
        res.ScanNamed(
            named.OptionalWithDefault("idT2", &user.idT2),
            named.OptionalWithDefault("Col1T2", &user.Col1T2),
            named.OptionalWithDefault("Col2T2", &user.Col2T2),
        )
        fmt.Printf("idT2=\"%d\", Col1T2=\"%s\", Col2T2=\"%s\"\n", user.idT2, user.Col1T2, user.Col2T2)
    }
    return res.Err() // for driver retry if not nil
    })
}