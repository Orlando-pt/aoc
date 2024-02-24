package days

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class Day4Test {

    private val dayOne = Day4()

    @Test
    fun testPartOne() {
        assertThat(dayOne.partOne()).isEqualTo(2)
    }

    @Test
    fun testPartTwo() {
        assertThat(dayOne.partTwo()).isEqualTo(2)
    }
}
