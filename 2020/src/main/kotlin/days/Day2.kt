package days

class Day2 : Day(2) {
    override fun partOne(): Any {
        return inputList.map {
            val lineParts = it.split(" ")

            val bounds = lineParts[0].split("-")
            PasswordPolicy(
                bounds[0].toInt(),
                bounds[1].toInt(),
                lineParts[1].first(),
                lineParts[2]
            )
        }.count {
            val letterOccurrences = it.password.count { pl -> pl == it.letter }

            letterOccurrences >= it.min && letterOccurrences <= it.max
        }
    }

    override fun partTwo(): Any {
        return inputList.map {
            val lineParts = it.split(" ")

            val bounds = lineParts[0].split("-")
            PasswordPolicy(
                bounds[0].toInt(),
                bounds[1].toInt(),
                lineParts[1].first(),
                lineParts[2]
            )
        }.count {
            (it.password[it.min - 1] == it.letter).xor(it.password[it.max - 1] == it.letter)
        }
    }

    data class PasswordPolicy(val min: Int, val max: Int, val letter: Char, val password: String)
}