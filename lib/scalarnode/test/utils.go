package test

import (
	"log"

	"gopkg.in/yaml.v3"
)

func stringToNodesArray(str string) []*yaml.Node {
	node := new(yaml.Node)

	err := yaml.Unmarshal([]byte(str), node)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return node.Content[0].Content
}

func ignoreLineColumn(nodes []*yaml.Node) {
	for _, node := range nodes {
		node.Column = 0
		node.Line = 0
	}
}

func addAnchor(nodes []*yaml.Node, arch string) {
	for _, node := range nodes {
		node.Anchor = arch
	}
}
