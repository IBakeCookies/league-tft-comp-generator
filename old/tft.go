package aa

import (
	"fmt"
	"math/rand"
	"time"
)

type Trait struct {
    name     string
    minCount int
}

type Champ struct {
    name   string
    traits []*Trait
}

func (t *Trait) isActive(count int) bool {
    if t.minCount == 1 {
        return false
    }
    return count >= t.minCount
}

var (
    fated      = &Trait{"Fated", 3}
    sniper     = &Trait{"Sniper", 2}
    umbral     = &Trait{"Umbral", 2}
    warden     = &Trait{"Warden", 2}
    invoker    = &Trait{"Invoker", 2}
    dryad      = &Trait{"Dryad", 2}
    trickshot  = &Trait{"Trickshot", 2}
    porcelain  = &Trait{"Porcelain", 2}
    dragonlord = &Trait{"Dragonlord", 2}
    sage       = &Trait{"Sage", 2}
    arcanist   = &Trait{"Arcanist", 2}
    ghostly    = &Trait{"Ghostly", 2}
    altruist   = &Trait{"Altruist", 2}
    heavenly   = &Trait{"Heavenly", 2}
    bruiser    = &Trait{"Bruiser", 2}
    mythic     = &Trait{"Mythic", 3}
    behemoth   = &Trait{"Behemoth", 2}
    duelist    = &Trait{"Duelist", 2}
    fortune    = &Trait{"Fortune", 3}
    inkshadow  = &Trait{"Inkshadow", 3}
    reaper     = &Trait{"Reaper", 2}
    storyweaver = &Trait{"Storyweaver", 3}
    artist     = &Trait{"Artist", 2}
    lovers     = &Trait{"Lovers", 1}
    spiritWalker = &Trait{"SpiritWalker", 1}
    greatSage  = &Trait{"Great Sage", 2}
)

var (
    ahri      = &Champ{"Ahri", []*Trait{arcanist, fated}}
    caitlyn   = &Champ{"Caitlyn", []*Trait{ghostly, sniper}}
    chogath   = &Champ{"Cho'Gath", []*Trait{behemoth, mythic}}
    darius    = &Champ{"Darius", []*Trait{duelist, umbral}}
    garen     = &Champ{"Garen", []*Trait{storyweaver, warden}}
    jax       = &Champ{"Jax", []*Trait{inkshadow, warden}}
    khazix    = &Champ{"Kha'Zix", []*Trait{heavenly, reaper}}
    kobuko    = &Champ{"Kobuko", []*Trait{bruiser, fortune}}
    kogmaw    = &Champ{"Kog'Maw", []*Trait{invoker, mythic, sniper}}
    malphite  = &Champ{"Malphite", []*Trait{behemoth, heavenly}}
    reksai    = &Champ{"Rek'Sai", []*Trait{bruiser, dryad}}
    sivir     = &Champ{"Sivir", []*Trait{storyweaver, trickshot}}
    yasuo     = &Champ{"Yasuo", []*Trait{duelist, fated}}

	aatrox  = &Champ{"Aatrox", []*Trait{bruiser, ghostly, inkshadow}}
    gnar    = &Champ{"Gnar", []*Trait{dryad, warden}}
    janna   = &Champ{"Janna", []*Trait{dragonlord, invoker}}
    kindred = &Champ{"Kindred", []*Trait{dryad, fated, reaper}}
    lux     = &Champ{"Lux", []*Trait{arcanist, porcelain}}
    neeko   = &Champ{"Neeko", []*Trait{arcanist, heavenly, mythic}}
    qiyana  = &Champ{"Qiyana", []*Trait{duelist, heavenly}}
    riven   = &Champ{"Riven", []*Trait{altruist, bruiser, storyweaver}}
    senna   = &Champ{"Senna", []*Trait{inkshadow, sniper}}
    shen    = &Champ{"Shen", []*Trait{behemoth, ghostly}}
    teemo   = &Champ{"Teemo", []*Trait{fortune, trickshot}}
    yorick  = &Champ{"Yorick", []*Trait{behemoth, umbral}}
    zyra    = &Champ{"Zyra", []*Trait{sage, storyweaver}}

	alune     = &Champ{"Alune", []*Trait{invoker, umbral}}
    amumu     = &Champ{"Amumu", []*Trait{porcelain, warden}}
    aphelios  = &Champ{"Aphelios", []*Trait{fated, sniper}}
    bard      = &Champ{"Bard", []*Trait{mythic, trickshot}}
    diana     = &Champ{"Diana", []*Trait{dragonlord, sage}}
    illaoi    = &Champ{"Illaoi", []*Trait{arcanist, ghostly, warden}}
    soraka    = &Champ{"Soraka", []*Trait{altruist, heavenly}}
    tahmKench = &Champ{"Tahm Kench", []*Trait{bruiser, mythic}}
    thresh    = &Champ{"Thresh", []*Trait{behemoth, fated}}
    tristana  = &Champ{"Tristana", []*Trait{duelist, fortune}}
    volibear  = &Champ{"Volibear", []*Trait{duelist, inkshadow}}
    yone      = &Champ{"Yone", []*Trait{reaper, umbral}}
    zoe       = &Champ{"Zoe", []*Trait{arcanist, fortune, storyweaver}}

    annie     = &Champ{"Annie", []*Trait{fortune, invoker}}
    ashe      = &Champ{"Ashe", []*Trait{porcelain, sniper}}
    galio     = &Champ{"Galio", []*Trait{bruiser, storyweaver}}
    kaisa     = &Champ{"Kai'Sa", []*Trait{inkshadow, trickshot}}
    kayn      = &Champ{"Kayn", []*Trait{ghostly, reaper}}
    leeSin    = &Champ{"Lee Sin", []*Trait{dragonlord, duelist}}
    lillia    = &Champ{"Lillia", []*Trait{invoker, mythic}}
    morgana   = &Champ{"Morgana", []*Trait{ghostly, sage}}
    nautilus  = &Champ{"Nautilus", []*Trait{mythic, warden}}
    ornn      = &Champ{"Ornn", []*Trait{behemoth, dryad}}
    sylas     = &Champ{"Sylas", []*Trait{bruiser, umbral}}
    syndra    = &Champ{"Syndra", []*Trait{arcanist, fated}}

    azir      = &Champ{"Azir", []*Trait{dryad, invoker}}
    hwei      = &Champ{"Hwei", []*Trait{artist, mythic}}
    irelia    = &Champ{"Irelia", []*Trait{duelist, storyweaver}}
    lissandra = &Champ{"Lissandra", []*Trait{arcanist, porcelain}}
    rakan     = &Champ{"Rakan", []*Trait{altruist, dragonlord, lovers}}
    sett      = &Champ{"Sett", []*Trait{fated, umbral, warden}}
    udyr      = &Champ{"Udyr", []*Trait{behemoth, inkshadow, spiritWalker}}
    wukong    = &Champ{"WuKong", []*Trait{greatSage, heavenly}}
    xayah     = &Champ{"Xayah", []*Trait{dragonlord, trickshot, lovers}}
)

