local splitStr = require("util.utils").splitStr
local Solution = {}

function Solution.part1(lines)
  local horizontal = 0
  local depth = 0

  for _, line in ipairs(lines) do
    local lineParts = splitStr(line, " ")

    if lineParts[1] == "forward" then
      horizontal = horizontal + tonumber(lineParts[2])
    elseif lineParts[1] == "down" then
      depth = depth + tonumber(lineParts[2])
    elseif lineParts[1] == "up" then
      depth = depth - tonumber(lineParts[2])
    end
  end

  return horizontal * depth
end

function Solution.part2(lines)
  local horizontal = 0
  local depth = 0
  local aim = 0

  for _, line in ipairs(lines) do
    local lineParts = splitStr(line, " ")

    if lineParts[1] == "forward" then
      horizontal = horizontal + tonumber(lineParts[2])
      depth = depth + (aim * tonumber(lineParts[2]))
    elseif lineParts[1] == "down" then
      aim = aim + tonumber(lineParts[2])
    elseif lineParts[1] == "up" then
      aim = aim - tonumber(lineParts[2])
    end

  end
  return horizontal * depth
end

return Solution
