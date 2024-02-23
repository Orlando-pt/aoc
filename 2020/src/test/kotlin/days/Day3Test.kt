package days

import org.assertj.core.api.Assertions
import org.junit.jupiter.api.Test

class Day3Test {
    private val day = Day3()

    @Test
    fun testPartOne() {
        Assertions.assertThat(day.partOne()).isEqualTo(7L)
    }

    @Test
    fun testPartTwo() {
        Assertions.assertThat(day.partTwo()).isEqualTo(336L)
    }
}