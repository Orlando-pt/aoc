package days

import kotlin.math.pow
import kotlin.math.sqrt

class Day8(val maxConnections: Int = 1000) : Day(8) {

    override fun partOne(): Any {
        val coordinates =
            inputList.map { line -> line.split(",").map { it.toInt() } }.map { Coordinate(it[0], it[1], it[2]) }
        val sortedDistances = getSortedDistances(coordinates)

        val circuits = mutableListOf<MutableSet<Coordinate>>()
        coordinates.forEach { circuits.add(mutableSetOf(it)) }

        var n = maxConnections
        while (n > 0) {
            n -= 1

            val (c1, c2) = sortedDistances.removeFirst()
            val i1 = circuits.indexOfFirst { it.contains(c1) }
            val i2 = circuits.indexOfFirst { it.contains(c2) }

            if (i1 != i2) {
                circuits[i1].addAll(circuits[i2])
                circuits.removeAt(i2)
            }
        }

        return circuits.sortedByDescending { it.size }.take(3).fold(1) { acc, circuit -> acc * circuit.size }
    }

    override fun partTwo(): Any {
        val coordinates =
            inputList.map { line -> line.split(",").map { it.toInt() } }.map { Coordinate(it[0], it[1], it[2]) }
        val sortedDistances = getSortedDistances(coordinates)

        val circuits = mutableListOf<MutableSet<Coordinate>>()
        coordinates.forEach { circuits.add(mutableSetOf(it)) }

        while (true) {

            val (c1,c2) = sortedDistances.removeFirst()

            val i1 = circuits.indexOfFirst { it.contains(c1) }
            val i2 = circuits.indexOfFirst { it.contains(c2) }

            if (i1 != i2) {
                circuits[i1].addAll(circuits[i2])
                circuits.removeAt(i2)
            }

            if (circuits.size == 1) return c1.x.toLong() * c2.x.toLong()
        }
    }

    private fun getSortedDistances(coordinates: List<Coordinate>): ArrayDeque<Triple<Coordinate, Coordinate, Double>> {
        val distances = mutableListOf<Triple<Coordinate, Coordinate, Double>>()

        for (i in coordinates.indices) {
            val coordinate = coordinates[i]
            for (j in i + 1..coordinates.lastIndex) {
                val other = coordinates[j]
                distances.add(Triple(coordinate, other, coordinate.distanceTo(other)))
            }
        }

        return ArrayDeque(distances.sortedBy { it.third })
    }

    data class Coordinate(val x: Int, val y: Int, val z: Int) {
        fun distanceTo(other: Coordinate): Double {
            val sum =
                (x - other.x).toDouble().pow(2) + (y - other.y).toDouble().pow(2) + (z - other.z).toDouble().pow(2)
            return sqrt(sum)
        }
    }
}
