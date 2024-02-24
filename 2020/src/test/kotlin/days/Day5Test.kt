package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day5Test {

    private val dayOne = Day5()

    @Test
    fun testPartOne() {
        assertThat(dayOne.partOne()).isEqualTo(820)
    }

    @Test
    fun testPartTwo() {
        assertThat(dayOne.partTwo()).isEqualTo(0)
    }
}
