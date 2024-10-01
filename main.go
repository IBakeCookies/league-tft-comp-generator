package main

import (
	"fmt"
	"main/champ"
	"main/trait"
	"math/rand/v2"
)

type T = trait.Trait
type B = trait.Buff
type C = champ.Champ

var (
	arcana       = T{Name: "Arcana", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 3, Weight: 3}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 5, Weight: 5}}}
	ascendant    = T{Name: "Ascendant", Buffs: []B{{MinChampsCount: 1, Weight: 3}}}
	bastion      = T{Name: "Bastion", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}, {MinChampsCount: 8, Weight: 8}}}
	batqueen     = T{Name: "BatQueen", Buffs: []B{{MinChampsCount: 1, Weight: 2}}}
	blaster      = T{Name: "Blaster", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	chrono       = T{Name: "Chrono", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	dragon       = T{Name: "Dragon", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 3, Weight: 3}}}
	druid        = T{Name: "Druid", Buffs: []B{{MinChampsCount: 1, Weight: 2}}}
	eldritch     = T{Name: "Eldritch", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 5, Weight: 5}, {MinChampsCount: 7, Weight: 7}, {MinChampsCount: 10, Weight: 10}}}
	bestFriends  = T{Name: "BestFriends", Buffs: []B{{MinChampsCount: 1, Weight: 1}}}
	faerie       = T{Name: "Faerie", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 5, Weight: 5}, {MinChampsCount: 7, Weight: 7}, {MinChampsCount: 9, Weight: 9}}}
	frost        = T{Name: "Frost", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 5, Weight: 5}, {MinChampsCount: 7, Weight: 7}, {MinChampsCount: 9, Weight: 9}}}
	honeymancy   = T{Name: "Honeymancy", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 5, Weight: 5}, {MinChampsCount: 7, Weight: 7}}}
	hunter       = T{Name: "Hunter", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	incantor     = T{Name: "Incantor", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}}}
	mage         = T{Name: "Mage", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 5, Weight: 5}, {MinChampsCount: 7, Weight: 7}, {MinChampsCount: 9, Weight: 9}}}
	multistriker = T{Name: "Multistriker", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 5, Weight: 5}, {MinChampsCount: 7, Weight: 7}, {MinChampsCount: 9, Weight: 9}}}
	portal       = T{Name: "Portal", Buffs: []B{{MinChampsCount: 3, Weight: 3}, {MinChampsCount: 6, Weight: 6}, {MinChampsCount: 8, Weight: 8}, {MinChampsCount: 10, Weight: 10}}}
	preserver    = T{Name: "Preserver", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 3, Weight: 3}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 5, Weight: 5}}}
	pyro         = T{Name: "Pyro", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 3, Weight: 3}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 5, Weight: 5}}}
	ravenous     = T{Name: "Ravenous", Buffs: []B{{MinChampsCount: 1, Weight: 1}}}
	scholar      = T{Name: "Scholar", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	shapeshifter = T{Name: "Shapeshifter", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}, {MinChampsCount: 8, Weight: 8}}}
	sugarcraft   = T{Name: "Sugarcraft", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	vanguard     = T{Name: "Vanguard", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	warrior      = T{Name: "Warrior", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}}}
	witchcraft   = T{Name: "Witchcraft", Buffs: []B{{MinChampsCount: 2, Weight: 2}, {MinChampsCount: 4, Weight: 4}, {MinChampsCount: 6, Weight: 6}, {MinChampsCount: 8, Weight: 8}}}
)

