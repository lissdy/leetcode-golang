package main

func main() {

}

func maxProfit1(k int, prices []int) int {
	profit := 0
	if k > len(prices)/2 { //无限次交易，贪心算法
		for i := 0; i < len(prices)-1; i++ {
			if prices[i] < prices[i+1] {
				profit += prices[i+1] - prices[i]
			}
		}
		return profit
	}
	//规定次数交易，动态规划
	profit_0 := make([][]int, len(prices)) //不持有股票
	profit_1 := make([][]int, len(prices)) //持有股票
	//创建二维切片
	for i := 0; i < len(prices); i++ {
		profit_0[i] = make([]int, k+1)
		profit_1[i] = make([]int, k+1)
	}
	//初始化第一天
	for j := 0; j < k; j++ {
		profit_0[0][j] = 0
		profit_1[0][j] = -prices[0]
	}
	//状态转移方程
	for x := 1; x < len(prices); x++ { //遍历每天股票价格
		for y := 1; y < k+1; y++ { //遍历每次交易
			profit_0[x][y] = max(profit_0[x-1][y], profit_1[x-1][y]+prices[x])
			if profit < profit_0[x][y] {
				profit = profit_0[x][y]
			}
			profit_1[x][y] = max(profit_1[x-1][y], profit_0[x-1][y-1]-prices[x])
		}
	}
	return profit
}
