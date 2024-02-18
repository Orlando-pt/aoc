package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day1Test {

    private val dayOne = Day1()

    @Test
    fun testPartOne() {
        assertThat(dayOne.partOne()).isEqualTo(514579)
    }

    @Test
    fun testPartTwo() {
        assertThat(dayOne.partTwo()).isEqualTo(241861950)
    }
}
