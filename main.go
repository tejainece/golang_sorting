package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Data struct {
	Expires time.Time
}

func main() {
	datas := make([]Data, 10)
	
	for i := range datas {
		datas[i].Expires = time.Now().Add(time.Hour * time.Duration(rand.Int31n(100)))
	}
	
	for _, d := range datas {
		fmt.Println(d)
	}
	
	sort.Slice(datas, func(i, j int) bool {
		return datas[i].Expires.After(datas[j].Expires)
	})

	fmt.Println("Sorted:")
	for _, d := range datas {
		fmt.Println(d)
	}
}
