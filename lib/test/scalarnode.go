package test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"updateConfig/lib/scalarnode"
)

// позитивный тест возвращаемого значения
func TestEqualScalarNode(t *testing.T) {
	// testArray: node, tmp, expected result
	testsArray := [...][3]string{
		[3]string{`name: rustam`, `name: `, `name: rustam`},
		[3]string{` name:     rustam`, `name: `, `name: rustam`},
		[3]string{`name: rustam`, `name: sergey`, `name: sergey`},
	}

	for _, test := range testsArray {
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
	// testArray: node, tmp, result
	testsArray := [...][3]string{
		[3]string{`name: rustam`, `name: `, `name: sergey`},
		[3]string{`name: rustam`, `name: `, `name: `},
		[3]string{`name: rustam`, `name: sergey`, `name: rustam`},
	}

	for _, test := range testsArray {
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
	// testArray: node, tmp
	testsArray := [...][2]string{
		[2]string{
			`
name:
- rustam
`,
			`
name: 
`,
		},
		[2]string{
			`
rustam:
 job: Developer
`,
			`
name: 
`,
		},
		[2]string{
			`
name:
- rustam
`,
			`
name: 
- rustam
`,
		},
		[2]string{
			`
rustam:
	job: Developer
`,
			`
rustam:
	job: Developer
`,
		},
		[2]string{
			`
name: rustam
`,
			`
name: 
- rustam
`,
		},
		[2]string{
			`
name: rustam
`,
			`
rustam:
	job: Developer
`,
		},
	}

	for _, test := range testsArray {
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
	// testArray: node, tmp
	testsArray := [...][2]string{
		[2]string{
			`
name: rustam
job: developer
`,
			`
name: 
`,
		},
		[2]string{
			`
name: rustam
`,
			`
name: 
job: 
`,
		},
	}

	for _, test := range testsArray {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])

		_, err := scalarnode.ResultScalarNode(sampleNodes, tmpNodes, arch)

		assert.Error(
			t,
			err,
		)
	}
}
