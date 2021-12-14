### Description
This is a tiny GO module for ***template text formatting like in C# or/and Python*** using:
- {n} , n here is a number to notes order of argument list to use
- {name} to notes arguments by name i.e. {name}, {last_name}, {address} and so on ...

### Examples

#### Format by arg order
i.e. you have following template:  Hello {0}, we are greeting you here: {1}!

if you call Format with args "manager" and "salesApp" : 

```go
formattedStr := Format("Hello {0}, we are greeting you here: {1}!", "manager", "salesApp")
```

you get string "Hello manager, we are greeting you here: salesApp!"

#### Format by arg key
i.e. you have following template: "Hello {user} what are you doing here {app} ?"

if you call FormatComplex with args "vpupkin" and "mn_console" FormatComplex("Hello {user} what are you doing here {app} ?", map[string]interface{}{"user":"vpupkin", "app":"mn_console"})

you get string "Hello vpupkin what are you doing here mn_console ?"

another example is: 

```go
    strFormatResult = FormatComplex("Current app settings are: ipAddr: {ipaddr}, port: {port}, use ssl: {ssl}.", map[string]interface{}{"ipaddr":"127.0.0.1", "port":5432, "ssl":false})
```
a result will be: "Current app settings are: ipAddr: 127.0.0.1, port: 5432, use ssl: false."

