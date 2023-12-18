# Day 7: Camel Cards

[Puzzle Link](https://adventofcode.com/2023/day/7).

## Reflections

This was quite challenging and my code got quite complex. This did prompt me to learn how to write unit tests in Go which was useful.

For Part 2 where you have to treat Joker cards as wildcards and work out the strongest hand, I couldn't think of another way to do it than work out every single combination of a hand and then rank them.

I think I could refactor this to shorten the lines of code as I repeat myself with some new near duplicate functions in Part 02.
