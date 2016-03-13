package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// An implementation of golang tutorial @ https://tour.golang.org/concurrency/8
// The function tree.New(k) constructs a randomly-structured binary tree holding the values k, 2k, 3k, ..., 10k.

// BST in-order traverse
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		// fmt.Printf("going left\n")
		Walk(t.Left, ch)
	}
	
	fmt.Printf("sending %d\n", t.Value)
	ch <- t.Value
	
	if t.Right != nil {
		// fmt.Printf("going right\n")
		Walk(t.Right, ch)
	}
	
	return
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// the 'append()' func will always append the value in the end of the slice
	// if using 'make([]int, 5, 5)' to initialize the slice 
	// the first 5 slots will always be untouched
	p1 := []int {}
	p2 := []int {}
	
	go func() {
		Walk(t1, ch1)
		close(ch1)
		fmt.Printf("Done ch1\n")
	} ()
	go func() {
		Walk(t2, ch2)
		close(ch2)
		fmt.Printf("Done ch2\n")
	} ()
	
	// Not sure if it is guarabteed that two channel return at the same time
	// for {
	// 	v1, ok1 := <-ch1
	// 	v2, ok2 := <-ch2

	// 	isSame = isSame && v1 == v2

	// 	if ok1 == false && ok2 == false {
	// 		break
	// 	}
	// }

	for {
		v1, ok1 := <-ch1
		if ok1 {
			fmt.Printf("ch1 receiving %d\n", v1)
			p1 = append(p1, v1)
		}
		v2, ok2 := <-ch2
		if ok2 {
			fmt.Printf("ch2 receiving %d\n", v2)
			p2 = append(p2, v2)
		}
		if ok1 == false && ok2 == false {
			break
		}
	}	

	isSame := true
	
	if len(p1) != len(p2) {
		isSame = false
	} else {
		for i:= 0; i < len(p1); i++ {
			fmt.Printf("p1[%d] = %d, p2[%d] = %d\n", i, p1[i], i, p2[i])
			isSame = isSame && p1[i] == p2[i]
		}
	}
	
	return isSame
}

func main() {
	test1 := Same(tree.New(1), tree.New(2)) // false
	test2 := Same(tree.New(1), tree.New(1)) // true
	
	fmt.Printf("test1 = %t\n", test1)
	fmt.Printf("test2 = %t\n", test2)
}
