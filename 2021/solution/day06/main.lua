local splitStr = require("util.utils").splitStr
local sum = require("util.utils").sum
local Solution = {}

local function parseLineToNumberArr(line)
  local numbers = splitStr(line, ",")
  local result = {}

  -- init result with 0
  for _ = 1, 9 do
    table.insert(result, 0)
  end

  for i = 1, #numbers do
    local number = tonumber(numbers[i])

    if not number then
      error("Invalid number: " .. numbers[i])
    end

    result[number + 1] = result[number + 1] + 1
  end

  return result
end

local function ageOneDay(numbers)
  local dead = numbers[1]

  for i = 1, #numbers - 1 do
    numbers[i] = numbers[i + 1]
  end

  numbers[7] = numbers[7] + dead
  numbers[9] = dead
end

function Solution.part1(lines)
  local numbers = parseLineToNumberArr(lines[1])
  for _ = 1, 80 do
    ageOneDay(numbers)
  end

  return sum(numbers)
end

function Solution.part2(lines)
  local numbers = parseLineToNumberArr(lines[1])
  for _ = 1, 256 do
    ageOneDay(numbers)
  end

  return sum(numbers)
end

return Solution
