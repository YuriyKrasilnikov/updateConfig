package main_test

import (
	"testing"
	"updateConfig/lib/scalarnode/test"
)

func TestAll(t *testing.T) {
	t.Run("Test for Error Len ScalarNode", test.TestErrorLenScalarNode)
	t.Run("Test for Error Kind ScalarNode", test.TestErrorKindScalarNode)
	t.Run("Test Equals of ScalarNode", test.TestEqualScalarNode)
	t.Run("Test Not Equal of ScalarNode", test.TestNotEqualScalarNode)

}
