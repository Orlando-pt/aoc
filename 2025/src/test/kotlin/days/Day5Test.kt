package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day5Test {

    private val day = Day5()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(3)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(14L)
    }
}
