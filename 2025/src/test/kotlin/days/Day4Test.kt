package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day4Test {

    private val day = Day4()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(13)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(43)
    }
}
