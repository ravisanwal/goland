package urlcounts

type Node struct {
	count int
	data []map[string]Node
}