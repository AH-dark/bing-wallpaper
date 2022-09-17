package util

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestRandInt(t *testing.T) {
	asserts := assert.New(t)

	nums := make([]int, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, RandInt(1, 100))
	}

	sort.Ints(nums)
	asserts.Equal(100, len(nums))
	for _, num := range nums {
		asserts.True(num >= 1 && num <= 100)
	}
}

func TestRandInt64(t *testing.T) {
	asserts := assert.New(t)

	nums := make([]int64, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, RandInt64(1, 100))
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	asserts.Equal(100, len(nums))
	for _, num := range nums {
		asserts.True(num >= 1 && num <= 100)
	}
}

func TestRandFloat64(t *testing.T) {
	asserts := assert.New(t)

	nums := make([]float64, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, RandFloat64(1, 100))
	}

	sort.Float64s(nums)
	asserts.Equal(100, len(nums))
	for _, num := range nums {
		asserts.True(num >= 1 && num <= 100)
	}
}

func TestRandString(t *testing.T) {
	asserts := assert.New(t)

	strs := make([]string, 0)
	for i := 0; i < 100; i++ {
		strs = append(strs, RandString(10))
	}

	asserts.Equal(100, len(strs))
	for _, str := range strs {
		asserts.Equal(10, len(str))
	}
}

func TestRandStringWithCharset(t *testing.T) {
	asserts := assert.New(t)

	strs := make([]string, 0)
	for i := 0; i < 100; i++ {
		strs = append(strs, RandStringWithCharset(10, "0123456789"))
	}

	asserts.Equal(100, len(strs))
	for _, str := range strs {
		asserts.Equal(10, len(str))
		for _, c := range str {
			asserts.True(c >= '0' && c <= '9')
		}
	}
}
