local Solution = {}

function Solution.part1(lines)
  local zeros = {}
  local ones = {}

  for _, line in ipairs(lines) do
    for i = 1, #line do
      local c = line:sub(i, i)
      if c == "0" then
        zeros[i] = (zeros[i] or 0) + 1
      elseif c == "1" then
        ones[i] = (ones[i] or 0) + 1
      end
    end
  end

  local gamma = ""
  local epsilon = ""
  for i = 1, #zeros do
    if zeros[i] > ones[i] then
      gamma = gamma .. "0"
      epsilon = epsilon .. "1"
    else
      gamma = gamma .. "1"
      epsilon = epsilon .. "0"
    end
  end

  return tonumber(gamma, 2) * tonumber(epsilon, 2)
end

local function findRating(lines, comparator)
  if not lines or #lines == 0 then
    return ""
  end

  local ratings = {}
  for _, line in ipairs(lines) do
    if line:sub(1, 1) == "0" then
      local records = ratings[1] or {}
      table.insert(records, line:sub(2))

      ratings[1] = records
    elseif line:sub(1, 1) == "1" then
      local records = ratings[2] or {}
      table.insert(records, line:sub(2))

      ratings[2] = records
    else
      return ""
    end
  end

  if comparator(ratings) then
    return "0" .. findRating(ratings[1], comparator)
  end

  return "1" .. findRating(ratings[2], comparator)
end

function Solution.part2(lines)
  local function oxygenCompare(ratings)
    if not ratings[1] or not ratings[2] then
      return ratings[2] or false
    end
    return #ratings[1] > #ratings[2]
  end

  local function carbonCompare(ratings)
    if not ratings[1] or not ratings[2] then
      return ratings[1] or false
    end
    return #ratings[1] <= #ratings[2]
  end

  local oxygen = findRating(lines, oxygenCompare)
  local carbon = findRating(lines, carbonCompare)
  return tonumber(oxygen, 2) * tonumber(carbon, 2)
end

return Solution
