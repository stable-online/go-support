# Support

Support is a Programming helpers powered by Golang [1.18](https://tip.golang.org/doc/go1.18)+ [generics](https://tip.golang.org/doc/go1.18#generics).

## Contents

- [Install](#install)
- [Features](#features)
  - [SliceStream](#SliceStream)
    - [OverviewCases](#reduce)
    - [NewS](#support.NewS)
    - [Get](#get)
    - [Map](#map)
    - [Filter](#filter)
    - [Reduce](#reduce)
  - [SliceFunction](#SliceFunction)
    - [SMapF](#SMapF)
    - [SFilterF](#SFilterF)
    - [SReduceF](#SReduceF)

## Install

Requires go 1.18+

```sh
go get github.com/stable-online/support
```

## Features

### SliceStream

#### OverviewCases

```go
support.NewS([]string{"a", "b", "c"}).Map(support.SMapF(func(key int, item string) string {
return item + "1"
})).Filter(support.SFilterF(func(i int, t string) bool {
return t != "a1"
})).Reduce(support.SReduceF(func(carry string, item string) string {
return carry + item
}, ""))

// => b1c1
```

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