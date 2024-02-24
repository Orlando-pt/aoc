package days

class Day5 : Day(5) {

    override fun partOne(): Any {
        return inputList.maxOfOrNull {
            parseLine(it)
        } ?: 0
    }

    override fun partTwo(): Any {
        // it was difficult to produce a test for this exercise
        if (inputList.size < 10) return 0

        val seatArray = Array(1000) { false }

        inputList.forEach() {
            seatArray[parseLine(it)] = true
        }

        for (i in 8 until seatArray.size - 9) {
            if (!seatArray[i] && seatArray[i - 1] && seatArray[i + 1]) return i
        }

        return 0
    }

    private fun parseLine(line: String): Int {
        var front = 0
        var back = 127

        for (i in 0..6) {
            val middle = (front + back) / 2

            if (line[i] == 'F') back = middle
            else front = middle
        }


        var left = 0
        var right = 7

        for (i in 7..9) {
            val middle = (left + right) / 2

            if (line[i] == 'R') left = middle
            else right = middle
        }

        return back * 8 + right
    }
}
