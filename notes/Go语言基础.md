

# Go语言学习

## 基础语法

### 注释

注释不会被编译，每一个包应该有相关注释。

单行注释是最常见的注释形式，你可以在任何地方使用以 // 开头的单行注释。多行注释也叫块注释，均已以 /* 开头，并以 */ 结尾。如：

```
// 单行注释
/*
 Author by 菜鸟教程
 我是多行注释
 */
```

### 标识符

标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~Z和a~z)数字(0~9)、下划线_组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

以下是有效的标识符：

```
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal
```

以下是无效的标识符：

- 1ab（以数字开头）

- case（Go 语言的关键字）

- a+b（运算符是不允许的）

  

### 关键字

下面列举了 Go 代码中会使用到的 25 个关键字或保留字：

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

除了以上介绍的这些关键字，Go 语言还有 36 个预定义标识符：

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

程序一般由关键字、常量、变量、运算符、类型和函数组成。

程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。

程序中可能会使用到这些标点符号：**.**、**,**、**;**、**:** 和 **…**。

### 格式化字符串

Go 语言中使用 **fmt.Sprintf** 或 **fmt.Printf** 格式化字符串并赋值给新串：

- **Sprintf** 根据格式化参数生成格式化的字符串并返回该字符串。
- **Printf** 根据格式化参数生成格式化的字符串并写入标准输出。

#### Sprintf 实例

```go
**package** main

**import** (
  "fmt"
)

func main() {
  *// %d 表示整型数字，%s 表示字符串*
  **var** stockcode=123
  **var** enddate="2020-12-31"
  **var** url="Code=%d&endDate=%s"
  **var** target_url=fmt.Sprintf(url,stockcode,enddate)
  fmt.Println(target_url)
}
```

输出结果为：

```
Code=123&endDate=2020-12-31
```



## 数据类型

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **布尔型** 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 |
| 2    | **数字类型** 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 |
| 3    | **字符串类型:** 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 |
| 4    | **派生类型:** 包括：(a) 指针类型（Pointer）(b) 数组类型(c) 结构化类型(struct)(d) Channel 类型(e) 函数类型(f) 切片类型(g) 接口类型（interface）(h) Map 类型 |

------

### 数字类型

Go 也有基于架构的类型，例如：int、uint 和 uintptr。

| 序号 | 类型和描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | **uint8** 无符号 8 位整型 (0 到 255)                         |
| 2    | **uint16** 无符号 16 位整型 (0 到 65535)                     |
| 3    | **uint32** 无符号 32 位整型 (0 到 4294967295)                |
| 4    | **uint64** 无符号 64 位整型 (0 到 18446744073709551615)      |
| 5    | **int8** 有符号 8 位整型 (-128 到 127)                       |
| 6    | **int16** 有符号 16 位整型 (-32768 到 32767)                 |
| 7    | **int32** 有符号 32 位整型 (-2147483648 到 2147483647)       |
| 8    | **int64** 有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

**浮点型**

| 序号 | 类型和描述                        |
| :--- | :-------------------------------- |
| 1    | **float32** IEEE-754 32位浮点型数 |
| 2    | **float64** IEEE-754 64位浮点型数 |
| 3    | **complex64** 32 位实数和虚数     |
| 4    | **complex128** 64 位实数和虚数    |

------

### 其他数字类型

以下列出了其他更多的数字类型：

| 序号 | 类型和描述                               |
| :--- | :--------------------------------------- |
| 1    | **byte** 类似 uint8                      |
| 2    | **rune** 类似 int32                      |
| 3    | **uint** 32 或 64 位                     |
| 4    | **int** 与 uint 一样大小                 |
| 5    | **uintptr** 无符号整型，用于存放一个指针 |

## 变量

声明变量的一般形式是使用 var 关键字：

```
var identifier type
```

可以一次声明多个变量：

```
var identifier1, identifier2 type
```

**简短形式，使用 := 赋值操作符**

我们知道可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了，因此我们可以将它们简写为 a := 50 或 b := false。

a 和 b 的类型（int 和 bool）将由编译器自动推断。

这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。

**注意事项**

如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的，编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。

如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。

如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误



## 常量

常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型。

常量的定义格式：

```
const identifier [type] = value
```

常量还可以用作枚举：

```
const (
    Unknown = 0
    Female = 1
    Male = 2
)
```

### iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

