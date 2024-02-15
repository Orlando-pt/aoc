local Reader = {}

function Reader.read_file(day, test)
  local filename
  if test then
    filename = "test/day" .. day .. "/input.txt"
  else
    filename = "solution/day" .. day .. "/input.txt"
  end

  local file = io.open(filename, "r")

  if file == nil then
    print("File not found: " .. filename)
    return nil
  end

  local lines = {}
  for line in file:lines() do
    table.insert(lines, line)
  end

  file:close()
  return lines
end

return Reader
