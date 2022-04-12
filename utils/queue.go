package utils

// mock 队列
type queue struct {
	list []interface{}
}

func InitQueue() *queue {
	res := make([]interface{}, 0)
	return &queue{list: res}
}

func (q *queue) Push(a interface{}) {
	q.list = append(q.list, a)
}

func (q *queue) Pop() interface{} {
	if len(q.list) == 0 {
		return 0
	}
	if len(q.list) == 1 {
		res := q.list[0]
		q.list = make([]interface{}, 0)
		return res
	}
	res := q.list[0]
	q.list = q.list[1:]
	return res
}

func (q *queue) Length() int {
	return len(q.list)
}
