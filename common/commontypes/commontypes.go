package commontypes

import "context"

type IDFinder func(context.Context, string) (int32, error)
type IDFinderWithIdAndName func(context.Context, string, int32) (int32, error)
type IDFinderUnknown func(context.Context, int32) (int32, error)
