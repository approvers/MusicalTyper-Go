package beatmap

import (
	"fmt"
	"strings"

	"musicaltyper-go/game/logger"
)

// Solve divides hiragana string to slice of Character
func Solve(HiraganaSentence string) []*Character {
	Result := make([]*Character, 0)

	Chars := strings.Split(HiraganaSentence, "")
	for i, c := range Chars {
		RomaStyles := make([]*RomaStyle, 0)

		for _, v := range GetRoma(c) {
			RomaStyles = append(RomaStyles, &RomaStyle{
				Forwards: 1,
				Roma:     v,
			})
		}

		if i < len(Chars)-1 {
			for _, v := range GetShortStyleRoma(c + Chars[i+1]) {
				if v == "" {
					break
				}
				RomaStyles = append(RomaStyles, &RomaStyle{
					Forwards: 2,
					Roma:     v,
				})
			}
		}

		Result = append(Result, &Character{
			Character:  c,
			RomaStyles: RomaStyles,
		})
	}

	//っ
	for i, r := range Result {
		if r.Character != "っ" {
			continue
		}

		if len(Result) <= i+1 {
			continue
		}

		for _, v := range Result[i+1].RomaStyles {
			if len(v.Roma) == 1 {
				continue
			}

			r.RomaStyles = append(r.RomaStyles, &RomaStyle{
				Forwards: v.Forwards + 1,
				Roma:     GetSmallTsuPattern(v.Roma),
			})
		}
	}

	return Result
}

func GetSmallTsuPattern(src string) string {
	runes := []rune(src)
	result := []rune{runes[0], runes[0]}

	for i := 1; i < len(runes); i++ {
		result = append(result, runes[i])
	}

	return string(result)
}

