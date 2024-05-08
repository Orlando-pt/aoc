package days

class Day10 : Day(10) {
    override fun partOne(): Any {
        val adapterChain = getAdapterChain()

        val ones =
            (1..adapterChain.lastIndex).fold(0) { acc, i ->
                if (adapterChain[i] - adapterChain[i - 1] == 1) acc + 1 else acc
            }
        val threes =
            (1..adapterChain.lastIndex).fold(0) { acc, i ->
                if (adapterChain[i] - adapterChain[i - 1] == 3) acc + 1 else acc
            }

        return ones * threes
    }

    override fun partTwo(): Any {
        val chain = getAdapterChain()
        val paths = mutableMapOf(0 to 1L)

        chain.drop(1).forEach { adapter ->
            paths[adapter] = (1..3).sumOf { paths.getOrDefault(adapter - it, 0) }
        }

        return paths[chain.last()] ?: 0
    }

    private fun getAdapterChain(): List<Int> {
        val numbers = inputList.map { it.toInt() }.toMutableList()
        var currentAdapter = 0
        val chain = mutableListOf(currentAdapter)

        while (numbers.isNotEmpty()) {
            currentAdapter = getPossibleAdapters(numbers, currentAdapter).min()

            numbers.remove(currentAdapter)
            chain.add(currentAdapter)
        }

        chain.add(currentAdapter + 3)
        return chain
    }

    private fun getPossibleAdapters(
        numbers: List<Int>,
        currentAdapter: Int,
    ) = numbers.filter { (1..3).contains(it - currentAdapter) }
}
