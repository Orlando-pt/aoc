package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day2Test {

    private val day = Day2()

    @Test
    fun testPartOne() {
        assertThat(day.partOne()).isEqualTo(1227775554L)
    }

    @Test
    fun testPartTwo() {
        assertThat(day.partTwo()).isEqualTo(4174379265L)
    }
}
