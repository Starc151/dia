package ydb

import (
	"context"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3/table"
	"github.com/ydb-platform/ydb-go-sdk/v3/table/types"
)

func Insert() {
  db, ctx, cancel := connect()
	defer cancel()
	defer db.Close(ctx)

  err := db.Table().DoTx( // Do retry operation on errors with best effort
    ctx, // context manages exiting from Do
    func(ctx context.Context, tx table.TransactionActor) (err error) { // retry operation
      res, err := tx.Execute(ctx, `
          DECLARE $Id_user AS Uint64;
          DECLARE $NName AS Utf8;
          DECLARE $FName AS Utf8;
          DECLARE $LName AS Utf8;
          INSERT INTO users ( NName, FName, LName )
          VALUES ( $NName, $FName, $LName );
        `,
        table.NewQueryParameters(
          table.ValueParam("$Id_user", types.Uint64Value(1)),
          table.ValueParam("$NName", types.UTF8Value("1")),
          table.ValueParam("$FName", types.UTF8Value("1")),
          table.ValueParam("$LName", types.UTF8Value("1")), // increment LName
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
