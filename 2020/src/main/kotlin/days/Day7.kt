package days

class Day7 : Day(7) {
    override fun partOne(): Any {
        val bags = getBags()

        return bags.count { canContainShinyGold(bags, it) }
    }

    override fun partTwo(): Any {
        val bags = getBags()

        return countBags(bags, bags.find { it.color == SHINY_GOLD })
    }

    private fun countBags(
        bags: List<Bag>,
        bag: Bag?,
    ): Int {
        if (bag == null || bag.contains.isEmpty()) {
            return 0
        }

        return bag.contains.sumOf { it.first + it.first * countBags(bags, bags.find { bag -> bag.color == it.second }) }
    }

    private fun getBags(): List<Bag> {
        return inputList.map {
            val (color, contains) = it.split(" bags contain ")
            val bags =
                contains.split(", ").map { inBag ->
                    val (count, inColor) = inBag.split(" ", limit = 2)
                    if (count != "no") {
                        Pair(count.toInt(), removeUnnecessaryWords(inColor))
                    } else {
                        Pair(0, "")
                    }
                }

            Bag(removeUnnecessaryWords(color), bags)
        }
    }

    private fun canContainShinyGold(
        bags: List<Bag>,
        bag: Bag?,
    ): Boolean {
        if (bag == null || bag.contains.isEmpty()) {
            return false
        }

        if (bag.contains.any { it.second == SHINY_GOLD }) {
            return true
        }

        return bag.contains.any { canContainShinyGold(bags, bags.find { bag -> bag.color == it.second }) }
    }

    private fun removeUnnecessaryWords(input: String): String {
        return input.replace(" bags", "").replace(" bag", "").replace(".", "").trim()
    }

    data class Bag(val color: String, val contains: List<Pair<Int, String>>)

    companion object {
        private const val SHINY_GOLD = "shiny gold"
    }
}
