package main

import (
	"context"

	"energy_tk/internal"
)

func main() {
	if err := run(context.Background()); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	return internal.RunApp(ctx)
}
