package dataStructure

type UFI interface {
	Unite(int, int)
	IsConnected(int, int) bool
	Find(int) int
}

type UF struct {
	Parent []int // 保存所有父节点 i的父节点为Parent[i]
	Size   []int // 保存父节点的子节点个数
}

func NewUF(n int) UFI {
	u := &UF{
		Parent: make([]int, n),
		Size:   make([]int, n),
	}
	for i := 0; i < n; i++ {
		u.Parent[i] = i // 初始化每个节点都是一个连通分量，其父节点为自己
		u.Size[i] = 1   // 初始化每个父节点都为自己，所以子节点个数都为1
	}
	return u
}

// Unite 连接x和y
func (u UF) Unite(x, y int) {
	rootX, rootY := u.Find(x), u.Find(y)
	if rootX == rootY {
		return
	} else if u.Size[rootX] < u.Size[rootY] {
		u.Parent[rootX] = rootY
		u.Size[rootY] += u.Size[rootX]
	} else {
		u.Parent[rootY] = rootX
		u.Size[rootX] += u.Size[rootY]
	}
}

// IsConnected 判断 x和y是否已连接
func (u UF) IsConnected(x, y int) bool {
	return u.Find(x) == u.Find(y)
}

// Find 返回x的根节点
func (u UF) Find(x int) int {
	for u.Parent[x] != x {
		u.Parent[x] = u.Parent[u.Parent[x]]
		x = u.Parent[x]
	}
	return x
}
