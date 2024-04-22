package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day6Test {

    private val day = Day6()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(11)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(6)
    }
}