```
const (
    a = iota
    b = iota
    c = iota
)
```

第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：

```
const (
    a = iota
    b
    c
)
```

### iota 用法

实例

```go
**package** main

**import** "fmt"

func main() {
  **const** (
      a = iota  *//0*
      b      *//1*
      c      *//2*
      d = "ha"  *//独立值，iota += 1*
      e      *//"ha"  iota += 1*
      f = 100   *//iota +=1*
      g      *//100  iota +=1*
      h = iota  *//7,恢复计数*
      i      *//8*
  )
  fmt.Println(a,b,c,d,e,f,g,h,i)
}

以上实例运行结果为：
```



```
0 1 2 ha ha 100 100 7 8
```



## 运算符

### 算术运算符

下表列出了所有Go语言的算术运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述 | 实例               |
| :----- | :--- | :----------------- |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |

### 关系运算符

下表列出了所有Go语言的关系运算符。假定 A 值为 10，B 值为 20。

| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |

### 逻辑运算符

下表列出了所有Go语言的逻辑运算符。假定 A 值为 True，B 值为 False。

| 运算符 | 描述                                                         | 实例               |
| :----- | :----------------------------------------------------------- | :----------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A \|\| B) 为 True |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 |                    |

### 位运算符

位运算符对整数在内存中的二进制位进行操作。

下表列出了位运算符 &, |, 和 ^ 的计算：

| p    | q    | p & q | p \| q | p ^ q |
| :--- | :--- | :---- | :----- | :---- |
| 0    | 0    | 0     | 0      | 0     |
| 0    | 1    | 0     | 1      | 1     |
| 1    | 1    | 1     | 1      | 0     |
| 1    | 0    | 0     | 1      | 1     |

| 运算符 | 描述                                                         | 实例                                   |
| :----- | :----------------------------------------------------------- | :------------------------------------- |
| &      | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。 | (A & B) 结果为 12, 二进制为 0000 1100  |
| \|     | 按位或运算符"\|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或 | (A \| B) 结果为 61, 二进制为 0011 1101 |
| ^      | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。 | (A ^ B) 结果为 49, 二进制为 0011 0001  |
| <<     | 左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。 | A << 2 结果为 240 ，二进制为 1111 0000 |
| >>     | 右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。 | A >> 2 结果为 15 ，二进制为 0000 1111  |

### 赋值运算符

下表列出了所有Go语言的赋值运算符。

| 运算符 | 描述                                           | 实例                                  |
| :----- | :--------------------------------------------- | :------------------------------------ |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C |
| +=     | 相加后再赋值                                   | C += A 等于 C = C + A                 |
| -=     | 相减后再赋值                                   | C -= A 等于 C = C - A                 |
| *=     | 相乘后再赋值                                   | C *= A 等于 C = C * A                 |
| /=     | 相除后再赋值                                   | C /= A 等于 C = C / A                 |
| %=     | 求余后再赋值                                   | C %= A 等于 C = C % A                 |
| <<=    | 左移后赋值                                     | C <<= 2 等于 C = C << 2               |
| >>=    | 右移后赋值                                     | C >>= 2 等于 C = C >> 2               |
| &=     | 按位与后赋值                                   | C &= 2 等于 C = C & 2                 |
| ^=     | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                 |
| \|=    | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2               |

### 其他运算符

下表列出了Go语言的其他运算符。

| 运算符 | 描述             | 实例                       |
| :----- | :--------------- | :------------------------- |
| &      | 返回变量存储地址 | &a; 将给出变量的实际地址。 |
| *      | 指针变量。       | *a; 是一个指针变量         |

## 条件语句

Go 语言提供了以下几种条件判断语句：

