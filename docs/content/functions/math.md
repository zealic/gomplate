---
title: math functions
menu:
  main:
    parent: functions
---

A set of basic math functions to be able to perform simple arithmetic operations with `gomplate`.

## `math.Add`

Adds two numbers together

### Usage
```go
math.Add x y
```
```go
y | math.Add x
```

### Example

```console
$ gomplate -i '{{ math.Add 1 2 }}
3
```

