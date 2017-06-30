package Array

import (
	"sort"
	"math"
	"fmt"
)

//Given a collection of intervals, merge all overlapping intervals.
//
//For example,
//Given [1,3],[2,6],[8,10],[15,18],
//return [1,6],[8,10],[15,18].

type Interval struct {
	Start int
	End   int
}

type ByInterval []Interval

func (a ByInterval) Len() int      { return len(a) }
func (a ByInterval) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByInterval) Less(i, j int) bool {
	return a[i].Start < a[j].Start || a[i].Start == a[j].Start && a[i].End < a[j].End
}

func merge(intervals []Interval) []Interval {
	inter := make([]Interval, 0)
	sort.Sort(ByInterval(intervals))
	s:=intervals[0]
	for i:=1 ; i<len(intervals);i++{
		if s.End>intervals[i].Start{

			s=union(s,intervals[i])
			fmt.Printf("new s")
			fmt.Println(s)
		}  else
		{
			//add s to the new array
			fmt.Printf("append")
			inter = append(inter,s)
			fmt.Printf("my inter")
			fmt.Println(inter)
			s=intervals[i]
		}
	}
	fmt.Printf("append my inter")
	inter = append(inter,s)
	fmt.Println(inter)

	return inter
}

func union(a Interval, b Interval) Interval{
	fmt.Printf("pairs to compare")
	fmt.Println(a)
	fmt.Println(b)
	vmin:=int(math.Min(float64(a.Start), float64(b.Start)))
	vmax:=int(math.Max(float64(a.End), float64(b.End)))

	return Interval{vmin,vmax}
}
