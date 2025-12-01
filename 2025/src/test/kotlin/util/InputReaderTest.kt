package util

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class InputReaderTest {

    @Test
    fun testReadInputAsString() {
        val testInputAsString = InputReader.getInputAsString(1)
        assertThat(testInputAsString).isEqualTo("L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82")
    }

    @Test
    fun testReadInputAsList() {
        val testInputAsList = InputReader.getInputAsList(1)
        assertThat(testInputAsList).contains("L68", "L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82")
    }
}
