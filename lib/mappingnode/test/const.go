package test

import (
	"gopkg.in/yaml.v3"
)

var arch = struct {
	Sample   string
	Temp     string
	Expected string
}{
	"Sample",
	"Temp",
	"Result",
}

var constData = map[string]*yaml.Node{
	"sample_scalar1_name": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "job",
		Anchor: arch.Sample,
	},
	"sample_scalar1_body": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "Developer",
		Anchor: arch.Sample,
	},
	"sample_scalar2_name": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "skill",
		Anchor: arch.Sample,
	},
	"sample_scalar2_body": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "Golang",
		Anchor: arch.Sample,
	},
	"sample_scalar_empty_body": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!null",
		Anchor: arch.Sample,
	},
	"sample_sequence_name": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "skills",
		Anchor: arch.Sample,
	},
	"sample_sequence_body": &yaml.Node{
		Kind:   yaml.SequenceNode,
		Tag:    "!!seq",
		Anchor: arch.Sample,
		Content: []*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "golang",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "c",
				Anchor: arch.Sample,
			},
		},
	},
	//temp
	"temp_scalar1_name": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "job",
		Anchor: arch.Temp,
	},
	"temp_scalar1_body": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "Developer",
		Anchor: arch.Temp,
	},
	"temp_scalar2_name": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "skill",
		Anchor: arch.Temp,
	},
	"temp_scalar2_body": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "C",
		Anchor: arch.Temp,
	},
	"temp_scalar_empty_body": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!null",
		Anchor: arch.Temp,
	},
	"temp_map_name": &yaml.Node{
		Kind:   yaml.ScalarNode,
		Tag:    "!!str",
		Value:  "skills",
		Anchor: arch.Temp,
	},
	"temp_map_body": &yaml.Node{
		Kind:   yaml.MappingNode,
		Tag:    "!!map",
		Anchor: arch.Temp,
		Content: []*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "golang",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "good",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "c",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!null",
				Anchor: arch.Temp,
			},
		},
	},
}

var arrayTestPositiveMappingNode = [...]struct {
	CaseName string
	Sample   []*yaml.Node
	Temp     []*yaml.Node
	Expected []*yaml.Node
}{
	//case1
	/*
			--Sample--
			&Sample rustam: &Sample
		 		&Sample job: &Sample Developer
			--Temp--
			&Temp rustam: &Temp
		 		&Temp job: &Temp Developer
			--Result--
			&Result rustam: &Result
				&Temp job: &Temp Developer
				&Sample job: &Sample Developer
	*/
	{
		"case1",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Expected,
			},
		},
	},
	//case2
	/*
			--Sample--
			&Sample rustam: &Sample
		 		&Sample job: &Sample Developer
				&Sample skill: &Sample Golang
			--Temp--
			&Temp rustam: &Temp
		 		&Temp job: &Temp Developer
			--Result--
			&Result rustam: &Result
				&Temp job: &Temp Developer
				&Sample job: &Sample Developer
				&Sample skill: &Sample Golang
	*/
	{
		"case2",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
					constData["sample_scalar2_name"],
					constData["sample_scalar2_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
					constData["sample_scalar2_name"],
					constData["sample_scalar2_body"],
				},
				Anchor: arch.Expected,
			},
		},
	},
	//case3
	/*
			--Sample--
			&Sample rustam: &Sample
		 		&Sample job: &Sample Developer
			--Temp--
			&Temp rustam: &Temp
		 		&Temp job: &Temp Developer
		 		&Temp skill: &Temp
			--Result--
			&Result rustam: &Result
				&Temp job: &Temp Developer
				&Temp skill: &Temp
				&Sample job: &Sample Developer
	*/
	{
		"case3",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["temp_scalar2_name"],
					constData["temp_scalar_empty_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["temp_scalar2_name"],
					constData["temp_scalar_empty_body"],
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Expected,
			},
		},
	},
	//case4
	/*
			--Sample--
			&Sample rustam: &Sample
		 		&Sample job: &Sample Developer
			--Temp--
			&Temp rustam: &Temp
		 		&Temp job: &Temp Developer
		 		&Temp skill: &Temp C
			--Result--
			&Result rustam: &Result
				&Temp job: &Temp Developer
				&Temp skill: &Temp C
				&Sample job: &Sample Developer
	*/
	{
		"case4",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["temp_scalar2_name"],
					constData["temp_scalar2_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["temp_scalar2_name"],
					constData["temp_scalar2_body"],
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Expected,
			},
		},
	},
	//case5
	/*
			--Sample--
			&Sample rustam: &Sample
		 		&Sample job: &Sample Developer
				&Sample skill: &Sample Golang
			--Temp--
			&Temp rustam: &Temp
		 		&Temp job: &Temp Developer
		 		&Temp skill: &Temp
			--Result--
			&Result rustam: &Result
				&Temp job: &Temp Developer
				&Temp skill: &Temp
				&Sample job: &Sample Developer
				&Sample skill: &Sample Golang
	*/
	{
		"case5",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
					constData["sample_scalar2_name"],
					constData["sample_scalar2_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["temp_scalar2_name"],
					constData["temp_scalar_empty_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
					constData["temp_scalar2_name"],
					constData["temp_scalar_empty_body"],
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
					constData["sample_scalar2_name"],
					constData["sample_scalar2_body"],
				},
				Anchor: arch.Expected,
			},
		},
	},
	//case6
	/*
		--Sample--
		&Sample rustam: &Sample
			&Sample job: &Sample Developer
			&Sample skills: &Sample
				- &Sample golang
				- &Sample c
		--Sample--
		&Temp rustam: &Temp
			&Temp skills: &Temp
				&Temp golang: &Temp good
				&Temp c: &Temp
		--Result--
		&Result rustam: &Result
			&Temp skills: &Temp
				&Temp golang: &Temp good
				&Temp c: &Temp
			&Sample job: &Sample Developer
			&Sample skills: &Sample
				- &Sample golang
				- &Sample c
	*/
	{
		"case6",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
					constData["sample_sequence_name"],
					constData["sample_sequence_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_map_name"],
					constData["temp_map_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_map_name"],
					constData["temp_map_body"],
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
					constData["sample_sequence_name"],
					constData["sample_sequence_body"],
				},
				Anchor: arch.Expected,
			},
		},
	},
}

