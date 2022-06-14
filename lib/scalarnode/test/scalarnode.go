package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"updateConfig/lib/scalarnode"
)

// позитивный тест возвращаемого значения
func TestEqualScalarNode(t *testing.T) {

	for _, test := range arrayTestEqualScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])
		expectedNodes := stringToNodesArray(test[2])

		//надо обнулять, а то проверяем еще и line и collumn
		ignoreLineColumn(expectedNodes)
		//добавим проверку на правильный Anchor
		addAnchor(expectedNodes, arch)

		resultNodes, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		//надо обнулять, а то проверяем еще и line и collumn
		ignoreLineColumn(resultNodes)

		assert.Equalf(
			t,
			resultNodes,
			expectedNodes,
			"Diff in \nInfo:\nsampleNodes %s,\ntmpNodes %s,\nexpectedNodes %s,\n resultNodes: %s",
			sampleNodes, tmpNodes, expectedNodes, resultNodes,
		)

	}
}

// негативный тест возвращаемого значения
func TestNotEqualScalarNode(t *testing.T) {

	for _, test := range arrayTestNotEqualScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])
		expectedNodes := stringToNodesArray(test[2])

		//надо обнулять, а то проверяем еще и line и collumn
		ignoreLineColumn(expectedNodes)
		//добавим проверку на правильный Anchor
		addAnchor(expectedNodes, arch)

		resultNodes, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)
		if err != nil {
			log.Fatalf("error: %v", err)
		}

		//надо обнулять, а то проверяем еще и line и collumn
		ignoreLineColumn(resultNodes)

		assert.NotEqualf(
			t,
			resultNodes,
			expectedNodes,
			"Diff in \nInfo:\nsampleNodes %s,\ntmpNodes %s,\nexpectedNodes %s,\n resultNodes: %s",
			sampleNodes, tmpNodes, expectedNodes, resultNodes,
		)
	}
}

// тест на ошибку при неправильном типе
func TestErrorKindScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorKindScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])

		_, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)

		assert.Error(
			t,
			err,
		)
	}
}

// тест на ошибку при неправильном количестве в массиве
func TestErrorLenScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorLenScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])

		_, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)

		assert.Error(
			t,
			err,
		)
	}
}
