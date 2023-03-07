package cli

import (
	"strconv"
	"strings"
	"task/dal/connection"
	"task/dal/tasks"
	"task/types"
	"time"

	"github.com/boltdb/bolt"
	"github.com/urfave/cli/v2"
)

var App = &cli.App{
	Name: "tasks",
	Commands: []*cli.Command{
		{
			Name:  "add",
			Usage: "Add a task",
			Action: func(ctx *cli.Context) error {
				args := ctx.Args()
				if err := verifyArgs(args); err != nil {
					return err
				}
				taskText := strings.Join(args.Slice(), " ")
				task := types.NewTask(taskText)
				return connection.Context(func(db *bolt.DB) error {
					return tasks.Save(db, task)
				})
			},
		}, {
			Name:  "list",
			Usage: "List all active tasks",
			Action: func(ctx *cli.Context) error {
				if err := verifyNoArgs(ctx.Args()); err != nil {
					return err
				}
				return connection.Context(func(db *bolt.DB) error {
					allTasks, err := tasks.GetActive(db)
					if err != nil {
						return err
					}
					printTasks(allTasks)
					return nil
				})
			},
		}, {
			Name:  "done",
			Usage: "Mark a task as done",
			Action: func(ctx *cli.Context) error {
				args := ctx.Args()
				if err := verifyArgs(args); err != nil {
					return err
				}
				taskId, err := strconv.Atoi(args.First())
				if err != nil {
					return err
				}
				return connection.Context(func(db *bolt.DB) error {
					task, err := tasks.GetByID(db, types.ID(taskId))
					if err != nil {
						return err
					}
					task.Complete()
					return tasks.Save(db, task)
				})
			},
		}, {
			Name:  "rm",
			Usage: "Removes a task",
			Action: func(ctx *cli.Context) error {
				args := ctx.Args()
				if err := verifyArgs(args); err != nil {
					return err
				}
				taskId, err := strconv.Atoi(args.First())
				if err != nil {
					return err
				}
				return connection.Context(func(db *bolt.DB) error {
					return tasks.Delete(db, types.ID(taskId))
				})
			},
		}, {
			Name:  "completed",
			Usage: "List completed tasks",
			Action: func(ctx *cli.Context) error {
				args := ctx.Args()
				if err := verifyNoArgs(args); err != nil {
					return err
				}
				return connection.Context(func(db *bolt.DB) error {
					allTasks, err := tasks.GetCompleted(db, time.Hour*24)
					if err != nil {
						return err
					}
					printTasks(allTasks)
					return nil
				})
			},
		},
	},
}
