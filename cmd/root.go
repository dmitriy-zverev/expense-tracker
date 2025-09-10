package cmd

import (
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type Command struct {
	Amount      float64
	Limit       float64
	ID          int
	Month       int
	WithDeleted bool
	Description string
	Cmd         string
	Category    string
	Output      string
	BudgetCmd   string
}

func (cmd *Command) Run() error {
	initCommands()

	command, ok := commands[cmd.Cmd]
	if !ok {
		return errors.New(cmd.Cmd + " is not found")
	}

	if err := command.Callback(*cmd); err != nil {
		return err
	}

	return nil
}

func help() {
	fmt.Printf("Usage: et <command> [-argument 1] [description 1] ...\n")
	fmt.Printf("	Example: et add --description \"Lunch\" --amount 20\n")
}

func ParseCommand(args []string) (Command, error) {
	if len(args) < 2 {
		help()
		return Command{}, errors.New("no command found")
	}

	cmd := Command{
		Cmd:         args[1],
		ID:          -1,
		Month:       -1,
		Amount:      -1.0,
		Limit:       -1.0,
		WithDeleted: false,
	}

	if slices.Contains(args, DESCRIPTION_PARAM) {
		idx := slices.Index(args, DESCRIPTION_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --description")
		}
		cmd.Description = args[idx+1]
	}

	if slices.Contains(args, AMOUNT_PARAM) {
		idx := slices.Index(args, AMOUNT_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --amount")
		}

		amount, err := strconv.Atoi(args[idx+1])
		if err != nil {
			return Command{}, errors.New("argument for --amount is not a number")
		}

		cmd.Amount = float64(amount)
	}

	if slices.Contains(args, ID_PARAM) {
		idx := slices.Index(args, ID_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --id")
		}

		id, err := strconv.Atoi(args[idx+1])
		if err != nil {
			return Command{}, errors.New("argument for --id is not a number")
		}

		cmd.ID = id
	}

	if slices.Contains(args, MONTH_PARAM) {
		idx := slices.Index(args, MONTH_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --month")
		}

		month, err := strconv.Atoi(args[idx+1])
		if err != nil {
			return Command{}, errors.New("argument for --month is not a number")
		}

		cmd.Month = month
	}

	if slices.Contains(args, CATEGORY_PARAM) {
		idx := slices.Index(args, CATEGORY_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --category")
		}

		cmd.Category = os.Args[idx+1]
	}

	if slices.Contains(args, WITH_DELETED_PARAM) {
		cmd.WithDeleted = true
	}

	if slices.Contains(args, OUTPUT_PARAM) {
		idx := slices.Index(args, OUTPUT_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --output")
		}

		cmd.Output = os.Args[idx+1]
	}

	if slices.Contains(args, LIMIT_PARAM) {
		idx := slices.Index(args, LIMIT_PARAM)
		if idx+1 > len(args) {
			return Command{}, errors.New("cannot find argument for --limit")
		}

		limit, err := strconv.Atoi(args[idx+1])
		if err != nil {
			return Command{}, errors.New("argument for --limit is not a number")
		}

		cmd.Limit = float64(limit)
	}

	if slices.Contains(args, BUDGET_SET_CMD) {
		cmd.BudgetCmd = BUDGET_SET_CMD
	}

	if slices.Contains(args, BUDGET_LIST_CMD) {
		cmd.BudgetCmd = BUDGET_LIST_CMD
	}

	if slices.Contains(args, BUDGET_REMOVE_CMD) {
		cmd.BudgetCmd = BUDGET_REMOVE_CMD
	}

	return cmd, nil
}
