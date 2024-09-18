package trait

import "testing"

func TestTrait(t *testing.T) {
	t.Run("A trait struct is created", func(t *testing.T) {
		got := Trait{Name:"Test Trait", MinCount: 2}
		wantName := "Test Trait"
		wantMinCount := 2

		if got.Name != wantName {
			t.Errorf("got %s want %s", got.Name, wantName)
		}

		if got.MinCount != wantMinCount {
			t.Errorf("got %d want %d", got.MinCount, wantMinCount)
		}
	})

	t.Run("Trait is active when count is greater than or equal to MinCount", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", MinCount: 2}

		got := trait.IsActive(2)
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("Trait is not active when count is smaller than MinCount", func(t *testing.T) {
		trait := Trait{Name: "Test Trait", MinCount: 4}

		got := trait.IsActive(2)
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}