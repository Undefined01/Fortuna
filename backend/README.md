# Fortuba -- backend

> Backend of Fortuna Score Query System (for TungWah Senior High School)
>
> Fortuna 成绩查询系统后端

#### 构建

编译时请将三个`go`文件同时编译，像下面这样

``` bash
go build main.go getter.go http.go
```

#### API

#####`getter.go`主要向后端提供API。

+ **GetExamList**
  获取考试列表
  没有参数，返回`[]string`，为考试列表。

+ **GetScoreList**
  获取每个人某一科的总分数
  三个参数，分别为`考试名称`，`班级`（要有前置零），`科目`。
  返回`map[string][3]float32`，map的键为每个人的名称，值分别为每个人该科的主观得分、客观得分和总分数。

+ **GetSubscore**

  获取每个人某一科的主观题小题分数
  三个参数，分别为`考试名称`，`班级`（要有前置零），`科目`。
  返回`map[string]map[int]float32`，外层map的键为每个人的名称，内层map为题目编号，值为该小题的得分。

##### `main.go`主要向前端提供API。

访问`/api`即可。

参数：

+ `type`：调用API的名称，共三个
  + `exam`：考试列表
  + `score`：单科得分
  + `subscore`：单科小题分
+ `exam`、`subject`、`class`：在查询*单科得分*和*单科小题分*时才会使用到，分别对应考试名称、班级（需要前置0）、科目