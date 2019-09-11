package findx

type WordInfo struct {
	Idx  int64
	Freq int64
}

func NewWordInfo(idx int64, freq int64) *WordInfo {
	return &WordInfo{
		//Word: word,
		Idx:  idx,
		Freq: freq,
	}
}