var (
	ashe       = C{Name: "Ashe", Traits: []T{eldritch, multistriker}}
	blitzcrank = C{Name: "Blitzcrank", Traits: []T{vanguard, honeymancy}}
	elise      = C{Name: "Elise", Traits: []T{eldritch, shapeshifter}}
	jax        = C{Name: "Jax", Traits: []T{chrono, multistriker}}
	jayce      = C{Name: "Jayce", Traits: []T{shapeshifter, portal}}
	lillia     = C{Name: "Lillia", Traits: []T{bastion, faerie}}
	nomsy      = C{Name: "Nomsy", Traits: []T{hunter, dragon}}
	poppy      = C{Name: "Poppy", Traits: []T{witchcraft, bastion}}
	seraphine  = C{Name: "Seraphine", Traits: []T{mage, faerie}}
	soraka     = C{Name: "Soraka", Traits: []T{mage, sugarcraft}}
	twitch     = C{Name: "Twitch", Traits: []T{frost, hunter}}
	warwick    = C{Name: "Warwick", Traits: []T{frost, vanguard}}
	ziggs      = C{Name: "Ziggs", Traits: []T{honeymancy, incantor}}
	zoe        = C{Name: "Zoe", Traits: []T{witchcraft, scholar, portal}}

	ahri       = C{Name: "Ahri", Traits: []T{scholar, arcana}}
	akali      = C{Name: "Akali", Traits: []T{warrior, multistriker, pyro}}
	cassiopeia = C{Name: "Cassiopeia", Traits: []T{witchcraft, incantor}}
	galio      = C{Name: "Galio", Traits: []T{vanguard, mage, portal}}
	kassadin   = C{Name: "Kassadin", Traits: []T{multistriker, portal}}
	kogmaw     = C{Name: "Kog'Maw", Traits: []T{hunter, honeymancy}}
	nilah      = C{Name: "Nilah", Traits: []T{warrior, eldritch}}
	nunu       = C{Name: "Nunu", Traits: []T{bastion, honeymancy}}
	rumble     = C{Name: "Rumble", Traits: []T{blaster, sugarcraft, vanguard}}
	shyvana    = C{Name: "Shyvana", Traits: []T{shapeshifter, dragon}}
	syndra     = C{Name: "Syndra", Traits: []T{incantor, eldritch}}
	tristana   = C{Name: "Tristana", Traits: []T{blaster, faerie}}
	zilean     = C{Name: "Zilean", Traits: []T{preserver, frost, chrono}}

	bard        = C{Name: "Bard", Traits: []T{scholar, preserver, sugarcraft}}
	ezreal      = C{Name: "Ezreal", Traits: []T{blaster, portal}}
	hecarim     = C{Name: "Hecarim", Traits: []T{multistriker, bastion, arcana}}
	hwei        = C{Name: "Hwei", Traits: []T{blaster, frost}}
	jinx        = C{Name: "Jinx", Traits: []T{hunter, sugarcraft}}
	katarina    = C{Name: "Katarina", Traits: []T{warrior, faerie}}
	mordekaiser = C{Name: "Mordekaiser", Traits: []T{vanguard, eldritch}}
	neeko       = C{Name: "Neeko", Traits: []T{witchcraft, shapeshifter}}
	shen        = C{Name: "Shen", Traits: []T{bastion, pyro}}
	swain       = C{Name: "Swain", Traits: []T{shapeshifter, frost}}
	veigar      = C{Name: "Veigar", Traits: []T{mage, honeymancy}}
	vex         = C{Name: "Vex", Traits: []T{mage, chrono}}
	wukong      = C{Name: "Wukong", Traits: []T{druid}}

	fiora     = C{Name: "Fiora", Traits: []T{witchcraft, warrior}}
	gwen      = C{Name: "Gwen", Traits: []T{warrior, sugarcraft}}
	kalista   = C{Name: "Kalista", Traits: []T{multistriker, faerie}}
	karma     = C{Name: "Karma", Traits: []T{incantor, chrono}}
	nami      = C{Name: "Nami", Traits: []T{mage, eldritch}}
	nasus     = C{Name: "Nasus", Traits: []T{shapeshifter, pyro}}
	olaf      = C{Name: "Olaf", Traits: []T{hunter, frost}}
	rakan     = C{Name: "Rakan", Traits: []T{preserver, faerie}}
	ryze      = C{Name: "Ryze", Traits: []T{scholar, portal}}
	tahmkench = C{Name: "Tahm Kench", Traits: []T{vanguard, arcana}}
	taric     = C{Name: "Taric", Traits: []T{bastion, portal}}
	varus     = C{Name: "Varus", Traits: []T{blaster, pyro}}

	briar   = C{Name: "Briar", Traits: []T{shapeshifter, ravenous, eldritch}}
	camille = C{Name: "Camille", Traits: []T{multistriker, chrono}}
	diana   = C{Name: "Diana", Traits: []T{bastion, frost}}
	milio   = C{Name: "Milio", Traits: []T{scholar, faerie}}
	morgana = C{Name: "Morgana", Traits: []T{witchcraft, preserver, batqueen}}
	norra   = C{Name: "Norra", Traits: []T{bestFriends, mage, portal}}
	smolder = C{Name: "Smolder", Traits: []T{blaster, dragon}}
	xerath  = C{Name: "Xerath", Traits: []T{ascendant, arcana}}
)

