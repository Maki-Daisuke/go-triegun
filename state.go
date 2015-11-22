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

func convertToDFA(s *state) {
	marked := map[int]bool{}
	var traverse func(*state)
	traverse = func(st *state) {
		if marked[st.Id] {
			return
		}
		marked[st.Id] = true
		mergeEdges(st)
		for _, e := range st.OutBounds {
			traverse(e.State)
		}
	}
	traverse(s)
}

// Destructively update `s`
func mergeEdges(s *state) {
	old := s.OutBounds
	new := []edge{}
	for len(old) != 0 {
		x := old[0]
		for i := len(old) - 1; i >= 1; i-- {
			y := old[i]
			if x.Key == y.Key {
				x = edge{Key: x.Key, State: mergeStates(x.State, y.State)}
				old = append(old[0:i], old[i+1:]...)
			}
		}
		new = append(new, x)
		old = old[1:]
	}
	s.OutBounds = new
}

func mergeStates(x, y *state) *state {
	if x == y {
		return x
	}
	z := newState()
	for _, e := range x.OutBounds {
		z.OutBounds = append(z.OutBounds, e)
	}
	for _, e := range y.OutBounds {
		z.OutBounds = append(z.OutBounds, e)
	}
	z.IsGoal = x.IsGoal || y.IsGoal
	return z
}