var arrayTestNegativeMappingNode = [...]struct {
	CaseName string
	Sample   []*yaml.Node
	Temp     []*yaml.Node
	Expected []*yaml.Node
}{
	//case1
	/*
		--Sample--
		&Sample rustam: &Sample
			&Sample job: &Sample Developer`
		--Temp--
		&Temp rustam: &Temp
			&Temp job: &Temp Developer`
		--Result--
		&Result rustam: &Result
			&Result job: &Result Developer
			&Result job: &Result Developer
	*/
	{
		"case1",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					&yaml.Node{
						Kind:   yaml.ScalarNode,
						Tag:    "!!str",
						Value:  "job",
						Anchor: arch.Expected,
					},
					&yaml.Node{
						Kind:   yaml.ScalarNode,
						Tag:    "!!str",
						Value:  "Developer",
						Anchor: arch.Expected,
					},
					&yaml.Node{
						Kind:   yaml.ScalarNode,
						Tag:    "!!str",
						Value:  "job",
						Anchor: arch.Expected,
					},
					&yaml.Node{
						Kind:   yaml.ScalarNode,
						Tag:    "!!str",
						Value:  "Developer",
						Anchor: arch.Expected,
					},
				},
				Anchor: arch.Expected,
			},
		},
	},
	//case2
	/*
		--Sample--
		&Sample rustam: &Sample
			&Sample job: &Sample Developer`
		--Temp--
		&Temp rustam: &Temp
			&Temp job: &Temp Developer`
		--Result--
		&Result rustam: &Result
	*/
	{
		"case2",
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Sample,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["sample_scalar1_name"],
					constData["sample_scalar1_body"],
				},
				Anchor: arch.Sample,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Temp,
			},
			&yaml.Node{
				Kind: yaml.MappingNode,
				Tag:  "!!map",
				Content: []*yaml.Node{
					constData["temp_scalar1_name"],
					constData["temp_scalar1_body"],
				},
				Anchor: arch.Temp,
			},
		},
		[]*yaml.Node{
			&yaml.Node{
				Kind:   yaml.ScalarNode,
				Tag:    "!!str",
				Value:  "rustam",
				Anchor: arch.Expected,
			},
			&yaml.Node{
				Kind:    yaml.MappingNode,
				Tag:     "!!map",
				Content: []*yaml.Node{},
				Anchor:  arch.Expected,
			},
		},
	},
}

var arrayTestErrorKindMappingNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{"case1",
		`&Sample name: &Sample
 - &Sample rustam`,
		`&Temp name: &Temp `,
	},
	{"case2",
		`&Sample rustam: &Sample
 &Sample job: &Sample Developer`,
		`&Temp name: &Temp`,
	}, {
		"case3",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp`,
	}, {
		"case4",
		`&Sample name: &Sample rustam`,
		`&Temp name: &Temp
 - &Temp rustam`,
	}, {
		"case6",
		`&Sample name: &Sample rustam`,
		`&Temp rustam: &Temp
 &Temp job: &Temp Developer`,
	},
}

var arrayTestErrorLenMappingNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{
		"case1: Sample has more then 1 map and temp has 1 map",
		`&Sample rustam: &Sample
 &Sample job: &Sample Developer
&Sample sergey: &Sample
 &Sample job: &Sample Developer`,
		`&Temp rustam: &Temp
 &Temp job: &Temp Developer`,
	},
	{"case2: Sample has 1 map and temp has more then 1 map",
		`&Sample rustam: &Sample
 &Sample job: &Sample Developer`,
		`&Temp rustam: &Temp
 &Temp job: &Temp Developer
&Temp sergey: &Temp
 &Temp job: &Temp Developer`,
	},
	{"case3: Sample has more then 1 map and temp has more then 1 map",
		`&Sample rustam: &Sample
 &Sample job: &Sample Developer
&Sample sergey: &Sample
 &Sample job: &Sample Developer`,
		`&Temp rustam: &Temp
 &Temp job: &Temp Developer
&Temp sergey: &Temp
 &Temp job: &Temp Developer`,
	},
}

var arrayTestErrorNameMappingNode = [...]struct {
	CaseName string
	Sample   string
	Temp     string
}{
	{
		"case1: Diff name for sample map and temp map",
		`&Sample job: &Sample
 &Sample rustam: &Sample Developer`,
		`&Temp rustam: &Temp
 &Temp job: &Temp Developer`,
	},
}
