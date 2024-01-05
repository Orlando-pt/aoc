package aoc.internal;

import java.io.BufferedReader;
import java.io.File;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.UncheckedIOException;
import java.util.List;
import java.util.stream.Collectors;

public class Reader {
    public static List<String> readFile(InputStream inputStream) {
        try (BufferedReader r = new BufferedReader(new InputStreamReader(inputStream))) {
            return r.lines().collect(Collectors.toList());
        } catch (IOException e) {
            throw new UncheckedIOException(e);
        }
    }

    public static InputStream getTestResourceAsStream(String day) {
        try {
            return new FileInputStream(new File("src/test/java/aoc/resources/day" + day + ".txt"));
        } catch (FileNotFoundException e) {
            throw new UncheckedIOException(e);
        }
    }
}
