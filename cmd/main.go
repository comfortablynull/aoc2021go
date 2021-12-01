package main

import (
	"context"
	"flag"
	"os"

	"github.com/comfortablynull/aoc2021go/internal/day"
	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newCmd("one", day.Func(day1.Run)), "days")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
