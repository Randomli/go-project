package leetcode_array

import (
	"logger"
)

var LOGFILE string = "./leetcode_array.log"
var LOGERRFILE string = "./leetcode_array.err.log"

func moveZeroes(nums []int) {

}

func Run() {
	logger := logger.InitLogger(LOGFILE, LOGERRFILE)
	defer logger.Sync()
	logger.Info("in leetcode")
	nums := []int{0, 0, 0, 30, 12}
	moveZeroes(nums)
	logger.Info(nums)
}
