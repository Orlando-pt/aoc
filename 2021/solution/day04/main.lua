local utils = require "util.utils"

local Solution = {}

local function addBool(boardLine)
  local boolLine = {}
  for _, n in ipairs(boardLine) do
    table.insert(boolLine, { n, false })
  end

  return boolLine
end

local function parse(lines)
  local numbers = utils.splitStr(lines[1], ",")

  local boards = {}
  local i = 2

  while i < #lines do
    i = i + 1
    local line = lines[i]
    local board = {}

    while line and line ~= "" do
      local lineParts = utils.splitStr(line, " ")
      table.insert(board, addBool(lineParts))

      i = i + 1
      line = lines[i]
    end

    if board[2] then
      table.insert(boards, board)
    end
  end

  return numbers, boards
end

local function updateBoards(boards, drawnNumber)
  local boardsWithNumber = {}
  for i, board in ipairs(boards) do
    local hasNumber = false
    for x, line in ipairs(board) do
      for y, cell in ipairs(line) do
        if cell[1] == drawnNumber then
          cell[2] = true
          hasNumber = true
          boardsWithNumber["" .. i] = { x, y }
          break
        end
      end
      if hasNumber then
        break
      end
    end
  end

  return boardsWithNumber
end

local function checkRow(row)
  for _, cell in ipairs(row) do
    if not cell[2] then
      return false
    end
  end

  return true
end

local function checkColumn(board, y)
  for _, line in ipairs(board) do
    if not line[y][2] then
      return false
    end
  end

  return true
end

local function hasBingo(boards, boardsWithNumber)
  local winners = {}
  for k, boardCoord in pairs(boardsWithNumber) do
    local boardNumber = tonumber(k)
    local x, y = boardCoord[1], boardCoord[2]

    -- check row
    if checkRow(boards[boardNumber][x]) then
      table.insert(winners, boardNumber)
      goto continue
    end

    -- check column
    if checkColumn(boards[boardNumber], y) then
      table.insert(winners, boardNumber)
    end
    ::continue::
  end

  return winners
end

local function sumUnmarked(board)
  local sum = 0
  for _, line in ipairs(board) do
    for _, cell in ipairs(line) do
      if not cell[2] then
        sum = sum + cell[1]
      end
    end
  end

  return sum
end

function Solution.part1(lines)
  local numbers, boards = parse(lines)

  for _, number in ipairs(numbers) do
    local luckyBoards = updateBoards(boards, number)

    local winners = hasBingo(boards, luckyBoards)

    if #winners ~= 0 then
      return sumUnmarked(boards[winners[1]]) * number
    end
  end

  return 0
end

function Solution.part2(lines)
  local numbers, boards = parse(lines)

  local winners = {}
  local lastWinnerNumber = 0
  local lastBoardNumber = 0

  for _, number in ipairs(numbers) do
    local luckyBoards = updateBoards(boards, number)

    local roundWinners = hasBingo(boards, luckyBoards)

    if #roundWinners ~= 0 then
      for _, boardNumber in ipairs(roundWinners) do
        local alreadyWon = winners[boardNumber] or false

        if not alreadyWon then
          winners[boardNumber] = true
          lastWinnerNumber = number
          lastBoardNumber = boardNumber
        end
      end
    end

    if utils.tableSize(winners) == #boards then
      break
    end
  end

  return sumUnmarked(boards[lastBoardNumber]) * lastWinnerNumber
end

return Solution
