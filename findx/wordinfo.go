package findx

type WordInfo struct {
	Word string
	Idx  int64
	Freq int64
}

func NewWordInfo(word string, idx int64, freq int64) *WordInfo {
	return &WordInfo{
		Word: word,
		Idx:  idx,
		Freq: freq,
	}
}
