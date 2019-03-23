package jmg

import (
	"strings"

	"github.com/pkg/errors"
)

const (
	metasep   = "\""
	n         = "\n"
	space     = " "
	colon     = ":"
	semicolon = ";"
	slash     = "/"
)

func (cl *Service) GetWords(in string) (ret []*Word, err error) {
	ws := strings.Split(cl.RawParse(in), n)
	for _, w := range ws {
		// TODO: 候補一覧
		if strings.HasPrefix(w, "@") {
			continue
		}
		if w == "EOS" {
			break
		}

		fs := strings.Split(w, space)
		if len(fs) < 12 {
			err = errors.Wrap(err, "mas length < 12")
			continue
		}

		meta := strings.SplitAfterN(w, metasep, -1)
		var mt *Meta
		if len(meta) > 2 {
			mt = &Meta{
				String: strings.TrimRight(meta[1], metasep),
			}
			fs = mt.tag(fs)
		}
		ret = append(ret, &Word{
			Surface:   fs[0],
			Sound:     fs[1],
			Lemma:     fs[2],
			Pos:       fs[3],
			PosDetail: fs[5],
			Conj:      fs[7],
			ConjGroup: fs[9],
			Meta:      mt,
		})
	}
	return
}
