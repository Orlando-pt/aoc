package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day7Test {

    private val day = Day7()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(21)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(40L)
    }
}
