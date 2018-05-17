# debug

Easy print debugging.

# feature

It outputs a string with the file name and line number to which the function was called.


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
It works like fmt.Println.

# func Printf

```
hoge := "piyo"
debug.Printf("%s", hoge) // => name.go:20 piyo
```

Printf prints a formatted message with a file name and a line number.
It works like fmt.Printf.
