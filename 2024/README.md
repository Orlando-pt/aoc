# Advent of Code 2024 in Go

Let's dominate AoC 2024 🎄!

- [Day01](./solution/day01/main.go) ★★
- [Day02](./solution/day02/main.go) ★★
- [Day03](./solution/day03/main.go) ★★
- [Day04](./solution/day04/main.go) ★★
- [Day05](./solution/day05/main.go) ★★
- [Day06](./solution/day06/main.go) ★★

★ = completed the solution

## Usage

I firstly test my solutions in the `test` folder. Here is an example:

```shell
make test d=01
```

or

```shell
go test ./test/day01
```

The solutions are executed from the root folder like this:

```shell
go run . 1
```

or with debug output:

```shell
go run . 1 debug
```

The output has the following structure(no spoilers):

```
===== 01 =====
Part 1: 5459
Part 1 took 53.584µs
Part 2: 5454
Part 2 took 11.4055ms
```

