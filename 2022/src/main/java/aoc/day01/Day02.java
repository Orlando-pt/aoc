package aoc.day01;

import aoc.Day;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;

public class Day02 implements Day {
  private static final HashMap<String, Integer> playPoints;
  private static final HashMap<String, Integer> playPoints2;
  private static final HashMap<String, String> winnerHand;
  private static final HashMap<String, String[]> winLoseHand;
  private static final HashMap<String, String> playMappings;

  static {
    playPoints = new HashMap<>();
    playPoints.put("X", 1);
    playPoints.put("Y", 2);
    playPoints.put("Z", 3);

    playPoints2 = new HashMap<>();
    playPoints2.put("A", 1);
    playPoints2.put("B", 2);
    playPoints2.put("C", 3);

    winnerHand = new HashMap<>();
    winnerHand.put("A", "Y");
    winnerHand.put("B", "Z");
    winnerHand.put("C", "X");

    winLoseHand = new HashMap<>();
    winLoseHand.put("A", new String[] {"B", "C"});
    winLoseHand.put("B", new String[] {"C", "A"});
    winLoseHand.put("C", new String[] {"A", "B"});

    playMappings = new HashMap<>();
    playMappings.put("A", "X");
    playMappings.put("B", "Y");
    playMappings.put("C", "Z");
  }

  @Override
  public String part1(List<String> input) {

    var games = new ArrayList<Game>();

    for (var line : input) {
      var split = line.split(" ");
      games.add(new Game(split[0], split[1]));
    }

    var points = games.stream().mapToInt(Game::points).sum();

    return String.valueOf(points);
  }

  @Override
  public String part2(List<String> input) {
    var games = new ArrayList<Game>();

    for (var line : input) {
      var split = line.split(" ");
      games.add(new Game(split[0], split[1]));
    }

    var points = games.stream().mapToInt(Game::pointsPart2).sum();

    return String.valueOf(points);
  }

  record Game(String opponent, String response) {
    public int points() {
      var points = playPoints.get(response);

      if (playMappings.get(opponent).equals(response)) {
        return points + 3;
      }

      if (winnerHand.get(opponent).equals(response)) {
        return points + 6;
      }

      return points;
    }

    public int pointsPart2() {
      String play =
          switch (response) {
            case "X" -> winLoseHand.get(opponent)[1];
            case "Z" -> winLoseHand.get(opponent)[0];
            default -> opponent;
          };

      var points = playPoints2.get(play);

      return switch (response) {
        case "X" -> points;
        case "Z" -> points + 6;
        default -> points + 3;
      };
    }
  }
}
