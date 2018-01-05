# todo
a command-line tool for TODO management

## Usage

### List todo

```
$ todo list
☐ 001: Send Mail
  send a mail to Mr. Yamada

☐ 002: Buy water
  buy water at the supermarket

☐ 003: Bake cake
  bake a cake
```

### Add new task

```
$ todo add
title:[] Write report
Detail:[] write a work report
```

```
$ todo list
 001: Send Mail
  send a mail to Mr. Yamada

☐ 002: Buy water
  buy water at the supermarket

☐ 003: Bake cake
  bake a cake

☐ 004: Write report
  write a work report
```

### Edit Task

```
$ todo edit 1
title:[Send Mail]
Detail:[send a mail to Mr. Yamada] send a mail to Mr. Tanaka
```

```
$ todo list
☐ 001: Send Mail
  send a mail to Mr. Tanaka

☐ 002: Buy water
  buy water at the supermarket

☐ 003: Bake cake
  bake a cake

☐ 004: Write report
  write a work report
```

### Delete task

```
$ todo delete 2
Delete task 002: Buy water ? (y/N) y
```

```
$ todo list
☐ 001: Send Mail
  send a mail to Mr. Tanaka

☐ 003: Bake cake
  bake a cake

☐ 004: Write report
  write a work report
```

### Done task

```
$ todo done 3
Done task 003: Bake cake ? (y/N) y
```

```
$ todo list
☐ 001: Send Mail
  send a mail to Mr. Tanaka

☑ 003: Bake cake
  bake a cake

☐ 004: Write report
  write a work report
```

### Undone task

```
$ todo undone 3
Undone task 003: Bake cake ? (y/N) y
```

```
$ todo list
☐ 001: Send Mail
  send a mail to Mr. Tanaka

☐ 003: Bake cake
  bake a cake

☐ 004: Write report
  write a work report
```

## Installation

```
$ go install github.com/ikmski/todo
```

