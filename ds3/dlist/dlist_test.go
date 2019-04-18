package ds3

import (
	"container/list"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestDlistUseCase(t *testing.T) {
	dlist := list.New() // returns a *list.List
	dlist.PushBack("claire")
	dlist.PushBack("dylan")
	dlist.PushBack("evelyn")
	dlist.PushBack("francis")
	printElements("Pushing to back repeatedly", dlist)

	dlist.PushFront("ginger")
	printElements("After pushing to front once", dlist)

	dlist.Init()
	printElements("After calling Init()", dlist)

	for i := 1; i <= 50; i++ {
		dlist.PushBack(strconv.Itoa(i))
	}
	printElements("After adding 1-50", dlist)

	for e := dlist.Front(); e != nil; {
		n := e.Next()            // intentionally getting next before e is threatened
		if rand.Int31n(2) == 0 { // 50% of the time
			dlist.Remove(e)
		}
		e = n
	}
	printElements("After Thanos had his way", dlist)
}

func printElements(title string, lst *list.List) {
	fmt.Printf("--- %v (%d elements) ---\n", title, lst.Len())
	printer := func(v interface{}) {
		fmt.Println(v.(string))
	}
	fooElements(lst, printer)
	fmt.Println("")
}

func fooElements(lst *list.List, f func(interface{})) {
	for e := lst.Front(); e != nil; e = e.Next() {
		f(e.Value)
	}
}
