# recurring-values

A simple project to play with spf13/cobra library used for building CLI tools.

Find recurring lines within specified files or standard input.

Usage:

- Build the project with
```
go build
```


- Find recurring words in standard input:

```
./recurring-values find text random1 random2 random1

```

- Find recurring lines in specified files:

```
 ./recurring-values find files words.txt words2.txt

```


### to do
- add file generator capability
