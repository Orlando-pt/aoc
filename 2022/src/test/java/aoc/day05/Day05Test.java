package aoc.day05;

import static org.junit.jupiter.api.Assertions.assertEquals;

import aoc.internal.Reader;
import java.util.List;
import org.junit.jupiter.api.Test;

public class Day05Test {

  @Test
  public void testPart1() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("05"));

    String result = new Day05().part1(input);

    assertEquals("CMZ", result);
  }

  @Test
  public void testPart2() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("05"));

    String result = new Day05().part2(input);

    assertEquals("MCD", result);
  }
}
