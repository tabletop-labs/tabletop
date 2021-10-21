package apps

import (
	"github.com/ergo-services/ergo"
	"github.com/ergo-services/ergo/gen"
	"github.com/ergo-services/ergo/node"
)

var (
	apps = []gen.ApplicationBehavior{}
)

func Add(app gen.ApplicationBehavior) {
	apps = append(apps, app)
}

func StartNode() (node.Node, error) {
	opts := node.Options{
		Applications: apps,
	}
	node, err := ergo.StartNode("nodename@localhost", "cookie", opts)
	if err != nil {
		return node, err
	}

	return node, nil
}
