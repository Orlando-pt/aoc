package aoc.day04;

import static org.junit.jupiter.api.Assertions.assertEquals;

import aoc.internal.Reader;
import java.util.List;
import org.junit.jupiter.api.Test;

public class Day04Test {

  @Test
  public void testPart1() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("04"));

    String result = new Day04().part1(input);

    assertEquals("2", result);
  }

  @Test
  public void testPart2() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("04"));

    String result = new Day04().part2(input);

    assertEquals("4", result);
  }
}
