package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/comfortablynull/aoc2021go/internal/day"
	"github.com/google/subcommands"
)

type cmd struct {
	name string
	day  day.Day
}

func newCmd(name string, day day.Day) cmd {
	return cmd{name: name, day: day}
}

func (c cmd) Name() string {
	return c.name
}

func (c cmd) Synopsis() string {
	return fmt.Sprintf("runs day %s", c.Name())
}

func (c cmd) Usage() string {
	return fmt.Sprintf("%s <file>", c.Name())
}

func (c cmd) SetFlags(set *flag.FlagSet) {}

func (c cmd) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if len(f.Args()) != 1 {
		return subcommands.ExitUsageError
	}
	t := time.Now()
	defer func(t time.Time) {
		fmt.Println("Duration:", time.Since(t))
	}(t)
	input, err := os.Open(f.Arg(0))
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	defer input.Close()
	first, second, err := c.day.Run(input)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}
	fmt.Printf("First: %v, Second: %v\n", first, second)
	return subcommands.ExitSuccess
}
