# debug

Easy print debugging.

# feature

Print the string to standartd output with the file name and line number where the method was called.


# func Print

```
debug.Print("hoge") // => name.go:15 hoge
```

Print works like fmt.Print.

# func Printf

```
hoge := "piyo"
debug.Printf("%s", hoge) // => name.go:20 piyo
```

Printf works like fmt.Printf.
