package test

const arch = "result"

// arrayTestEqualScalarNode: node, tmp, expected result
var arrayTestEqualScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{"case1", `name: rustam`, `name: `, `name: rustam`},
	{"case2", ` name:     rustam`, `name: `, `name: rustam`},
	{"case3", `name: rustam`, `name: sergey`, `name: sergey`},
}

// arrayTestNotEqualScalarNode: node, tmp, result
var arrayTestNotEqualScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{"case1", `name: rustam`, `name: `, `name: sergey`},
	{"case2", `name: rustam`, `name: `, `name: `},
	{"case3", `name: rustam`, `name: sergey`, `name: rustam`},
}

// arrayTestErrorKindScalarNode: node, tmp
var arrayTestErrorKindScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{"case1",
		`
name:
- rustam
`,
		`
name: 
`,
	},
	{"case2",
		`
rustam:
 job: Developer
`,
		`
name: 
`,
	},
	{"case3",
		`
name:
- rustam
`,
		`
name: 
- rustam
`,
	},
	{"case4",
		`
rustam:
	job: Developer
`,
		`
rustam:
	job: Developer
`,
	},
	{"case5",
		`
name: rustam
`,
		`
name: 
- rustam
`,
	},
	{"case6",
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
var arrayTestErrorLenScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{"case1",
		`
name: rustam
job: developer
`,
		`
name: 
`,
	},
	{"case2",
		`
name: rustam
`,
		`
name: 
job: 
`,
	},
}
