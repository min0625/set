package set

import "golang.org/x/exp/maps"

type None = struct{}

type Set[Elem comparable] map[Elem]None

func (s Set[Elem]) Store(elems ...Elem) {
	for _, elem := range elems {
		s[elem] = None{}
	}
}

func (s Set[Elem]) Delete(elems ...Elem) {
	for _, elem := range elems {
		delete(s, elem)
	}
}

func (s Set[Elem]) Range(f func(elem Elem) bool) {
	for elem := range s {
		if !f(elem) {
			return
		}
	}
}

func (s Set[Elem]) Has(elem Elem) (has bool) {
	_, has = s[elem]
	return has
}

func (s Set[Elem]) Clone() Set[Elem] {
	return maps.Clone(s)
}

func (s Set[Elem]) Equal(s2 Set[Elem]) (equal bool) {
	return maps.Equal(s, s2)
}

func (s Set[Elem]) Len() (length int) {
	return len(s)
}

func (s Set[Elem]) Slice() []Elem {
	return maps.Keys(s)
}

func (s Set[Elem]) Clear() {
	maps.Clear(s)
}

func (s Set[Elem]) Union(s2 Set[Elem]) Set[Elem] {
	s3 := make(Set[Elem], s.Len()+s2.Len())

	maps.Copy(s3, s)
	maps.Copy(s3, s2)

	return s3
}

func (s Set[Elem]) Difference(s2 Set[Elem]) Set[Elem] {
	s3 := make(Set[Elem], s.Len())

	s.Range(func(elem Elem) bool {
		if !s2.Has(elem) {
			s3.Store(elem)
		}

		return true
	})

	return s3
}

func (s Set[Elem]) Intersection(s2 Set[Elem]) Set[Elem] {
	n := s.Len()
	if n2 := s2.Len(); n > n2 {
		n = n2
	}

	s3 := make(Set[Elem], n)
	s.Range(func(elem Elem) bool {
		if s2.Has(elem) {
			s3.Store(elem)
		}

		return true
	})

	return s3
}

func (s Set[Elem]) IsSubset(s2 Set[Elem]) (isSubset bool) {
	isSubset = true

	s2.Range(func(elem Elem) bool {
		isSubset = s.Has(elem)
		return isSubset
	})

	return isSubset
}
