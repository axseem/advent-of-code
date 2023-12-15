![logo](./assets/logo.png)

# Advent of Code

Solutions for the annual programming competition "Advent of Code"

# Efficiency

All the benchmarks were executed on the MacBook Air M1.

The way I check the efficiency of the solutions:

```sh
cd go
go test -bench=. -count=10 ./... > bench.txt
benchstat bench.txt
```

## 2023

| Go                              | Part One | Part Two |
| ------------------------------- | -------- | -------- |
| [day 01](./go/2023/01/day01.go) | `11 µs`  | `66 μs`  |
| [day 02](./go/2023/02/day02.go) | `2.5 µs` | `3.6 μs` |