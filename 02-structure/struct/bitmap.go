package util

type Bitmap struct {
	words  []uint16
	length int
}

func NewBitmap() *Bitmap {
	return &Bitmap{}
}

func (bitmap *Bitmap) Has(num int) bool {
	word, bit := num/16, uint(num%16)
	return word < bitmap.length && (bitmap.words[word]&(1<<bit)) != 0
}

func (bitmap *Bitmap) Add(num int) {
	word, bit := num/16, uint(num%16)
	for word >= len(bitmap.words) {
		bitmap.words = append(bitmap.words, 0)
	}
	// 判断num是否已经存在bitmap中
	if bitmap.words[word]&(1<<bit) == 0 {
		bitmap.words[word] |= 1 << bit
		bitmap.length++
	}
}

func (bitmap *Bitmap) Len() int {
	return bitmap.length
}