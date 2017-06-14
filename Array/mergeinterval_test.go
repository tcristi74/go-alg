package Array

import (
	"testing"
	"fmt"
)

func TestMergeInterval(t *testing.T){
	inter := make([]Interval,0)
	inter = append(inter,Interval{Start:1,End:5})
	inter = append(inter,Interval{Start:1,End:3})
	inter = append(inter,Interval{Start:8,End:10})
	inter = append(inter,Interval{Start:15,End:18})
	inter = append(inter,Interval{Start:2,End:6})
	inter = append(inter,Interval{Start:15,End:19})

	//[1,3],[2,6],[8,10],[15,18],
	ret :=merge(inter)

	if len(ret)==3 && ret[0].Start==1 && ret[0].End==6 {
		fmt.Println("goood")
	} else {
		fmt.Println(ret)
		t.Errorf("bad")
	}
}