| 语句              | 描述                                                         |
| :---------------- | :----------------------------------------------------------- |
| [if 语句]         | **if 语句** 由一个布尔表达式后紧跟一个或多个语句组成。       |
| [if...else 语句]( | **if 语句** 后可以使用可选的 **else 语句**, else 语句中的表达式在布尔表达式为 false 时执行。 |
| [if 嵌套语句](    | 你可以在 **if** 或 **else if** 语句中嵌入一个或多个 **if** 或 **else if** 语句。 |
| [switch 语句]     | **switch** 语句用于基于不同条件执行不同动作。                |
| [select 语句]     | **select** 语句类似于 **switch** 语句，但是select会随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。 |

### switch语句

switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。

switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。

switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 **fallthrough** 。

Go 编程语言中 switch 语句的语法如下：

```
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var grade string = "B"
   var marks int = 90

   switch marks {
      case 90: grade = "A"
      case 80: grade = "B"
      case 50,60,70 : grade = "C"
      default: grade = "D"  
   }

   switch {
      case grade == "A" :
         fmt.Printf("优秀!\n" )     
      case grade == "B", grade == "C" :
         fmt.Printf("良好\n" )      
      case grade == "D" :
         fmt.Printf("及格\n" )      
      case grade == "F":
         fmt.Printf("不及格\n" )
      default:
         fmt.Printf("差\n" );
   }
   fmt.Printf("你的等级是 %s\n", grade );      
}
```

### select语句

select 是 Go 中的一个控制结构，类似于 switch 语句。

select 语句只能用于**通道操作**，每个 case 必须是一个通道操作，要么是发送要么是接收。

select 语句会监听所有指定的通道上的操作，一旦其中一个通道准备好就会执行相应的代码块。

如果多个通道都准备好，那么 select 语句会随机选择一个通道执行。如果所有通道都没有准备好，那么执行 default 块中的代码。

Go 编程语言中 select 语句的语法如下：

```go
**select** {
 **case** <- channel1:
  *// 执行的代码*
 **case** value := <- channel2:
  *// 执行的代码*
 **case** channel3 <- value:
  *// 执行的代码*

  *// 你可以定义任意数量的 case*

 **default**:
  *// 所有通道都没有准备好，执行的代码*
}
```

## 循环语句

### for循环

or 循环是一个循环控制结构，可以执行指定次数的循环。

Go 语言的 For 循环有 3 种形式，只有其中的一种使用分号。

和 C 语言的 for 一样：

```
for init; condition; post { }
```

和 C 的 while 一样：

```
for condition { }
```

和 C 的 for(;;) 一样：

```
for { }
```

- init： 一般为赋值表达式，给控制变量赋初值；
- condition： 关系表达式或逻辑表达式，循环控制条件；
- post： 一般为赋值表达式，给控制变量增量或减量。

**循环嵌套**：在for循环中嵌套一个或者多个for循环

## 函数

### 函数定义

Go 语言函数定义格式如下：

```
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

### 函数返回多个值

Go 函数可以返回多个值，例如：

```go
**package** main

**import** "fmt"

func swap(x, y string) (string, string) {
  **return** y, x
}

func main() {
  a, b := swap("Google", "Runoob")
  fmt.Println(a, b)
}
```

## 数组

### 声明数组

Go 语言数组声明需要指定元素类型及元素个数，语法格式如下：

```
var arrayName [size]dataType
```

其中，**arrayName** 是数组的名称，**size** 是数组的大小，**dataType** 是数组中元素的数据类型。

以下定义了数组 balance 长度为 10 类型为 float32：

```
var balance [10]float32
```

### 初始化数组

以下演示了数组初始化：

以下实例声明一个名为 numbers 的整数数组，其大小为 5，在声明时，数组中的每个元素都会根据其数据类型进行默认初始化，对于整数类型，初始值为 0。

```
var numbers [5]int
```

还可以使用初始化列表来初始化数组的元素：

```
var numbers = [5]int{1, 2, 3, 4, 5}
```

以上代码声明一个大小为 5 的整数数组，并将其中的元素分别初始化为 1、2、3、4 和 5。

另外，还可以使用 **:=** 简短声明语法来声明和初始化数组：

```
numbers := [5]int{1, 2, 3, 4, 5}
```

以上代码创建一个名为 numbers 的整数数组，并将其大小设置为 5，并初始化元素的值。

**注意：**在 Go 语言中，数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 **[5]int** 和 **[10]int** 是不同的类型。

## 指针

### 什么是指针

一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

```
var var_name *var-type
```

var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针。以下是有效的指针声明：

```
var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

本例中这是一个指向 int 和 float32 的指针。

### 如何使用指针

指针使用流程：

- 定义指针变量。
- 为指针变量赋值。
- 访问指针变量中指向地址的值。

在指针类型前面加上 * 号（前缀）来获取指针所指向的内容。

```go
**package** main

**import** "fmt"

func main() {
  **var** a int= 20  */\* 声明实际变量 \*/*
  **var** ip *int     */\* 声明指针变量 \*/*

  ip = &a  */\* 指针变量的存储地址 \*/*

  fmt.Printf("a 变量的地址是: %x**\n**", &a  )

  */\* 指针变量的存储地址 \*/*
  fmt.Printf("ip 变量储存的指针地址: %x**\n**", ip )

  */\* 使用指针访问值 \*/*
  fmt.Printf("*ip 变量的值: %d**\n**", *ip )
}
```

以上实例执行输出结果为：

```
a 变量的地址是: 20818a220
ip 变量储存的指针地址: 20818a220
*ip 变量的值: 20
```

## 结构体

### 定义结构体

结构体定义需要使用 type 和 struct 语句。struct 语句定义一个新的数据类型，结构体中有一个或多个成员。type 语句设定了结构体的名称。结构体的格式如下：

```
type struct_variable_type struct {
   member definition
   member definition
   ...
   member definition
}
```

一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：

```
variable_name := structure_variable_type {value1, value2...valuen}
或
variable_name := structure_variable_type { key1: value1, key2: value2..., keyn: valuen}
```

### 访问结构体成员

如果要访问结构体成员，需要使用点号 **.** 操作符，格式为：

```
结构体.成员名"
```

### 结构体指针

你可以定义指向结构体的指针类似于其他指针变量，格式如下：

```
var struct_pointer *Books
```

以上定义的指针变量可以存储结构体变量的地址。查看结构体变量地址，可以将 & 符号放置于结构体变量前：

```
struct_pointer = &Book1
```

使用结构体指针访问结构体成员，使用 "." 操作符：

```
struct_pointer.title
```

## 切片（slice）

Go 语言切片是对数组的抽象。

Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go 中提供了一种灵活，功能强悍的内置类型切片("动态数组")，与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。

------

### 定义切片

你可以声明一个未指定大小的数组来定义切片：

```
var identifier []type
```

或使用 **make()** 函数来创建切片:

```
var slice1 []type = make([]type, len)

也可以简写为

slice1 := make([]type, len)
```

## 范围（range）

Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。

for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：

```
for key, value := range oldMap {
    newMap[key] = value
}
```

以上代码中的 key 和 value 是可以省略。

如果只想读取 key，格式如下：

```
for key := range oldMap
```

或者这样：

for key, _ := range oldMap

如果只想读取 value，格式如下：

```
for _, value := range oldMap
```

## Map

Map 是一种无序的键值对的集合。

Map 最重要的一点是通过 key 来快速检索数据，key 类似于索引，指向数据的值。

Map 是一种集合，所以我们可以像迭代数组和切片那样迭代它。不过，Map 是无序的，遍历 Map 时返回的键值对的顺序是不确定的。

在获取 Map 的值时，如果键不存在，返回该类型的零值，例如 int 类型的零值是 0，string 类型的零值是 ""。

Map 是引用类型，如果将一个 Map 传递给一个函数或赋值给另一个变量，它们都指向同一个底层数据结构，因此对 Map 的修改会影响到所有引用它的变量。

### 定义 Map

可以使用内建函数 make 或使用 map 关键字来定义 Map:

```
/* 使用 make 函数 */
map_variable := make(map[KeyType]ValueType, initialCapacity)
```

其中 KeyType 是键的类型，ValueType 是值的类型，initialCapacity 是可选的参数，用于指定 Map 的初始容量。Map 的容量是指 Map 中可以保存的键值对的数量，当 Map 中的键值对数量达到容量时，Map 会自动扩容。如果不指定 initialCapacity，Go 语言会根据实际情况选择一个合适的值。

## Defer

```go
defer fmt.Println("111") //defer的注册
```

“后注册先执行”

如果defer注册的是一条go语句，语句中包含的值在注册时就计算好

如果defer注册的是一个匿名函数，函数中的值在defer执行时才进行计算

![image-20250310163355206](https://my-typora329.oss-cn-beijing.aliyuncs.com/image-20250310163355206.png)



## 接口

接口（interface）是 Go 语言中的一种类型，用于定义行为的集合，它通过描述类型必须实现的方法，规定了类型的行为契约。写大型项目必用。

Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口。

Go 的接口设计简单却功能强大，是实现多态和解耦的重要工具。

接口可以让我们将不同的类型绑定到一组公共的方法上，从而实现多态和灵活的设计。

```
type Interface_name interface{
	
}
```



## 时间相关函数

time包

```go
import(
	"time"
)
t:=time.Now() //数据类型：time.Time
fmt.Println(t0.Unix()) //输出时间戳
time.sleep(50*time.Millisecond) //程序暂停50ms
t1:=time.Now()
diff:=t1.Sub(t0)//计算时间差
fmt.Println(diff.Milliseconds())
//Time - Time = Duration

//Time + Duration = Time
d:=time.Duration(2*time.Second)
t2:=t0.Add(d)
fmt.Println(t2.Unix())
```

时间的格式化：

```
GO语言中的日期和时间遵循以下格式：
DATE="2006-01-02"
TIME="2006-01-02 15:04:05"
```

输出：

```go
fmt.Println(t0.Format(DATE))//转换成string
fmt.Println(t0.Format(TIME))

t3:=time.Parse(TIME,s)
```

指定时区：

```go
loc:=time.LoadLocation("Asia/Shanghai")//东八区
t3:=time.ParseInLocation(TIME,s,loc)
```



## 读写文件

os包

```go
import(
	"os"
)
```

```go
//读文件
file:=os.Open("   ") //打开文件
content:=make([]byte,100)
n,err:=file.Read(content)//读文件 err会显示读出多少
file.Close() //关闭文件
```

```go
//写文件
file,err:=os.OpenFile("b.txt",
	os.O_CREATE|os.O_TRUNC|os.O_WEONLY,0644)//文件不存在则创建|覆盖原有新写入|只写只读不可执行
	os.O_APPEND//在已有的内容后追加
n,err:=file.Write([]byte(content))
```

```go
package main

import(
	"fmt"
	"os"
)

func main(){
	file,err:=os.Open("b.txt")
	if err!=nil{
		fmt.Println("Open file failed",err)
		return
	}
	defer file.Close()
	reader:=bufio.NewReader(file)//缓存
    for{
        line,err:=reader.ReadString('\n')//读文件，遇到换行符停一下
        if err!=nil{
            if err==io.EOF{ //END OF FILE
                break
            }else{
                fmt.Println("Read file failed",err)
                return
            }
        }else{
            fmt.Print(line)
        }
    }
}
```

## JSON

JSON （JavaScript Object Notation）是一种比XML更轻量级的数据交换格式，在易于人们阅读和编写的同时，也易于程序解析和生成。尽管JSON是JavaScript的一个子集，但JSON采用完全独立于编程语言的文本格式，且表现为键/值对集合的文本描述形式（类似一些编程语言中的字典结构），这使它成为较为理想的、跨平台、跨语言的数据交换语言。

### 编码JSON（通过结构体或者map）

```go
//导入json包
import("encoding/json")
//使用json.Marshal 通过结构体生成JSON
func Marshal(v interface{})([]byte,error) 

// MarshalIndent 很像 Marshal，只是用缩进对输出进行格式化
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
```

### 解码JSON

```go
func Unmarshal(data []byte, v interface{}) error
```

## 单元测试

```go
//导入testing包
import("testing")
```

单元测试是为了检测函数是否正确

### 测试覆盖率

`go test -cover $dir` 只能给出`$dir`目录的整体单测覆盖率

`go test ./project_prepare -coverprofile=data/test_cover` （在data目录下生成个文件）

或`go test ./project_prepare -coverprofile=data/test_cover -covermode=count`

`covermode` 的3个取值：

​	• get:每个语句是否执行，默认值

​	• count:每个语句执行了几次，鼠标悬停在语句上能显示执行的次数

​	• atomic:类似于count，但表示的是并行程序中的精确计数 

用data目录下生成的文件执行以下命令：

`go tool cover -func=data/test_cover` 输出每一个函数的覆盖率

`go tool cover -html=data/test_cover` 细化到每一行代码的覆盖情况



## init()

在Go程序运行时，每个包都会自动调用它的`init()`函数（如果有的话），用来做初始化操作。

•**每个包可以有多个`init()`函数**，它们会按照代码出现的顺序执行。

•**init函数优先于main函数执行**。也就是说，所有的包级变量初始化和所有的`init()`都会在`main()`运行前完成。

•**导入包的顺序决定了`init()`的执行顺序**（先初始化依赖包，再初始化本包）。



## 并发编程

### 什么是并发

![image-20250314220210464](https://my-typora329.oss-cn-beijing.aliyuncs.com/image-20250314220210464.png)

进程：每一个独立执行的程序（程序、数据集、进程控制块）

线程：轻量级进程，程序执行的最小单元，一个进程包含多个线程（程序执行的每个步骤）

协程：比线程更加轻量，用户态的轻量级线程，也称为微线程，完全由用户控制调度

### Go语言的协程：Goroutine

与线程相比，创建Goroutine的成本很小，Go应用程序可以并发运行数千个Goroutines。

#### 主Goroutine

封装main函数的goroutine就是主goroutine

主goroutine设置每一个goroutine所能申请空间的最大尺寸，接着进行一系列的初始化工作

![image-20250314223049266](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20250314223049266.png)

#### 如何使用Goroutine

在函数或者方法调用前面加上关键字go，就会运行一个新的goroutine

```go
package main

import (
	"fmt"
)

func Print_Num() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func main() {
	go Print_Num()
	for i := 1; i <= 100; i++ {
		fmt.Println("Hello World")
	}
	fmt.Println("main over")
}
/*运行结果为打印数字和打印Hello World交替进行
主goroutine结束后子goroutine再继续执行一段时间后结束*/
```

#### runtime包

**`runtime` 调度器是个非常有用的东西，关于 `runtime` 包几个方法:**

- **NumCPU**：返回当前系统的 `CPU` 核数量

- **GOMAXPROCS**：设置最大的可同时使用的 `CPU` 核数

  通过runtime.GOMAXPROCS函数，应用程序何以在运行期间设置运行时系统中得P最大数量。但这会引起“Stop the World”。所以，应在应用程序最早的调用。并且最好是在运行Go程序之前设置好操作程序的环境变量GOMAXPROCS，而不是在程序中调用runtime.GOMAXPROCS函数。

  无论我们传递给函数的整数值是什么值，运行时系统的P最大值总会在1~256之间。

\> go1.8后，默认让程序运行在多个核上,可以不用设置了 > go1.8前，还是要设置一下，可以更高效的利益cpu

- **Gosched**：让当前线程让出 `cpu` 以让其它线程运行,它不会挂起当前线程，因此当前线程未来会继续执行

  这个函数的作用是让当前 `goroutine` 让出 `CPU`，当一个 `goroutine` 发生阻塞，`Go` 会自动地把与该 `goroutine` 处于同一系统线程的其他 `goroutine` 转移到另一个系统线程上去，以使这些 `goroutine` 不阻塞。

- **Goexit**：退出当前 `goroutine`(但是`defer`语句会照常执行)

- **NumGoroutine**：返回正在执行和排队的任务总数

  runtime.NumGoroutine函数在被调用后，会返回系统中的处于特定状态的Goroutine的数量。这里的特指是指Grunnable\Gruning\Gsyscall\Gwaition。处于这些状态的Groutine即被看做是活跃的或者说正在被调度。

  注意：垃圾回收所在Groutine的状态也处于这个范围内的话，也会被纳入该计数器。

- **GOOS**：目标操作系统

- **runtime.GC**:会让运行时系统进行一次强制性的垃圾收集

  1. 强制的垃圾回收：不管怎样，都要进行的垃圾回收。
  2. 非强制的垃圾回收：只会在一定条件下进行的垃圾回收（即运行时，系统自上次垃圾回收之后新申请的堆内存的单元（也成为单元增量）达到指定的数值）。

- **GOROOT** :获取goroot目录

- **GOOS** : 查看目标操作系统 很多时候，我们会根据平台的不同实现不同的操作，就而已用GOOS了：

## Go module

模块（module）由模块路径标识，模块路径就是模块的规范名称

模块名称在go.mod文件里

### 模块名称不是源码仓库地址

先向模块路径发送一个GET请求（go-get=1）获得源码仓库地址，然后通过git clone去下载源码

![image-20250516104356740](https://my-typora329.oss-cn-beijing.aliyuncs.com/image-20250516104356740.png)

![image-20250516104541550](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20250516104541550.png)

### GOSUMDB

源码校验

![image-20250516105914127](https://my-typora329.oss-cn-beijing.aliyuncs.com/image-20250516105914127.png)



### 开发私有模块

![image-20250516121213461](https://my-typora329.oss-cn-beijing.aliyuncs.com/image-20250516121213461.png)

### 部署私有代理

![image-20250516122020302](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20250516122020302.png)

### go module 版本

go module 每个版本以v开头，后面跟语义版本

当vcs为git时，通常情况下模块的版本就是git tag的版本