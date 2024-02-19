package days

import org.assertj.core.api.Assertions
import org.junit.jupiter.api.Test

class Day2Test {
    private val day = Day2()

    @Test
    fun testPartOne() {
        Assertions.assertThat(day.partOne()).isEqualTo(2)
    }

    @Test
    fun testPartTwo() {
        Assertions.assertThat(day.partTwo()).isEqualTo(1)
    }
}