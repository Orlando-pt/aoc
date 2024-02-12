package aoc.day01;

import static org.junit.jupiter.api.Assertions.assertEquals;

import aoc.internal.Reader;
import java.util.List;
import org.junit.jupiter.api.Test;

public class Day03Test {

  @Test
  public void testPart1() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("03"));

    String result = new Day03().part1(input);

    assertEquals("157", result);
  }

  @Test
  public void testPart2() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("03"));

    String result = new Day03().part2(input);

    assertEquals("70", result);
  }
}
