package main

import "fmt"

type Bread struct {
	val string
}

type StrawberryJam struct {
	opened bool
}

type SpoonOfStrawberryJam struct {
}

type Sandwitch struct {
	val string
}

func GetBreads(num int) []*Bread {
	breads := make([]*Bread, 2)
	for i := 0; i < num; i++ {
		breads[i] = &Bread{val: "bread"}
	}
	return breads
}

func OpenStrawberryJam(jam *StrawberryJam) {
	jam.opened = true
}

func GetOneSpoon(jam *StrawberryJam) *SpoonOfStrawberryJam {
	return &SpoonOfStrawberryJam{}
}

func PutJamOnBread(bread *Bread, jam *SpoonOfStrawberryJam) {
	bread.val += " + Strawberry Jam"
}

func MakeSandwitch(breads []*Bread) *Sandwitch {
	sandwitch := &Sandwitch{}
	for i := 0; i < len(breads); i++ {
		sandwitch.val += breads[0].val + " + "
	}
	return sandwitch
}

func main() {
	// 1. 빵 두개를 꺼낸다.
	breads := GetBreads(2)

	jam := &StrawberryJam{}
	// 2. 딸기잼 뚜껑을 연다.
	OpenStrawberryJam(jam)

	// 3. 딸기잼을 한스푼 뜬다.
	spoon := GetOneSpoon(jam)

	// 4. 딸기잼을 빵에 바른다.
	PutJamOnBread(breads[0], spoon)

	// 5. 빵을 덮는다.
	sandwitch := MakeSandwitch(breads)

	// 6. 완성
	fmt.Println(sandwitch.val)
}
