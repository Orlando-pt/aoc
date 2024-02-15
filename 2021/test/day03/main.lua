local dayStr = "03"

local reader = require("internal.reader")
local day = require("solution.day"..dayStr ..".main")

local lines = reader.read_file(dayStr, true)

local function part1Test()
  local expected = 198
  local result = day.part1(lines)

  assert(result == expected, "Expected " .. expected .. ", but got " .. result)
end

local function part2Test()
  local expected = 230
  local result = day.part2(lines)

  assert(result == expected, "Expected " .. expected .. ", but got " .. result)
end

part1Test()
part2Test()
