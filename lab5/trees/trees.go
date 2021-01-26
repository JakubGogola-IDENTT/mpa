package trees

import (
	"fmt"
	"math"
	"strings"
)

// Generate creates random tree of given number of nodes
func (t *Tree) Generate(n int) {
	seq := randomArray(n)

	var prefix, sufix, word []int
	partialSum := 0

	for _, s := range seq {
		word = append(word, s)

		partialSum += s

		if partialSum != 0 {
			continue
		}

		if s == -1 {
			prefix = append(prefix, word...)
		} else {
			prefix = append(prefix, 1)

			var nextSufix []int
			nextSufix = append(nextSufix, -1)

			for _, w := range word[1 : len(word)-1] {
				nextSufix = append(nextSufix, -w)
			}

			nextSufix = append(nextSufix, sufix...)
			sufix = nextSufix
		}

		word = nil
	}

	t.Seq = append(prefix, sufix...)
}

// IsValid checks if generated sequence of braces represents valid binary tree.
func (t *Tree) IsValid() bool {
	partialSum := 0

	for _, s := range t.Seq {
		if s != rightBracket && s != leftBracket {
			return false
		}

		partialSum += s

		if partialSum < 0 {
			return false
		}
	}

	return partialSum == 0
}

// Stats creates BST representation using generated sequence
func (t *Tree) Stats() (int, int) {
	idx := 0
	min := math.MaxInt32
	max := math.MinInt32

	var rec func(level int) *Node
	rec = func(level int) *Node {
		if idx >= len(t.Seq) || t.Seq[idx] == rightBracket {
			idx++
			return nil
		}

		idx++

		node := &Node{
			left:  rec(level + 1),
			right: nil,
		}

		if idx >= len(t.Seq) || t.Seq[idx] == rightBracket {
			if node.left == nil && node.right == nil {
				min = minInt(min, level)
				max = maxInt(max, level)
			}

			idx++
			return node
		}

		node.right = rec(level + 1)
		return node
	}

	rec(0)

	return min, max
}

// SeqToString returns stringified sequence representing tree
func (t *Tree) SeqToString() string {
	mapped := make([]string, len(t.Seq))

	for i, v := range t.Seq {
		mapped[i] = fmt.Sprint(v)
	}

	return strings.Join(mapped, "")
}
