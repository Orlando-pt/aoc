local Utils = {}

function Utils.splitStr(inputstr, sep)
  if sep == nil then
    sep = "%s"
  end
  local t = {}
  for str in string.gmatch(inputstr, "([^" .. sep .. "]+)") do
    table.insert(t, str)
  end
  return t
end

function Utils.tableSize(t)
  local count = 0
  for _ in pairs(t) do
    count = count + 1
  end
  return count
end

function Utils.sum(arr)
  local sum = 0
  for _, n in ipairs(arr) do
    sum = sum + n
  end

  return sum
end

return Utils
