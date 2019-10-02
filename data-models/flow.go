package data_models

// lists all flows currently configured. Is useful in determining if a
// source event belongs to the head node of any existing flow
var activeFlows []flow

// flow represents how a flow can be accessed. Since it's represented as a tree
// all nodes can be accessed from the head node
type flow struct {
	head     node
	isActive bool
	id       string
}
