package aut

import (
	"context"
	"evap1/pkg/fsm"
	"evap1/pkg/utils"
	"strings"
)

type AutContext struct {
	Alphabet []rune
	Initials []rune
	Id       []rune
	Input    []rune
	Index    int
}

type State = fsm.State[AutContext]

func New(fullName string, id string) AutContext {

	fullName = strings.ToLower(fullName)
	id = strings.ToLower(id)

	var aut = AutContext{
		Alphabet: utils.GetAlphabet(fullName, id),
		Initials: utils.GetInitials(fullName),
		Id:       []rune(id),
	}

	// Obtener automatas dinamicos
	initialsStates = getContinuousCharStates(aut.Initials, StateQ1, StateQ4)
	idStates = getContinuousCharStates(aut.Id, StateQ4, StateQ6)

	return aut
}

func Eval(aut AutContext, input string) bool {

	aut.Index = 0
	aut.Input = []rune(input)

	return fsm.Run(context.Background(), aut, StateQ0)
}
