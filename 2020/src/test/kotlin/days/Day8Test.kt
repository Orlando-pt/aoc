package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day8Test {
    private val day = Day8()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(5)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(8)
    }
}
