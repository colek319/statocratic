package democracy

import (
	"gonum.org/v1/gonum/mat"
	"math/rand"
	"strconv"
)

// adding some test comments

type Democracy struct {
	States     []State
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
		States:     states,
		population: peoplePerState * numStates, // Note that population isn't necessarily what we input
	}
}

type State struct {
	C mat.Matrix // citizens matrix
	P mat.Vector // state policy vector
}

func NewStateWithCitizens(population, numPolicy int) State {
	citizens := make([]Citizen, population)
	for i := 0; i < population; i++ {
		citizens = append(citizens, NewCitizenWithPolicy(numPolicy))
	}
	// TODO: add random entries
	C := mat.NewDense(population, numPolicy, nil)
	P := mat.NewVecDense(numPolicy, nil)
	return State{
		C: C,
		P: P,
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
