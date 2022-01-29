package minigli

import (
	"os"
)

type CliPack struct {
	Cmd string
	Args []string
	Opts map[string]string
	ShortOpts map[string]string
	Paths []string
}

func Pack() CliPack {
	paths, longs, shorts := parseInputFrom(os.Args)
	return CliPack{
		Cmd : paths[0],
		Args: paths[1:],
		Opts: longs,
		Paths: paths,
		ShortOpts: shorts}
}

type argKind = int

const (
	shortOptKey argKind = iota
	shortOptPair
	longOptKey
	longOptPair
	valueOrArg
)

func parseInputFrom(args []string) ([]string,map[string]string,map[string]string) {

	paths := make([]string,0,len(args))
	sOpts := make(map[string]string)
	lOpts := make(map[string]string)

	i := 0
	lim := len(args)
	for i < lim {
		arg := args[i]
		// Judge an option or argument.
		if arg[0] == '-' {
			// Option (Key or Pair).
			// Judge long or short option.
			if arg[1] == '-' {
				// Long Option.
				// Judge a KeyValue or Key only.
				pos := findColonOrEqual(arg)
				if pos > -1 {
					lOpts[arg[2:pos]] = arg[pos+1:]
				} else {
					// check if the last
					if (i+1 < lim){
						lOpts[arg[2:]] = args[i+1]
						i++ // extra add
					} else {
						lOpts[arg[2:]] = ""
					}
				}
			} else {
				// Short Option
				// Judge KeyValue or Key only
				pos := findColonOrEqual(arg)
				if pos > -1 {
					sOpts[arg[1:pos]] = arg[pos+1:]
				} else {
					// check if the last
					if (i+1 < lim){
						sOpts[arg[1:]] = args[i+1]
						i++ // extra add
					} else {
						sOpts[arg[1:]] = ""
					}
				}
			}
		} else {
			paths = append(paths,arg)
		}
		i++ // for loop add
	}
	return paths, lOpts, sOpts
}


func findColonOrEqual(arg string) int {
	// For fail, return -1.
	result := -1
	i := 1
	for i<len(arg) {
		if arg[i] == ':' || arg[i] == '=' {
			result = i
			break
		}
		i++
	}
	return result
}
