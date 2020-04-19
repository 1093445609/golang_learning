~~~~
字符串和数值类型的相互转换
包 strconv 下的方法  
~~~~

~~~~
string到int
int,err:=strconv.Atoi(string)


string到int64
int64, err := strconv.ParseInt(string, 10, 64)


int到string
string:=strconv.Itoa(int)


int64到string
string:=strconv.FormatInt(int64,10)




~~~~