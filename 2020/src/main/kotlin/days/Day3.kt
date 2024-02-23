package days

class Day3 : Day(3) {
    private val xLen = inputList[0].length
    override fun partOne(): Any {
        return getTreesInSlope(3)
    }

    override fun partTwo(): Any {
        val x1 = getTreesInSlope(1)
        val x3 = getTreesInSlope(3)
        val x5 = getTreesInSlope(5)
        val x7 = getTreesInSlope(7)
        val d2 = getTreesInSlope(1, 2)

        return x1 * x3 * x5 * x7 * d2
    }

    private fun getTreesInSlope(xOffset: Int, yOffset: Int = 1): Long {

        var x = 0;
        return inputList
            .filterIndexed { index, _ -> index % yOffset == 0 }
            .count {
                val isTree = it[x] == '#'

                x = (x + xOffset) % xLen

                isTree
            }.toLong()
    }
}