local day01 = require("solution.day01.day01")
local day02 = require("solution.day02.day02")
local reader = require("internal.reader")

local solutions = {}
solutions["1"] = day01
solutions["2"] = day02

local function main()
  if arg[1] == nil then
    print("No day specified. Run with 'lua main.lua <day>' where <day> is the day number.")
    return
  end

  local solution = solutions[arg[1]]
  if solution == nil then
    print("No solution for day " .. arg[1])
    return
  end

  local lines = reader.read_file("01")
  print("Day " .. arg[1] .. ":")
  print("-----------------")
  local start = os.clock()
  print("Part 1: ", solution.part1(lines))
  print("Time: ", os.clock() - start)
  print("-----------------")
  start = os.clock()
  print("Part 2: ", solution.part2(lines))
  print("Time: ", os.clock() - start)
  print("-----------------")
end

main()
