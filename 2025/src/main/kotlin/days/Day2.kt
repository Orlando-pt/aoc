package days

class Day2 : Day(2) {

    override fun partOne(): Any {
        val ids = getRanges()

        var ret = 0L
        ids.forEach { (r1, r2) ->
            (r1..r2).forEach { num ->
                val idString = num.toString()
                val part1 = idString.substring(0, idString.length / 2)
                val part2 = idString.substring(idString.length / 2)

                if (part1 == part2) {
                    ret += num
                }
            }
        }

        return ret
    }

    override fun partTwo(): Any {
        val ids = getRanges()

        var ret = 0L
        ids.forEach { (r1, r2) ->
            (r1..r2).forEach { num ->
                val idString = num.toString()

                for (i in 1..idString.length / 2) {
                    if (idString.length % i != 0) {
                        continue
                    }

                    val subStrings = (0..idString.length - 1 step i).map { idString.substring(it, it + i) }

                    if (subStrings.distinct().size == 1) {
                        ret += num
                        break
                    }
                }
            }
        }

        return ret
    }

    private fun getRanges(): List<Pair<Long, Long>> {
        return inputList.first().split(",").map {
            val (r1, r2) = it.split("-")
            Pair(r1.toLong(), r2.toLong())
        }
    }
}
