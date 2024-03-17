local splitStr = require("util.utils").splitStr
local max = require("util.utils").max

local Solution = {}

local function parseLine(line)
  local parts = splitStr(line, ",")
  local numbers = {}

  for _, part in ipairs(parts) do
    table.insert(numbers, tonumber(part))
  end

  return numbers
end

local function arithmeticProgression(n)
  return (n * (n + 1)) / 2
end

function Solution.part1(lines)
  local numbers = parseLine(lines[1])

  local fuel = math.maxinteger
  for _, n in ipairs(numbers) do
    local jumps = 0

    for _, nn in ipairs(numbers) do
      jumps = jumps + math.abs(n - nn)
    end

    if jumps < fuel then
      fuel = jumps
    end
  end

  return fuel
end

function Solution.part2(lines)
  local numbers = parseLine(lines[1])
  local maxNumbers = max(numbers)

  local fuel = math.maxinteger
  for i = 1, maxNumbers do
    local fuelSpent = 0

    for _, n in ipairs(numbers) do
      fuelSpent = fuelSpent + arithmeticProgression(math.abs(n - i))
    end

    if fuelSpent < fuel then
      fuel = fuelSpent
    end
  end

  return fuel
end

return Solution
