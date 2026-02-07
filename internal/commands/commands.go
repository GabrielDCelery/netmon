package commands

import "context"

type Command interface {
	Run(ctx context.Context) (string, error)
	PrintCommandAsStr() string
}
