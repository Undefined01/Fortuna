package utils

// 前端所需要的主要数据格式
type TableData struct {
	/* 数据表标题 */
	Title string
	/* 有序的数据键 */
	Cols []string
	/* 数值数组，每行的每一个元素都与在同一位置上键对应 */
	Data [][]interface{}
}

// 某一科目的得分
type ScoreOfSubject struct {
	/* 学号 */
	Sid int
	/* 姓名 */
	Name string
	/* 主观得分 */
	Object float32
	/* 客观得分 */
	Subject float32
	/* 总分数 */
	Total float32
	/* 级排名 */
	Rank int
}
