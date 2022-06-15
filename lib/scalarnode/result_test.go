package scalarnode

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

const arch = "result"

// arrayTestEqualScalarNode: node, tmp, expected result
var arrayTestEqualScalarNode = []struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{"case1", "name: rustam", "name: ", "name: rustam"},
	{"case2", " name:     rustam", "name: ", "name: rustam"},
	{"case3", "name: rustam", "name: sergey", "name: sergey"},
}

// arrayTestNotEqualScalarNode: node, tmp, result
var arrayTestNotEqualScalarNode = [...][3]string{
	[3]string{`name: rustam`, `name: `, `name: sergey`},
	[3]string{`name: rustam`, `name: `, `name: `},
	[3]string{`name: rustam`, `name: sergey`, `name: rustam`},
}

// arrayTestErrorKindScalarNode: node, tmp
var arrayTestErrorKindScalarNode = [...][2]string{
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

// arrayTestErrorLenScalarNode: node, tmp
var arrayTestErrorLenScalarNode = [...][2]string{
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

// TestEqualScalarNode позитивный тест возвращаемого значения
func TestEqualScalarNode(t *testing.T) {

	for _, test := range arrayTestEqualScalarNode {
		t.Run(test.CaseName, func(t *testing.T) {
			// arrange
			sampleNodes := stringToNodesArray(test.Sample)
			tmpNodes := stringToNodesArray(test.Temp)
			expectedNodes := stringToNodesArray(test.Expected)

			// надо обнулять, а то проверяем еще и line и collumn
			ignoreLineColumn(expectedNodes)
			// добавим проверку на правильный Anchor
			addAnchor(expectedNodes, arch)

			resultNodes, err := ResultScalarNode(sampleNodes, tmpNodes, arch)
			if err != nil {
				log.Fatalf("error: %v", err)
			}

			// надо обнулять, а то проверяем еще и line и collumn
			// act
			ignoreLineColumn(resultNodes)

			// assert
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

// TestNotEqualScalarNode негативный тест возвращаемого значения
func TestNotEqualScalarNode(t *testing.T) {

	for _, test := range arrayTestNotEqualScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])
		expectedNodes := stringToNodesArray(test[2])

		//надо обнулять, а то проверяем еще и line и collumn
		ignoreLineColumn(expectedNodes)
		//добавим проверку на правильный Anchor
		addAnchor(expectedNodes, arch)

		resultNodes, err := ResultScalarNode(sampleNodes, tmpNodes, arch)
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

// TestErrorKindScalarNode тест на ошибку при неправильном типе
func TestErrorKindScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorKindScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])

		_, err := ResultScalarNode(sampleNodes, tmpNodes, arch)

		assert.Error(
			t,
			err,
		)
	}
}

// TestErrorLenScalarNode тест на ошибку при неправильном количестве в массиве
func TestErrorLenScalarNode(t *testing.T) {

	for _, test := range arrayTestErrorLenScalarNode {
		sampleNodes := stringToNodesArray(test[0])
		tmpNodes := stringToNodesArray(test[1])

		_, err := ResultScalarNode(sampleNodes, tmpNodes, arch)

		assert.Error(
			t,
			err,
		)
	}
}
