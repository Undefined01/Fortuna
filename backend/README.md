# Fortuna -- backend

> Backend of Fortuna Score Query System (for TungWah Senior High School)
>
> Fortuna 成绩查询系统后端

#### 构建

请使用`make`编译。

``` bash
make
```

#### 文档

请使用`godoc`查看。

[![GoDoc](https://godoc.org/github.com/Undefined01/Fortuna/backend?status.svg)](https://godoc.org/github.com/Undefined01/Fortuna/backend)

#### 前端API。

访问`/api`即可。

参数：

+ `type`：调用API的名称，共三个
  + `exam`：考试列表
  + `score`：单科得分
  + `subscore`：单科小题分
  + `examdata`：一场考试某一班级的全部数据
+ `exam`、`class`、`subject`：在查询*单科得分*和*单科小题分*时才会使用到，分别对应考试名称、班级（需要前置0）、科目
+ 在调用`examdata`的时候不需要传递`subject`参数。
