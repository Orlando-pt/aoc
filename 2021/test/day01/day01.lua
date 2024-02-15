local reader = require("internal.reader")
local day01 = require("solution.day01.day01")

local lines = reader.read_file("01", true)

local function part1Test()
  local expected = 7
  local result = day01.part1(lines)

  assert(result == expected, "Expected " .. expected .. ", but got " .. result)
end

local function part2Test()
  local expected = 5
  local result = day01.part2(lines)

  assert(result == expected, "Expected " .. expected .. ", but got " .. result)
end

part1Test()
part2Test()
