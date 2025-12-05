package days

import kotlin.math.max
import kotlin.math.min

class Day5 : Day(5) {

    override fun partOne(): Any {
        val emptyLine = inputList.indexOfFirst { it.isEmpty() }
        val ingredientRanges = getIngredientRanges()
        val ingredientIds = inputList.subList(emptyLine + 1, inputList.size).map { it.toLong() }

        return ingredientIds.filter {
            ingredientRanges.any { range -> it >= range.first && it <= range.second }
        }.size
    }

    override fun partTwo(): Any {
        var res = 0L
        val ingredientRanges = getIngredientRanges().sortedBy { it.first }

        var prevRange = ingredientRanges.first()
        for (i in 1 until ingredientRanges.size) {
            val currRange = ingredientRanges[i]

            if (prevRange.first <= currRange.first && currRange.first <= prevRange.second) {
                prevRange = Pair(min(prevRange.first, currRange.first), max(prevRange.second, currRange.second))
            } else {
                res += prevRange.second - prevRange.first + 1
                prevRange = currRange
            }
        }

        res += prevRange.second - prevRange.first + 1
        return res
    }

    private fun getIngredientRanges(): List<Pair<Long, Long>> {
        val emptyLine = inputList.indexOfFirst { it.isEmpty() }
        return inputList.subList(0, emptyLine).map {
            val (start, end) = it.split("-")
            Pair(start.toLong(), end.toLong())
        }
    }
}
