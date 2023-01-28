package brokennode

const (
	working      = 'W'
	broken       = 'B'
	undetermined = '?'
)

type result struct {
	buffer []rune
}

func FindBrokenNodes(brokenNodes int, reports []bool) string {

	res := result{buffer: make([]rune, 0)}

	res.calculate(brokenNodes, 1, reports, []rune{working})
	res.calculate(brokenNodes-1, 1, reports, []rune{broken})

	return string(res.buffer)
}

func (r *result) calculate(brokenNodes, i int, reportedCorrect []bool, buf []rune) {
	if brokenNodes < 0 {
		return
	}

	last := buf[len(buf)-1]
	if i == len(reportedCorrect) {
		if brokenNodes != 0 {
			return
		}

		first := buf[0]

		if last == broken || reportedCorrect[i-1] && first == working || !reportedCorrect[i-1] && first == broken {
			r.merge(buf)
		}
		return
	}

	if last == broken || reportedCorrect[i-1] {
		r.calculate(brokenNodes, i+1, reportedCorrect, append(buf, working))
	}

	if last == broken || !reportedCorrect[i-1] {
		r.calculate(brokenNodes-1, i+1, reportedCorrect, append(buf, broken))
	}
}

func (r *result) merge(upd []rune) {
	if len(r.buffer) == 0 {
		r.buffer = make([]rune, len(upd))
		copy(r.buffer, upd)
		return
	}

	for k, v := range upd {
		if v != r.buffer[k] {
			r.buffer[k] = undetermined
		}
	}
}
