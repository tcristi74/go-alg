package InsertDelete

import (
	"fmt"
	"math/rand"
	"time"
)

//Design a data structure that supports all following operations in average O(1) time.
//
//insert(val): Inserts an item val to the set if not already present.
//remove(val): Removes an item val from the set if present.
//getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.

type RandomizedSet struct {
	M map[int]int
	P map[int]int
	//Gt int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	fmt.Println("Contruction")
	var r RandomizedSet
	r.M = make(map[int]int)
	r.P = make(map[int]int)
	//	r.Gt = 0
	return r
}

/** Inserts a value to the set. Returns true if the set did not already
//contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	_, ok := this.M[val]
	if !ok {
		gt := len(this.P)
		this.P[gt] = val
		this.M[val] = gt
		//this.Gt++
		return true
	} else {
		return false
	}
}

/** Removes a value from the set. Returns true if the set
//contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	vm, ok := this.M[val]
	if ok {
		//get the last from p
		gt := len(this.P)
		vHighp, _ := this.P[gt-1]
		this.P[vm] = vHighp

		delete(this.P, gt-1)
		this.M[vHighp] = vm

		//delete from first
		delete(this.M, val)
		return true
	}
	return false

}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.M) == 0 {
		return -1
	} else if len(this.M) == 1 {
		for _, v := range this.M {
			return v
		}
	}

	tn := time.Now()
	// use a random seed
	s1 := rand.NewSource(tn.UnixNano())
	i := rand.New(s1).Intn(len(this.P))
	//fmt.Println("index of P-",i)
	//fmt.Println("value of P",this.P[i])
	ret := this.P[i]
	return ret
}
