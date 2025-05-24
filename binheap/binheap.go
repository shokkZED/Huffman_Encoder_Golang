package binheap

import "math"

type (
	BinaryHeap []Tdata
	Tdata      int32
)

const MinData = math.MinInt32

func (b *BinaryHeap) Init(data []Tdata) {
	for _, x := range data {
		*b = append(*b, x)
	}
	b.Heapify()
}

func (b BinaryHeap) pushUp(place int) {
	if place <= 0 || place >= len(b) {
		return
	}
	x := b[place]
	parent := (place - 1) / 2
	for place > 0 && b[parent] > x {
		b[place] = b[parent]
		place = parent
		parent = (place - 1) / 2
	}
	b[place] = x
}

func (b BinaryHeap) pushDown(place int) {
	if place < 0 || place >= len(b) {
		return
	}
	x := b[place]
	for {
		left := 2*place + 1
		if left >= len(b) {
			break
		}
		minson := left
		rson := left + 1
		if rson < len(b) && b[rson] < b[minson] {
			minson = rson
		}
		if b[minson] >= x {
			break
		}
		b[place] = b[minson]
		place = minson
	}
	b[place] = x
}

func (b *BinaryHeap) Add(value Tdata) {
	*b = append(*b, value)
	bVal := *b
	bVal.pushUp(len(bVal) - 1)
}

func (b BinaryHeap) GetMax() Tdata {
	if len(b) == 0 {
		return MinData
	}
	return b[0]
}

func (b *BinaryHeap) ExtractMin() Tdata {
	if len(*b) == 0 {
		return MinData
	}
	max := (*b)[0]
	last := len(*b) - 1
	(*b)[0] = (*b)[last]
	*b = (*b)[:last]
	(*b).pushDown(0)
	return max
}

func (b *BinaryHeap) Delete(place int) {
	n := len(*b)
	if place < 0 || place >= n {
		return
	}
	last := n - 1
	x := (*b)[last]
	*b = (*b)[:last]
	bVal := *b
	bVal.Change(place, x)
}

func (b BinaryHeap) Change(place int, value Tdata) {
	if place < 0 || place >= len(b) {
		return
	}
	b[place] = value
	if place > 0 && value < b[(place-1)/2] {
		b.pushUp(place)
	} else {
		b.pushDown(place)
	}
}

func (b BinaryHeap) Heapify() {
	for k := (len(b)/2 - 1); k >= 0; k-- {
		b.pushDown(k)
	}
}

type Lmnt struct {
	Index int
	Value Tdata
}

type LocatorBinaryHeap struct {
	heap    []Lmnt
	locator []int
}

func (b *LocatorBinaryHeap) Init(data []Tdata) {
	for i, x := range data {
		b.heap = append(b.heap, Lmnt{i, x})
		b.locator = append(b.locator, i)
	}
	b.Heapify()
}

func (b LocatorBinaryHeap) pushUp(place int) {
	if place <= 0 || place >= len(b.heap) {
		return
	}
	t := b.heap[place]
	parent := (place - 1) / 2
	for place > 0 && b.heap[parent].Value < t.Value {
		b.heap[place] = b.heap[parent]
		b.locator[b.heap[parent].Index] = place
		place = parent
		parent = (place - 1) / 2
	}
	b.heap[place] = t
	b.locator[t.Index] = place
}

func (b LocatorBinaryHeap) pushDown(place int) {
	if place < 0 || place >= len(b.heap) {
		return
	}
	t := b.heap[place]
	for {
		left := 2*place + 1
		if left >= len(b.heap) {
			break
		}
		maxson := left
		rson := left + 1
		if rson < len(b.heap) && b.heap[rson].Value > b.heap[maxson].Value {
			maxson = rson
		}
		if b.heap[maxson].Value <= t.Value {
			break
		}
		b.heap[place] = b.heap[maxson]
		b.locator[b.heap[maxson].Index] = place
		place = maxson
	}
	b.heap[place] = t
	b.locator[t.Index] = place
}

func (b *LocatorBinaryHeap) Add(item Lmnt) {
	if item.Index < 0 || item.Index >= len(b.locator) || b.locator[item.Index] >= 0 {
		return
	}
	b.heap = append(b.heap, item)
	pos := len(b.heap) - 1
	b.locator[item.Index] = pos
	b.pushUp(pos)
}

func (b LocatorBinaryHeap) GetMax() Lmnt {
	if len(b.heap) == 0 {
		return Lmnt{-1, MinData}
	}
	return b.heap[0]
}

func (b *LocatorBinaryHeap) ExtractMax() Lmnt {
	if len(b.heap) == 0 {
		return Lmnt{-1, MinData}
	}
	max := b.heap[0]
	b.locator[max.Index] = -1
	last := len(b.heap) - 1
	b.heap[0] = b.heap[last]
	b.heap = b.heap[:last]
	b.pushDown(0)
	return max
}

func (h *LocatorBinaryHeap) Size() int {
	return len(h.heap)
}

func (b *LocatorBinaryHeap) Delete(index int) {
	if index < 0 || index >= len(b.locator) || b.locator[index] == -1 {
		return
	}
	place := b.locator[index]
	b.locator[index] = -1
	last := len(b.heap) - 1
	t := b.heap[last]
	b.heap[place] = t
	b.locator[t.Index] = place
	b.heap = b.heap[:last]
	b.Change(t)
}

func (b LocatorBinaryHeap) Change(item Lmnt) {
	if item.Index < 0 || item.Index >= len(b.locator) || b.locator[item.Index] == -1 {
		return
	}
	place := b.locator[item.Index]
	old := b.heap[place].Value
	b.heap[place] = item
	if item.Value > old {
		b.pushUp(place)
	} else {
		b.pushDown(place)
	}
}

func (b LocatorBinaryHeap) Heapify() {
	for k := (len(b.heap)/2 - 1); k >= 0; k-- {
		b.pushDown(k)
	}
}
