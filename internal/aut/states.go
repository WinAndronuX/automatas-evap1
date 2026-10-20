package aut

import (
	"context"
	"errors"
	"slices"
)

var initialsStates []State
var idStates []State


func getPrefixFunction(input []rune) []int {

	var pi = make([]int, len(input))

	for k := 1; k < len(input); k++ {

		j := pi[k-1]
		for j > 0 && input[k] != input[j] {
			j = pi[j-1]
		}

		if input[k] == input[j] {
			j++
		}

		pi[k] = j
	}

	return pi
}

func getContinuousCharStates(input []rune, prev State, end State) []State {

	var result []State

	var pi = getPrefixFunction(input)

	for i := len(input) - 1; i >= 0; i-- {

		next := end
		if i < len(input)-1 {
			next = result[len(input)-i-2]
		}

		var st = func(ctx context.Context, data AutContext) (AutContext, State, error) {
			if data.Index >= len(data.Input) {
				return data, nil, ErrIncomplete
			}

			ch := data.Input[data.Index]

			switch ch {
			case input[i]:
				data.Index++
				return data, next, nil
			default:
				if i > 0 && pi[i-1] > 0 {
					return data, result[len(input)-1-pi[i-1]], nil
				}

				return data, prev, nil
			}
		}

		result = append(result, st)
	}

	return result
}

/*
 * Estados
 */

var ErrInvalidInput = errors.New("Invalid input")

var ErrIncomplete = errors.New("Incomplete input")

func StateQ0(ctx context.Context, data AutContext) (AutContext, State, error) {

	if data.Index >= len(data.Input) {
		return data, nil, ErrIncomplete
	}

	ch := data.Input[data.Index]
	data.Index++

	if !slices.Contains(data.Id, ch) {
		return data, nil, ErrInvalidInput
	}

	return data, StateQ1, nil
}

func StateQ1(ctx context.Context, data AutContext) (AutContext, State, error) {

	if data.Index >= len(data.Input) {
		return data, nil, ErrIncomplete
	}

	ch := data.Input[data.Index]
	data.Index++

	if !slices.Contains(data.Alphabet, ch) {
		return data, nil, ErrInvalidInput
	}

	switch ch {
	case '.':
		return data, StateQ2, nil
	case data.Initials[0]:
		return data, StateQ3, nil
	default:
		return data, StateQ1, nil
	}
}

func StateQ2(ctx context.Context, data AutContext) (AutContext, State, error) {

	if data.Index >= len(data.Input) {
		return data, nil, ErrIncomplete
	}

	ch := data.Input[data.Index]
	data.Index++

	if !slices.Contains(data.Alphabet, ch) {
		return data, nil, ErrInvalidInput
	}

	switch ch {
	case '.':
		return data, nil, ErrInvalidInput
	case data.Initials[0]:
		return data, StateQ3, nil
	default:
		return data, StateQ1, nil
	}
}

func StateQ3(ctx context.Context, data AutContext) (AutContext, State, error) {

	if data.Index >= len(data.Input) {
		return data, nil, ErrIncomplete
	}

	ch := data.Input[data.Index]

	if !slices.Contains(data.Alphabet, ch) {
		return data, nil, ErrInvalidInput
	}

	return initialsStates[len(data.Initials)-2](ctx, data)
}

func StateQ4(ctx context.Context, data AutContext) (AutContext, State, error) {
	if data.Index >= len(data.Input) {
		return data, nil, ErrIncomplete
	}

	ch := data.Input[data.Index]
	data.Index++

	if !slices.Contains(data.Alphabet, ch) {
		return data, nil, ErrInvalidInput
	}

	switch ch {
	case '.':
		return data, StateQ5, nil
	default:
		return data, StateQ4, nil
	}
}

func StateQ5(ctx context.Context, data AutContext) (AutContext, State, error) {
	if data.Index >= len(data.Input) {
		return data, nil, ErrIncomplete
	}

	ch := data.Input[data.Index]
	data.Index++

	if !slices.Contains(data.Alphabet, ch) {
		return data, nil, ErrInvalidInput
	}

	switch ch {
	case '.':
		return data, nil, ErrInvalidInput
	case data.Id[0]:
		data.Index--
		return idStates[len(data.Id)-1](ctx, data)
	default:
		return data, StateQ4, nil
	}
}

func StateQ6(ctx context.Context, data AutContext) (AutContext, State, error) {
	if data.Index >= len(data.Input) {
		return data, nil, nil
	}

	ch := data.Input[data.Index]
	data.Index++

	if !slices.Contains(data.Alphabet, ch) {
		return data, nil, ErrInvalidInput
	}

	switch ch {
	case '.':
		return data, StateQ5, nil
	default:
		return data, StateQ4, nil
	}
}
