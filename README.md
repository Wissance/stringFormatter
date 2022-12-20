# StringFormatter

A set of a ***high performance string tools*** that helps to build strings from templates and process text that 
faster than `fmt`!!!.

![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/wissance/stringFormatter?style=plastic) 
![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/wissance/stringFormatter?style=plastic) 
![GitHub issues](https://img.shields.io/github/issues/wissance/stringFormatter?style=plastic)
![GitHub Release Date](https://img.shields.io/github/release-date/wissance/stringFormatter) 
![GitHub release (latest by date)](https://img.shields.io/github/downloads/wissance/stringFormatter/v1.0.5/total?style=plastic)

![String Formatter: a convenient string formatting tool](/img/sf_cover.png)

## 1. Features

1. Text formatting with template using traditional for C#, Python programmers style - {0}, {name} that faster then fmt does:
![String Formatter: a convenient string formatting tool](/img/benchmarks2.png)
2. Additional text utilities:
   - convert ***map to string*** using one of predefined formats (see `text_utils.go`)

### 1. Text formatting from templates

#### 1.1 Description

This is a GO module for ***template text formatting in syntax like in C# or/and Python*** using:
- `{n}` , n here is a number to notes order of argument list to use i.e. `{0}`, `{1}`
- `{name}` to notes arguments by name i.e. `{name}`, `{last_name}`, `{address}` and so on ...

#### 1.2 Examples

##### 1.2.1 Format by arg order

i.e. you have following template:  `"Hello {0}, we are greeting you here: {1}!"`

if you call Format with args "manager" and "salesApp" :

```go
formattedStr := Format("Hello {0}, we are greeting you here: {1}!", "manager", "salesApp")
```

you get string `"Hello manager, we are greeting you here: salesApp!"`

##### 1.2.2 Format by arg key

i.e. you have following template: `"Hello {user} what are you doing here {app} ?"`

if you call `FormatComplex` with args `"vpupkin"` and `"mn_console"` `FormatComplex("Hello {user} what are you doing here {app} ?", map[string]interface{}{"user":"vpupkin", "app":"mn_console"})`

you get string `"Hello vpupkin what are you doing here mn_console ?"`

another example is:

```go
    strFormatResult = FormatComplex("Current app settings are: ipAddr: {ipaddr}, port: {port}, use ssl: {ssl}.", 
                                    map[string]interface{}{"ipaddr":"127.0.0.1", "port":5432, "ssl":false})
```
a result will be: `"Current app settings are: ipAddr: 127.0.0.1, port: 5432, use ssl: false."``

#### 1.2.3 Benchmarks of the Format and FormatComplex functions

benchmark could be running using following commands from command line:
* to see `Format` result - `go test -bench=Format -benchmem -cpu 1`
* to see `fmt` result - `go test -bench=Fmt -benchmem -cpu 1`

### 2. Text utilities

Map to string function allow to convert map to string using one of predefined line format:
* `key => value`
* `key : value`
* `value`

For example see code from test (`text_utils_test.go`):
```go
options := map[string]interface{}{
		"connectTimeout": 1000,
		"useSsl":         true,
		"login":          "sa",
		"password":       "sa",
	}

	str := MapToString(&options, KeyValueWithSemicolonSepFormat, ", ")
	assert.True(t, len(str) > 0)
	assert.Equal(t, "connectTimeout : 1000, useSsl : true, login : sa, password : sa", str)
```

### 3. Contributors

<a href="https://github.com/Wissance/stringFormatter/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Wissance/stringFormatter" />
</a>
