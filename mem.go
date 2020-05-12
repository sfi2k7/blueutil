package blueutil

type Mem struct {
	m      map[string]string
	Parent *Mem
}

func (s *Mem) Add(k, v string) {
	s.m[k] = v
}

func (s *Mem) Del(k string) {
	delete(s.m, k)
}

func (s *Mem) Get(k string, remove ...bool) string {
	v, ok := s.m[k]
	if !ok {
		if s.Parent == nil {
			return ""
		}
		return s.Parent.Get(k)
	}

	if len(remove) > 0 {
		if remove[0] == true {
			delete(s.m, k)
		}
	}

	return v
}

func (s *Mem) Close() {
	s.m = nil
	s.Parent = nil
}

func NewMem() *Mem {
	return &Mem{m: make(map[string]string)}
}

func NewMemWithParent(s *Mem) *Mem {
	ns := NewMem()
	ns.Parent = s
	return ns
}
