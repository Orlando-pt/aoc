package days

class Day3 : Day(3) {
    override fun partOne(): Any {
        return inputList.sumOf { line ->
            val lineSorted = line.toSortedSet().reversed()
            val highestPosition = line.indexOf(lineSorted.first())
            val decimal = if (highestPosition == line.length - 1) lineSorted[1] else lineSorted.first()

            val decimalPos = line.indexOf(decimal)
            val secondHalf = line.substring(decimalPos + 1, line.length)
            val secondHalfSorted = secondHalf.toSortedSet().reversed()

            val joltage = decimal.toString() + secondHalfSorted.first()

            joltage.toInt()
        }
    }

    override fun partTwo(): Any {
        return inputList.sumOf { line ->
            var remove = line.length - BATTERIES
            val stack = ArrayDeque<Char>()

            for (digit in line) {
                while (remove > 0 && stack.isNotEmpty() && stack.last() < digit) {
                    stack.removeLast()
                    remove -= 1
                }
                stack.addLast(digit)
            }

            val joltage = stack.joinToString("").take(BATTERIES)
            joltage.toLong()
        }
    }

    private companion object {
        const val BATTERIES = 12
    }
}
