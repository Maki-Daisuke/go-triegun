package triematcher

func allowSubmatch(start *state) *state {
	for _, state := range allStates(start)[1:] { // Ignoring the start state
		for k, n := range start.Nexts {
			if next := state.Nexts[k]; next != nil {
				bridge(next, n)
			}
		}
	}
	return start
}

/*
What `bridge` does is as follows:

For example, when we want to match "sushi" or "sukiyaski", the initial state
(DFA) is like this:

       s     u     k     i     y     a     k     i
START --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> OK
                 \ s     h     i
                  \-> 8 --> 9 --> OK

`bridge` modifies it as follow:

       s     u     k     i     y     a     k     i
START --> 1 --> 2 --> 3 --> 4 --> 5 --> 6 --> 7 --> OK
                ^\ s     h     i
                | \-> 8 --> 9 --> OK
                +---/
                  u

Connect state8 and state2, because 's' is already found during matcing "sushi".
*/
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
