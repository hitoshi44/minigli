# minigli

`minigli` is a tiny command argument parser for Go.

# Usage 

`minigli.Pack()` returns MiniGli (and error check bool).

``` go
type MiniGli struct {
	Cmd	string           // Command 
	Subs	[]string         // Sub Commands 
	Longs	map[string]string// Options start with "--"
	Shorts	map[string]string// Options start with "-"
}
```

 - Arguments consist of Options and Commands.
 - Options are key-value pair given with 1 or 2 "-"s.
 - Commands are the rest.
 - `Cmd` is the first of Commands
 - `Subs` are the rest.
 
e.g)

``` go
// your-cli verb target --option value -another:value target2
mg, ok := minigli.Pack()
if !ok {
    // Handle Invalid Argument.
    // Only "invalid option" occurs invalid argument.
    // "invalid option" will be explained later.
}

mg.Cmd // "verb"
mg.Subs// ["target", "target2"]
mg.Longs// {"option" : "value"}
mg.Shorts//{"another": "value"}
```

# Options rule

Long option can be specified as below:

 - `--option value`
 - `--option:value`
 - `--option=value`

Short option can be as same.

Key-only option is specified like this:

 - `--opt:`
 - `-opts=`

Only last of arguments can be key only option withou ":" or "=".

```
your-cli command sub commands -o file.txt --validKeyOnlyOption
```

Belows will be treated as invalid options and `Pack()` fails to parse then returns `false`.

 - `--:SomeString`
 - `-:ThisIsToo`
