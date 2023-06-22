package ydb

import (
	"context"
	"fmt"
	"time"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func Insert(res map[string]float32) {
	db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)
	t := uint32(time.Now().Unix())+10800 //поправка на часовой пояс Мск
	err := db.Table().DoTx(ctx,
		func(ctx context.Context, tx table.TransactionActor) (err error) {
			res, err := tx.Execute(ctx, `
			DECLARE $date AS Datetime;
			DECLARE $bolus AS Float;
			DECLARE $glucose AS Float;
			DECLARE $xe AS Float;
			INSERT INTO res ( date, bolus, glucose, xe )
			VALUES ( $date, $bolus, $glucose, $xe );
		`,
				table.NewQueryParameters(
					table.ValueParam("$date", types.DatetimeValue(t)),
					table.ValueParam("$bolus", types.FloatValue(res["bolus"])),
					table.ValueParam("$glucose", types.FloatValue(res["glucose"])),
					table.ValueParam("$xe", types.FloatValue(res["xe"])),
				),
			)
			if err != nil {
				return err
			}
			if err = res.Err(); err != nil {
				return err
			}
			return res.Close()
		}, table.WithIdempotent(),
	)
	if err != nil {
		fmt.Printf("unexpected error: %v", err)
	}
}
