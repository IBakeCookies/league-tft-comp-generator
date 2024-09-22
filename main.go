package main

import (
	"fmt"
	"main/champ"
	"main/trait"
	"math/rand"
	"time"
)

type T = trait.Trait
type C = champ.Champ

var (
    fated      = T{Name:"Fated", MinCount: 3, Weight: 2}
    sniper      = T{Name: "Sniper", MinCount: 2, Weight: 2}
    umbral      = T{Name: "Umbral", MinCount: 2, Weight: 2}
    warden      = T{Name: "Warden", MinCount: 2, Weight: 2}
    invoker     = T{Name: "Invoker", MinCount: 2, Weight: 2}
    dryad       = T{Name: "Dryad", MinCount: 2, Weight: 2}
    trickshot   = T{Name: "Trickshot", MinCount: 2, Weight: 2}
    porcelain   = T{Name: "Porcelain", MinCount: 2, Weight: 2}
    dragonlord  = T{Name: "Dragonlord", MinCount: 2, Weight: 2}
    sage        = T{Name: "Sage", MinCount: 2, Weight: 2}
    arcanist    = T{Name: "Arcanist", MinCount: 2, Weight: 2}
    ghostly     = T{Name: "Ghostly", MinCount: 2, Weight: 2}
    altruist    = T{Name: "Altruist", MinCount: 2, Weight: 2}
    heavenly    = T{Name: "Heavenly", MinCount: 2, Weight: 2}
    bruiser     = T{Name: "Bruiser", MinCount: 2, Weight: 2}
    mythic      = T{Name: "Mythic", MinCount: 3, Weight: 2}
    behemoth    = T{Name: "Behemoth", MinCount: 2, Weight: 2}
    duelist     = T{Name: "Duelist", MinCount: 2, Weight: 2}
    fortune     = T{Name: "Fortune", MinCount: 3, Weight: 2}
    inkshadow   = T{Name: "Inkshadow", MinCount: 3, Weight: 2}
    reaper      = T{Name: "Reaper", MinCount: 2, Weight: 2}
    storyweaver = T{Name: "Storyweaver", MinCount: 3, Weight: 2}
    artist      = T{Name: "Artist", MinCount: 2, Weight: 2}
    lovers      = T{Name: "Lovers", MinCount: 1, Weight: 2}
    spiritWalker = T{Name: "SpiritWalker", MinCount: 1, Weight: 2}
    greatSage   = T{Name: "Great Sage", MinCount: 2, Weight: 2}
)

var (
    ahri      = C{Name: "Ahri", Traits: []T{arcanist, fated}}
    caitlyn   = C{Name: "Caitlyn", Traits: []T{ghostly, sniper}}
    chogath   = C{Name: "Cho'Gath", Traits: []T{behemoth, mythic}}
    darius    = C{Name: "Darius", Traits: []T{duelist, umbral}}
    garen     = C{Name: "Garen", Traits: []T{storyweaver, warden}}
    jax       = C{Name: "Jax", Traits: []T{inkshadow, warden}}
    khazix    = C{Name: "Kha'Zix", Traits: []T{heavenly, reaper}}
    kobuko    = C{Name: "Kobuko", Traits: []T{bruiser, fortune}}
    kogmaw    = C{Name: "Kog'Maw", Traits: []T{invoker,mythic, sniper}}
    malphite  = C{Name: "Malphite", Traits: []T{behemoth, heavenly}}
    reksai    = C{Name: "Rek'Sai", Traits: []T{bruiser, dryad}}
    sivir     = C{Name: "Sivir", Traits: []T{storyweaver, trickshot}}
    yasuo     = C{Name: "Yasuo", Traits: []T{duelist, fated}}

	aatrox  = C{Name: "Aatrox", Traits: []T{bruiser, ghostly, inkshadow}}
	gnar    = C{Name: "Gnar", Traits: []T{dryad, warden}}
	janna   = C{Name: "Janna", Traits: []T{dragonlord, invoker}}
	kindred = C{Name: "Kindred", Traits: []T{dryad, fated, reaper}}
	lux     = C{Name: "Lux", Traits: []T{arcanist, porcelain}}
	neeko   = C{Name: "Neeko", Traits: []T{arcanist, heavenly, mythic}}
	qiyana  = C{Name: "Qiyana", Traits: []T{duelist, heavenly}}
	riven   = C{Name: "Riven", Traits: []T{altruist, bruiser, storyweaver}}
	senna   = C{Name: "Senna", Traits: []T{inkshadow, sniper}}
	shen    = C{Name: "Shen", Traits: []T{behemoth, ghostly}}
	teemo   = C{Name: "Teemo", Traits: []T{fortune, trickshot}}
	yorick  = C{Name: "Yorick", Traits: []T{behemoth, umbral}}
	zyra    = C{Name: "Zyra", Traits: []T{sage, storyweaver}}

	alune     = C{Name: "Alune", Traits: []T{invoker, umbral}}
	amumu     = C{Name: "Amumu", Traits: []T{porcelain, warden}}
	aphelios  = C{Name: "Aphelios", Traits: []T{fated, sniper}}
	bard      = C{Name: "Bard", Traits: []T{mythic, trickshot}}
	diana     = C{Name: "Diana", Traits: []T{dragonlord, sage}}
	illaoi    = C{Name: "Illaoi", Traits: []T{arcanist, ghostly, warden}}
	soraka    = C{Name: "Soraka", Traits: []T{altruist, heavenly}}
	tahmKench = C{Name: "Tahm Kench", Traits: []T{bruiser, mythic}}
	thresh    = C{Name: "Thresh", Traits: []T{behemoth, fated}}
	tristana  = C{Name: "Tristana", Traits: []T{duelist, fortune}}
	volibear  = C{Name: "Volibear", Traits: []T{duelist, inkshadow}}
	yone      = C{Name: "Yone", Traits: []T{reaper, umbral}}
	zoe       = C{Name: "Zoe", Traits: []T{arcanist, fortune, storyweaver}}

	annie     = C{Name: "Annie", Traits: []T{fortune, invoker}}
	ashe      = C{Name: "Ashe", Traits: []T{porcelain, sniper}}
	galio     = C{Name: "Galio", Traits: []T{bruiser, storyweaver}}
	kaisa     = C{Name: "Kai'Sa", Traits: []T{inkshadow, trickshot}}
	kayn      = C{Name: "Kayn", Traits: []T{ghostly, reaper}}
	leeSin    = C{Name: "Lee Sin", Traits: []T{dragonlord, duelist}}
	lillia    = C{Name: "Lillia", Traits: []T{invoker, mythic}}
	morgana   = C{Name: "Morgana", Traits: []T{ghostly, sage}}
	nautilus  = C{Name: "Nautilus", Traits: []T{mythic, warden}}
	ornn      = C{Name: "Ornn", Traits: []T{behemoth, dryad}}
	sylas     = C{Name: "Sylas", Traits: []T{bruiser, umbral}}
	syndra    = C{Name: "Syndra", Traits: []T{arcanist, fated}}

	azir      = C{Name: "Azir", Traits: []T{dryad, invoker}}
	hwei      = C{Name: "Hwei", Traits: []T{artist, mythic}}
	irelia    = C{Name: "Irelia", Traits: []T{duelist, storyweaver}}
	lissandra = C{Name: "Lissandra", Traits: []T{arcanist, porcelain}}
	rakan     = C{Name: "Rakan", Traits: []T{altruist, dragonlord, lovers}}
	sett      = C{Name: "Sett", Traits: []T{fated, umbral, warden}}
	udyr      = C{Name: "Udyr", Traits: []T{behemoth, inkshadow, spiritWalker}}
	wukong    = C{Name: "WuKong", Traits: []T{greatSage, heavenly}}
	xayah     = C{Name: "Xayah", Traits: []T{dragonlord, trickshot, lovers}}
)

