package hiragana

import (
	"math/rand/v2"

	"github.com/charmbracelet/bubbles/key"
)

func getHiraganaMap() map[string]string {
	return map[string]string{
		"あ": "a",
		"い": "i",
		"う": "u",
		"え": "e",
		"お": "o",
		"か": "ka",
		"き": "ki",
		"く": "ku",
		"け": "ke",
		"こ": "ko",
		"さ": "sa",
		"し": "shi",
		"す": "tsu",
		"せ": "se",
		"そ": "so",
		"た": "ta",
		"ち": "chi",
		"つ": "tsu",
		"て": "te",
		"と": "to",
	}
}

func getRandomKana(pool map[string]string, except ...string) (string, string) {
	exceptMap := make(map[string]bool)
	for _, k := range except {
		exceptMap[k] = true
	}

	for k := range pool {
		if !exceptMap[k] {
			return k, pool[k]
		}
	}
	return "", ""
}

func getNewKana(pool map[string]string) (string, []string) {
	kana, romaji := getRandomKana(pool)
	except := []string{kana}
	alts := []string{romaji}

	n := 2
	for i := 0; i < n; i++ {
		k, r := getRandomKana(pool, except...)
		alts = append(alts, r)
		except = append(except, k)
	}

	rand.Shuffle(len(alts), func(i, j int) { alts[i], alts[j] = alts[j], alts[i] })
	return kana, alts
}

func (m Model) helpView() string {
	return "\n" + m.help.ShortHelpView([]key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
	})
}

func (m Model) KeyList() []key.Binding {
	return []key.Binding{
		m.keymap.up,
		m.keymap.down,
		m.keymap.choose,
	}
}
