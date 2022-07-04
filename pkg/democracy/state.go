package democracy

import (
	"math/rand"
	"strconv"
)

type Democracy struct {
	states     []State
	population int
}

func NewDemocracyWithStates(population, numStates, numPolicy int) Democracy {
	// For now, all states start with equal population
	peoplePerState := numStates / population
	states := make([]State, numStates)
	for i := 0; i < numStates; i++ {
		states = append(states, NewStateWithCitizens(peoplePerState, numPolicy))
	}
	return Democracy{
		states:     states,
		population: peoplePerState * numStates, // Note that population isn't necessarily what we input
	}
}

type State struct {
	Citizens         []Citizen
	PolicyPreference Policy
}

func NewStateWithCitizens(population, numPolicy int) State {
	citizens := make([]Citizen, population)
	for i := 0; i < population; i++ {
		citizens = append(citizens, NewCitizenWithPolicy(numPolicy))
	}
	return State{
		Citizens:         citizens,
		PolicyPreference: NewPolicy(numPolicy),
	}
}

type Citizen struct {
	PolicyPreference Policy
}

func NewCitizenWithPolicy(n int) Citizen {
	return Citizen{PolicyPreference: NewPolicy(n)}
}

type Policy map[string]float64

func NewPolicy(n int) Policy {
	p := Policy{}
	for i := 0; i < n; i++ {
		p[strconv.Itoa(i)] = rand.Float64()
	}
	return p
}