var champs = []*Champ{
    rakan,
    riven,
    soraka,
    ahri,
    illaoi,
	sett,
    lissandra,
    lux,
    neeko,
    syndra,
    zoe,
    hwei,
    chogath,
    malphite,
    ornn,
    shen,
    thresh,
    udyr,
    yorick,
    aatrox,
    galio,
    kobuko,
    reksai,
    sylas,
    tahmKench,
    darius,
    irelia,
    leeSin,
    qiyana,
    tristana,
    volibear,
    yasuo,
    wukong,
    alune,
    annie,
    azir,
    janna,
    kogmaw,
    lillia,
    kayn,
    khazix,
    kindred,
    yone,
    diana,
    morgana,
    zyra,
    aphelios,
    ashe,
    caitlyn,
    senna,
    bard,
    kaisa,
    sivir,
    teemo,
    xayah,
    amumu,
    garen,
    gnar,
    jax,
    nautilus,
}

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

func random(min, max int) int {
    return rng.Intn(max-min+1) + min
}

const teamSize = 8

func a() (activeTraits []string, team []*Champ) {
    team = []*Champ{}
    teamMap := make(map[string]bool)
    for _, champ := range team {
        teamMap[champ.name] = true
    }

    copy := make([]*Champ, 0)
    for _, champ := range champs {
        if !teamMap[champ.name] {
            copy = append(copy, champ)
        }
    }

    hasRakan := false
    hasXayah := false

    for i := len(team); i < teamSize; i++ {
        randomNum := random(0, len(copy)-1)
        curr := copy[randomNum]

        if (curr.name == "Rakan" && hasXayah) || (curr.name == "Xayah" && hasRakan) {
            i--
            continue
        }

        team = append(team, curr)

        if curr.name == "Rakan" {
            hasRakan = true
        } else if curr.name == "Xayah" {
            hasXayah = true
        }

        copy = append(copy[:randomNum], copy[randomNum+1:]...)
    }

    traits := make(map[string]*struct {
        count    int
        isActive bool
    })

    for _, champ := range team {
        for _, trait := range champ.traits {
            if _, ok := traits[trait.name]; !ok {
                traits[trait.name] = &struct {
                    count    int
                    isActive bool
                }{1, trait.isActive(1)}
            } else {
                traits[trait.name].count++
                traits[trait.name].isActive = trait.isActive(traits[trait.name].count)
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
    var bestTeam []*Champ
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
        bestTeamNames[i] = champ.name
    }

    fmt.Println(bestTeamNames, traits)
}