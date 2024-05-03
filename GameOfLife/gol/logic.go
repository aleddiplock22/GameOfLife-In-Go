package gameoflife

type State [][]bool

var ConwaysGameOfLife = RuleSet{
	LiveSurvives: LiveSurvivesRule{2, 3},
	DeadToLife:   DeadToLifeRule{3},
}

func getEmptyState() State {
	states := make(State, NUM_LINES)
	for col := range states {
		states[col] = make([]bool, NUM_LINES)
	}
	return states
}

func DefineState(tiles []*Tile) *State {
	state := getEmptyState()
	for _, tile := range tiles {
		state[tile.x][tile.y] = true
	}
	return &state
}

type LiveSurvivesRule struct {
	LowerBound int
	UpperBound int
}

type DeadToLifeRule struct {
	NumLiveNeighbours int
}

type RuleSet struct {
	LiveSurvives LiveSurvivesRule
	DeadToLife   DeadToLifeRule
}

func getNeighbours(x, y int) (neighbours [][]int) {
	// returns N valid neighbours [[x_1, y_1], [x_2, y_2], ..., [x_N, y_N]]
	for _, dx := range []int{-1, 0, 1} {
		for _, dy := range []int{-1, 0, 1} {
			new_x := x + dx
			new_y := y + dy
			if new_x == x && new_y == y {
				continue
			}
			if 0 <= new_x && new_x <= NUM_LINES-1 && 0 <= new_y && new_y <= NUM_LINES-1 {
				neighbours = append(neighbours, []int{new_x, new_y})
			}
		}
	}
	return neighbours
}

func getNumActiveNeighbours(x, y int, state *State) int {
	neighbours := getNeighbours(x, y)
	num_active := 0
	for _, neighbour := range neighbours {
		x, y := neighbour[0], neighbour[1]
		if (*state)[x][y] {
			num_active++
		}
	}
	return num_active
}

func (state *State) PerformGeneration(rules RuleSet) {
	new_state := getEmptyState()
	for x, col := range *state {
		for y, active := range col {
			num_active_neighbours := getNumActiveNeighbours(x, y, state)
			if active {
				// alive
				if num_active_neighbours >= rules.LiveSurvives.LowerBound && num_active_neighbours <= rules.LiveSurvives.UpperBound {
					new_state[x][y] = true // survives
				} else {
					new_state[x][y] = false // dies (over or under population)
				}
			} else {
				// dead
				if num_active_neighbours == rules.DeadToLife.NumLiveNeighbours {
					new_state[x][y] = true // brought to life (reproduction)
				} else {
					new_state[x][y] = false
				}
			}
		}
	}
	copy(*state, new_state)
}
