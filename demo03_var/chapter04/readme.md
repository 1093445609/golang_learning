~~~~
Go 语言数据类型  因为比较多内容 建议百度
在 Go 编程语言中，数据类型用于声明函数和变量。

数据类型的出现是为了把数据分成所需内存大小不同的数据，编程的时候需要用大数据的时候才需要申请大内存，就可以充分利用内存。

Go 语言按类别有以下几种数据类型：
1.布尔型
    布尔型的值只可以是常量 true 或者 false(不会隐式转换为0 1)。一个简单的例子：var b bool = true。
2.数字类型
    整型: int uint8 uint16 uint32 uint64 
             int8   int16  int32  int64
    
    浮点 : float32 float64        复数: complex64 complex128
    
    其他数字类型: byte rune  uint int uintptr
3. 字符串
    utf8格式 占用字节1-4字节,节省内存
    
    
值类型和引用类型
所有像 int、float、bool 和 string 这些基本类型都属于值类型，
使用这些类型的变量直接指向存在内存中的值：

Go 语言常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量const的定义格式：
 const LENGTH int = 10
 const a, b, c = 1, false, "str" //多重赋值












~~~~