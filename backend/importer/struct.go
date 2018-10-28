package importer

import "sort"

// 前端所需要的主要数据格式
type Table struct {
	Title string
	Cols []string
	Data [][]interface{}
}

// 中间层，用于数据的再加工
type TableMap struct {
	Cols []string
	Data map[int][]interface{}
}

func (this TableMap) Transform(title string) Table {
	data := make([][]interface{}, 0, 30)
	for _, v := range this.Data {
		data = append(data, v)
	}
	return Table{
		title,
		this.Cols,
		data,
	}
}

// 某一科目的得分
type Score struct {
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

type ScoreList []Score
func (this ScoreList) Len() int           { return len(this) }
func (this ScoreList) Less(i, j int) bool { return this[i].Total > this[j].Total }
func (this ScoreList) Swap(i, j int)      { this[i], this[j] = this[j], this[i] }
func (this ScoreList) ReRank() {
	if len(this) == 0 {
		return
	}
	sort.Sort(this)
	this[0].Rank = 1
	for i, _ := range this {
		if i!=0 {
			if this[i].Total == this[i-1].Total {
				this[i].Rank = this[i-1].Rank
			} else {
				this[i].Rank = i+1
			}
		}
	}
}

type ScoreMap map[int]*Score
func (this ScoreMap) Add(list ScoreList) {
	for _, v := range list {
		_, ok := this[v.Sid]
		if !ok {
			newNode := v
			this[v.Sid] = &newNode
		} else {
			this[v.Sid].Total += v.Total
		}
	}
}
func (this ScoreMap) Transform() ScoreList {
	list := make(ScoreList, 0, len(this))
	for _, v := range this {
		list = append(list, *v)
	}
	list.ReRank()
	return list
}

// 某一科目的小题得分
type SubscoreMap TableMap
type Subscore Table

func (this *SubscoreMap) Transform(title string) *Subscore {
	if this == nil {
		return nil
	}
	table := TableMap(*this).Transform(title)
	return (*Subscore)(&table)
}
