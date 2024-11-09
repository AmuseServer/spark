package utils

const defaultSortStep = 100

type SortIteration struct {
	initSort int64 // 初始化的sort值
	sort     int64 // 当前sort值
	Iter     int64 // 第几次next
	step     int64 // 每次排序的步长
}

func NewSortIteration(sort int64) *SortIteration {
	return &SortIteration{
		sort: sort,
	}
}

func (iteration *SortIteration) Next(step ...int64) int64 {
	iteration.Iter++
	if len(step) > 0 {
		iteration.step = step[0]
	} else {
		iteration.step = defaultSortStep
	}
	iteration.sort = iteration.sort + iteration.step
	return iteration.sort
}

func (iteration *SortIteration) GetSort() int64 {

	return iteration.sort
}

func (iteration *SortIteration) Clean() {
	iteration.sort = iteration.initSort
	iteration.Iter = 0
	return
}
