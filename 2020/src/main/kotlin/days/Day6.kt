package days

class Day6 : Day(6) {
    override fun partOne(): Any {
        var res = 0
        val differentLetters = mutableSetOf<Char>()

        inputList.forEach {
            if (it.isNotEmpty()) {
                differentLetters.addAll(it.toList())
            } else {
                res += differentLetters.size
                differentLetters.clear()
            }
        }
        return res + differentLetters.size
    }

    override fun partTwo(): Any {
        var res = 0
        val groupAnswers = mutableListOf<String>()

        inputList.forEach { answer ->
            if (answer.isEmpty()) {
                res += getSameAnswers(groupAnswers).length
                groupAnswers.clear()
            } else {
                groupAnswers.add(answer)
            }
        }

        return res + getSameAnswers(groupAnswers).length
    }

    private fun getSameAnswers(groupAnswers: List<String>): String {
        return groupAnswers.reduce { acc, s -> acc.filter { s.contains(it) } }
    }
}
