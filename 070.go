package main

import(
    "fmt"
    "code.google.com/p/go-tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    
    if t.Left != nil {
	    Walk(t.Left, ch)
    }
    
    ch <- t.Value
    
    if t.Right != nil {
    	Walk(t.Right, ch)
    }
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
    cht1 := make(chan int)
    cht2 := make(chan int)
    go Walk(t1, cht1)
    go Walk(t2, cht2)
    t1vals := []int{}
    t2vals := []int{}
    compareLen := 10
  	var value1, value2 int
    for {
        select {
        case value1 = <- cht1:
            if len(t1vals) <= compareLen {
	            t1vals = append(t1vals, value1)
            }
        case value2 = <- cht2:
            if len(t2vals) <= compareLen {
	            t2vals = append(t2vals, value2)
            }
        }
        if len(t1vals) >= compareLen && len(t2vals) >= compareLen {
            break
        }
    }
    for i := 0; i < compareLen; i++ {
        if t1vals[i] != t2vals[i] {
            return false
        }
    }
    
	return true
}

func main() {
    if !Same(tree.New(1), tree.New(1)) {
        fmt.Println("fail!! 1 && 1")
        return
    }
    if Same(tree.New(1), tree.New(2)) {
        fmt.Println("fail!! 1 && 2")
        return
    }
    fmt.Println("all succeed!")
}
