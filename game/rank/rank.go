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
		{200, "Wow"},
		{150, "Unexpected"},
		{125, "Very God"},
		{100, "God"},
		{99.5, "Pro"},
		{99, "Genius"},
		{98, "Geki-tsuyo"},
		{97, "tsuyotsuyo"},
		{94, "AAA"},
		{90, "AA"},
		{80, "A"},
		{60, "B"},
		{40, "C"},
		{20, "D"},
		{10, "E"},
		{0, "F"}}
)
