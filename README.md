# minigli

`minigli` is a tiny command argument parser for Go.

# Usage 

## Start to use

Just use `minigli.Pack()` which returns object variant named MiniGli (and error check bool).

(TODO : change this book into err?)

```
import "github.com/hitoshi44/minigli"

mg, ok := minigli.Pack()
if !ok {...}
```

And this is MiniGli object defined.

``` go
type MiniGli struct {
	Cmds	string           // Command 
	Longs	map[string]string// Options start with "--"
	Shorts	map[string]string// Options start with "-"
}
```

Each property is read from Arguments. This is a rule.

 - Arguments consist of Options and Commands.
 - Options are key-value pair given with 1 or 2 "-"s. `Longs`, `Shorts`
 - Commands are the rest. `Cmds`
 
e.g)

``` go
// your-cli verb target --option value -another:value target2
mg, ok := minigli.Pack()
if !ok {
    // Handle Invalid Argument.
    // Only "invalid option" occurs invalid argument.
    // "invalid option" will be explained later.
}

mg.Cmds // ["verb", "target", "target2"]
mg.Longs// {"option" : "value"}
mg.Shorts//{"another": "value"}
```

## Helper Methods

### GetOption(string, bool) (string, bool)
``` go
value, found := mg.GetOption("option", false)
if exist {
  process(value)
}
```
`GetOption` take option's name and bool(fullMatch), returns options value if exist, and boolean represents found or not.
If pass argument fullMatch=true, this method find only full match option, ignore 1 char match.

e.g) 

 - `GetOption("option", false)` matches "--option", "-option", "-o"
 - `GetOption("optionF, true")` matches "--option", "-option"



# Options rule

Long option can be specified as below:

 - `--option value`
 - `--option:value`
 - `--option=value`

Short option can be as same.

Key-only option is specified like this:

 - `--opt:`
 - `-opts=`

Only last of arguments can be key-only option withou ":" or "=".

```
your-cli command sub commands -o file.txt --validKeyOnlyOption
```

Belows will be treated as invalid options and `Pack()` fails to parse then returns `false`.

 - `--:SomeString`
 - `-:ThisIsToo`
