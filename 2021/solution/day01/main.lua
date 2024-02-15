local Solution = {}

function Solution.part1(lines)
  local last = math.maxinteger
  local measurements = 0

  for _, line in ipairs(lines) do
    local measure = tonumber(line)

    if measure then
      if measure > last then
        measurements = measurements + 1
      end
      last = measure
    end
  end

  return measurements
end

function Solution.part2(lines)
  local bucketMeasurements = {}

  for i, line in ipairs(lines) do
    if i + 2 > #lines then
      break
    end

    bucketMeasurements[i] = tonumber(line)

    for j = 1, 2 do
      local idx = i + j
      if not lines[idx] then
        break
      end
      bucketMeasurements[i] = bucketMeasurements[i] + tonumber(lines[idx])
    end
  end

  local measurements = 0
  for i = 1, #bucketMeasurements - 1 do
    if bucketMeasurements[i] < bucketMeasurements[i + 1] then
      measurements = measurements + 1
    end
  end

  return measurements
end

return Solution
