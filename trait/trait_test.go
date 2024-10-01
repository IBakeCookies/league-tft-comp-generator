package trait

import (
	"reflect"
	"testing"
)

func TestTrait(t *testing.T) {
	buff1 := Buff{MinChampsCount: 2, Weight: 2}
	buff2 := Buff{MinChampsCount: 3, Weight: 3}
	buff3 := Buff{MinChampsCount: 5, Weight: 5}

	t.Run("A trait is not active", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", Buffs: []Buff{buff1, buff2}}
		got := trait.IsActive(1)
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Trait is active when count is equal to the min champ count", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", Buffs: []Buff{buff1, buff2}}
		got := trait.IsActive(2)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Trait is active when count is bigger than the min champ count", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", Buffs: []Buff{buff1, buff2}}
		got := trait.IsActive(5)
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Should get the lowest value buff", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", Buffs: []Buff{buff1, buff2}}
		got := trait.ActiveBuff(2)
		want := buff1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Should get the highest value buff", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", Buffs: []Buff{buff1, buff2, buff3}}
		got := trait.ActiveBuff(5)
		want := buff3

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
