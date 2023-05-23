# Support

Support 是通过Golang 1.18+版本以上,提供的开箱即用的功能工具助手.    

Support is a Programming helpers powered by Golang [1.18](https://tip.golang.org/doc/go1.18)+ [generics](https://tip.golang.org/doc/go1.18#generics).

## Contents

- [Install](#install)
- [OverviewCases](#OverviewCases)
  - [SliceCase](#SliceCase)
  - [MapCase](#MapCase)
- [Features](#features)
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

因为要用到泛型, 需要 go 版本为 1.18+  

Requires go 1.18+

```sh
go get github.com/stable-online/support
```

## OverviewCases

## Slice

```go
support.NewS([]string{"a", "b", "c"}).Map(support.SMapF(func(key int, item string) string {
// Concatenate the string '1' to the value of each item.
// 切片内每一个元素后拼接一个 1 字符串
return item + "1"
})).Filter(support.SFilterF(func(i int, t string) bool {
// Display to standard output the values of each item that is not equal to "a1".
// 如果不等于1, 返回true(将元素返回)	
return t != "a1"
})).Reduce(support.SReduceF(func(carry string, item string) string {
// Concatenate the value of each item.
// 拼接切片内的每一个元素. 	
return carry + item
}, ""))

// => b1c1
```

## Map

```go
support.NewM(map[string]int{"a": 1, "b": 2, "c": 3}).Map(support.MMapF(func(key string, v int) int {
// Add 1 to the value of each item.
// 字典内的每一个元素 + 1	
return v + 1
})).Filter(support.MFilterF(func(k string, t int) bool {
// Display to standard output the values of each item that is not equal to 2.
// 如果不等于2, 返回true(将元素返回)	
return t != 2
})).Reduce(support.MReduceF(func(carry int, key string, item int) int {
// Add the result to the value of each item.
// 将字典内的每一个元素值相加.
return carry + item
}, 0))

// => b1c1
```
## Features

### SliceStream

#### NewS

```go
support.NewS([]string{"a", "b", "c"})

// => [a b c]
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

#### NewM

```go
support.NewM(map[string]string{"a": "1", "b": "2", "c": "3"})

// => map[a:1 b:2 c:3]
```

#### Get

```go
support.NewM(map[string]string{"a": "1", "b": "2", "c": "3"}).Get()

// => map[a:1 b:2 c:3]
```

#### Map

```go
support.NewM(map[string]string{"a": "1", "b": "2", "c": "3"}).Map(support.MMapF(func(key string, val string) (res string) { return val + "1" })).Get()

// => map[a:11 b:21 c:31]
```

#### Filter

```go
support.NewM(map[string]string{"a": "1", "b": "2", "c": "3"}).Filter(support.MFilterF(func(k string, t string) bool { return t != "1" })).Get()

// => map[b:2 c:3]
```
#### Reduce

```go
support.NewM(map[string]string{"a": "1", "b": "2", "c": "3"}).Reduce(support.MReduceF(func(carry string, key string, item string) string { return carry + key + item }, ""))

// tips : Because the output of maps is unordered, the output may be different each time.
// => abc 
```

### MapFunction

#### MMapF

```go
support.MMapF(func(key string, val string) (res string) { return key + val })(map[string]string{"a": "1", "b": "2", "c": "3"})

// => map[a:a1 b:b2 c:c3]
```

#### MFilterF

```go
support.MFilterF(func(key string, val string) bool { return val != "2" })(map[string]string{"a": "1", "b": "2", "c": "3"})

// => map[a:1 c:3]
```

#### MReduceF

```go
support.MReduceF(func(carry string, key string, item string) string { return carry + key + item }, "")(map[string]string{"a": "1", "b": "2", "c": "3"})

// => a1b2c3
```