local splitStr = require("util.utils").splitStr
local Solution = {}

local function initMatrix(x, y)
  local matrix = {}
  for i = 1, y do
    matrix[i] = {}
    for j = 1, x do
      matrix[i][j] = 0
    end
  end
  return matrix
end

local function coordsAreHorizontal(c1, c2)
  return c1[2] == c2[2]
end

local function coordsAreVert(c1, c2)
  return c1[1] == c2[1]
end

local function sortCoords(c1, c2)
  if c1 > c2 then
    return c2, c1
  end
  return c1, c2
end

local function drawHorizontal(matrix, c1, c2)
  local x1, x2 = sortCoords(c1[1], c2[1])
  for i = x1, x2 do
    matrix[c1[2] + 1][i + 1] = matrix[c1[2] + 1][i + 1] + 1
  end
end

local function drawVertical(matrix, c1, c2)
  local y1, y2 = sortCoords(c1[2], c2[2])
  for i = y1, y2 do
    matrix[i + 1][c1[1] + 1] = matrix[i + 1][c1[1] + 1] + 1
  end
end

local function findNrIntersections(matrix)
  local sum = 0
  for y = 1, #matrix do
    for x = 1, #matrix[y] do
      if matrix[y][x] > 1 then
        sum = sum + 1
      end
    end
  end
  return sum
end

local function parseInput(lines)
  local x, y = 0, 0
  local coords = {}

  for _, line in ipairs(lines) do
    local lineParts = splitStr(line, " ")
    local c1 = splitStr(lineParts[1], ",")
    local c2 = splitStr(lineParts[3], ",")

    local coord1 = { tonumber(c1[1]), tonumber(c1[2]) }
    local coord2 = { tonumber(c2[1]), tonumber(c2[2]) }

    if coord1[1] > x then
      x = coord1[1]
    end
    if coord2[1] > x then
      x = coord2[1]
    end
    if coord1[2] > y then
      y = coord1[2]
    end
    if coord2[2] > y then
      y = coord2[2]
    end

    table.insert(coords, { coord1, coord2 })
  end

  return x + 1, y + 1, coords
end

local function coordsAreDiagonal(c1, c2)
  return math.abs(c1[1] - c2[1]) == math.abs(c1[2] - c2[2])
end

local function drawDiagonal(matrix, c1, c2)
  if c1[1] < c2[1] then
    if c1[2] < c2[2] then
      for i = 0, c2[1] - c1[1] do
        matrix[c1[2] + i + 1][c1[1] + i + 1] = matrix[c1[2] + i + 1][c1[1] + i + 1] + 1
      end
    else
      for i = 0, c2[1] - c1[1] do
        matrix[c1[2] - i + 1][c1[1] + i + 1] = matrix[c1[2] - i + 1][c1[1] + i + 1] + 1
      end
    end
  else
    if c1[2] < c2[2] then
      for i = 0, c1[1] - c2[1] do
        matrix[c1[2] + i + 1][c1[1] - i + 1] = matrix[c1[2] + i + 1][c1[1] - i + 1] + 1
      end
    else
      for i = 0, c1[1] - c2[1] do
        matrix[c1[2] - i + 1][c1[1] - i + 1] = matrix[c1[2] - i + 1][c1[1] - i + 1] + 1
      end
    end
  end
end

function Solution.part1(lines)
  local x, y, coords = parseInput(lines)
  local matrix = initMatrix(x, y)

  for _, line in ipairs(coords) do
    local coord1 = line[1]
    local coord2 = line[2]
    local horizontal = coordsAreHorizontal(coord1, coord2)
    local vertical = coordsAreVert(coord1, coord2)

    if horizontal then
      drawHorizontal(matrix, coord1, coord2)
    elseif vertical then
      drawVertical(matrix, coord1, coord2)
    end
  end

  return findNrIntersections(matrix)
end

function Solution.part2(lines)
  local x, y, coords = parseInput(lines)
  local matrix = initMatrix(x, y)

  for _, line in ipairs(coords) do
    local coord1 = line[1]
    local coord2 = line[2]

    if coordsAreHorizontal(coord1, coord2) then
      drawHorizontal(matrix, coord1, coord2)
    elseif coordsAreVert(coord1, coord2) then
      drawVertical(matrix, coord1, coord2)
    elseif coordsAreDiagonal(coord1, coord2) then
      drawDiagonal(matrix, coord1, coord2)
    end
  end

  return findNrIntersections(matrix)
end

return Solution
