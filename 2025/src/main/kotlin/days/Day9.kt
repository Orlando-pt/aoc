package days

import kotlin.math.abs

class Day9 : Day(9) {

    override fun partOne(): Any {
        val coordinates = inputList.map {
            val (x, y) = it.split(",")
            Point(x.toLong(), y.toLong())
        }

        var maxArea = 0L
        for (i in coordinates.indices) {
            val point = coordinates[i]

            for (j in i + 1..coordinates.lastIndex) {
                val otherPoint = coordinates[j]
                val area = point.area(otherPoint)

                if (area > maxArea) maxArea = area
            }
        }

        return maxArea
    }

    override fun partTwo(): Any {
        val coordinates = inputList.map {
            val (x, y) = it.split(",")
            Point(x.toLong(), y.toLong())
        }

        val edges = mutableListOf<Pair<Point, Point>>()
        val sizes = mutableListOf<Triple<Long, Point, Point>>()

        for (i in coordinates.indices) {
            val s = getSortedEdge(coordinates[i], coordinates.getOrElse(i - 1) { coordinates.last() })
            edges.add(s)
            for (j in i + 1..coordinates.lastIndex) {
                val (c1, c2) = getSortedEdge(coordinates[i], coordinates[j])
                sizes.add(Triple(c1.area(c2), c1, c2))
            }
        }

        edges.sortByDescending { it.first.area(it.second) }
        sizes.sortByDescending { it.first }

        for (size in sizes) {
            val (area, p1, p2) = size
            val (x1, y1) = p1
            val (x2, y2) = p2

            val (ys1, ys2) = listOf(y1, y2).sorted()

            if (!edges.any { (p3, p4) -> p4.x > x1 && p3.x < x2 && p4.y > ys1 && p3.y < ys2 }) {
                return area
            }
        }

        return 0
    }

    private fun getSortedEdge(p1: Point, p2: Point): Pair<Point, Point> {
        return if (p1.x < p2.x) Pair(p1, p2)
        else if (p1.x == p2.x) if (p1.y < p2.y) Pair(p1, p2) else Pair(p2, p1)
        else Pair(p2, p1)
    }

    data class Point(val x: Long, val y: Long) {
        fun area(point: Point): Long = (abs(x - point.x) + 1) * (abs(y - point.y) + 1)
    }
}
