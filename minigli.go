package minigli

import (
	"os"
)

type MiniGli struct {
	Cmds   []string
	Longs  map[string]string
	Shorts map[string]string
}

func Pack() (MiniGli, bool) {
	paths, longs, shorts, ok := parseInputFrom(os.Args[1:])
	return MiniGli{
			Cmds:   paths,
			Longs:  longs,
			Shorts: shorts},
		ok
}
func (mg *MiniGli) GetOption(option string, fullMatch bool) (value string, exist bool) {
	value, exist = mg.Longs[option]
	if exist {
		return value, exist
	}
	value, exist = mg.Shorts[option]
	if exist {
		return value, exist
	}
	if fullMatch {
		return value, exist
	}
	value, exist = mg.Shorts[option[:1]] // very first rune
	return value, exist
}
func parseInputFrom(args []string) ([]string, map[string]string, map[string]string, bool) {

	paths := make([]string, 0, len(args))
	sOpts := make(map[string]string)
	lOpts := make(map[string]string)
	ok := true

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
				if pos == 2 {
					// Invalid option "--:*"
					ok = false
					break
				}
				if pos > -1 {
					lOpts[arg[2:pos]] = arg[pos+1:]
				} else {
					// check if the last
					if i+1 < lim {
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
				if pos == 1 {
					// Invalid option "-:*"
					ok = false
					break
				}
				if pos > -1 {
					sOpts[arg[1:pos]] = arg[pos+1:]
				} else {
					// check if the last
					if i+1 < lim {
						sOpts[arg[1:]] = args[i+1]
						i++ // extra add
					} else {
						sOpts[arg[1:]] = ""
					}
				}
			}
		} else {
			paths = append(paths, arg)
		}
		i++ // for loop add
	}
	return paths, lOpts, sOpts, ok
}
func findColonOrEqual(arg string) int {
	// For fail, return -1.
	result := -1
	i := 1
	for i < len(arg) {
		if arg[i] == ':' || arg[i] == '=' {
			result = i
			break
		}
		i++
	}
	return result
}
