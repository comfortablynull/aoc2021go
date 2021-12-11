package main

import (
	"context"
	"flag"
	"os"

	"github.com/comfortablynull/aoc2021go/internal/day"
	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day10"
	"github.com/comfortablynull/aoc2021go/internal/day/day11"
	"github.com/comfortablynull/aoc2021go/internal/day/day2"
	"github.com/comfortablynull/aoc2021go/internal/day/day3"
	"github.com/comfortablynull/aoc2021go/internal/day/day4"
	"github.com/comfortablynull/aoc2021go/internal/day/day5"
	"github.com/comfortablynull/aoc2021go/internal/day/day6"
	"github.com/comfortablynull/aoc2021go/internal/day/day7"
	"github.com/comfortablynull/aoc2021go/internal/day/day8"
	"github.com/comfortablynull/aoc2021go/internal/day/day9"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newCmd("1", day.Func(day1.Run)), "days")
	subcommands.Register(newCmd("2", day.Func(day2.Run)), "days")
	subcommands.Register(newCmd("3", &day3.Runner{}), "days")
	subcommands.Register(newCmd("4", day.Func(day4.Run)), "days")
	subcommands.Register(newCmd("5", day.Func(day5.Run)), "days")
	subcommands.Register(newCmd("6", day.Func(day6.Run)), "days")
	subcommands.Register(newCmd("7", day.Func(day7.Run)), "days")
	subcommands.Register(newCmd("8", day.Func(day8.Run)), "days")
	subcommands.Register(newCmd("9", day.Func(day9.Run)), "days")
	subcommands.Register(newCmd("10", day.Func(day10.Run)), "days")
	subcommands.Register(newCmd("11", day.Func(day11.Run)), "days")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
