package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day9Test {

    private val day = Day9()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(50L)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(24L)
    }
}
