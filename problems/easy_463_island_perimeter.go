package problems

// IslandPerimeter represents
// Easy 463. Island Perimeter
// https://leetcode.com/problems/island-perimeter/
//
// 這題需要用中文解釋一下，
// 給定一個二維陣列，內值只有 0 與 1，1 是島，0 是海水，
// 島必定相鄰，島與島中間不會有海水，要你算出「島鏈」的外邊長。
// 以這張圖 (https://leetcode.com/static/images/problemset/island.png) 舉例，
// 要算出黃色邊有多少，圖例的答案是 16。
//
// 怎麼解呢？
// 一個島有 4 邊，如果 A 島右邊相鄰一個 B 島，兩個島的外邊等於 4 * 2 - 2，
// 減 2 是因為兩個島各少一邊。
// 如果 A 島右邊相鄰一個 B 島，下面又相鄰一個 C 島，三個島外邊為 4 * 3 - 2 - 2 = 8，
// 減兩次 2 是因為 A 右與 A 下各少一邊。
// 這樣大概能想出攻略了，先決定要取「左上」或「右下」來當標的，
// 決定標的後解題順序如下（以「左上」舉例）：
// 遞迴給定的二維陣列，
// 如果遇到「島」，先加 4，
// 再檢查「左」是否也是島，是的話就減 2，
// 再檢查「上」是否也是島，是的話就減 2，
// 遞迴到結束，就能得到答案囉。
//
// * 為什麼要取「左上」或「右下」？
// 因為取上下左右四邊的話，你遇到 A 個島檢查四邊，遇到相鄰的 B 島也檢查四邊，這樣就重複扣了！
//
func IslandPerimeter(grid [][]int) int {

	var res int
	if grid == nil {
		return res
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res -= 2
				}
				if j > 0 && grid[i][j-1] == 1 {
					res -= 2
				}
			}
		}
	}

	return res

}
