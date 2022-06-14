package test

const arch = "result"

// arrayTestEqualScalarNode: node, tmp, expected result
var arrayTestEqualScalarNode = [...][3]string{
	[3]string{`name: rustam`, `name: `, `name: rustam`},
	[3]string{` name:     rustam`, `name: `, `name: rustam`},
	[3]string{`name: rustam`, `name: sergey`, `name: sergey`},
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
