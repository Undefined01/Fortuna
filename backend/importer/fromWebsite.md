### GetExamList

#### 数据格式

```html
<select>
<option value="考试名称" >考试名称 </option>
<option value="考试名称" >考试名称 </option>
</select>
<select>
<option value="班级" >班级 </option>
<option value="班级" >班级 </option>
</select>
```

#### 查找依据

“`<select>`”“`</select>`”获取考试名称的列表

“`<option value=\"`”“`\" >`”获取考试名称

忽略最后三行平均分等信息

### GetScoreList

#### 数据格式

```html
<table 
  <tr 
    <p align="center">标题</
  </tr>
</table>
<table 
  <tr 
    表头
  </tr>
  <tr>
    <p align="center">学生考号</
    <p align="center">姓名</
    <p align="center">原基础名次</
    <p align="center">本次基础名次重排</
    <p align="center">本次考试名次</
    <p align="center">全级排名</
    <p align="center">文理排名</
    <p align="center">主观分数</
    <p align="center">客观分数</
    <p align="center">总分数</
  </tr>
</table>
```

#### 查找依据

“`<tr>`”“`</tr>`”获取一行数据

“`<p align="center">`”“`</`”获取该列数据

忽略最后三行平均分等信息

### GetSubscore

#### 数据格式

```html
<table 
  <tr 
    <p align="center">标题</
  </tr>
</table>
<table 
  <tr 
    <p align=center>学号</
    <p align=center>姓名</
    <p align=center>小题1</
    <p align=center>小题2</
  <tr>
  <tr>
    <p align="center">学号</
    <p align="center">姓名</
    <p align="center">小题1得分</
    <p align="center">小题2得分</
  <tr  class='alt'>
    <p align="center">学号</
    <p align="center">姓名</
    <p align="center">小题1得分</
    <p align="center">小题2得分</
</table>
<table 
  <tr 
    <p align="center">试题分析</
  </tr>
</table>
<table 
  <tr>
    <p align="center">试题号</
    <p align="center"><font color="#FF0000">分数</
    <p align="center">读卡</
    <p align="center">得分</
    <p align="center">得分率</
    <p align="center">平均分</
  </tr>
  <tr>（问号数据格式=_+）
    align="center">(空格或换行)1(空格或换行)</
    <font color="#FF0000">(空格或换行)3(空格或换行)</
    align="center">48(空格或换行)</
    align="center">144(空格或换行)</
    <p align="center">100%</
    align="center">3
  </tr>
</table>
```

#### 查找依据

“`</table>”`跳过标题

“`<tr`”“`<tr>`”获取表头

“`<p align="center">`”“`</`”获取表头数据

“`<tr`”“`<tr`”获取一行数据

“`<p align="center">`”“`</`”获取数据内容

“`<tr`”“`\n`”获取最后一行数据

“`<p align="center">`”“`</`”获取数据内容
