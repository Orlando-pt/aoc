package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day10Test {
    private val day = Day10()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(220)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(19208L)
    }
}
