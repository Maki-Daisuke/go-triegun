package triematcher

type state struct {
	Id        int
	OutBounds []edge
	IsGoal    bool
}

type edge struct {
	Key   string
	State *state
}

var state_id_seq = 0

func newState() *state {
	state_id_seq++
	return &state{Id: state_id_seq - 1, OutBounds: []edge{}}
}

func initMap(inputs []string) *state {
	start_s := newState()
	for _, str := range inputs {
		addString(start_s, str)
	}
	return start_s
}

func addString(st *state, str string) {
	for len(str) == 0 {
		st.IsGoal = true
		return
	}
	next := newState()
	addString(next, str[1:])
	st.OutBounds = append(st.OutBounds, edge{Key: str[0:1], State: next})
}
