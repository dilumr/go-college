package bst

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type zipToCity struct {
	zipCode  int
	cityName string
}

func (first *zipToCity) LessThan(second interface{}) bool {
	return first.zipCode < second.(*zipToCity).zipCode
}

func TestZipCodeUseCase(t *testing.T) {
	tree := New() // note -- using bst.New() would cause problems.
	tree.Insert(&zipToCity{75205, "Dallas"})
	tree.Insert(&zipToCity{90210, "Beverly Hills"})
	tree.Insert(&zipToCity{78759, "Austin"})
	tree.Insert(&zipToCity{91945, "San Don't Go"})
	tree.Insert(&zipToCity{91945, "San Diego"})
	found, ok := tree.Find(&zipToCity{78759, ""})
	assert.True(t, ok)
	fzc := *(found.(*zipToCity))
	assert.Equal(t, 78759, fzc.zipCode)
	assert.Equal(t, "Austin", fzc.cityName)

	_, gtz, ltok, gtok := tree.FindNearest(&zipToCity{10658, ""})
	assert.False(t, ltok)
	assert.True(t, gtok)
	fzc = *(gtz.(*zipToCity))
	assert.Equal(t, 75205, fzc.zipCode)
}

type IntBody int

func (first *IntBody) LessThan(second interface{}) bool {
	return (*first) < (*second.(*IntBody))
}

func TestInsertion(t *testing.T) {
	tree := New()
	assert.Equal(t, 0, tree.nodeCount())

	data := uniqueRandomInts(100)
	for _, n := range data {
		v := IntBody(n)
		tree.Insert(&v)
	}
	assert.Equal(t, len(data), tree.nodeCount())

	expected := append(data[0:0], data...) // clone data
	sort.Sort(sort.IntSlice(expected))

	actual := make([]int, 0, len(data))
	appender := func(body interface{}) {
		value := body.(*IntBody)
		actual = append(actual, int(*value))
	}
	tree.VisitInOrder(appender)

	assert.Equal(t, expected, actual)
}

func uniqueRandomInts(count int) []int {
	upper := int32(count) * 4
	m := make(map[int]bool)
	for len(m) < count {
		m[int(rand.Int31n(upper))] = true
	}
	uniques := make([]int, 0, count)
	for n := range m {
		uniques = append(uniques, n)
	}
	return uniques
}
