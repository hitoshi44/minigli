package minigli

import (
	"os"
)

type MiniGli struct {
	Cmd		string
	Subs	[]string
	Longs	map[string]string
	Shorts	map[string]string
}

func Pack() MiniGli {
	paths, longs, shorts := parseInputFrom(os.Args)
	return MiniGli{
		Cmd : paths[0],
		Subs: paths[1:],
		Longs:  longs,
		Shorts: shorts}
}

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
