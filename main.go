package main

import (
	"fmt"
	"huffman-encoding-project/binheap"
)

type (
	huffmanNode struct {
		Char  rune
		Freq  int
		Left  *huffmanNode
		Right *huffmanNode
	}

	PriorityQueue struct {
		Heap  *binheap.BinaryHeap
		Nodes map[int]*huffmanNode
		Index int
	}
)

func buildFrequencyMap(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}
	return freq
}

func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		Heap:  &binheap.BinaryHeap{},
		Nodes: make(map[int]*huffmanNode),
		Index: 0,
	}
}

func (pq *PriorityQueue) Push(node *huffmanNode) {
	pq.Nodes[pq.Index] = node
	pq.Heap.Add(binheap.Tdata(node.Freq))
	pq.Index++
}

func (pq *PriorityQueue) Pop() *huffmanNode {
	if len(*pq.Heap) == 0 {
		return nil
	}

	freq := pq.Heap.ExtractMin()
	for i, node := range pq.Nodes {
		if node.Freq == int(freq) {
			delete(pq.Nodes, i)
			return node
		}
	}
	return nil
}

func buildHuffmanTree(freqMap map[rune]int) *huffmanNode {
	pq := NewPriorityQueue()

	for char, freq := range freqMap {
		pq.Push(&huffmanNode{Char: char, Freq: freq})
	}

	for len(*pq.Heap) > 1 {
		left := pq.Pop()
		right := pq.Pop()
		newNode := &huffmanNode{
			Char:  0,
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}
		pq.Push(newNode)
	}

	return pq.Pop()
}

func generateCodes(node *huffmanNode, prefix string, codes map[rune]string) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		codes[node.Char] = prefix
	}
	generateCodes(node.Left, prefix+"0", codes)
	generateCodes(node.Right, prefix+"1", codes)
}

func getHuffmanCodes(root *huffmanNode) map[rune]string {
	codes := make(map[rune]string)
	generateCodes(root, "", codes)
	return codes
}

func huffmanEncode(s string) map[rune]string {
	freqMap := buildFrequencyMap(s)
	root := buildHuffmanTree(freqMap)
	return getHuffmanCodes(root)
}

func printHuffmanCodes(codes map[rune]string) {
	for char, code := range codes {
		fmt.Printf("%c:%s\n", char, code)
	}
}

func encodeStringWithHuffman(s string, codes map[rune]string) string {
	result := ""
	for _, char := range s {
		result += codes[char]
	}
	return result
}

func main() {
	s := "aaaabbbccx"
	codes := huffmanEncode(s)
	fmt.Println("Коды символов:")
	printHuffmanCodes(codes)

	encoded := encodeStringWithHuffman(s, codes)
	fmt.Println("\nСлитный код:", encoded)
}
