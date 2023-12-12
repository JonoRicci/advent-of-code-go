# Jono's Advent Of Code GoLang Solutions <!-- omit in toc -->

My attempts at the problems from [Advent of Code](https://adventofcode.com/). Solutions are organised by year and problem.

## Table of Contents <!-- omit in toc -->

- [Solutions](#solutions)
- [Usage](#usage)
  - [Testing](#testing)

## Solutions

| Day | 2023 |
|---|---|
| 01 | [Trebuchet?!][23d01] |
| 02 | [Cube Conundrum][23d02] |
| 03 | [Gear Ratios][23d03] |
| 04 |  |
| 05 |  |
| 06 |  |
| 07 |  |
| 08 |  |
| 09 |  |
| 10 |  |
| 11 |  |
| 12 |  |
| 13 |  |
| 14 |  |
| 15 |  |
| 16 |  |
| 17 |  |
| 18 |  |
| 19 |  |
| 20 |  |
| 21 |  |
| 22 |  |
| 23 |  |
| 24 |  |
| 25 |  |

## Usage

Navigate to problem directory and run `go run main.go`.

```shell
[19:59:13] ➜  advent-of-code-go git:(main) pwd
/Users/jono/repos/github/jonoricci/advent-of-code-go
[20:01:37] ➜  advent-of-code-go git:(main) ✗ cd 2023/day_01
[20:01:43] ➜  day_01 git:(main) ✗ go run main.go
2023/12/11 20:01:46 [INFO] Part 1 took: 65.667µs
2023/12/11 20:01:46 [INFO] Part 2 took: 17.475625ms
2023/12/11 20:01:46 [INFO] Part 1: 54597
2023/12/11 20:01:46 [INFO] Part 2: 54504
```

### Testing

The `ADVENT_OF_CODE_TEST` environment variable will dictate what puzzle input to use.

It is included in each problem's `main` function:

```go
// Set env var which dictates what input to use
// Options are "", "PART_01", "PART_02"
err := os.Setenv("ADVENT_OF_CODE_TEST", "PART_01")
if err != nil {
  fmt.Println("Error setting environment variable:", err)
}
```

However if the above is not present you can set it from your shell like so:

```shell
export "ADVENT_OF_CODE_TEST=" # Use regular puzzle "input.txt"
export "ADVENT_OF_CODE_TEST=PART_01" # Use "part01_test.txt"
export "ADVENT_OF_CODE_TEST=PART_02" # Use "part02_test.txt"
```

<!-- Links -->

[23d01]: 2023/day_01/
[23d02]: 2023/day_02/
[23d03]: 2023/day_03/
