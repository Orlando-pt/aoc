package days

class Day9 : Day(9) {
    private val preamble = if (inputList.size > 25) 25 else 5

    override fun partOne(): Any {
        return findNumber().first
    }

    override fun partTwo(): Any {
        val (_, pos) = findNumber()

        val numbers = findContiguousListToGoal(pos)
        return numbers.min() + numbers.max()
    }

    private fun findContiguousListToGoal(goalPos: Int): List<Long> {
        val previousNumbers = inputList.slice(0..<goalPos).map { it.toLong() }
        val goal = inputList[goalPos].toLong()

        var currentSum: Long
        for (left in 0..<goalPos) {
            currentSum = previousNumbers[left]

            for (right in left + 1..<previousNumbers.size) {
                currentSum += previousNumbers[right]

                if (currentSum > goal) {
                    break
                } else if (currentSum == goal) {
                    return previousNumbers.slice(left..right)
                }
            }
        }
        return emptyList()
    }

    private fun findNumber(): Pair<Long, Int> {
        val numbers = inputList.map { it.toLong() }
        var currentGoal = preamble

        while (currentGoal < numbers.size) {
            val res = hasPreambleSum(numbers, currentGoal).not()
            if (res) {
                return Pair(numbers[currentGoal], currentGoal)
            }
            currentGoal++
        }

        return Pair(0, 0)
    }

    private fun hasPreambleSum(
        numbers: List<Long>,
        currentGoal: Int,
    ): Boolean {
        val left = currentGoal - preamble

        for (l in left..<left + currentGoal) {
            for (r in currentGoal - 1 downTo left) {
                if (l - currentGoal == preamble) return false
                if (numbers[l] + numbers[r] == numbers[currentGoal]) return true
            }
        }

        return false
    }
}
