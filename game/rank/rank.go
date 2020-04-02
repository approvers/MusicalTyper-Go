package rank

// Rank expresses player rank
type Rank struct {
	bottom float64
	text   string
}

// FromAchievementRate decides player rank
func FromAchievementRate(rate float64) Rank {
	for _, rank := range ranks {
		if rank.bottom < rate*100 {
			return rank
		}
	}
	return ranks[len(ranks)-1]
}

func (r Rank) Text() string {
	return r.text
}

var (
	ranks = [...]Rank{
		Rank{200, "Wow"},
		Rank{150, "Unexpected"},
		Rank{125, "Very God"},
		Rank{100, "God"},
		Rank{99.5, "Pro"},
		Rank{99, "Genius"},
		Rank{98, "Geki-tsuyo"},
		Rank{97, "tsuyotsuyo"},
		Rank{94, "AAA"},
		Rank{90, "AA"},
		Rank{80, "A"},
		Rank{60, "B"},
		Rank{40, "C"},
		Rank{20, "D"},
		Rank{10, "E"},
		Rank{0, "F"}}
)
