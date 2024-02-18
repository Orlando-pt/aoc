package days

class Day1 : Day(1) {

    override fun partOne(): Any {
        val numbers = inputList.map { it.toInt() }

        for (n1 in numbers.indices) {
            for (n2 in n1 until numbers.size) {
                if (numbers[n1] + numbers[n2] == 2020) {
                    return numbers[n1] * numbers[n2]
                }
            }
        }

        return 0
    }

    override fun partTwo(): Any {
        val numbers = inputList.map { it.toInt() }

        for (n1 in numbers.indices) {
            for (n2 in n1 until numbers.size) {
                for (n3 in n2 until numbers.size) {
                    if (numbers[n1] + numbers[n2] + numbers[n3] == 2020) {
                        return numbers[n1] * numbers[n2] * numbers[n3]
                    }
                }
            }
        }

        return 0
    }
}
