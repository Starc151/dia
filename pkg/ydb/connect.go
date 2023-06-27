package ydb

import (
	"context"
	"fmt"

	"github.com/ydb-platform/ydb-go-sdk/v3"
	yc "github.com/ydb-platform/ydb-go-yc"
)

func connect() (*ydb.Driver, context.Context, context.CancelFunc){
	ctx, cancel := context.WithCancel(context.Background())
	
	db, err := ydb.Open(ctx,
		"grpcs://ydb.serverless.yandexcloud.net:2135/ru-central1/b1gvs9nokmiitnhv21jt/etn1rjbm3bjs0a8emhjd",
		yc.WithServiceAccountKeyFileCredentials("token/starc.json"),
	)
	if err != nil {
		fmt.Println("No connect")
	} else {
		fmt.Println("CONNECTED!")
	}
	return db, ctx, cancel
}

