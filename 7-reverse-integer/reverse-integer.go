func reverse(x int) int {
    var current_num int64 = 0
    for math.Abs(float64(x)) > 0 {
        current_num = int64(current_num * 10) + int64(x % 10)
        x = x /10
    }
    if current_num > math.MaxInt32 || math.MinInt32 > current_num {
        return 0
    }
    return int(current_num)
}