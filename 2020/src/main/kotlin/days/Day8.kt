package days

class Day8 : Day(8) {
    override fun partOne(): Any {
        val (accumulator, _) = getAccumulator(inputList)
        return accumulator
    }

    override fun partTwo(): Any {
        inputList.forEachIndexed { idx, line ->
            val (op, arg) = line.split(" ")
            val newOp =
                when (op) {
                    "jmp" -> "nop"
                    "nop" -> "jmp"
                    else -> null
                }

            if (newOp != null) {
                val newInput = inputList.toMutableList()
                newInput[idx] = "$newOp $arg"

                val (accumulator, terminated) = getAccumulator(newInput)
                if (terminated) return accumulator
            }
        }
        return 0
    }

    private fun getAccumulator(input: List<String>): Pair<Int, Boolean> {
        var idx = 0
        var accumulator = 0
        val visited = mutableSetOf<Int>()

        while (visited.contains(idx).not()) {
            visited.add(idx)

            if (idx == input.size) return accumulator to true

            val (op, arg) = input[idx].split(" ")
            when (op) {
                "acc" -> {
                    accumulator += arg.toInt()
                    idx++
                }

                "jmp" -> idx += arg.toInt()
                "nop" -> idx++
            }
        }

        return accumulator to false
    }
}
