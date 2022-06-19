package test

const arch = "Result"

var arrayTestPositiveScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{"case1: Sample full and temp empty",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp`,
		`&Result name: &Result rustam`,
	},
	{"case2: Sample empty and temp full",
		`&Sample name: &Sample`,
		`&Temp name: &Temp rustam`,
		`&Result name: &Result rustam`,
	},
	{"case3: Sample full and temp full",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp sergey`,
		`&Result name: &Result sergey`,
	},
	{"case3: Sample empty and temp empty",
		`&Sample name: &Sample`,
		`&Temp name: &Temp`,
		`&Result name: &Result`,
	},
}

var arrayTestNegativeScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
	Expected string
}{
	{
		"case1: Sample full and temp full, but result incorrect",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp sergey`,
		`&Result name: &Result rustam`,
	},
}

var arrayTestErrorKindScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{
		"case1: Sample is sequence and temp is scalar",
		`
&Sample name: &Sample
 - &Sample rustam
`,
		`
&Temp name: &Temp 
`,
	}, {
		"case2: Sample is map and temp is scalar",
		`
&Sample rustam: &Sample
 &Sample job: &Sample Developer
`,
		`
&Temp name: &Temp
`,
	}, {
		"case3: Sample is sequence and temp is sequence",
		`
&Sample name: &Sample
 - &Sample rustam
`,
		`
&Temp name: &Temp
 - &Temp rustam
`,
	}, {
		"case4: Sample is map and temp is map",
		`
&Sample rustam: &Sample
 &Sample job: &Sample Developer
`,
		`
&Temp rustam: &Temp
 &Temp job: &Temp Developer
`,
	}, {
		"case5: Sample is scalar and temp is full sequence",
		`
&Sample name: &Sample rustam
`,
		`
&Temp name: &Temp
 - &Temp rustam
`,
	}, {
		"case6: Sample is scalar and temp is full map",
		`
&Sample name: &Sample rustam
`,
		`
&Temp rustam: &Temp
 &Temp job: &Temp Developer
`,
	},
}

var arrayTestErrorLenScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{"case1: Sample has more then 1 scalar and temp has 1 scalar",
		`&Sample name: &Sample rustam
&Sample job: &Sample developer`,
		`&Temp name: &Temp`,
	},
	{"case2: Sample has 1 scalar and temp has more then 1 scalar",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp
&Temp job: &Temp`,
	},
	{"case3: Sample has more then 1 scalar and temp has more then 1 scalar",
		`&Sample name: &Sample rustam
&Sample job: &Sample developer`,
		`&Temp name: &Temp
&Temp job: &Temp`,
	},
}

var arrayTestErrorNameScalarNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{
		"case1: Diff name for sample full scalar and temp empty scalar",
		`&Sample name: &Sample rustam`,
		`&Temp rustam: &Temp`,
	}, {
		"case2: Diff name for sample empty scalar and temp full scalar",
		`&Sample name: &Sample`,
		`&Temp rustam: &Temp name`,
	}, {
		"case3: Diff name for sample full scalar and temp full scalar",
		`&Sample rustam: &Sample name`,
		`&Temp sergey: &Temp name`,
	},
}
