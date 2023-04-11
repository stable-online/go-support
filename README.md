# Support

Support is a Programming helpers powered by Golang [1.18](https://tip.golang.org/doc/go1.18)+ [generics](https://tip.golang.org/doc/go1.18#generics).

## Contents

- [Install](#install)
- [Features](#features)
  - [OverviewCases](#OverviewCases)
    - [SliceCase](#SliceCase)
    - [MapCase](#MapCase)
  - [SliceStream](#SliceStream)
    - [NewS](#NewS)
    - [Get](#get)
    - [Map](#map)
    - [Filter](#filter)
    - [Reduce](#reduce)
  - [SliceFunction](#SliceFunction)
    - [SMapF](#SMapF)
    - [SFilterF](#SFilterF)
    - [SReduceF](#SReduceF)
  - [MapStream](#MapStream)
    - [NewM](#NewS)
    - [Get](#get)
    - [Map](#map)
    - [Filter](#filter)
    - [Reduce](#reduce)
  - [MapFunction](#MapFunction)
    - [MMapF](#SMapF)
    - [MFilterF](#SFilterF)
    - [MReduceF](#SReduceF)
  

## Install

Requires go 1.18+

```sh
go get github.com/stable-online/support
```

## Features

### OverviewCases

### SliceCase

```go
support.NewS([]string{"a", "b", "c"}).Map(support.SMapF(func(key int, item string) string {
// Concatenate the string "1" to the value of each item.
return item + "1"
})).Filter(support.SFilterF(func(i int, t string) bool {
// Display to standard output the values of each item that is not equal to "a1".
return t != "a1"
})).Reduce(support.SReduceF(func(carry string, item string) string {
// Concatenate the value of each item.
return carry + item
}, ""))

// => b1c1
```

### MapCase

```go
support.NewM(map[string]int{"a": 1, "b": 2, "c": 3}).Map(support.MMapF(func(key string, v int) int {
// Add 1 to the value of each item.
return v + 1
})).Filter(support.MFilterF(func(k string, t int) bool {
// Display to standard output the values of each item that is not equal to 2.
return t != 2
})).Reduce(support.MReduceF(func(carry int, key string, item int) int {
// Add the result to the value of each item.
return carry + item
}, 0))

// => b1c1
```

### SliceStream

#### NewS

```go
support.NewS([]string{"a", "b", "c"})

// => &{[a b c]}
```

#### Get

```go
support.NewS([]string{"a", "b", "c"}).Get()

// => [a b c]
```

#### Map

```go
support.NewS([]string{"a", "b", "c"}).Map(support.SMapF(func(i int, t string) string { return t + "h" })).Get()

// => [ah bh ch]
```

#### Filter

```go
support.NewS([]string{"a", "b", "c"}).Filter(support.SFilterF(func(k int, v string) bool {return v != "a" })).Get()

// => &{[b c]}
```
#### Reduce

```go
support.NewS([]string{"a", "b", "c"}).Reduce(support.SReduceF(func(carry string, item string) string {return carry + item }, ""))

// => abc
```

### SliceFunction

#### SMapF

```go
support.SMapF(func(key int, item string) string { return item + "1" })([]string{"a", "b", "c"})

// => [a1 b1 c1]
```

#### SFilterF

```go
support.SFilterF(func(i int, t string) bool { return t != "a" })([]string{"a", "b", "c"})

// => [b c]
```

#### SReduceF

```go
support.SReduceF(func(i string, t string) string { return i + t }, "")([]string{"a", "b", "c"})

// => abc
```

### MapStream

#### NewS

```go
support.NewS([]string{"a", "b", "c"})

// => &{[a b c]}
```

#### Get

```go
support.NewS([]string{"a", "b", "c"}).Get()

// => [a b c]
```

#### Map

```go
support.NewS([]string{"a", "b", "c"}).Map(support.SMapF(func(i int, t string) string { return t + "h" })).Get()

// => [ah bh ch]
```

#### Filter

```go
support.NewS([]string{"a", "b", "c"}).Filter(support.SFilterF(func(k int, v string) bool {return v != "a" })).Get()

// => &{[b c]}
```
#### Reduce

```go
support.NewS([]string{"a", "b", "c"}).Reduce(support.SReduceF(func(carry string, item string) string {return carry + item }, ""))

// => abc
```

### MapFunction

#### MMapF

```go
support.SMapF(func(key int, item string) string { return item + "1" })([]string{"a", "b", "c"})

// => [a1 b1 c1]
```

#### MFilterF

```go
support.SFilterF(func(i int, t string) bool { return t != "a" })([]string{"a", "b", "c"})

// => [b c]
```

#### MReduceF

```go
support.SReduceF(func(i string, t string) string { return i + t }, "")([]string{"a", "b", "c"})

// => abc
```