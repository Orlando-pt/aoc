# Advent of Code 2022 in Java

Year 2022 is the year to consolidate Java knowledge. The following content is from the original repository
that is used as a template for this year's journey.

I created a custom [Makefile](./Makefile) to simplify the build and run process. Feel free to use it or
adapt it to your needs.

## Solutions

- [Day01](./src/main/java/aoc/day01/Day01.java) ★★
- [Day02](./src/main/java/aoc/day02/Day02.java) ★★
- [Day03](./src/main/java/aoc/day03/Day03.java) ★★
- [Day04](./src/main/java/aoc/day04/Day04.java) ★★
- [Day05](./src/main/java/aoc/day05/Day05.java) ★★

★ = completed the solution
☆ = had to get help

## Architecture

The `Main` method in `App` will execute the code for a given day based on input in `src/main/resources/day[number].txt`

The given `Day` class will process the input and return a result, which `App` will print to `stdout`.

## Assumptions

1. Each day's puzzle will have an input consisting of one or more lines of plain text.
1. Each day's puzzle will have two "parts" requiring separate implementations, but using the same input.
1. Puzzle solutions will be a single line of text.
1. Each puzzle description will provide at least one example input/output.

## How to add a puzzle

The shell for day 1 is already there. These are the steps to add day 2, etc. For day 1 you can start by writing the test
cases once the puzzle posts.

### Part 1

1. Create `aoc.day02.Day02.java implements Day`
1. Add a line to the static initialization block in `App.java` to add your nwe `Day`
   implementation: `DAYS.put(2, new Day02())`
1. Create `Day02Test.java` and use the sample input/output from adventofcode.com/2019/day/2 to write tests (they will
   initially fail).
1. Implement `Day02.part1()` until the tests pass (Test Driven Development!). You can run `./gradlew --continuous check`
   to have the tests run every time you save your code. Alternatively, you can run `make test`.
1. When the example input passes, get your solution from adventofcode.com/2019/day/2/input and save it
   as `src/main/resources/day02.txt`
1. Run `./gradlew clean build install && ./build/install/advent-of-code/bin/advent-of-code 2` to get your output.
   Alternatively, you can run `make run day=2`.
    - Notice the day number ("2") is passed to the start script as an argument.
1. Check your solution on adventofcode.com!

### Part 2

1. Add new tests for part 2 to `Day02Test.java` using the sample input/output from adventofcode.com/2019/day/2
1. Implement `Day02.part2()` until the tests pass.
1. When the example input passes,
   run `./gradlew clean build install && ./build/install/advent-of-code/bin/advent-of-code 2 2` to get your output.
   Alternatively, you can run `make run day=2 part=2`.
    - Notice the second argument to the script is the part number. It defaults to 1, so now you have to pass it as 2.
1. Check your solution on adventofcode.com!
