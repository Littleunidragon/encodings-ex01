package base85

var extraLenDecode = [5]int{0: 0, 2: 1, 3: 2, 4: 3}

func Encode(src []byte) (dst []byte) {
	if len(src) == 0 {
		return nil
	}
	var remove int
	for 4-(len(src)%4) != 4 {
		src = append(src, 0)
		remove++
	}
	var pos int
	dst = make([]byte, len(src)/4*5)
	for ; len(src) != 0; src = src[4:] {
		x := uint32(src[0])<<24 | uint32(src[1])<<16 | uint32(src[2])<<8 | uint32(src[3])
		for i := 4; i >= 0; i-- {
			dst[pos+i] = byte(x%85) + '!'
			x /= 85
		}
		pos += 5
	}
	dst = dst[:pos-remove]
	return dst
}
func Decode(src []byte) (dst []byte) {
	if len(src) == 0 {
		return nil
	}
	var pos, nb int
	var x uint32
	dst = make([]byte, len(src)/5*4+extraLenDecode[len(src)%5])
	for _, b := range src {
		x = x*85 + uint32(b-'!')
		nb++
		if nb == 5 {
			dst[pos] = byte(x >> 24)
			dst[pos+1] = byte(x >> 16)
			dst[pos+2] = byte(x >> 8)
			dst[pos+3] = byte(x)
			pos += 4
			x = 0
			nb = 0
		}
	}
	for i := nb; i < 5; i++ {
		x = x*85 + 'u'
	}
	for i := 1; i < nb; i++ {
		dst[pos] = byte(x >> 24)
		x <<= 8
		pos++
	}
	return dst
}
