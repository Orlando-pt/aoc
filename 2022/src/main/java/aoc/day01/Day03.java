package aoc.day01;

import aoc.Day;
import java.util.ArrayList;
import java.util.List;

public class Day03 implements Day {
  @Override
  public String part1(List<String> input) {
    var ruckSacks =
        input.stream()
            .map(
                s -> {
                  var c1 = s.substring(0, s.length() / 2);
                  var c2 = s.substring(s.length() / 2);
                  return new RuckSack(c1, c2);
                })
            .toList();

    return String.valueOf(ruckSacks.stream().mapToInt(RuckSack::charValue).sum());
  }

  @Override
  public String part2(List<String> input) {
    var ruckSacks = new ArrayList<RuckSackPart2>();

    for (int i = 0; i < input.size(); i += 3) {
      ruckSacks.add(
          new RuckSackPart2(new String[] {input.get(i), input.get(i + 1), input.get(i + 2)}));
    }

    return String.valueOf(ruckSacks.stream().mapToInt(RuckSackPart2::charValue).sum());
  }

  record RuckSackPart2(String[] sacks) {
    public int charValue() {
      var commonChar = commonChar();
      var res = commonChar - 'a' + 1;

      if (res < 0) {
        res = commonChar - 'A' + 27;
      }

      return res;
    }

    private char commonChar() {
      for (char c : sacks[0].toCharArray()) {
        if (sacks[1].contains(String.valueOf(c)) && sacks[2].contains(String.valueOf(c))) return c;
      }

      return ' ';
    }
  }

  record RuckSack(String c1, String c2) {
    public int charValue() {
      var commonChar = commonChar();
      var res = commonChar - 'a' + 1;

      if (res < 0) {
        res = commonChar - 'A' + 27;
      }

      return res;
    }

    private char commonChar() {
      for (char c : c1.toCharArray()) {
        if (c2.contains(String.valueOf(c))) {
          return c;
        }
      }

      return ' ';
    }
  }
}
