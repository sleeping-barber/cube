package node

type Node struct {
	Cores           int
	Disk            int
	DiskAllocated   int
	Ip              string
	Memory          int
	MemoryAllocated int
	Name            string
	Role            string
	TasksCount      int
}
