package commands

import "context"

type Command[T any] interface {
	Run(ctx context.Context) (T, error)
	PrintCommandAsStr() string
}
