package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"updateConfig/lib/scalarnode"
)

// позитивный тест возвращаемого значения
func TestPositiveScalarNode(t *testing.T) {

	for _, test := range arrayTestPositiveScalarNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)
			expectedNodes := stringToNodesArray(test.Expected)

			//надо обнулять, а то проверяем еще и line и collumn
			ignoreLineColumn(expectedNodes)

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
		})
	}
}

// негативный тест возвращаемого значения
func TestNegativeScalarNode(t *testing.T) {

	for _, test := range arrayTestNegativeScalarNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)
			expectedNodes := stringToNodesArray(test.Expected)

			//надо обнулять, а то проверяем еще и line и collumn
			ignoreLineColumn(expectedNodes)

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
		})
	}
}

// тест на ошибку при неправильном типе
func TestErrorKindScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorKindScalarNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)

			_, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)

			assert.Error(
				t,
				err,
			)
		})
	}
}

// тест на ошибку при неправильном количестве в массиве
func TestErrorLenScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorLenScalarNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)

			_, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)

			assert.Error(
				t,
				err,
			)
		})
	}
}

// тест на ошибку при неправильных именах в массиве
func TestErrorNameScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorNameScalarNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)

			_, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)

			assert.Error(
				t,
				err,
			)
		})
	}
}
