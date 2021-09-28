package dola_test

import (
	"context"

	"github.com/ydm/dola"
)

func ExampleKeep() {
	keep, _ := dola.NewKeepBuilder().Build()
	keep.Root.Add("verbose", dola.VerboseStrategy{}) //nolint:exhaustivestruct
	keep.Run(context.Background())
}
