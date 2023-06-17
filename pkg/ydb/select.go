package ydb

import (
	"context"
	"fmt"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
)

func Select(query string) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

    //столбцы таблицы
    var user struct{
        idT1    uint64
        Col1T1  string
        Col2T1  time
    }

    db.Table().Do(ctx, func(ctx context.Context, s table.Session) (err error) {
    _, res, _ := s.Execute(ctx, table.DefaultTxControl(), query, nil)

    defer res.Close()

    if err = res.NextResultSetErr(ctx); err != nil {
        return err
    }

    for res.NextRow() {
        res.ScanNamed(
            named.OptionalWithDefault("id_user", &user.id_user),
            named.OptionalWithDefault("name_user", &user.name_user),
            named.OptionalWithDefault("tst", &user.tst),
        )
        fmt.Printf("id_user=\"%d\", name_user=\"%s\", tst=\"%s\"\n", user.id_user, user.name_user, user.tst)
    }
    return res.Err() // for driver retry if not nil
    })
}