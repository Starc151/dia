package ydb

import (
	"context"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/result/named"
)
type Table struct{
    Date    time.Time
    Bolus   float32
    Glucose float32
    Xe      float32
} 

func Select() []Table {
    query := "SELECT * FROM res;"
    loc, _ := time.LoadLocation("Europe/London")
    time.Local = loc
    resList := []Table{}
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

    db.Table().Do(ctx, func(ctx context.Context, s table.Session) (err error) {
        _, res, _ := s.Execute(ctx, table.DefaultTxControl(), query, nil)

        defer res.Close()

        if err = res.NextResultSetErr(ctx); err != nil {
            return err
        }
        var Table Table
        for res.NextRow() {
            res.ScanNamed(
                named.OptionalWithDefault("date", &Table.Date),
                named.OptionalWithDefault("bolus", &Table.Bolus	),
                named.OptionalWithDefault("glucose", &Table.Glucose),
                named.OptionalWithDefault("xe", &Table.Xe),
            )
            resList = append(resList, Table)
        }
        return res.Err()
    })
    return resList
}