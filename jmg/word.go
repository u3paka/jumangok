package jmg

import (
	"errors"
	"strings"
)

type Word struct {
	Surface   string `json:"surface"`
	Sound     string `json:"sound"`
	Lemma     string `json:"lemma"`
	Pos       string `json:"pos"`
	PosDetail string `json:"pos_detail"`
	Conj      string `json:"conj"`
	ConjGroup string `json:"conj_group"`
	Meta      *Meta  `json:"meta"`
}

//HasDomain returns if a word struct has domain meta data.
func (w *Word) HasDomain(ds ...string) bool {
	if w.Meta == nil {
		return false
	}
	for _, d := range w.Meta.Domain {
		for _, v := range ds {
			if strings.Contains(d, v) {
				return true
			}
		}
	}
	return false
}

//HasCategory returns if a word struct has domain meta data.
func (w *Word) HasCategory(cs ...string) bool {
	if w.Meta == nil {
		return false
	}
	for _, d := range w.Meta.Category {
		for _, v := range cs {
			if strings.Contains(d, v) {
				return true
			}
		}
	}
	return false
}

//HasMetatag returns if a word struct has a certain meta tag.
func (w *Word) HasMetatag(cs ...string) bool {
	if w.Meta == nil {
		return false
	}
	for _, t := range w.Meta.SingleTags {
		for _, v := range cs {
			if strings.Contains(t, v) {
				return true
			}
		}
	}
	return false
}

//Extract returns word datas extracted with a boolean function.
func Extract(ws []*Word, f func(w *Word) bool) (rs []*Word) {
	for _, w := range ws {
		if f(w) {
			rs = append(rs, w)
		}
	}
	return
}

// CSVtoW converts csv string into worddata format.
func CSVtoW(csv string) (*Word, error) {
	csv = strings.Replace(csv, "\t", ",", 1)
	fs := strings.Split(csv, ",")
	if len(fs) < 7 {
		return &Word{}, errors.New("too short csv length")
	}
	w := &Word{
		Surface:   fs[0],
		Sound:     fs[1],
		Lemma:     fs[2],
		Pos:       fs[3],
		PosDetail: fs[4],
		Conj:      fs[5],
		ConjGroup: fs[6],
		Meta:      &Meta{},
	}
	if len(fs) > 8 {
		w.Meta.String = fs[7]
		w.Meta.tag(fs)
	}
	return w, nil
}

// WtoCSV converts a worddata into csv format.
func WtoCSV(w *Word) string {
	return strings.Join([]string{
		w.Surface,
		w.Sound,
		w.Lemma,
		w.Pos,
		w.PosDetail,
		w.Conj,
		w.ConjGroup,
		w.Meta.String,
	}, ",")
}
