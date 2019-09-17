package sutil

func BKDRHash(str string) uint64 {
	var hashcode uint64 = 0
	var seed uint64 = 131
	for i := 0; i < len(str); i++ {
		hashcode = hashcode*seed + uint64(str[i])
	}
	return hashcode
}
