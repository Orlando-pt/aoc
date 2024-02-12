package aoc.day01;

import static org.junit.jupiter.api.Assertions.assertEquals;

import aoc.internal.Reader;
import java.util.List;
import org.junit.jupiter.api.Test;

public class Day02Test {

  @Test
  public void testPart1() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("02"));

    String result = new Day02().part1(input);

    assertEquals("15", result);
  }

  @Test
  public void testPart2() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("02"));

    String result = new Day02().part2(input);

    assertEquals("12", result);
  }
}
