package main

import (
	"container/heap"
	"strconv"
)

type State struct {
	pos    string
	weight int
	index  int
	prev   *State
}

type PriorityQueue []*State

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(item interface{}) {
	n := len(*pq)
	state := item.(*State)
	state.index = n
	*pq = append(*pq, state)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *State) {
	heap.Fix(pq, item.index)
}

func solve23(fileName string) int {
	lines := readLines(fileName)

	pos := ""
	for _, line := range lines {
		for i, c := range line {
			if c != '#' {
				pos += line[i : i+1]
			}
		}
	}

	bests := map[string]*State{}
	pq := make(PriorityQueue, 0)
	initial := State{pos, 0, -1, nil}
	pq.Push(&initial)
	bests[initial.pos] = &initial
	heap.Init(&pq)

	c := 0
	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*State)
		if state.pos[11:15] == "ABCD" {
			for s := state; s != nil; s = s.prev {
				s.print()
			}
			return state.weight
		}
		for _, next := range state.next() {
			next = next.greedy()
			best, ok := bests[next.pos]
			if !ok {
				heap.Push(&pq, next)
				bests[next.pos] = next
			} else if best.weight > next.weight {
				best.weight = next.weight
				best.prev = next.prev
				pq.update(best)
			}
		}
		c++
	}

	return 0
}

func (state *State) greedy() *State {
	pos := state.pos
	for i := range pos[:11] {
		if pos[i] != '.' {
			target := 11 + int(pos[i]-'A')
			if pos[target] != '.' {
				continue
			}
			to := (target - 10) * 2
			free := true
			dj := intSign(to - i)
			for j := i + dj; j != to; j += dj {
				free = free && pos[j] == '.'
			}
			if free {
				for ; target < len(pos)-4; target += 4 {
				}
				for ; pos[target] == pos[i]; target -= 4 {
				}
				if pos[target] == '.' {
					return state.move(i, target).greedy()
				}
			}
		}
	}
	return state
}

func (state *State) next() (result []*State) {
	for i := range state.pos {
		if state.pos[i] != '.' {
			if 11 <= i && i < 15 {
				result = state.exit(result, i, ((i-11)%4+1)*2)
			} else if 15 <= i && state.pos[i-4] == '.' {
				result = state.exit(result, i, ((i-11)%4+1)*2)
			}
		}
	}
	return result
}

func (state *State) exit(result []*State, from int, target int) []*State {
	for to := target; 0 <= to && state.pos[to] == '.'; to-- {
		if to == 0 || to == 10 || (to&1) == 1 {
			result = append(result, state.move(from, to))
		}
	}
	for to := target; to < 11 && state.pos[to] == '.'; to++ {
		if to == 0 || to == 10 || (to&1) == 1 {
			result = append(result, state.move(from, to))
		}
	}
	return result
}

func coords(index int) (int, int) {
	if index < 11 {
		return 0, index
	} else {
		return (index - 7) / 4, 2 * ((index-11)%4 + 1)
	}
}

var weights map[uint8]int

func (state *State) move(from int, to int) *State {
	cost := weights[state.pos[from]]

	oldR, oldC := coords(from)
	newR, newC := coords(to)
	weight := state.weight + cost*(intAbs(newR-oldR)+intAbs(newC-oldC))
	return &State{set(set(state.pos, from, '.'), to, rune(state.pos[from])), weight, -1, state}
}

func set(pos string, i int, value rune) string {
	return pos[:i] + string(value) + pos[i+1:]
}

func (state State) str() string {
	str := state.pos[0:11]
	for i := 11; i < len(state.pos); i += 4 {
		str += " " + state.pos[i:i+4]
	}
	return str + " " + strconv.Itoa(state.weight)
}

func (state State) print() {
	println(state.str())
}

func main() {
	weights = map[uint8]int{'A': 1, 'B': 10, 'C': 100, 'D': 1000}

	println(solve23("23_1.in"))
	println(solve23("23_2.in"))
	println(solve23("23_3.in"))
	println(solve23("23_4.in"))
}
