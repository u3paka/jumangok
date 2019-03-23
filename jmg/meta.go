package jmg

import "strings"

// Meta is meta data struct
type Meta struct {
	String       string   `json:"string"`
	SingleTags   []string `json:"tags"`
	TypicalWords []*Word  `json:"typical"`
	Category     []string `json:"category"`
	Domain       []string `json:"domain"`
	Name         string   `json:"name"`
	Place        string   `json:"place"`
	Source       string   `json:"source"`
	Antonym      []*Word  `json:"antonym"`
	Undefined    bool     `json:"is_undefined"`
	TFscore      int64    `json:"tf_score"`
}

// TODO: to improve
func (mt *Meta) tag(fs []string) []string {
	ms := strings.Split(mt.String, space)
	for k, m := range ms {
		vs := strings.SplitN(m, colon, 2)
		// 単一タグの場合
		if len(vs) != 2 {
			mt.SingleTags = append(mt.SingleTags, m)
			continue
		}

		// e.g. カテゴリ: 施設; 場所
		vss := strings.Split(vs[1], semicolon)
		switch vs[0] {
		case "カテゴリ":
			mt.Category = vss
		case "ドメイン":
			mt.Domain = vss
		case "地名":
			mt.Place = vs[1]
		case "人名":
			mt.Name = vs[1]
		case "代表表記":
			for _, vv := range vss {
				v2 := strings.Split(vv, slash)
				ww := &Word{}
				switch len(v2) {
				case 0:
					continue
				case 1:
					ww = &Word{
						Surface: v2[0],
					}
				default:
					ww = &Word{
						Surface: v2[0],
						Sound:   v2[1],
					}
				}
				if len(ms) > k+1 {
					if !strings.Contains(ms[k+1], colon) {
						ww.Pos = ms[k+1]
					}
				}
				if ww.Pos == "" {
					ww.Pos = fs[3]
					ww.PosDetail = fs[5]
				}
				mt.TypicalWords = append(mt.TypicalWords, ww)
			}
		case "反義":
			for _, vv := range vss {
				v22 := strings.Split(vv, colon)
				v2 := strings.Split(v22[1], slash)
				mt.Antonym = append(mt.Antonym, &Word{
					Surface: v2[0],
					Sound:   v2[1],
					Pos:     v22[0],
				})
			}
		case "品詞推定":
			mt.Undefined = true
			fs[3] = vs[1]
		case "自動獲得":
			mt.Source = vs[1]
		}
	}
	return fs
}
