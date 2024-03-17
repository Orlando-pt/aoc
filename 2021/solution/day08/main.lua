local utils = require "util.utils"
local inspect = require "inspect"
local Solution = {}

function Solution.part1(lines)
  local validSegmentLevel = {2, 3, 4, 7}
  local res = 0

  for _, line in ipairs(lines) do
    local segments = utils.splitStr(line, "|")

    local outValues = utils.splitStr(segments[2], " ")
    for _, number in ipairs(outValues) do
      if utils.hasValue(validSegmentLevel, number:len()) then
        res = res + 1
      end
    end
  end

  -- print(inspect(segments))
  -- print(inspect(outValues))
  return res
end

function Solution.part2(lines)
  return 0
end

return Solution
