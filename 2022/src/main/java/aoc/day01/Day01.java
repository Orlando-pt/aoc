package aoc.day01;

import aoc.Day;

import java.util.Arrays;
import java.util.List;

public class Day01 implements Day {

    @Override
    public String part1(List<String> input) {
        int ret = 0;

        int currentElfCalories = 0;
        for (String line : input) {
            if (!line.isBlank()) {
                currentElfCalories += Integer.parseInt(line);

                if (currentElfCalories > ret) {
                    ret = currentElfCalories;
                }
            } else {
                currentElfCalories = 0;
            }

        }
        return String.valueOf(ret);
    }

    @Override
    public String part2(List<String> input) {
        int[] ret = new int[3];

        int currentElfCalories = 0;
        for (String line : input) {
            if (!line.isBlank()) {
                currentElfCalories += Integer.parseInt(line);
            } else {
                updateElfCalories(ret, currentElfCalories);
                currentElfCalories = 0;
            }
        }
        updateElfCalories(ret, currentElfCalories);
        return String.valueOf(Arrays.stream(ret).sum());
    }

    private void updateElfCalories(int[] elfCalories, int calories) {
        // get the index of the elf with less calories
        int lowestIndex = 0;
        for (int i = 1; i < elfCalories.length; i++)
            if (elfCalories[i] < elfCalories[lowestIndex]) {
                lowestIndex = i;
            }

        if (elfCalories[lowestIndex] < calories) {
            elfCalories[lowestIndex] = calories;
        }
    }

}
