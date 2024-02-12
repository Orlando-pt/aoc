package aoc.day01;

import static org.junit.jupiter.api.Assertions.assertEquals;

import aoc.internal.Reader;
import java.util.List;
import org.junit.jupiter.api.Test;

public class Day01Test {

  @Test
  public void testPart1() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("01"));

    String result = new Day01().part1(input);

    assertEquals("24000", result);
  }

  @Test
  public void testPart2() {
    List<String> input = Reader.readFile(Reader.getTestResourceAsStream("01"));

    String result = new Day01().part2(input);

    assertEquals("45000", result);
  }
}
