package days

class Day1 : Day(1) {
    private val dialMax = 100

    override fun partOne(): Any {
        val instructions = getInstructions()
        var dial = 50
        var dialZero = 0

        instructions.forEach {
            if (it.direction == Direction.L) {
                val newDial = dial - it.distance

                dial = if (newDial < 0) {
                    dialMax + newDial
                } else {
                    newDial
                }
            } else {
                dial += it.distance
            }
            dial = dial % dialMax

            if (dial == 0) {
                dialZero += 1
            }
        }

        return dialZero
    }

    override fun partTwo(): Any {
        val instructions = getInstructions()
        var dial = 50
        var dialZero = 0

        instructions.forEach {
            val remainder = it.distance % dialMax
            dialZero += it.distance / dialMax

            if (it.direction == Direction.L) {
                val newDial = dial - remainder

                if (newDial < 0) {
                    if (dial != 0) {
                        dialZero += 1
                    }

                    dial = dialMax + newDial
                } else {
                    dial = newDial
                }
            } else {
                dial += remainder

                if (dial > dialMax) {
                    dialZero += 1
                }
            }
            dial = dial % dialMax

            if (dial == 0) {
                dialZero += 1
            }
        }
        return dialZero
    }
    enum class Direction {
        L, R
    }

    data class Instruction(val direction: Direction, val distance: Int)

    private fun getInstructions(): List<Instruction> {
        return inputList.map {
            val direction = Direction.valueOf(it[0].toString())
            val distance = it.substring(1).toInt()
            Instruction(direction, distance)
        }
    }
}
