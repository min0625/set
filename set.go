package set

import "golang.org/x/exp/maps"

type None = struct{}

type Set[Element comparable] map[Element]None

func (s Set[Element]) Store(elements ...Element) {
	for _, element := range elements {
		s[element] = None{}
	}
}

func (s Set[Element]) Delete(elements ...Element) {
	for _, element := range elements {
		delete(s, element)
	}
}

func (s Set[Element]) Range(f func(element Element) bool) {
	for element := range s {
		if !f(element) {
			return
		}
	}
}

func (s Set[Element]) Has(element Element) (has bool) {
	_, has = s[element]
	return has
}

func (s Set[Element]) Clone() Set[Element] {
	return maps.Clone(s)
}

func (s Set[Element]) Equal(s2 Set[Element]) (equal bool) {
	return maps.Equal(s, s2)
}

func (s Set[Element]) Len() (length int) {
	return len(s)
}

func (s Set[Element]) Slice() []Element {
	return maps.Keys(s)
}

func (s Set[Element]) Clear() {
	maps.Clear(s)
}

func (s Set[Element]) Union(s2 Set[Element]) Set[Element] {
	s3 := make(Set[Element], s.Len()+s2.Len())

	maps.Copy(s3, s)
	maps.Copy(s3, s2)

	return s3
}

func (s Set[Element]) Difference(s2 Set[Element]) Set[Element] {
	s3 := make(Set[Element], s.Len())

	s.Range(func(element Element) bool {
		if !s2.Has(element) {
			s3.Store(element)
		}

		return true
	})

	return s3
}

func (s Set[Element]) Intersection(s2 Set[Element]) Set[Element] {
	length := s.Len()

	if length2 := s2.Len(); length > length2 {
		length = length2
	}

	s3 := make(Set[Element], length)

	s.Range(func(element Element) bool {
		if s2.Has(element) {
			s3.Store(element)
		}

		return true
	})

	return s3
}

func (s Set[Element]) IsSubset(s2 Set[Element]) (isSubset bool) {
	isSubset = true

	s2.Range(func(element Element) bool {
		isSubset = s.Has(element)
		return isSubset
	})

	return isSubset
}
