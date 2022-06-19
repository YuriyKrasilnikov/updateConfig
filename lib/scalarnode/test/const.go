package test

const arch = "Result"

// arrayTestEqualScalarNode: node, tmp, expected result
var arrayTestEqualScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{
		"case1",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp`,
		`&Result name: &Result rustam`,
	}, {
		"case2",
		`&Sample name: &Sample`,
		`&Temp name: &Temp rustam`,
		`&Result name: &Result rustam`,
	}, {
		"case3",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp sergey`,
		`&Result name: &Result sergey`,
	},
}

// arrayTestNotEqualScalarNode: node, tmp, result
var arrayTestNotEqualScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{
		"case1",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp`,
		`&Result name: &Result sergey`,
	}, {
		"case2",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp`,
		`&Result name: &Result`,
	}, {
		"case3",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp sergey`,
		`&Result name: &Result rustam`,
	},
}

// arrayTestErrorKindScalarNode: node, tmp
var arrayTestErrorKindScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{
		"case1",
		`
&Sample name: &Sample
	- &Sample rustam
`,
		`
&Temp name: &Temp 
`,
	}, {
		"case2",
		`
&Sample rustam: &Sample
	&Sample job: &Sample Developer
`,
		`
&Temp name: &Temp
`,
	}, {
		"case3",
		`
&Sample name: &Sample
	- &Sample rustam
`,
		`
&Temp name: &Temp
	- &Temp rustam
`,
	}, {
		"case4",
		`
&Sample rustam: &Sample
	&Sample job: &Sample Developer
`,
		`
&Temp rustam: &Temp
	&Temp job: &Temp Developer
`,
	}, {
		"case5",
		`
&Sample name: &Sample rustam
`,
		`
&Temp name: &Temp
	- &Temp rustam
`,
	}, {
		"case6",
		`
&Sample name: &Sample rustam
`,
		`
&Temp rustam: &Temp
	&Temp job: &Temp Developer
`,
	},
}

// arrayTestErrorLenScalarNode: node, tmp
var arrayTestErrorLenScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{
		"case1",
		`
&Sample name: &Sample rustam
&Sample job: &Sample developer
`,
		`
&Temp name: &Temp
`,
	}, {
		"case2",
		`
&Sample name: &Sample rustam
`,
		`
&Temp name: &Temp
&Temp job: &Temp
`,
	},
}
