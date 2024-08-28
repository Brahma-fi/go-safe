package utils

import "errors"

var (
	ErrInvalidRange = errors.New("invalid range")
)

func Slice[T any](slice []T, from int, to int) ([]T, error) {
	if from > len(slice) || to > len(slice) || from < 0 || to < 0 || from > to {
		return nil, ErrInvalidRange
	}
	return slice[from:to], nil
}

func ReadNAndShift[T any](data []T, baseOffset int, n int) ([]T, int, error) {
	if subData, err := Slice(data, baseOffset, baseOffset+n); err != nil {
		return nil, baseOffset, err
	} else {
		return subData, baseOffset + n, nil
	}
}
