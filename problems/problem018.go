package problems

import "github.com/ihrk/euler-go/utils"

func P18() int {
	paths := make([][]int, 0, 15)
	paths = append(paths, []int{75})
	paths = append(paths, []int{95, 64})
	paths = append(paths, []int{17, 47, 82})
	paths = append(paths, []int{18, 35, 87, 10})
	paths = append(paths, []int{20, 4, 82, 47, 65})
	paths = append(paths, []int{19, 1, 23, 75, 3, 34})
	paths = append(paths, []int{88, 2, 77, 73, 7, 63, 67})
	paths = append(paths, []int{99, 65, 4, 28, 6, 16, 70, 92})
	paths = append(paths, []int{41, 41, 26, 56, 83, 40, 80, 70, 33})
	paths = append(paths, []int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29})
	paths = append(paths, []int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14})
	paths = append(paths, []int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57})
	paths = append(paths, []int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48})
	paths = append(paths, []int{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31})
	paths = append(paths, []int{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23})
	vals := make([][]int, len(paths))
	for i := range paths {
		vals[i] = utils.CopyArr(paths[i])
	}
	for i := 1; i < len(paths); i++ {
		for j := 0; j < len(paths[i]); j++ {
			if j == 0 {
				vals[i][j] = vals[i-1][j] + paths[i][j]
			} else if j == len(paths[i])-1 {
				vals[i][j] = vals[i-1][j-1] + paths[i][j]
			} else {
				vals[i][j] = utils.Max(vals[i-1][j], vals[i-1][j-1]) + paths[i][j]
			}

		}
	}
	return utils.Maximum(vals[len(vals)-1])
}
