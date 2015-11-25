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
		start_s.addString(str)
	}
	return start_s
}

func (st *state) addString(str string) {
	for len(str) == 0 {
		st.IsGoal = true
		return
	}
	next := newState()
	next.addString(str[1:])
	st.addOutbound(str[0:1], next)
}

func (st *state) addOutbound(key string, next *state) {
	st.OutBounds = append(st.OutBounds, edge{Key: key, State: next})
}

func (st *state) getNextByKey(k string) *state {
	for _, edg := range st.OutBounds {
		if edg.Key == k {
			return edg.State
		}
	}
	return nil
}

func (st *state) HasOutboundKey(k string) bool {
	return st.getNextByKey(k) != nil
}
