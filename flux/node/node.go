package node

type Node interface {
	Put(kid Node) error
	Compile() (string, error)
}
