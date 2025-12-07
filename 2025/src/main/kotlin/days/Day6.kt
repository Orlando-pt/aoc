package days

class Day6 : Day(6) {

    override fun partOne(): Any {
        val operations = inputList.last().split(" ").filter { it.isNotEmpty() }
        val columns = LongArray(operations.size) { 0 }

        operations.forEachIndexed { i, operation -> if (operation == "*") columns[i] = 1 }

        for (i in 0..inputList.lastIndex - 1) {
            val row = inputList[i].split(" ").filter { it.isNotEmpty() }

            row.forEachIndexed { j, value ->
                if (operations[j] == "*") {
                    columns[j] *= value.toLong()
                } else if (operations[j] == "+") {
                    columns[j] += value.toLong()
                } else {
                    error("Unknown operation")
                }
            }
        }

        return columns.sum()
    }

    override fun partTwo(): Any {
        var res = 0L
        val operations = inputList.last().split(" ").filter { it.isNotEmpty() }
        val maxLineLength = inputList.maxOf { it.length } - 1
        val numbersList = inputList.subList(0, inputList.size - 1)

        var operationIndex = operations.lastIndex
        val columnValues = mutableListOf<Long>()
        for (i in maxLineLength downTo 0) {
            var number = ""
            for (line in numbersList) {
                if (i >= line.length) continue
                if (line[i] == ' ') continue

                number += line[i]
            }

            if (number.isNotEmpty()) {
                columnValues.add(number.toLong())
            } else {
                res += getColumnSum(columnValues, operations[operationIndex])

                columnValues.clear()
                operationIndex--
            }
        }

        res += getColumnSum(columnValues, operations.first())

        return res
    }

    private fun getColumnSum(column: List<Long>, operation: String): Long {
        return when (operation) {
            "+" -> {
                column.sum()
            }
            "*" -> {
                column.reduce { acc, l -> acc * l }
            }
            else -> {
                error("Unknown operation")
            }
        }
    }
}
