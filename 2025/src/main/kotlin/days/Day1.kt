package days

class Day1 : Day(1) {

    override fun partOne(): Any {
        val instructions = getInstructions()
        var dial = 50
        var dialZero = 0

        instructions.forEach {
            if (it.direction == Direction.L) {
                val newDial = dial - it.distance

                dial = if (newDial < 0) {
                    DIAL_MAX + newDial
                } else {
                    newDial
                }
            } else {
                dial += it.distance
            }
            dial = dial % DIAL_MAX

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
            val remainder = it.distance % DIAL_MAX
            dialZero += it.distance / DIAL_MAX

            if (it.direction == Direction.L) {
                val newDial = dial - remainder

                if (newDial < 0) {
                    if (dial != 0) {
                        dialZero += 1
                    }

                    dial = DIAL_MAX + newDial
                } else {
                    dial = newDial
                }
            } else {
                dial += remainder

                if (dial > DIAL_MAX) {
                    dialZero += 1
                }
            }
            dial = dial % DIAL_MAX

            if (dial == 0) {
                dialZero += 1
            }
        }
        return dialZero
    }

    private companion object {
        const val DIAL_MAX = 100
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
