package domain

import "context"

type (
	// Profiler interface
	Profiler interface {
		CPUProfile(ctx context.Context) ([]byte, error)
	}
)
