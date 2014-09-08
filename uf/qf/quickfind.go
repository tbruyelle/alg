package quickfind

type QuickFind []int

func New(N int) QuickFind {
	qf := make([]int, N)
	for i := 0; i < N; i++ {
		qf[i] = i
	}
	return qf
}

func (qf QuickFind) Connected(p, q int) bool {
	return qf[p] == qf[q]
}

func (qf QuickFind) Union(p, q int) {
	pid, qid := qf[p], qf[q]
	for i := 0; i < len(qf); i++ {
		if qf[i] == pid {
			qf[i] = qid
		}
	}
}
