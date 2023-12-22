# Jono's Advent Of Code Golang Solutions <!-- omit in toc -->

My attempts at the problems from [Advent of Code][url_aoc]. Solutions are organised by year and problem.

## Table of Contents <!-- omit in toc -->

- [Disclaimer](#disclaimer)
- [Solutions](#solutions)
- [Usage](#usage)
  - [Go Version](#go-version)
  - [Config File](#config-file)
  - [Unit Tests](#unit-tests)

## Disclaimer

The solutions presented in this repository reflect my ongoing learning process and, as such, may not always be the most efficient or idiomatic approaches to solving these problems in Go. I welcome any feedback or suggestions that would help enhance my understanding of Go and improve the quality of these solutions.

## Solutions

| Day | 2023 |
|---|---|
| 01 | [Trebuchet?!][23d01] |
| 02 | [Cube Conundrum][23d02] |
| 03 | [Gear Ratios][23d03] |
| 04 | [Scratchcards][23d04] |
| 05 | [If You Give A Seed A Fertilizer][23d05] |
| 06 | [Wait For It][23d06] |
| 07 | [Camel Cards][23d07] |
| 08 | [Haunted Wasteland][23d08] |
| 09 | [Mirage Maintenance][23d09] |
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
2023-12-14T13:28:09.981Z   info   day_01/main.go:95    Part 1 took: 102.291µs
2023-12-14T13:28:09.997Z   info   day_01/main.go:125   Part 2 took: 15.032625ms
2023-12-14T13:28:09.997Z   info   day_01/main.go:58    Part 1: 54597
2023-12-14T13:28:09.997Z   info   day_01/main.go:59    Part 2: 54504
```

Running commands from the repo root directory or any other directory won't work as the config expects relative directories from the main file.

### Go Version

I'm using `1.21.4` throughout the repo as that was the latest available.

I'm using [goenv][url_goenv] to manage Go versions in my development environment. This places a `.go-version` file in my root directory.

You do not need to have [goenv][url_goenv] installed to run this code, you just need to have the same version of Go available to use on your system.

### Config File

Each day has it's own `config.yaml` config file which can be used to modify some behaviours.

- `inputFile`: relative path to the puzzle input, can switch between test and real input.
- `logLevel`: [zap][url_zap] logging levels, handy to switch between `Debug` and `Info`.

### Unit Tests

Some problems have a unit test file in their directory usually named `main_test.go`. You can run this test by navigating to the directory and running `go test`.

```shell
[16:59:03] ➜  day_07 git:(main) ✗ pwd
/Users/jono/repos/github/jonoricci/advent-of-code-go/2023/day_07
[16:59:05] ➜  day_07 git:(main) ✗ ls
config.yaml    input.txt      main.go        main_test.go   test_input.txt
[16:59:09] ➜  day_07 git:(main) ✗ go test
PASS
ok     jonoricci/advent-of-code-go/2023/day_07   0.234s
```

<!-- Links -->

[23d01]: 2023/day_01/
[23d02]: 2023/day_02/
[23d03]: 2023/day_03/
[23d04]: 2023/day_04/
[23d05]: 2023/day_05/
[23d06]: 2023/day_06/
[23d07]: 2023/day_07/
[23d08]: 2023/day_08/
[23d09]: 2023/day_09/

[url_aoc]: https://adventofcode.com/
[url_zap]: https://github.com/uber-go/zap
[url_goenv]: https://github.com/go-nv/goenv
