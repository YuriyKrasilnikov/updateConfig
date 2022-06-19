package main_test

import (
	"testing"
	"updateConfig/lib/scalarnode/test"
)

func TestAll(t *testing.T) {
	t.Run("Test for ScalarNode", TestScalarNode)
}

func TestScalarNode(t *testing.T) {
	t.Run("Test for Error Name ScalarNode", test.TestErrorLenScalarNode)
	t.Run("Test for Error Len ScalarNode", test.TestErrorLenScalarNode)
	t.Run("Test for Error Kind ScalarNode", test.TestErrorKindScalarNode)
	t.Run("Test Equals of ScalarNode", test.TestPositiveScalarNode)
	t.Run("Test Not Equal of ScalarNode", test.TestNegativeScalarNode)
}
