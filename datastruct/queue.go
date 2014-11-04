package datastruct

type Queue struct {
	a []interface{}
}

func NewQueue() *Queue {
	q := &Queue{}
	q.a = make([]interface{}, 0)
	return q
}

func (q *Queue) Push(x interface{}) {
	q.a = append(q.a, x)
}

func (q *Queue) Pull() interface{} {
	if len(q.a) == 0 {
		return nil
	}
	x := q.a[0]
	q.a = q.a[1:]
	return x
}
