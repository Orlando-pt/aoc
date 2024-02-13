package aoc.day04;

import aoc.Day;
import java.util.List;

public class Day04 implements Day {
  @Override
  public String part1(List<String> input) {
    var cleanupPairs = parseInput(input);

    return String.valueOf(cleanupPairs.stream().mapToInt(p -> p.pairContains() ? 1 : 0).sum());
  }

  @Override
  public String part2(List<String> input) {
    var cleanupPairs = parseInput(input);
    return String.valueOf(cleanupPairs.stream().mapToInt(p -> p.pairsOverlap() ? 1 : 0).sum());
  }

  private List<CleanupPair> parseInput(List<String> input) {
    return input.stream()
        .map(
            s -> {
              var parts = s.split(",");
              var p1Parts = parts[0].split("-");
              var p2Parts = parts[1].split("-");

              return new CleanupPair(
                  new int[] {Integer.parseInt(p1Parts[0]), Integer.parseInt(p1Parts[1])},
                  new int[] {Integer.parseInt(p2Parts[0]), Integer.parseInt(p2Parts[1])});
            })
        .toList();
  }

  record CleanupPair(int[] p1, int[] p2) {
    public boolean pairContains() {
      return (p1[0] <= p2[0] && p1[1] >= p2[1]) || (p2[0] <= p1[0] && p2[1] >= p1[1]);
    }

    public boolean pairsOverlap() {
      return (p1[0] <= p2[0] && p1[1] >= p2[0]) || (p2[0] <= p1[0] && p2[1] >= p1[0]);
    }
  }
}
