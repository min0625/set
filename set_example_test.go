package set_test

import (
	"fmt"

	"github.com/min0625/set"
)

func ExampleSet() {
	s := make(set.Set[int])

	s.Store(1, 2, 3)
	fmt.Println(s.Has(1))

	s.Delete(1)
	fmt.Println(s.Has(1))

	// Output:
	// true
	// false
}

func ExampleSet_string() {
	s := make(set.Set[string])

	s.Store("a")
	s.Store("b", "c")

	s.Range(func(elem string) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// a
	// b
	// c
}

func ExampleSet_Delete() {
	s := make(set.Set[int])

	s.Store(1, 2, 3)

	s.Delete(2)

	s.Range(func(elem int) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// 1
	// 3
}

func ExampleSet_Range() {
	s := make(set.Set[int])

	s.Store(1, 2, 3)

	s.Range(func(elem int) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleSet_Range_break() {
	s := make(set.Set[string])

	s.Store("a", "b", "c")

	var cnt int
	s.Range(func(elem string) bool {
		cnt++
		return false
	})

	fmt.Println(cnt)

	// Output:
	// 1
}

func ExampleSet_Has() {
	s := make(set.Set[int])

	s.Store(1, 2)

	fmt.Println(s.Has(1))
	fmt.Println(s.Has(2))
	fmt.Println(s.Has(3))
	fmt.Println(s.Has(4))

	// Unordered output:
	// true
	// true
	// false
	// false
}

func ExampleSet_Clone() {
	s1 := make(set.Set[int])
	s1.Store(1, 2, 3)

	s2 := s1.Clone()

	s1.Store(4, 5)

	s2.Range(func(elem int) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleSet_Equal() {
	s1 := make(set.Set[int])
	s1.Store(1, 2, 3)

	s2 := make(set.Set[int])
	s2.Store(1, 2, 3)

	s3 := make(set.Set[int])
	s3.Store(1, 2)

	fmt.Println(s1.Equal(s2))
	fmt.Println(s1.Equal(s3))

	// Output:
	// true
	// false
}

func ExampleSet_Len() {
	s := make(set.Set[int])

	fmt.Println(s.Len())
	s.Store(1, 2)
	fmt.Println(s.Len())
	s.Delete(2)
	fmt.Println(s.Len())

	// Output:
	// 0
	// 2
	// 1
}

func ExampleSet_Slice() {
	s := make(set.Set[int])
	s.Store(1, 2, 3)

	for _, e := range s.Slice() {
		fmt.Println(e)
	}

	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleSet_Clear() {
	s := make(set.Set[int])
	s.Store(1, 2, 3)

	s.Clear()

	fmt.Println(s.Len())

	// Output:
	// 0
}

func ExampleSet_Union() {
	s1 := make(set.Set[int])
	s1.Store(1, 2)

	s2 := make(set.Set[int])
	s2.Store(2, 3)

	s3 := s1.Union(s2)

	s3.Range(func(elem int) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// 1
	// 2
	// 3
}

func ExampleSet_Difference() {
	s1 := make(set.Set[int])
	s1.Store(1, 2)

	s2 := make(set.Set[int])
	s2.Store(2, 3)

	s3 := s1.Difference(s2)

	s3.Range(func(elem int) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// 1
}

func ExampleSet_Intersection() {
	s1 := make(set.Set[int])
	s1.Store(1, 2)

	s2 := make(set.Set[int])
	s2.Store(2, 3)

	s3 := s1.Intersection(s2)

	s3.Range(func(elem int) bool {
		fmt.Println(elem)
		return true
	})

	// Unordered output:
	// 2
}

func ExampleSet_IsSubset() {
	s1 := make(set.Set[int])
	s1.Store(1, 2)

	s2 := make(set.Set[int])
	s2.Store(1)

	s3 := make(set.Set[int])
	s3.Store(3)

	s4 := make(set.Set[int])
	s4.Store(1, 2, 3)

	fmt.Println(s1.IsSubset(s1))
	fmt.Println(s1.IsSubset(s2))
	fmt.Println(s1.IsSubset(s3))
	fmt.Println(s1.IsSubset(s4))

	// Output:
	// true
	// true
	// false
	// false
}
