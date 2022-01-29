package minigli

import (
	"testing"
)

var testCEWords = []struct {
	input	string
	want	int
}{
	{"--option:value",8},
	{"-option:value",7},
	{"--flag:",6},
	{"--NoColon",-1},
	{"orJustArg",-1},
	{":::",1}, //Does not care position 0.
}
func Test_findColonOrEqual(t *testing.T) {
	for _, test := range testCEWords {
		actual := findColonOrEqual(test.input)
		if test.want != actual {
			t.Errorf("  %q => %d ...want:%d",test.input, actual, test.want)
		}
	}
}

var testArguments = []struct {
	// Test Input
	input	[]string
	// Test Want
	path	[]string
	longs	map[string]string
	shorts	map[string]string
}{
	{
		[]string{"command","some","path","--longOpt","optValue","-s=some","--flag:"},
		[]string{"command","some","path"},
		map[string]string{"longOpt":"optValue", "flag":""},
		map[string]string{"s":"some"},
	},
	{
		[]string{"readwith","--json:","-f=filename.csv","options","inner"},
		[]string{"readwith","options","inner"},
		map[string]string{"json":""},
		map[string]string{"f":"filename.csv"},
	},
	{
		[]string{"-tst:","--options:","-only"},
		[]string{},
		map[string]string{"options":""},
		map[string]string{"tst":"","only":""},
	},
}
func Test_parseInputsFrom(t *testing.T) {
	const fmtp =   "%s\ninput:%s\ngot  :%s\nwant :%s\n"
	const fmtmap = "%s\ninput:%s\ngot  :%s => %s\nwant :%s => %s\n"
	for _, test := range testArguments {

		path, longs, shorts:= parseInputFrom(test.input)

		// check Command-Path
		tag := "Command-Path"
		if len(test.path) == len(path) {
			i := 0
			for i <len(path) {
				if path[i] != test.path[i] {
					t.Errorf(fmtp, tag, test.input, path, test.path)
				}
				i++
			}
		} else {
			t.Errorf(fmtp, tag, test.input, path, test.path)
		}

		// check long options
		tag = "Long Options"
		if len(test.longs) == len(longs) {
			for key,want := range test.longs {
				value, keyExists := longs[key]
				if !(keyExists && (value == want)) {
					t.Errorf(fmtmap, tag, test.input, key,value,key,want)
				}
			}
		} else {
			t.Errorf(fmtp, tag + " Length", test.input, longs,test.longs)
		}

		// check short options
		tag = "Short Options"
		if len(test.shorts) == len(shorts) {
			for key,want := range test.shorts {
				value, keyExists := shorts[key]
				if !(keyExists && (value == want)) {
					t.Errorf(fmtmap, tag, test.input, key,value,key,want)
				}
			}
		} else {
			t.Errorf(fmtp, tag + " Length", test.input, shorts, test.shorts)
		}
	}
}
