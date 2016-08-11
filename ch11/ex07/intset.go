package intset

// OptimizedIntSet

type OptimizedIntSet struct {
	words []uint64
}

func (s *OptimizedIntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *OptimizedIntSet) Add(x int) {
	word, bit := x>>6, uint(x&((1<<6)-1))
	if len(s.words) <= word {
		new := make([]uint64, 2*(word+1), 2*(word+1))
		copy(new, s.words)
		s.words = new
	}
	s.words[word] |= 1 << bit
}

func (s *OptimizedIntSet) UnionWith(t *OptimizedIntSet) {
	tlen := len(t.words)
	if len(s.words) <= tlen {
		new := make([]uint64, 2*(tlen+1), 2*(tlen+1))
		copy(new, s.words)
		s.words = new
	}

	for i, tword := range t.words {
		s.words[i] |= tword
	}
}

// IntSetByMap

type IntSetByMap struct {
	m map[int]struct{}
}

func NewIntSetByMap() *IntSetByMap {
	return &IntSetByMap{map[int]struct{}{}}
}

func (s *IntSetByMap) Has(x int) bool {
	_, ok := s.m[x]
	return ok
}

func (s *IntSetByMap) Add(x int) {
	s.m[x] = struct{}{}
}

func (s *IntSetByMap) UnionWith(t *IntSetByMap) {
	for x := range t.m {
		s.m[x] = struct{}{}
	}
}

// IntSetByBoolMap

type IntSetByBoolMap struct {
	m map[int]bool
}

func NewIntSetByBoolMap() *IntSetByBoolMap {
	return &IntSetByBoolMap{map[int]bool{}}
}

func (s *IntSetByBoolMap) Has(x int) bool {
	_, ok := s.m[x]
	return ok
}

func (s *IntSetByBoolMap) Add(x int) {
	s.m[x] = true
}

func (s *IntSetByBoolMap) UnionWith(t *IntSetByBoolMap) {
	for x := range t.m {
		s.m[x] = true
	}
}

// IntSetByUint64Slice

type IntSetByUint64Slice struct {
	words []uint64
}

func (s *IntSetByUint64Slice) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSetByUint64Slice) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSetByUint64Slice) UnionWith(t *IntSetByUint64Slice) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntSetByUint32Slice

type IntSetByUint32Slice struct {
	words []uint32
}

func (s *IntSetByUint32Slice) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSetByUint32Slice) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSetByUint32Slice) UnionWith(t *IntSetByUint32Slice) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
