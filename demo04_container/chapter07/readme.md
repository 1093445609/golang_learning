~~~~
map 概念
~~~~
~~~~
map 是引用类型，可以使用如下方式声明：
map[keyType]       键对应值类型> int为value类型
var mapList map[string] int



map 容量

map 可以根据新增的 key-value 动态的伸缩，因此它不存在固定长度或者最大限制，
但是也可以选择标明 map 的初始容量 capacity，格式如下：
map2 := make(map[string]float, 100)


用切片作为 map 的值
value可以是切片 这样value就可以有多个了
mp1 := make(map[int][]int)
mp2 := make(map[int]*[]int)

~~~~