package main

import (
	"context"
	"flag"
	"os"

	"github.com/comfortablynull/aoc2021go/internal/day"
	"github.com/comfortablynull/aoc2021go/internal/day/day1"
	"github.com/comfortablynull/aoc2021go/internal/day/day10"
	"github.com/comfortablynull/aoc2021go/internal/day/day11"
	"github.com/comfortablynull/aoc2021go/internal/day/day12"
	"github.com/comfortablynull/aoc2021go/internal/day/day13"
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
	subcommands.Register(newCmd[int, int]("1", day.Func[int, int](day1.Run)), "days")
	subcommands.Register(newCmd[int, int]("2", day.Func[int, int](day2.Run)), "days")
	subcommands.Register(newCmd[int, int]("3", &day3.Runner{}), "days")
	subcommands.Register(newCmd[int, int]("4", day.Func[int, int](day4.Run)), "days")
	subcommands.Register(newCmd[int, int]("5", day.Func[int, int](day5.Run)), "days")
	subcommands.Register(newCmd[int, int]("6", day.Func[int, int](day6.Run)), "days")
	subcommands.Register(newCmd[int, int]("7", day.Func[int, int](day7.Run)), "days")
	subcommands.Register(newCmd[int, int]("8", day.Func[int, int](day8.Run)), "days")
	subcommands.Register(newCmd[int, int]("9", day.Func[int, int](day9.Run)), "days")
	subcommands.Register(newCmd[int, int]("10", day.Func[int, int](day10.Run)), "days")
	subcommands.Register(newCmd[int, int]("11", day.Func[int, int](day11.Run)), "days")
	subcommands.Register(newCmd[int, int]("12", day.Func[int, int](day12.Run)), "days")
	subcommands.Register(newCmd[int, string]("13", day.Func[int, string](day13.Run)), "days")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