func GetRoma(Character string) []string {
	switch Character {
	case "あ":
		return []string{"a"}
	case "い":
		return []string{"i"}
	case "う":
		return []string{"u"}
	case "え":
		return []string{"e"}
	case "お":
		return []string{"o"}

	case "か":
		return []string{"ka"}
	case "き":
		return []string{"ki"}
	case "く":
		return []string{"ku"}
	case "け":
		return []string{"ke"}
	case "こ":
		return []string{"ko"}

	case "が":
		return []string{"ga"}
	case "ぎ":
		return []string{"gi"}
	case "ぐ":
		return []string{"gu"}
	case "げ":
		return []string{"ge"}
	case "ご":
		return []string{"go"}

	case "さ":
		return []string{"sa"}
	case "し":
		return []string{"si", "shi"}
	case "す":
		return []string{"su"}
	case "せ":
		return []string{"se"}
	case "そ":
		return []string{"so"}

	case "ざ":
		return []string{"za"}
	case "じ":
		return []string{"zi", "ji"}
	case "ず":
		return []string{"zu"}
	case "ぜ":
		return []string{"ze"}
	case "ぞ":
		return []string{"zo"}

	case "た":
		return []string{"ta"}
	case "ち":
		return []string{"ti", "chi"}
	case "つ":
		return []string{"tu", "tsu"}
	case "て":
		return []string{"te"}
	case "と":
		return []string{"to"}

	case "だ":
		return []string{"da"}
	case "ぢ":
		return []string{"di"}
	case "づ":
		return []string{"du"}
	case "で":
		return []string{"de"}
	case "ど":
		return []string{"do"}

	case "な":
		return []string{"na"}
	case "に":
		return []string{"ni"}
	case "ぬ":
		return []string{"nu"}
	case "ね":
		return []string{"ne"}
	case "の":
		return []string{"no"}

	case "は":
		return []string{"ha"}
	case "ひ":
		return []string{"hi"}
	case "ふ":
		return []string{"fu", "hu"}
	case "へ":
		return []string{"he"}
	case "ほ":
		return []string{"ho"}

	case "ば":
		return []string{"ba"}
	case "び":
		return []string{"bi"}
	case "ぶ":
		return []string{"bu"}
	case "べ":
		return []string{"be"}
	case "ぼ":
		return []string{"bo"}

	case "ぱ":
		return []string{"pa"}
	case "ぴ":
		return []string{"pi"}
	case "ぷ":
		return []string{"pu"}
	case "ぺ":
		return []string{"pe"}
	case "ぽ":
		return []string{"po"}

	case "ま":
		return []string{"ma"}
	case "み":
		return []string{"mi"}
	case "む":
		return []string{"mu"}
	case "め":
		return []string{"me"}
	case "も":
		return []string{"mo"}

	case "や":
		return []string{"ya"}
	case "ゆ":
		return []string{"yu"}
	case "よ":
		return []string{"yo"}

	case "ら":
		return []string{"ra"}
	case "り":
		return []string{"ri"}
	case "る":
		return []string{"ru"}
	case "れ":
		return []string{"re"}
	case "ろ":
		return []string{"ro"}

	case "わ":
		return []string{"wa"}
	case "を":
		return []string{"wo"}
	case "ん":
		return []string{"nn"}

	case "ぁ":
		return []string{"la", "xa"}
	case "ぃ":
		return []string{"li", "xi"}
	case "ぅ":
		return []string{"lu", "xu"}
	case "ぇ":
		return []string{"le", "xe"}
	case "ぉ":
		return []string{"lo", "xo"}
	case "っ":
		return []string{"ltu", "xtu"}
	case "ゃ":
		return []string{"lya", "xya"}
	case "ゅ":
		return []string{"lyu", "xyu"}
	case "ょ":
		return []string{"lyo", "xyo"}

	case "ゔ":
		return []string{"vu"}

	case "、":
		return []string{","}
	case "。":
		return []string{"."}
	case "ー":
		return []string{"-"}
	case " ", "　":
		return []string{" "}

	default:
		switch char := ([]rune(Character))[0]; {
		case char >= 'A' && char <= 'Z':
			return []string{strings.ToLower(string(char))}

		case char >= 'a' && char <= 'z':
			return []string{string(char)}

		case char >= '0' && char <= '9':
			return []string{string(char)}
		}

		log := logger.NewLogger("GetRoma")
		log.FatalError(fmt.Sprintf("fixme: Unknown charcter \"%s\"", Character))
		return nil
	}
}

