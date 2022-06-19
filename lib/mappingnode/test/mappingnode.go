package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"updateConfig/lib/mappingnode"
)

// позитивный тест возвращаемого значения
func TestPositiveMappingNode(t *testing.T) {

	for _, test := range arrayTestPositiveMappingNode {
		t.Run(test.CaseName, func(t *testing.T) {

			resultNodes, err := mappingnode.ResultMappingNode(test.Sample, test.Temp, arch.Expected)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			//printTestCase(test)

			assert.Equal(
				t,
				resultNodes,
				test.Expected,
			)
		})
	}
}

// негативный тест возвращаемого значения
func TestNegativeMappingNode(t *testing.T) {

	for _, test := range arrayTestNegativeMappingNode {
		t.Run(test.CaseName, func(t *testing.T) {

			resultNodes, err := mappingnode.ResultMappingNode(test.Sample, test.Temp, arch.Expected)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			//printTestCase(test)

			assert.NotEqual(
				t,
				resultNodes,
				test.Expected,
			)
		})
	}
}

// тест на ошибку при неправильном типе
func TestErrorKindMappingNode(t *testing.T) {

	for _, test := range arrayTestErrorKindMappingNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)

			_, err := mappingnode.ResultMappingNode(sampleNodes, tmpNodes, arch.Expected)

			assert.Error(
				t,
				err,
			)
		})
	}
}

// тест на ошибку при неправильном количестве в массиве
func TestErrorLenMappingNode(t *testing.T) {

	for _, test := range arrayTestErrorLenMappingNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)

			_, err := mappingnode.ResultMappingNode(sampleNodes, tmpNodes, arch.Expected)

			assert.Error(
				t,
				err,
			)
		})
	}
}

// тест на ошибку при неправильных именах в массиве
func TestErrorNameMappingNode(t *testing.T) {

	for _, test := range arrayTestErrorNameMappingNode {
		t.Run(test.CaseName, func(t *testing.T) {
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)

			_, err := mappingnode.ResultMappingNode(sampleNodes, tmpNodes, arch.Expected)

			assert.Error(
				t,
				err,
			)
		})
	}
}
