package unionfindset


type treeFindset struct {
	rank []int
	parents []int
}

func NewTreeFindset(n int) *treeFindset {
	t := &treeFindset{}
	t.parents = make([]int, n)
	t.rank = make([]int, n)
	for i := 0; i < n; i++ {
		t.parents[i] = i
		t.rank[i] = 1
	}

	return t
}


func (t *treeFindset) findParent(i int) int {
	if t.parents[i] != i {
		t.parents[i] = t.findParent(t.parents[i])
	}

	return t.parents[i]
}

func (t *treeFindset) Find(i, j int) bool {
	return t.findParent(i) == t.findParent(j)
}

func (t *treeFindset) Union(i, j int) {
	iParent, jParent := t.findParent(i), t.findParent(j)
	if t.rank[iParent] < t.rank[jParent] {
		t.parents[iParent] = jParent
	} else if t.rank[iParent] > t.rank[jParent] {
		t.parents[jParent] = iParent
	} else {
		t.parents[iParent] = jParent
		t.rank[iParent]++
	}
}