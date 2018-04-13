# debug

Easy print debugging.

# feature

Print the string to standartd output with the file name and line number where the method was called.


# func Print

```
debug.Print("hoge") // => name.go:15 hoge
```

Print prints a message with a file name and a line number.
It works like fmt.Print.

# func Println

```
debug.Println("hoge") // => name.go:17 hoge
```
Println prints a message with a file name and a line number.
It works like fmt.Print.

# func Printf

```
hoge := "piyo"
debug.Printf("%s", hoge) // => name.go:20 piyo
```

Printf prints a formatted message with a file name and a line number.
It works like fmt.Printf.
