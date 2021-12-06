package main

import (
	"context"
	"flag"
	"os"

	"github.com/comfortablynull/aoc2021go/internal/day"
	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day2"
	"github.com/comfortablynull/aoc2021go/internal/day/day3"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newCmd("one", day.Func(day1.Run)), "days")
	subcommands.Register(newCmd("two", day.Func(day2.Run)), "days")
	subcommands.Register(newCmd("three", &day3.Runner{}), "days")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
