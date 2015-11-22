package triematcher

func squash(st *state) *state {
	st, _ = squash_aux(st, "")
	return st
}

func squash_aux(st *state, edg string) (*state, string) {
	switch len(st.OutBounds) {
	case 0:
		return st, edg
	case 1:
		return squash_aux(st.OutBounds[0].State, edg+st.OutBounds[0].Key)
	default:
		for i, e := range st.OutBounds {
			st.OutBounds[i].State, st.OutBounds[i].Key = squash_aux(e.State, e.Key)
		}
		return st, edg
	}
}
