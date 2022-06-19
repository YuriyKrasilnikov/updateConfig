package main_test

import (
	"testing"
	mappingnode "updateConfig/lib/mappingnode/test"
	scalarnode "updateConfig/lib/scalarnode/test"
)

func TestScalarNode(t *testing.T) {
	t.Run("Test for Error Name ScalarNode", scalarnode.TestErrorLenScalarNode)
	t.Run("Test for Error Len ScalarNode", scalarnode.TestErrorLenScalarNode)
	t.Run("Test for Error Kind ScalarNode", scalarnode.TestErrorKindScalarNode)
	t.Run("Test Equals of ScalarNode", scalarnode.TestPositiveScalarNode)
	t.Run("Test Not Equal of ScalarNode", scalarnode.TestNegativeScalarNode)
}

func TestMappingNode(t *testing.T) {
	t.Run("Test for Error Name MappingNode", mappingnode.TestErrorLenMappingNode)
	t.Run("Test for Error Len MappingNode", mappingnode.TestErrorLenMappingNode)
	t.Run("Test for Error Kind MappingNode", mappingnode.TestErrorKindMappingNode)
	t.Run("Test Equals of MappingNode", mappingnode.TestPositiveMappingNode)
	t.Run("Test Not Equal of MappingNode", mappingnode.TestNegativeMappingNode)
}
