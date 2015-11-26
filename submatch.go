package triematcher

func allowSubmatch(start *state) *state {
	for _, st := range allStates(start) {
		bridge(st, start)
	}
	return start
}

func bridge(src, dst *state) {
	marked := map[int]bool{}
	n := 0
	var traverse func(*state, *state)
	traverse = func(src, dst *state) {
		n++
		if n > 10 {
			return
		}
		if marked[src.Id] {
			return
		}
		marked[src.Id] = true
		for k, n := range dst.Nexts {
			next := src.Nexts[k]
			if next == nil {
				src.Nexts[k] = n
			} else {
				traverse(next, n)
			}
		}
	}
	traverse(src, dst)
}
