# WC Tool 
Word Count Tool written in Golang as part of reading the book [https://pragprog.com/titles/rggo/powerful-command-line-applications-in-go/](Powerful Command-Line Applications in Go)

## Build from Source
### Build for Windows
```sh
GOOS=windows go build
```

### Build for Linux
```sh
GOOS=linux go build
```

### Build for Darwin OS(MacOS)
```sh
GOOS=darwin go build
```


## Usage
### Count Number of words in a file
```sh
cat main.go | wc
```

### Count Number of Lines in a file
```sh
cat main.go | wc -l
```

### Count Number of Bytes in a File
```sh
cat main.go | wc -b
```

