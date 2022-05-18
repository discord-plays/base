package utils

import "time"

type GameTimer struct {
	time.Duration
	infinite bool
}

func (gt GameTimer) IsInfinite() bool {
	return gt.infinite
}

func ParseGameTimer(value string) (GameTimer, error) {
	if value == "" {
		return GameTimer{0, true}, nil
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return GameTimer{}, err
	}
	return GameTimer{duration, false}, nil
}
