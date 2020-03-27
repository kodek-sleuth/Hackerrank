package main

import (
	"fmt"
	"sort"
)

type Ranker struct {
	Score int
	Rank int
}

var results = []int{}

func ranker(scores []int, alice int, ranker []Ranker){
	rank := 1
	for i, _ := range scores {
		if i == len(scores) - 1 {
			break
		}
		if scores[0] == alice {
			rank += 1
			x := Ranker{Score:alice, Rank:1}
			ranker = append(ranker, x)
			break
		}

		if scores[i] == scores[i+1]{
			rank += 0
			x := Ranker{Score:scores[i], Rank:rank}
			j :=  Ranker{Score:scores[i+1], Rank:rank}
			ranker = append(ranker, x, j)
		}

		if scores[i] > scores[i+1] {
			rank += 1
			x := Ranker{Score:scores[i+1], Rank:rank}
			ranker = append(ranker, x)
		}
	}

	for _, l := range ranker {
		if l.Score == alice {
			results = append(results, l.Rank)
			break
		}
	}
}

func climbingLeaderboard(scores []int, alice []int) {
	for _, b := range alice {
		scores = append(scores, b)
		sort.Slice(scores, func(i, j int) bool {
			return scores[i] > scores[j]
		})
		ranker(scores, b, []Ranker{})
	}
}

func main(){
	climbingLeaderboard([]int{100, 100, 50, 40, 40, 20, 10}, []int{5, 25, 50, 120})
	for _, v := range results{
		fmt.Println(v)
	}
}
