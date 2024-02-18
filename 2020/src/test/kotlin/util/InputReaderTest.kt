package util

import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test

class InputReaderTest {

    @Test
    fun testReadInputAsString() {
        val testInputAsString = InputReader.getInputAsString(1)
        assertThat(testInputAsString).isEqualTo("1721\n979\n366\n299\n675\n1456")
    }

    @Test
    fun testReadInputAsList() {
        val testInputAsList = InputReader.getInputAsList(1)
        assertThat(testInputAsList).contains("1721", "979", "366", "299", "675", "1456")
    }
}
