package democracy

type Democracy struct {
	states []State
}

type State struct {
	Citizens         []Citizen
	PolicyPreference Policy
}

type Citizen struct {
	PolicyPreference Policy
}

type Policy map[string]float32
