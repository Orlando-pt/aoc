package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day1Test {

    private val day = Day1()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(514579)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(241861950)
    }
}
