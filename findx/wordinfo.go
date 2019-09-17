package findx

// WordInfo use to record the position and frequency of string
type WordInfo struct {
	Idx  int64
	Freq int64
}

// NewWordInfo return point of a WordInfo struct
func NewWordInfo(idx int64, freq int64) *WordInfo {
	return &WordInfo{
		//Word: word,
		Idx:  idx,
		Freq: freq,
	}
}
