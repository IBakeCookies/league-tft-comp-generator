package trait

type Trait struct {
	Name     string
	MinCount int
	Weight int
}

func (t Trait) IsActive(count int) bool {
	if t.MinCount == 1 {
		return false
	}

	return count >= t.MinCount
}