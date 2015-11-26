package triematcher

type state struct {
	Id     int
	Nexts  map[byte]*state
	IsGoal bool
}

var state_id_seq = 0

func newState() *state {
	state_id_seq++
	return &state{Id: state_id_seq - 1, Nexts: map[byte]*state{}}
}

func initMap(inputs []string) *state {
	start_s := newState()
	for _, str := range inputs {
		start_s.addString(str)
	}
	return start_s
}

func (st *state) addBytes(bytes []byte) {
	for len(bytes) == 0 {
		st.IsGoal = true
		return
	}
	var next = st.Nexts[bytes[0]]
	if next == nil {
		next := newState()
		st.Nexts[bytes[0]] = next
	}
	next.addBytes(bytes[1:])
}

func (st *state) addString(str string) {
	for len(str) == 0 {
		st.IsGoal = true
		return
	}
	next := st.Nexts[str[0]]
	if next == nil {
		next = newState()
		st.Nexts[str[0]] = next
	}
	next.addString(str[1:])
}
