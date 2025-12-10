package days

class Day7 : Day(7) {

    override fun partOne(): Any {
        var res = 0
        val startPost = inputList.first().indexOfFirst { it == 'S' }
        val beams = BooleanArray(inputList.first().length) { false }
        beams[startPost] = true

        for (row in inputList.subList(2, inputList.size)) {
            val splitters = row.mapIndexedNotNull { index, elem -> index.takeIf { elem == '^' } }
            splitters.forEach {
                if (beams[it]) {
                    beams[it] = false
                    beams[it - 1] = true
                    beams[it + 1] = true

                    res += 1
                }
            }
        }

        return res
    }

    override fun partTwo(): Any {
        var res = 0L
        val startPost = inputList.first().indexOfFirst { it == 'S' }
        var currentBeams = mutableMapOf(Pair(0, startPost) to 1L)

        while (currentBeams.isNotEmpty()) {
            val nextBeams = mutableMapOf<Pair<Int, Int>, Long>()

            for ((beam, count) in currentBeams) {
                val (row, col) = beam
                if (row + 1 > inputList.lastIndex) {
                    res += count
                    continue
                }

                if (inputList[row + 1][col] == '^') {
                    val left = Pair(row + 1, col - 1)
                    val right = Pair(row + 1, col + 1)
                    val newCountLeft = nextBeams.getOrDefault(left, 0L) + count
                    val newCountRight = nextBeams.getOrDefault(right, 0) + count

                    nextBeams[left] = newCountLeft
                    nextBeams[right] = newCountRight
                } else {
                    val center = Pair(row + 1, col)
                    val newCountCenter = nextBeams.getOrDefault(center, 0) + count

                    nextBeams[center] = newCountCenter
                }
            }
            currentBeams = nextBeams
        }

        return res
    }
}
