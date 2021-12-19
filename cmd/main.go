package main

import (
	"context"
	"flag"
	"os"

	"github.com/comfortablynull/aoc2021go/internal/day"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	subcommands.Register(newCmd[int, int]("1", day.One()), "days")
	subcommands.Register(newCmd[int, int]("2", day.Two()), "days")
	subcommands.Register(newCmd[int, int]("3", day.Three()), "days")
	subcommands.Register(newCmd[int, int]("4", day.Four()), "days")
	subcommands.Register(newCmd[int, int]("5", day.Five()), "days")
	subcommands.Register(newCmd[int, int]("6", day.Six()), "days")
	subcommands.Register(newCmd[int, int]("7", day.Seven()), "days")
	subcommands.Register(newCmd[int, int]("8", day.Eight()), "days")
	subcommands.Register(newCmd[int, int]("9", day.Nine()), "days")
	subcommands.Register(newCmd[int, int]("10", day.Ten()), "days")
	subcommands.Register(newCmd[int, int]("11", day.Eleven()), "days")
	subcommands.Register(newCmd[int, int]("12", day.Twelve()), "days")
	subcommands.Register(newCmd[int, string]("13", day.Thirteen()), "days")
	subcommands.Register(newCmd[int, int]("14", day.Fourteen()), "days")
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
