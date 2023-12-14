# Jono's Advent Of Code Golang Solutions <!-- omit in toc -->

My attempts at the problems from [Advent of Code](https://adventofcode.com/). Solutions are organised by year and problem.

## Table of Contents <!-- omit in toc -->

- [Solutions](#solutions)
- [Usage](#usage)
- [Config File](#config-file)

## Solutions

| Day | 2023 |
|---|---|
| 01 | [Trebuchet?!][23d01] |
| 02 | [Cube Conundrum][23d02] |
| 03 | [Gear Ratios][23d03] |
| 04 | [Scratchcards][23d04] |
| 05 | [If You Give A Seed A Fertilizer][23d05] |
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
[13:26:13] ➜  advent-of-code-go git:(main) pwd
/Users/jono/repos/github/jonoricci/advent-of-code-go
[13:27:53] ➜  advent-of-code-go git:(main) ✗ cd 2023/day_01
[13:28:07] ➜  day_01 git:(main) ✗ go run main.go
2023-12-14T13:28:09.981Z	info	day_01/main.go:95	Part 1 took: 102.291µs
2023-12-14T13:28:09.997Z	info	day_01/main.go:125	Part 2 took: 15.032625ms
2023-12-14T13:28:09.997Z	info	day_01/main.go:58	Part 1: 54597
2023-12-14T13:28:09.997Z	info	day_01/main.go:59	Part 2: 54504
```

Running commands from the root won't work as the config expects relative directories.

## Config File

Each day has it's own config file which can be used to modify some behaviours.

- `inputFile`: relative path to the puzzle input, can switch between test and real input.
- `logLevel`: [zap][url_zap] logging levels, handy to switch between `Debug` and `Info`.

<!-- Links -->

[23d01]: 2023/day_01/
[23d02]: 2023/day_02/
[23d03]: 2023/day_03/
[23d04]: 2023/day_04/
[23d05]: 2023/day_05/

[url_zap]: https://github.com/uber-go/zap
