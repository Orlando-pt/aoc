package day09

import (
	"github.com/orlando-pt/aoc/2024/utils"
)

const FREESPACE = -1

type DiskFree struct {
	start int
	size  int
}

type DiskFile struct {
	id    int
	start int
	size  int
}

func Part1(lines []string) (sum int) {
	diskLayout := parseInput(lines)

	left, right := 0, len(diskLayout)-1

	for left <= right {
		for diskLayout[left] != FREESPACE {
			left++
		}

		for diskLayout[right] == FREESPACE {
			right--
		}

		if left < right {
			diskLayout[left], diskLayout[right] = diskLayout[right], FREESPACE
		}
	}
	utils.Debug(diskLayout)

	for i := 0; diskLayout[i] != FREESPACE; i++ {
		sum += i * diskLayout[i]
	}
	return
}

func Part2(lines []string) (sum int) {
	diskLayout := parseInput(lines)
	files, freeSpace := getFilesAndFreeSpace(diskLayout)

	utils.Debug("Disk layout: ", diskLayout)
	utils.Debug("Files: ", files)
	utils.Debug("Free space: ", freeSpace)

	for i := len(files) - 1; i >= 0; i-- {
		file := files[i]
		for j := range freeSpace {
			if freeSpace[j].size >= file.size && freeSpace[j].start <= file.start {
				for k := 0; k < file.size; k++ {
					diskLayout[freeSpace[j].start+k], diskLayout[file.start+k] = file.id, FREESPACE
				}

				freeSpace[j].start += file.size
				freeSpace[j].size -= file.size
				break
			}
		}
	}

	for i := 0; i < len(diskLayout); i++ {
		if diskLayout[i] != FREESPACE {
			sum += i * diskLayout[i]
		}
	}
	return
}

func parseInput(lines []string) (diskLayout []int) {
	fileCounter := 0
	for i, nString := range lines[0] {
		n := utils.StrToInt(string(nString))

		if i%2 == 0 {
			diskLayout = append(diskLayout, getRepeatedArr(fileCounter, n)...)
			fileCounter++
		} else {
			diskLayout = append(diskLayout, getRepeatedArr(FREESPACE, n)...)
		}
	}

	utils.Debug(diskLayout)
	return
}

func getFilesAndFreeSpace(diskLayout []int) (files []DiskFile, freeSpace []DiskFree) {
	for i := 0; i < len(diskLayout); i++ {
		size := 0
		if diskLayout[i] == FREESPACE {
			for j := i; j < len(diskLayout) && diskLayout[j] == FREESPACE; j++ {
				size++
			}
			freeSpace = append(freeSpace, DiskFree{i, size})
		} else {
			for j := i; j < len(diskLayout) && diskLayout[j] == diskLayout[i]; j++ {
				size++
			}
			files = append(files, DiskFile{diskLayout[i], i, size})
		}
		i += size - 1
	}
	return
}

func getRepeatedArr(val, n int) (result []int) {
	result = make([]int, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, val)
	}
	return
}
