package trait

type Buff struct {
	MinChampsCount int
	Weight         int
}

type Trait struct {
	Name  string
	Buffs []Buff
}

func (t Trait) IsActive(count int) bool {
	for _, points := range t.Buffs {
		if count >= points.MinChampsCount {
			return true
		}
	}

	return false
}

func (t Trait) ActiveBuff(chmapsWithSelectedBuffcount int) Buff {
	for i := len(t.Buffs) - 1; i >= 0; i-- {
		if chmapsWithSelectedBuffcount >= t.Buffs[i].MinChampsCount {
			return t.Buffs[i]
		}
	}

	return Buff{}
}
