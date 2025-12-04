package days

class Day4 : Day(4) {

    override fun partOne(): Any {
        var res = 0
        for (y in inputList.indices) {
            for (x in inputList[y].indices) {
                if (inputList[y][x] != '@') {
                    continue
                }

                val paperCounter = getSurrondings(y, x, inputList)

                if (paperCounter < 4) {
                    res += 1
                }
            }
        }

        return res
    }

    override fun partTwo(): Any {
        var res = 0
        val toRemoveList = mutableListOf<Pair<Int, Int>>()
        val inputListMutable = inputList.toMutableList()

        do {
            toRemoveList.clear()
            for (y in inputListMutable.indices) {
                for (x in inputListMutable[y].indices) {
                    if (inputListMutable[y][x] != '@') {
                        continue
                    }

                    val paperCounter = getSurrondings(y, x, inputListMutable)

                    if (paperCounter < 4) {
                        toRemoveList.add(Pair(x, y))
                    }
                }
            }

            toRemoveList.forEach {
                val guilhermeGay = inputListMutable[it.second].replaceRange(it.first, it.first + 1, ".")
                inputListMutable[it.second] = guilhermeGay
            }
            res += toRemoveList.size
        } while (toRemoveList.isNotEmpty())

        return res
    }

    private fun getSurrondings(y: Int, x: Int, matrix: List<String>): Int {
        var paperCounter = 0

        if (x - 1 >= 0 && y - 1 >= 0 && matrix[y - 1][x - 1] == '@') {
            paperCounter += 1
        }
        if (y - 1 >= 0 && matrix[y - 1][x] == '@') {
            paperCounter += 1
        }
        if (y - 1 >= 0 && x + 1 < matrix.size && matrix[y - 1][x + 1] == '@') {
            paperCounter += 1
        }
        if (x - 1 >= 0 && matrix[y][x - 1] == '@') {
            paperCounter += 1
        }
        if (x + 1 < matrix.size && matrix[y][x + 1] == '@') {
            paperCounter += 1
        }
        if (y + 1 < matrix.size && x - 1 >= 0 && matrix[y + 1][x - 1] == '@') {
            paperCounter += 1
        }
        if (y + 1 < matrix.size && matrix[y + 1][x] == '@') {
            paperCounter += 1
        }
        if (y + 1 < matrix.size && x + 1 < matrix.size && matrix[y + 1][x + 1] == '@') {
            paperCounter += 1
        }

        return paperCounter
    }
}
