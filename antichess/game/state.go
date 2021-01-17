package game

// State = State struct for antichess
type State struct {
	player int
}

// Newstate = returns a ptr to a new instance of the State struct
func Newstate(player int) *State {
	s := State{player: player}
	return &s
}

// Copystate = copies an existing state
func Copystate(state State) *State {
	return &state
}
