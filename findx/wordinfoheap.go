package findx

// WordInfoHeap ...
type WordInfoHeap []*WordInfo

// NewWordInfoHeap ...
func NewWordInfoHeap() WordInfoHeap {
	return make([]*WordInfo, 0)
}

func (wh WordInfoHeap) Len() int {
	return len(wh)
}

func (wh WordInfoHeap) Less(i, j int) bool {
	return wh[i].Freq < wh[j].Freq || (wh[i].Freq == wh[j].Freq && wh[i].Idx < wh[j].Idx)
}

func (wh WordInfoHeap) Swap(i, j int) {
	wh[i], wh[j] = wh[j], wh[i]
}

// Pop ...
func (wh *WordInfoHeap) Pop() interface{} {
	old := *wh
	n := len(old)
	x := old[n-1]
	*wh = old[0 : n-1]
	return x
}

// Push ...
func (wh *WordInfoHeap) Push(x interface{}) {
	*wh = append(*wh, x.(*WordInfo))
}
