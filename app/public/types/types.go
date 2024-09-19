package types

import (
	"context"
)

type IDFinder func(context.Context, string) (int32, error)
