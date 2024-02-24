local reader = require("internal.reader")

local function main()
  if arg[1] == nil then
    print("No day specified. Run with 'lua main.lua <day>' where <day> is the day number.")
    return
  end

  local dayStr = string.format("%02d", arg[1])

  local status, day = pcall(require, "solution.day" .. dayStr .. ".main")
  if not status then
    print("No solution for day " .. arg[1])
    return
  end

  local lines = reader.read_file(dayStr)

  print("Day " .. arg[1] .. ":")
  print("-----------------")

  local start = os.clock()

  print("Part 1:  " .. day.part1(lines))
  print("Time:    " .. os.clock() - start .. "s")

  print("-----------------")

  start = os.clock()

  print("Part 2:  " .. day.part2(lines))
  print("Time:    " .. os.clock() - start .. "s")
  print("-----------------")
end

main()
