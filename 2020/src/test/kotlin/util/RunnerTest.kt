package util

import com.github.stefanbirkner.systemlambda.SystemLambda.tapSystemErr
import com.github.stefanbirkner.systemlambda.SystemLambda.tapSystemOut
import org.assertj.core.api.Assertions.assertThat
import org.junit.jupiter.api.Test
import kotlin.time.ExperimentalTime

@OptIn(ExperimentalTime::class)
class RunnerTest {

    @Test
    fun testPrintDay() {
        val day1 = tapSystemOut { Runner.main(arrayOf("1")) }
        assertThat(day1).matches(
            """
            
            === DAY 1 ===
            Part 1: \d+\s+\(.*\)
            Part 2: \d+\s+\(.*\)
            
            """.trimIndent()
        )
    }

    @Test
    fun testPrintErrors() {
        val dayNotAString = tapSystemErr { Runner.main(arrayOf("one")) }
        assertThat(dayNotAString).isEqualTo("\n=== ERROR ===\nDay argument must be an integer\n")

        val dayNotExists = tapSystemErr { Runner.main(arrayOf("99")) }
        assertThat(dayNotExists).isEqualTo("\n=== ERROR ===\nDay 99 not found\n")
    }
}