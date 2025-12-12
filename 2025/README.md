# Advent of Code in Kotlin

Let's dominate AoC 2025 🎄!

- [Day01](./src/main/kotlin/days/Day1.kt) ★★
- [Day02](./src/main/kotlin/days/Day2.kt) ★★
- [Day03](./src/main/kotlin/days/Day3.kt) ★☆
- [Day04](./src/main/kotlin/days/Day4.kt) ★★
- [Day05](./src/main/kotlin/days/Day5.kt) ★★
- [Day06](./src/main/kotlin/days/Day6.kt) ★★
- [Day07](./src/main/kotlin/days/Day7.kt) ★☆
- [Day08](./src/main/kotlin/days/Day8.kt) ☆☆
- [Day09](./src/main/kotlin/days/Day9.kt) ★☆

★ = completed the solution
☆ = had to get help

### Running

Project is already setup with gradle. To run the app:

* Navigate to top-level directory on the command line
* Run `./gradlew run` to run all days
* Run `./gradlew run --args $DAY` where `$DAY` is an integer to run a specific day
* Optional (running with Make): `make run d=1`

### Testing

Project includes Junit and Hamcrest and a stub unit test to get you going. To run all tests:

* Navigate to top-level directory on the command line
* Run `./gradlew test`
* Add `--info`, `--debug` or `--stacktrace` flags for more output
* Optional (running with Make): `make test`

##### Test input

By default, instantiations of `Day` classes in tests will use the input files in `src/test/resources`, _not_ those
in `src/main/resources`.
This hopefully gives you flexibility - you could either just copy the real input into `src/test/resources` if you want
to test
the actual answers, or you could add a file of test data based on the examples given on the Advent of Code description
for the day.
The stub `Day1Test` class shows a test of the functionality of `Day1` where the test input differs from the actual
input.

### Architecture

* Inputs go into `src/main/resources` and follow the naming convention `input_day_X.txt`
* Solutions go into `src/main/kotlin/days` and extend the `Day` abstract class, calling its constructor with their day
  number
* Solutions follow the naming convention `DayX`
* It is assumed all solutions will have two parts but share the same input
* Input is exposed in the solution classes in two forms - `inputList` and `inputString`
* Day 1 solution class and input file are stubbed as a guide on how to extend the project,
  and how you can use the `inputList` and `inputString` mentioned above
* To get started simply replace `src/main/input_day_1.txt` with the real input and the solutions in `Day1` with your own
* A Day 1 test class also exists, mostly to show a few hamcrest matchers, and how test input files can differ from
  actual ones (see **Test input** section above).
  To get started with testing you can edit this class, and the input file at `src/test/resources/input_day_1.txt`

