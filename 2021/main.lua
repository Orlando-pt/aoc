local reader = require("internal.reader")
local day01 = require("solution.day01.main")
local day02 = require("solution.day02.main")
local day03 = require("solution.day03.main")
local day04 = require("solution.day04.main")

local solutions = {}
solutions["1"] = day01
solutions["2"] = day02
solutions["3"] = day03
solutions["4"] = day04

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

  local lines = reader.read_file(string.format("%02d", arg[1]))
  print("Day " .. arg[1] .. ":")
  print("-----------------")
  local start = os.clock()
  print("Part 1:  " .. solution.part1(lines))
  print("Time:    " .. os.clock() - start .. "s")
  print("-----------------")
  start = os.clock()
  print("Part 2:  " .. solution.part2(lines))
  print("Time:    " .. os.clock() - start .. "s")
  print("-----------------")
end

main()