func GetShortStyleRoma(Character string) []string {
	switch Character {
	case "きゃ":
		return []string{"kya"}
	case "きぃ":
		return []string{"kyi"}
	case "きぅ":
		return []string{"kyu"}
	case "きぇ":
		return []string{"kye"}
	case "きょ":
		return []string{"kyo"}

	case "くぁ":
		return []string{"qa"}
	case "くぃ":
		return []string{"qi"}
	case "くぅ":
		return []string{"qwu"}
	case "くぇ":
		return []string{"qe"}
	case "くぉ":
		return []string{"qo"}

	case "ぎゃ":
		return []string{"gya"}
	case "ぎぃ":
		return []string{"gyi"}
	case "ぎゅ":
		return []string{"gyu"}
	case "ぎぇ":
		return []string{"gye"}
	case "ぎょ":
		return []string{"gyo"}

	case "ぐぁ":
		return []string{"gwa"}
	case "ぐぃ":
		return []string{"gwi"}
	case "ぐぅ":
		return []string{"gwu"}
	case "ぐぇ":
		return []string{"gwe"}
	case "ぐぉ":
		return []string{"gwo"}

	case "しゃ":
		return []string{"sya", "sha"}
	case "しぃ":
		return []string{"swi"}
	case "しゅ":
		return []string{"syu", "shu"}
	case "しぇ":
		return []string{"sye", "she"}
	case "しょ":
		return []string{"syo", "sho"}

	case "すぁ":
		return []string{"swa"}
	case "すぃ":
		return []string{"swi"}
	case "すぅ":
		return []string{"swu"}
	case "すぇ":
		return []string{"swe"}
	case "すぉ":
		return []string{"swo"}

	case "じゃ":
		return []string{"ja", "zya"}
	case "じぃ":
		return []string{"zyi"}
	case "じゅ":
		return []string{"ju", "zyu"}
	case "じぇ":
		return []string{"je", "zye"}
	case "じょ":
		return []string{"jo", "zyo"}

	case "ちゃ":
		return []string{"tya", "cha"}
	case "ちぃ":
		return []string{"tyi"}
	case "ちゅ":
		return []string{"tyu", "chu"}
	case "ちぇ":
		return []string{"tye", "che"}
	case "ちょ":
		return []string{"tyo", "cho"}

	case "てゃ":
		return []string{"tha"}
	case "てぃ":
		return []string{"thi"}
	case "てゅ":
		return []string{"thu"}
	case "てぇ":
		return []string{"the"}
	case "てょ":
		return []string{"tho"}

	case "とぁ":
		return []string{"twa"}
	case "とぃ":
		return []string{"twi"}
	case "とぅ":
		return []string{"twu"}
	case "とぇ":
		return []string{"twe"}
	case "とぉ":
		return []string{"two"}

	case "ぢゃ":
		return []string{"dya"}
	case "ぢぃ":
		return []string{"dyi"}
	case "ぢゅ":
		return []string{"dyu"}
	case "ぢぇ":
		return []string{"dye"}
	case "ぢょ":
		return []string{"dyo"}

	case "でゃ":
		return []string{"dhi"}
	case "でぃ":
		return []string{"dhi"}
	case "でゅ":
		return []string{"dhu"}
	case "でぇ":
		return []string{"dhe"}
	case "でょ":
		return []string{"dho"}

	case "どぁ":
		return []string{"dwa"}
	case "どぃ":
		return []string{"dwi"}
	case "どぅ":
		return []string{"dwu"}
	case "どぇ":
		return []string{"dwe"}
	case "どぉ":
		return []string{"dwo"}

	case "にゃ":
		return []string{"nya"}
	case "にぃ":
		return []string{"nyi"}
	case "にゅ":
		return []string{"nyu"}
	case "にぇ":
		return []string{"nye"}
	case "にょ":
		return []string{"nyo"}

	case "ひゃ":
		return []string{"hya"}
	case "ひぃ":
		return []string{"hyi"}
	case "ひゅ":
		return []string{"hyu"}
	case "ひぇ":
		return []string{"hye"}
	case "ひょ":
		return []string{"hyo"}

	case "ふぁ":
		return []string{"fa"}
	case "ふぃ":
		return []string{"fi"}
	case "ふぅ":
		return []string{"fwu"}
	case "ふぇ":
		return []string{"fe"}
	case "ふぉ":
		return []string{"fo"}

	case "びゃ":
		return []string{"bya"}
	case "びぃ":
		return []string{"byi"}
	case "びゅ":
		return []string{"byu"}
	case "びぇ":
		return []string{"bye"}
	case "びょ":
		return []string{"byo"}

	case "ぴゃ":
		return []string{"pya"}
	case "ぴぃ":
		return []string{"pyi"}
	case "ぴゅ":
		return []string{"pyu"}
	case "ぴぇ":
		return []string{"pye"}
	case "ぴょ":
		return []string{"pyo"}

	case "みゃ":
		return []string{"mya"}
	case "みぃ":
		return []string{"myi"}
	case "みゅ":
		return []string{"myu"}
	case "みぇ":
		return []string{"mye"}
	case "みょ":
		return []string{"myo"}

	case "りゃ":
		return []string{"rya"}
	case "りぃ":
		return []string{"ryi"}
	case "りゅ":
		return []string{"ryu"}
	case "りぇ":
		return []string{"rye"}
	case "りょ":
		return []string{"ryo"}

	case "うぁ":
		return []string{"wha"}
	case "うぃ":
		return []string{"wi"}
	case "うぇ":
		return []string{"we"}
	case "うぉ":
		return []string{"who"}

	default:
		return []string{""}
	}
}