var champs = []C{
	ashe,
	blitzcrank,
	elise,
	jax,
	jayce,
	lillia,
	nomsy,
	poppy,
	seraphine,
	soraka,
	twitch,
	warwick,
	ziggs,
	zoe,
	ahri,
	akali,
	cassiopeia,
	galio,
	kassadin,
	kogmaw,
	nilah,
	nunu,
	rumble,
	shyvana,
	syndra,
	tristana,
	zilean,
	bard,
	ezreal,
	hecarim,
	hwei,
	jinx,
	katarina,
	mordekaiser,
	neeko,
	shen,
	swain,
	veigar,
	vex,
	wukong,
	fiora,
	gwen,
	kalista,
	karma,
	nami,
	nasus,
	olaf,
	rakan,
	ryze,
	tahmkench,
	taric,
	varus,
	briar,
	camille,
	diana,
	milio,
	morgana,
	norra,
	smolder,
	xerath,
}

func random(min, max int) int {
	return rand.IntN(max-min+1) + min
}

const teamSize = 8

func GetTeam() []C {
	team := []C{}
	champsCopy := make([]C, len(champs))

	copy(champsCopy, champs)

	// teamMap := make(map[string]bool)
	// for _, champ := range team {
	// 	teamMap[champ.Name] = true
	// }

	// champsCopy := make([]C, 0)
	// for _, champ := range champs {
	// 	if !teamMap[champ.Name] {
	// 		champsCopy = append(champsCopy, champ)
	// 	}
	// }

	for i := 0; i < teamSize; i++ {
		randomNum := random(0, len(champsCopy)-1)
		curr := champsCopy[randomNum]
		team = append(team, curr)
		champsCopy = append(champsCopy[:randomNum], champsCopy[randomNum+1:]...)
	}

	return team
}

type TraitMap struct {
	count    int
	isActive bool
	weight   int
}

func main() {
	bestTeams := [][]C{}
	bestWeights := []int{}
	bestTraits := [][]string{}

	// var wg sync.WaitGroup
	// numGoroutines := 10

	// for j := 0; j < numGoroutines; j++ {
	// wg.Add(1)
	// go func(id int) {
	// defer wg.Done()

	for i := 0; i < 1000000; i++ {
		// team := []C{akali, ahri, ashe, hecarim, jax, camille, kassadin, xerath}
		team := GetTeam()
		traits := make(map[string]TraitMap)
		currWeight := 0

		for _, champ := range team {
			for _, trait := range champ.Traits {
				newCount := traits[trait.Name].count + 1
				traits[trait.Name] = TraitMap{
					count:    newCount,
					isActive: trait.IsActive(newCount),
					weight:   trait.ActiveBuff(newCount).Weight,
				}
			}
		}

		for _, value := range traits {
			currWeight += value.weight
		}

		if currWeight > 18 {
			bestTeams = append(bestTeams, team)
			bestWeights = append(bestWeights, currWeight)
			traitsHolder := []string{}

			for key := range traits {
				if traits[key].isActive {
					traitsHolder = append(traitsHolder, key)
				}
			}

			bestTraits = append(bestTraits, traitsHolder)
		}
	}

	// }(j)
	// }

	// wg.Wait()

	for i, team := range bestTeams {
		bestTeamNames := make([]string, teamSize)

		for i, champ := range team {
			bestTeamNames[i] = champ.Name
		}

		// fmt.Println(bestTraits)
		fmt.Println(bestTeamNames, bestTraits[i], bestWeights[i])

	}
}
