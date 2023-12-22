# Day 8: Haunted Wasteland

[Puzzle Link](https://adventofcode.com/2023/day/8).

## Reflections

Part 1 was okay for this challenge, but Part 2 I found very difficult and it took a number of days for me to work out how to solve this problem.

My initial strategy was to brute force through each step until we reached one where all nodes ended with "Z". My program was taking an hour to run at least and after finding my answer through a third party online tool I realised my program was going to take literal years to reach the answer.

As a result I was forced to change my solution for something more efficient.

I read up about using the Least Common Multiple (LCM) method, where I would tackle each search through the network seperately. Then I would compare the number of steps each path took to reach a node that ends in Z, and find the lowest common multiple between these paths to get my answer.

My previous solution was going to take over three years to run, where as this method took 5.1 milliseconds so I can safely say this is a bit more efficient.
