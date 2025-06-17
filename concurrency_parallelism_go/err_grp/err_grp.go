package errgrp

import (
	"context"
	"errors"
	"fmt"

	errgroup "golang.org/x/sync/errgroup"
)

var (
	errFail = errors.New("fail from print with error")
)

func Run() error {
	ctx := context.Background()
	g := new(errgroup.Group)
	g.Go(func() error {
		return printWithError(ctx)
	})
	g.Go(func() error {
		return printWithoutError(ctx)
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to operate: %w", err)
	}

	return nil
}

func RunWithCtx() error {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return printWithError(ctx)
	})
	g.Go(func() error {
		return printWithoutError(ctx)
	})

	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to operate: %w", err)
	}

	return nil
}

func printWithError(ctx context.Context) error {
	for i := 0; i <= 20; i++ {
		fmt.Println("printWithError i:", i)
	}

	return errFail
}

func printWithoutError(ctx context.Context) error {
	for i := 0; i <= 1000; i++ {
		fmt.Println("printWithoutError k:", i)
	}

	return nil
}
