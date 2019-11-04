package tree

import (
	"goContainer"
	"goContainer/heap/bheap"
	"goContainer/tree/btree"
	"goContainer/tree/rbtree"
	"goContainer/tree/splaytree"
	"goContainer/utils"
	"testing"
)

func BenchmarkTree_InsertOne(b *testing.B) {
	b.Run("Red-Black Tree: ", func(b *testing.B) {
		rbTree := new(rbtree.RBTree)
		rbTree.Comparator = container.IntComparator
		data := make([]int, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = utils.GenerateRandomInt()
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			rbTree.Insert(data[i], data[i])
		}
	})

	b.Run("Binary-Heap: ", func(b *testing.B) {
		minH := new(bheap.MinHeap)
		minH.Comparator = container.IntComparator
		minH.Init()
		data := make([]int, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = utils.GenerateRandomInt()
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			minH.Push(data[i])
		}
	})

	b.Run("B Tree: ", func(b *testing.B) {
		bTree := btree.NewBTree(10, container.IntComparator)
		data := make([]int, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = utils.GenerateRandomInt()
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			bTree.Insert(&btree.Item{
				Key:   data[i],
				Value: data[i],
			})
		}
	})

	b.Run("Splay Tree: ", func(b *testing.B) {
		spTree := splaytree.NewSplayTree(container.IntComparator)
		data := make([]int, b.N)
		for i := 0; i < b.N; i++ {
			data[i] = utils.GenerateRandomInt()
		}

		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			_ = spTree.Insert(data[i], data[i])
		}
	})
}

/*
BenchmarkTree_InsertOne/Red-Black_Tree:_-8         	 1000000	      1446 ns/op
BenchmarkTree_InsertOne/Binary-Heap:_-8            	10000000	       136 ns/op
BenchmarkTree_InsertOne/B_Tree:_-8                 	 1000000	      1482 ns/op
BenchmarkTree_InsertOne/Splay_Tree:_-8             	 1000000	      1720 ns/op
*/
