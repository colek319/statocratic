package democracy

type Simulation struct {
	democracy     Democracy
	actions       []Action
	numIterations int
}

func NewSimulation(population, numStates, numPolicy int, actions ...Action) Simulation {
	return Simulation{
		democracy:     NewDemocracyWithStates(population, numStates, numPolicy),
		actions:       actions,
		numIterations: 0,
	}
}

func (s Simulation) Tick() {
	for _, a := range s.actions {
		s.democracy = a(s.democracy)
	}
}

// Action takes a democracy and affects a democracy
type Action func(democracy Democracy) Democracy
