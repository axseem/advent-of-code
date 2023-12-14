![logo](./assets/logo.png)

# Advent of Code

Solutions for the annual programming competition "Advent of Code"

# Efficiency

All the benchmarks were executed on the MacBook Air M1.

The method I used to check the efficiency of the solutions:

```sh
go test -bench=. -count=10 ./... > bench.txt && benchstat bench.txt
```

## 2023

| Go                              | Part One | Part Two |
| ------------------------------- | -------- | -------- |
| [day 01](./go/2023/01/day01.go) | `11 µs`  | `66 μs`  |