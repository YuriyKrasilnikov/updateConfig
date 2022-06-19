package utils

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v3"
)

func PrintTrees(nodes []*yaml.Node) {
	PrintTree(
		yaml.Node{
			Kind:    yaml.MappingNode,
			Tag:     "!!map",
			Content: nodes,
		},
	)
}

// напечатать дерево
func PrintTree(node yaml.Node) {
	out, err := yaml.Marshal(&node)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

func PrintTreeNodes(nodes []*yaml.Node) {
	for _, node := range nodes {
		printTreeNode(*node, 0)
	}
}

func PrintTreeNode(node yaml.Node) {
	printTreeNode(node, 0)
}

func printTreeNode(node yaml.Node, i int) {
	fmt.Printf("%d> %+v\n", i, node)
	for _, n := range node.Content {
		printTreeNode(*n, i+1)
	}
}
