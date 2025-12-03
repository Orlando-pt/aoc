package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day3Test {

    private val day = Day3()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(357)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(3121910778619L)
    }
}
