package triematcher

func allStates(st *state) []*state {
	marked := map[int]bool{}
	states := []*state{}

	var traverse func(*state)
	traverse = func(s *state) {
		if marked[s.Id] {
			return
		}
		states = append(states, s)
		marked[s.Id] = true
		for _, next := range s.Nexts {
			traverse(next)
		}
	}
	traverse(st)

	return states
}
