package aoc.day05;

import aoc.Day;
import java.util.*;

public class Day05 implements Day {
  @Override
  public String part1(List<String> input) {
    var supplyStacks = parseCrateStack(input, false);
    supplyStacks.makeMoves(input);

    return supplyStacks.getTopCrates();
  }

  @Override
  public String part2(List<String> input) {
    var supplyStacks = parseCrateStack(input, true);
    supplyStacks.makeMoves(input);

    return supplyStacks.getTopCrates();
  }

  private SupplyStacks parseCrateStack(List<String> input, boolean part2) {
    SupplyStacks supplyStacks = new SupplyStacks(part2);

    var stacksInput = new ArrayList<String>();
    int i = 0;
    String currentInput = input.get(0);

    while (!currentInput.isBlank()) {
      stacksInput.add(currentInput);

      currentInput = input.get(++i);
    }
    var indexes = stacksInput.remove(stacksInput.size() - 1);

    supplyStacks.addStacks(stacksInput, indexes);
    supplyStacks.setMovesIdx(i);
    return supplyStacks;
  }

  static class SupplyStacks {
    private final HashMap<Integer, Stack<Character>> stacks;
    private int movesIdx;
    private final boolean part2;

    public SupplyStacks(boolean part2) {
      stacks = new HashMap<>();
      movesIdx = 0;
      this.part2 = part2;
    }

    public void setMovesIdx(int movesIdx) {
      this.movesIdx = movesIdx;
    }

    public void makeMoves(List<String> input) {
      var currentIdx = movesIdx;
      while (currentIdx < input.size() - 1) {
        var move = input.get(++currentIdx);
        var moveParts = move.split(" ");

        if (!part2) move(moveParts[1], moveParts[3], moveParts[5]);
        else movePart2(moveParts[1], moveParts[3], moveParts[5]);
      }
    }

    public String getTopCrates() {
      var result = new StringBuilder();

      for (var stack : stacks.values()) {
        result.append(stack.peek());
      }

      return result.toString();
    }

    public void move(String crateNumber, String from, String to) {
      var crates = Integer.parseInt(crateNumber);
      var fromStack = stacks.get(Integer.parseInt(from));
      var toStack = stacks.get(Integer.parseInt(to));

      for (int i = 0; i < crates; i++) {
        var crate = fromStack.pop();
        toStack.push(crate);
      }
    }

    public void movePart2(String crateNumber, String from, String to) {
      var crates = Integer.parseInt(crateNumber);
      var fromStack = stacks.get(Integer.parseInt(from));
      var toStack = stacks.get(Integer.parseInt(to));

      var tempStack = new Stack<Character>();
      for (int i = 0; i < crates; i++) {
        var crate = fromStack.pop();
        tempStack.push(crate);
      }

      for (int i = 0; i < crates; i++) {
        var crate = tempStack.pop();
        toStack.push(crate);
      }
    }

    public void addStacks(List<String> stacksList, String indexes) {
      var indexesPositions =
          Arrays.stream(indexes.split(" "))
              .filter(s -> !s.isBlank())
              .map(indexes::indexOf)
              .toList();

      initStacks(indexesPositions.size());

      for (int i = stacksList.size() - 1; i >= 0; i--) {
        for (int j = 0; j < indexesPositions.size(); j++) {
          var stackString = stacksList.get(i);
          var pos = indexesPositions.get(j);

          if (pos >= stackString.length()) break;

          var crate = stackString.charAt(pos);

          if (crate != ' ') getStack(j).push(crate);
        }
      }
    }

    private Stack<Character> getStack(int index) {
      return stacks.get(index + 1);
    }

    private void initStacks(int size) {
      for (int i = 1; i <= size; i++) {
        stacks.put(i, new Stack<>());
      }
    }
  }
}
