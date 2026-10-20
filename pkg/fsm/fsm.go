package fsm

import "context"

type State[T any] func(ctx context.Context, args T) (T, State[T], error)

func Run[T any](ctx context.Context, data T, q0 State[T]) bool {

	currentState := q0

	var err error

	for currentState != nil {

		if err := ctx.Err(); err != nil {
			return false
		}

		data, currentState, err = currentState(ctx, data)
		if err != nil {
			return false
		}

		if currentState == nil {
			return true
		}
	}

	return false
}
