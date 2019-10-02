package data_models

// node defines the tree structure for a node in this application
// when users configure flows, we store the configured data as a tree of nodes
type node struct {
	device device
	// events contains the events the user has configured for this device in the flow
	events []abstractEvent
	// children is the list of child nodes for every node in the tree
	// events would be executed first in the parent node, then in children
	children []node
}
