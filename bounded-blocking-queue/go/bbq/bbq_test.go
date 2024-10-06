package bbq

import (
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBBQ(t *testing.T) {
	t.Run("bbq of ints", func(t *testing.T) {
		q := NewBBQ[int](10)
		assert.NotNil(t, q)
	})

	t.Run("bbq of strings", func(t *testing.T) {
		q := NewBBQ[string](10)
		assert.NotNil(t, q)
	})
}

func TestLen(t *testing.T) {
	q := NewBBQ[int](10)

	for i := range 10 {
		assert.Equal(t, q.Len(), i)
		q.Push(i)
	}
}

func TestPushPop(t *testing.T) {
	q := NewBBQ[int](10)

	for i := range 10 {
		q.Push(i)
	}
	for i := range 10 {
		val := q.Pop()
		assert.Equal(t, val, i)
	}
}

func TestConcurrency(t *testing.T) {
	q := NewBBQ[int](10)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := range 5 {
			q.Push(i)
		}
		wg.Done()
	}()

	go func() {
		for i := range 5 {
			q.Push(i + 5)
		}
		wg.Done()
	}()

	wg.Wait()

	nums := make([]int, 0, 10)
	for range 10 {
		val := q.Pop()
		nums = append(nums, val)
	}
	sort.Ints(nums)
	for i := range 10 {
		assert.Equal(t, i, nums[i])
	}
}
