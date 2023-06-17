package main

import (
	// "github.com/Starc151/dia/pkg/fyne"
	"github.com/Starc151/dia/pkg/ydb"
)

func main() {
	// fyne.Show()
	ydb.Select(`SELECT * FROM test2;`)
}