var champs = []C{
	ahri, caitlyn,  chogath, darius, garen, jax, khazix, kobuko, kogmaw, malphite, reksai, sivir, yasuo, aatrox, gnar, janna, kindred, lux, neeko, qiyana, riven, senna, shen, teemo, yorick, zyra, alune, amumu, aphelios, bard, diana, illaoi, soraka, tahmKench, thresh, tristana, volibear, yone, zoe, annie, ashe, galio, kaisa, kayn, leeSin, lillia, morgana, nautilus, ornn, sylas, syndra, azir, hwei, irelia, lissandra, rakan, sett, udyr, wukong, xayah,
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func random(min, max int) int {
    return rng.Intn(max - min + 1) + min
}

const teamSize = 8

func a() (activeTraits []string, team []C) {
    team = []C{}
    teamMap := make(map[string]bool)
    for _, champ := range team {
        teamMap[champ.Name] = true
    }

    copy := make([]C, 0)
    for _, champ := range champs {
        if !teamMap[champ.Name] {
            copy = append(copy, champ)
        }
    }

    hasRakan := false
    hasXayah := false

    for i := len(team); i < teamSize; i++ {
        randomNum := random(0, len(copy)-1)
        curr := copy[randomNum]

        if (curr.Name == "Rakan" && hasXayah) || (curr.Name == "Xayah" && hasRakan) {
            i--
            continue
        }

        team = append(team, curr)

        if curr.Name == "Rakan" {
            hasRakan = true
        } else if curr.Name == "Xayah" {
            hasXayah = true
        }

        copy = append(copy[:randomNum], copy[randomNum+1:]...)
    }

    traits := make(map[string]*struct {
        count    int
        isActive bool
    })

    for _, champ := range team {
        for _, trait := range champ.Traits {
            if _, ok := traits[trait.Name]; !ok {
                traits[trait.Name] = &struct {
                    count    int
                    isActive bool
                }{1, trait.IsActive(1)}
            } else {
                traits[trait.Name].count++
                traits[trait.Name].isActive = trait.IsActive(traits[trait.Name].count)
            }
        }
    }

    for trait, t := range traits {
        if t.isActive {
            activeTraits = append(activeTraits, trait)
        }
    }

    return
}

func main() {
    var bestTeam []C
    var traits []string

    for i := 0; i < 1000000; i++ {
        activeTraits, team := a()

        if len(activeTraits) > len(traits) {
            traits = activeTraits
            bestTeam = team
        }
    }

    bestTeamNames := make([]string, len(bestTeam))
    for i, champ := range bestTeam {
        bestTeamNames[i] = champ.Name
    }

    fmt.Println(bestTeamNames, traits)
}
