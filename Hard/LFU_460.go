package Hard

import (
	"fmt"

)

//Design and implement a data structure for Least Frequently Used (LFU) cache.
// It should support the following operations: get and put.
//get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
//put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity,
// it should invalidate the least frequently used item before inserting a new item.
// For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency),
// the least recently used key would be evicted.
//Could you do both operations in O(1) time complexity?
//LFUCache cache = new LFUCache( 2 /* capacity */ );
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // returns 1
//cache.put(3, 3);    // evicts key 2
//cache.get(2);       // returns -1 (not found)
//cache.get(3);       // returns 3.
//cache.put(4, 4);    // evicts key 1.
//cache.get(1);       // returns -1 (not found)
//cache.get(3);       // returns 3
//cache.get(4);       // returns 4

type LFUCache struct {

	Ht map[int]int
	Capacity int
	CurrentCount int
}


func Constructor(capacity int) LFUCache {

	return LFUCache{ Ht:make(map[int]int) , Capacity:capacity}
}


func (this *LFUCache) Get(key int) int {
   if v,ok:= this.Ht[key]; ok {
	   return v
   } else {
	   return -1
   }

}


func (this *LFUCache) Put(key int, value int)  {
	if this.CurrentCount<this.Capacity{
		this.CurrentCount++
	}
	this.Ht[key]=value
}

func MyCache(cap int) {
	obj := Constructor(cap)
	fmt.Println(obj)
}
/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */