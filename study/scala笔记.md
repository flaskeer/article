如果想用数组的值做可变实参，可以将数组展开成离散值
用 :_*可以做到
val numbers = Array(34,3,434,3)
println(max(numbers:_*))

T <: Pet 表示T所代表的类派生自Pet  有上界的语法

D:>S 有下界
+T告诉scala允许协变  在类型检查期间，让scala接收某个类型
或者其基类型
-T就可以让scala支持类型逆变

/: 是foldLeft()方法

创建伴生类的实例不需要new关键字 这是apply()替我们完成的

Scala 不支持＋＋i,i++ 运算符，因此需要使用 i＋＝1 来加一。 这段代码看起来和 Java 代码差不多，实际上 while 也是一个函数，你自动可以利用 scala 语言的扩展性，实现while 语句，使它看起来和 Scala 语言自带的关键字一样调用。Scala访问数组的语法是使用()而非[]。
val greetStrings =new Array[String](3)
greetStrings(0)="Hello"
greetStrings(1)=","
如果一个函数只有一个参数并且只包含一个表达式，那么你无需明确指明参数。
args.foreach( println)

在 Scala 中你可以使用 new 来实例化一个类。当你创建一个对象的实例时，你可以使用数值或类型参数。如果使用类型参数，它的作用类似 Java 或 .Net 的 Generic 类型。所不同的是 Scala 

cala 还是提供了初始化数组的简单的方法，比如什么的例子数组可以使用如下代码：

val greetStrings =Array("Hello",",","World\n")
这里使用（）其实还是调用 Array 类的关联对象 Array 的 apply 方法，也就是

val greetStrings =Array.apply("Hello",",","World\n")

 Scala 中，数组和其它普遍的类定义一样，没有什么特别之处，当你在某个值后面使用（）时，Scala 将其翻译成对应对象的 apply 方法。因此本例中 greetStrings(1) 其实调用 greetString.apply(1) 方法。这种表达方法不仅仅只限于数组，对于任何对象，如果在其后面使用（）,都将调用该对象的 apply 方法。同样的如果对某个使用（）的对象赋值，比如：

greetStrings(0)="Hello"
Scala 将这种赋值转换为该对象的 update 方法， 也就是 greetStrings.update(0,”hello”)。因此上面的例子，使用传统的方法调用可以写成：

val greetStrings =new Array[String](3)
greetStrings.update(0,"Hello")
greetStrings.update(1,",")
greetStrings.update(2,"world!\n")
for(i <- 0 to 2)
  print(greetStrings.apply(i))
从这点来说，数组在 Scala 中并不某种特殊的数据类型，和普通的类没有什么不同。

通过:::操作符（其实为:::方法）将两个列表链接起来。实际上由于 List 的不可以修改特性，Scala 创建了一个新的 List 对象 oneTwoThreeFour 来保存两个列表连接后的值。

List 也提供了一个::方法用来向 List 中添加一个元素，::方法（操作符）是右操作符，也就是使用::右边的对象来调用它的::方法，Scala 中规定所有以：开头的操作符都是右操作符，因此如果你自己定义以：开头的方法（操作符）也是右操作符。

如下面使用常量创建一个列表：

val oneTowThree = 1 :: 2 ::3 :: Nil
println(oneTowThree)
调用空列表对象 Nil 的 ::方法 也就是

val oneTowThree =  Nil.::(3).::(2).::(1)

Scala 中另外一个很有用的容器类为 Tuples，和 List 不同的 Tuples 可以包含不同类型的数据，而 List 只能包含同类型的数据。Tuples 在方法需要返回多个结果时非常有用。（ Tuple 对应到数学的矢量的概念）。

一旦定义了一个元组，可以使用 ._和索引来访问员组的元素（矢量的分量，注意和数组不同的是，元组的索引从 1 开始）。

val pair=(99,"Luftballons")
println(pair._1)
println(pair._2)
元祖的实际类型取决于它的分量的类型，比如上面 pair 的类型实际为 Tuple2[Int,String]，而 (‘u’,’r’,”the”,1,4,”me”) 的类型为 Tuple6[Char,Char,String,Int,Int,String]。

目前 Scala 支持的元祖的最大长度为 22。如果有需要，你可以自己扩展更长的元祖。

使用 Set 的基本方法如下：

var jetSet = Set ("Boeing","Airbus")
jetSet +="Lear"
println(jetSet.contains("Cessna"))
缺省情况 Set 为 Immutable Set，如果你需要使用可修改的集合类（ Set 类型），你可以使用全路径来指明 Set，比如 scala.collection.mutalbe.Set 。

Map 的基本用法如下（ Map 类似于其它语言中的关联数组如 PHP ）

val romanNumeral = Map ( 1 -> "I" , 2 -> "II",
  3 -> "III", 4 -> "IV", 5 -> "V")
println (romanNumeral(4))

def printArgs ( args: Array[String]) : Unit ={
    for( arg <- args)
      println(arg)
}
或者更简化为：

def printArgs ( args: Array[String]) : Unit ={
    args.foreach(println)
}
这个例子也说明了尽量少用 vars 的好处，代码更简洁和明了，从而也可以减少错误的发生。因此 Scala 编程的一个基本原则上，能不用 Vars，尽量不用 vars,能不用 mutable变量，尽量不用 mutable变量，能避免函数的副作用，尽量不产生副作用。

Scala 的 singleton 对象不仅限于作为静态对象的容器，它在 Scala 中也是头等公民，但仅仅定义 Singleton 对象本身不会创建一个新的类型，你不可以使用 new 再创建一个新的 Singleton 对象（这也是 Singleton 名字的由来），此外和类定义不同的是，singleton 对象不可以带参数（类定义参数将在后面文章介绍）。

Scala 为 Singleton 对象的 main 定义了一个 App trait 类型，因此上面的例子可以简化为：

object HelloWorld extends App{
   println("Hello, world!")
}
这段代码就不能作为脚本运行，Scala 的脚本要求代码最后以表达式结束。因此运行这段代码，需要先编译这段代码：

scalac Helloworld.scala
编译好之后，运行该应用

scala HelloWord
注意： Scala 提供了一个快速编译代码的辅助命令 fsc (fast scala compliler) ，使用这个命令，只在第一次使用 fsc时启动 JVM，之后 fsc 在后台运行，这样就避免每次使用 scalac 时都要载入相关库文件，从而提高编译速度。

实际上类 Int 定义了多个+方法的重载方法（以支持不同的数据类型）比如和 Long 类型相加。 +符号为一运算符，为一中缀运算符。 在 Scala 中你可以定义任何方法为一操作符。 比如 String 的 IndexOf 方法也可以使用操作符的语法来书写。 例如：

scala> val s ="Hello, World"
s: String = Hello, World
scala> s indexOf 'o'
res0: Int = 4
由此可以看出运算符在 Scala 中并不是什么特殊的语法，任何 Scala 方法都可以作为操作符来使用。是否是操作符取决于你如何使用这个方法，当你使用 s.indexOf(‘o’) indexOf 不是一个运算符。 而你写成 s indexOf ‘o’ ,indexOf 就是一个操作符，因为你使用了操作符的语法。

除了类似+的中缀运算符（操作符在两个操作符之间），还可以有前缀运算符和后缀运算符。顾名思义前缀运算符的操作符在操作数前面，比如 -7 的“-”。后缀运算符的运算符在操作数的后面，比如 7 toLong 中的 toLong。 前缀和后缀操作符都使用一个操作数，而中缀运算符使用前后两个操作数。Scala 在实现前缀和后缀操作符的方法，其方法名都以 unary_-开头。比如:表达式 -2.0 实际上调用 （2.0）.unary_-方法。

如果你需要定义前缀方法，你只能使用+,-,! 和～四个符号作为前缀操作符。

后缀操作符在不使用.和括号调用时不带任何参数。在 Scala 中你可以省略掉没有参数的方法调用的空括号。按照惯例，如果你调用方法是为了利用方法的“副作用”，此时写上空括号，如果方法没有任何副作用（没有修改其它程序状态），你可以省略掉括号。比如：

scala> val s ="Hello, World"
s: String = Hello, World
scala> s toLowerCase
res0: String = hello, world

对象恒等比较
如果需要比较两个对象是否相等，可以使用==和!=操作符

scala> 1 == 2
res6: Boolean = false
scala> 1 !=2
res7: Boolean = true
scala> List(1,2,3) == List (1,2,3)
res8: Boolean = true
scala> ("he"+"llo") == "hello"
res9: Boolean = true
Scala 的==和 Java 不同，scala 的==只用于比较两个对象的值是否相同。而对于引用类型的比较使用另外的操作符 eq 和 ne。

操作符的优先级和左右结合性
Scala 的操作符的优先级和 Java 基本相同，如果有困惑时，可以使用（）改变操作符的优先级。 操作符一般为左结合，Scala 规定了操作符的结合性由操作符的最后一个字符定义。对于以“：”结尾的操作符都是右结合，其它的操作符多是左结合。例如：

ab 为　a.(b)　而　a:::b 为 b.:::(a),而 a:::b:::c = a::: (b ::: c) , abc= (ab)c

前提条件检查
前面说过有理数可以表示为 n/d (其中 d,n 为证书，而 d 不能为 0），对于前面的 Rational 定义，我们如果使用 0，也是可以的

scala> new Rational(5,0)
res0: Rational = 5/0
怎么解决分母不能为 0 的问题了，面向对象编程的一个优点是实现了数据的封装，你可以确保在其生命周期过程中是有效的。对于有理数的一个前提条件是分母不可以为 0，Scala 中定义为传入构造函数和方法的参数的限制范围，也就是调用这些函数或方法的调用者需要满足的条件。Scala 中解决这个问题的一个方法是使用 require 方法（require 方法为 Prede f对象的定义的一个方法，Scala 环境自动载入这个类的定义，因此无需使用 import 引入这个对象），因此修改 Rational 定义如下：

scala> class Rational (n:Int, d:Int) {
     |    require(d!=0)
     |    override def toString = n + "/" +d
     | }
defined class Rational
scala> new Rational(5,0)
java.lang.IllegalArgumentException: requirement failed
  at scala.Predef$.require(Predef.scala:211)
  ... 33 elided
可以看到如果再使用 0 作为分母，系统将抛出 IllegalArgumentException 异常。

class Rational (n:Int, d:Int) {
   require(d!=0)
   override def toString = n + "/" +d
   def add(that:Rational) : Rational =
     new Rational(n*that.d + that.n*d,d*that.d)
}
实际上编译器会给出如下编译错误：

<console>:11: error: value d is not a member of Rational
            new Rational(n*that.d + that.n*d,d*that.d)
                                ^
<console>:11: error: value d is not a member of Rational
            new Rational(n*that.d + that.n*d,d*that.d)
这是为什么呢？尽管类参数在新定义的函数的访问范围之内，但仅限于定义类的方法本身(比如之前定义的 toString 方法，可以直接访问类参数），但对于 that 来说，无法使用 that.d 来访问 d. 因为 that 不在定义的类可以访问的范围之内。此时需要定类的成员变量。（注：后面定义的 case class 类型编译器自动把类参数定义为类的属性，这是可以使用 that.d 等来访问类参数）。

修改 Rational 定义，使用成员变量定义如下：

class Rational (n:Int, d:Int) {
   require(d!=0)
   val number =n
   val denom =d 
   override def toString = number + "/" +denom 
   def add(that:Rational)  =
     new Rational(
       number * that.denom + that.number* denom,
       denom * that.denom
     )
}
要注意的我们这里定义成员变量都使用了 val ，因为我们实现的是“immutable”类型的类定义。number 和 denom 以及 add 都可以不定义类型，Scala 编译能够根据上下文推算出它们的类型。

scala> val oneHalf=new Rational(1,2)
oneHalf: Rational = 1/2
scala> val twoThirds=new Rational(2,3)
twoThirds: Rational = 2/3
scala> oneHalf add twoThirds
res0: Rational = 7/6
scala> oneHalf.number
res1: Int = 1
可以看到，这是就可以使用 .number 等来访问类的成员变量。

辅助构造函数
在定义类时，很多时候需要定义多个构造函数，在 Scala 中，除主构造函数之外的构造函数都称为辅助构造函数（或是从构造函数），比如对于 Rational 类来说，如果定义一个整数，就没有必要指明分母，此时只要整数本身就可以定义这个有理数。我们可以为 Rational 定义一个辅助构造函数，Scala 定义辅助构造函数使用 this(…)的语法，所有辅助构造函数名称为 this。

def this(n:Int) = this(n,1)
所有 Scala 的辅助构造函数的第一个语句都为调用其它构造函数，也就是 this(…)，被调用的构造函数可以是主构造函数或是其它构造函数（最终会调用主构造函数），这样使得每个构造函数最终都会调用主构造函数，从而使得主构造函数称为创建类单一入口点。在 Scala 中也只有主构造函数才能调用基类的构造函数，这种限制有它的优点，使得 Scala 构造函数更加简洁和提高一致性。

隐式类型转换
上面我们定义 Rational 的加法，并重载+以支持整数，r + 2 ,当如果我们需要 2 + r 如何呢？ 下面的例子：

scala> val x =new Rational(2,3)
x: Rational = 2/3
scala> val y = new Rational(3,7)
y: Rational = 3/7
scala> val z = 4
z: Int = 4
scala> x + z
res0: Rational = 14/3
scala> x + 3
res1: Rational = 11/3
scala> 3 + x
<console>:10: error: overloaded method value + with alternatives:
  (x: Double)Double <and>
  (x: Float)Float <and>
  (x: Long)Long <and>
  (x: Int)Int <and>
  (x: Char)Int <and>
  (x: Short)Int <and>
  (x: Byte)Int <and>
  (x: String)String
 cannot be applied to (Rational)
              3 + x
                ^
可以看到 x+3 没有问题，3 + x 就报错了，这是因为整数类型不支持和 Rational 相加。我们不可能去修改 Int 的定义（除非你重写 Scala 的 Int 定义）以支持 Int 和 Rational 相加。如果你写过 .Net 代码，这可以通过静态扩展方法来实现，Scala 提供了类似的机制来解决这种问题。如果 Int 类型能够根据需要自动转换为 Rational 类型，那么 3 + x 就可以相加。Scala 通过 implicit def 定义一个隐含类型转换，比如定义由整数到 Rational 类型的转换如下：

implicit def intToRational(x:Int) = new Rational(x)
再重新计算 r+2 和 2 + r 的例子：

scala> val r = new Rational(2,3)
r: Rational = 2/3
scala> r + 2
res0: Rational = 8/3
scala> 2 + r
res1: Rational = 8/3
其实此时 Rational 的一个+重载方法是多余的， 当 Scala 计算 2+ r，发现 2（Int)类型没有 可以和 Rational 对象相加的方法，Scala 环境就检查 Int 的隐含类型转换方法是否有合适的类型转换方法，类型转换后的类型支持+r，一检查发现定义了由 Int 到 Rational 的隐含转换方法，就自动调用该方法，把整数转换为 Rational 数据类型，然后调用 Rational 对象的+ 方法。从而实现了 Rational 类或是 Int 类的扩展。关于 implicit def 的详细介绍将由后面的文章来说明，隐含类型转换在设计 Scala 库时非常有用。

过滤
某些时候，你不想枚举集合中的每个元素，而是只迭代某些符合条件的元素，在 Scala 中，你可以为 for 表达式添加一个过滤器–在 for 的括号内添加一个 if 语句，例如：修改前面枚举文件的例子，改成只列出 .scala 文件如下：

val filesHere = (new java.io.File(".")).listFiles
for( file   println(file)
如果有必要的话，你可以使用多个过滤器，只要添加多个 if 语句即可，比如，为保证前面列出的文件不是目录，可以添加一个 if，如下面代码：

val filesHere = (new java.io.File(".")).listFiles
for( file <-filesHere
   if file.isFile
   if file.getName.endsWith(".scala")
)  println(file)

枚举集合元素
这是使用 for 表示式的一个基本用法，和 Java 的 for 非常类似，比如下面的代码可以枚举当前目录下所有文件：

val filesHere = (new java.io.File(".")).listFiles
for( file <-filesHere)
  println(file)
其中如 file < – filesHere 的语法结构，在 Scala 中称为“生成器 (generator)”。 本例中 filesHere 的类型为 Array[File]，每次迭代 变量 file 会初始化为该数组中一个元素， File 的 toString()为文件的文件名，因此 println(file)打印出文件名。 Scala 的 for 表达式支持所有类型的集合类型，而不仅仅是数组，

生成新集合
for 表达式也可以用来生产新的集合，这是 Scala 的 for 表达式比 Java 的 for 语句功能强大的地方。它的基本语法如下：

for clauses yield body
关键字 yield 放在 body 的前面，for 没迭代一次，产生一个 body，yield 收集所有的 body 结果，返回一个 body 类型的集合。比如，前面列出所有 .scala 文件，返回这些文件的集合：

def scalaFiles =
  for {
    file     if file.getName.endsWith(".scala")
  } yield file

  注意：和 Java 异常处理不同的一点是，Scala 不需要你捕获 checked 的异常，这点和 C# 一样，也不需要使用 throw 来声明某个异常，当然如果有需要还是可以通过 @throw 来声明一个异常，但这不是必须的。

finally语句
Scala 也支持 finally 语句，你可以在 finally 块中添加一些代码，这些代码不管 try 块是否抛出异常，都会执行。比如，你可以在 finally 块中添加代码保证关闭已经打开的文件，而不管前面代码中是否出现异常。

import java.io.FileReader
val file = new FileReader("input.txt")
try {
  //use the file
} finally {
  file.close()
}
生成返回值
和大部分Scala 控制结构一样，Scala 的 try-catch-finally 也生成某个值，比如下面的例子尝试分析一个 URL，如果输入的 URL 无效，则使用缺省的 URL 链接地址：

import java.net.URL
import java.net.MalformedURLException
def urlFor(path:String) =
  try {
    new URL(path)
  } catch {
    case e: MalformedURLException =>
      new URL("http://www.scala-lang.org")
  }
通常情况下，finally 块用来做些清理工作，而不应该产生结果，但如果在 finally 块中使用 return 来返回某个值，这个值将覆盖 try-catch 产生的结果，比如：

scala> def f(): Int = try { return 1 } finally { return 2}
f: ()Int
scala> f
res0: Int = 2
而下面的代码：

scala> def g() :Int = try 1 finally 2
g: ()Int
scala> g
res0: Int = 1
结果却是 1，上面两种情况常常使得程序员产生困惑，因此关键的一点是避免在 finally 生成返回值，而只用来做些清理工作，比如关闭文件。

cala 的 Match 表达式支持从多个选择中选取其一，类似其它语言中的 switch 语句。通常来说，Scala 的 match 表达式支持任意的匹配模式，这种基本模式将在后面介绍，本篇介绍类似 switch 用法的 match 表达式，也是在多个选项中选择其一。

例如下面的例子从参数中读取食品的名称，然后根据食品的名称，打印出该和该食品搭配的食品，比如输入 ”salt”，与之对应的食品为”pepper”。如果是”chips”，那么搭配的就是“salsa”等等。

val firstArg = if (args.length >0 ) args(0) else ""
firstArg match {
  case "salt" => println("pepper")
  case "chips" => println("salsa")
  case "eggs" => println("bacon")
  case _ => println("huh?")
}
这段代码和 Java 的 switch 相比有几点不同：
一是任何类型的常量都可以用在 case 语句中，而不仅仅是 int 或是枚举类型。
二是每个 case 语句无需使用 break，Scala不支持“fall through”。
三是 Scala 的缺省匹配为”_”,其作用类似 java 中的 default。

而最关键的一点是 scala 的 match 表达式有返回值，上面的代码使用的是 println 打印，而实际上你可以使用表达式，比如修改上面的代码如下：

val firstArg = if (args.length >0 ) args(0) else ""
val friend = firstArg match {
  case "salt" => "pepper" 
  case "chips" => "salsa" 
  case "eggs" => "bacon" 
  case _ => "huh?" 
}
这段代码和前面的代码是等效的，不同的是后面这段代码 match 表达式返回结果。

使用纯函数化编程的一个方法是去掉 var 变量的使用，递归函数（回溯函数）的使用是通常使用的一个方法来去除循环结构中使用 var 变量。

也就是说可以在函数的内部再定义函数，如同定义一个局部变量，例如，修改前面的 ProcessFile 的例子如下：

import scala.io.Source
object LongLines {
  def processFile(filename: String, width: Int) {
    def processLine(filename:String,
     width:Int, line:String){
     if(line.length > width)
       println(filename + ":" +line.trim)
   }
    val source= Source.fromFile(filename)
    for (line <- source.getLines())
      processLine(filename,width,line)
   }
}
这个例子不私有成员函数 processLine 移动到 processFile 内部，成为一个局部函数，也正因为 processLine 变成了 processFile 的一个局部函数，因此 processLine 可以直接访问到 processFile 的参数 filename 和 width,因此代码可以进一步优化如下：

import scala.io.Source
object LongLines {
  def processFile(filename: String, width: Int) {
    def processLine(line:String){
     if(line.length > width)
       println(filename + ":" +line.trim)
   }
    val source= Source.fromFile(filename)
    for (line <- source.getLines())
      processLine(line)
   }
}
代码变得更简洁，是不是，局部函数的作用域和局部变量作用域一样，局部函数访问包含该函数的参数是非常常见的一种嵌套函数的用法。

这个函数可以进一步去掉参数的括号，这里的括号不起什么作用：

 x => x +1
Scala 还可以进一步简化，Scala 允许使用“占位符”下划线”_”来替代一个或多个参数，只要这个参数值函数定义中只出现一次，Scala编译器可以推断出参数。比如：

scala> val someNumbers = List ( -11, -10, - 5, 0, 5, 10)
someNumbers: List[Int] = List(-11, -10, -5, 0, 5, 10)
scala> someNumbers.filter(_ >0)
res0: List[Int] = List(5, 10)
可以看到简化后的函数定义为 > 0，你可以这样来理解，就像我们以前做过的填空题，“”为要填的空，Scala 来完成这个填空题，你来定义填空题。

有时，如果你使用_来定义函数，可能没有提供足够的信息给 Scala 编译器，此时 Scala 编译器将会报错，比如，定义一个加法函数如下：

scala> val f = _ + _
<console>:7: error: missing parameter type for expanded function ((x$1, x$2) => x$1.$plus(x$2))
       val f = _ + _
               ^
<console>:7: error: missing parameter type for expanded function ((x$1: <error>, x$2) => x$1.$plus(x$2))
       val f = _ + _
Scala 编译器无法推断出_的参数类型，就报错了，但如果你给出参数的类型，依然可以使用_来定义函数，比如：

scala> val f = (_ :Int ) + ( _ :Int)
f: (Int, Int) => Int = <function2>
scala> f (5,10)
res1: Int = 15
因为替代的参数在函数体中只能出现一次，因此多个“_”代表多个参数。第一个“_”代表第一个参数，第二个“\”代表第二个参数，以此类推。

函数–部分应用的函数

前面例子中我们使用“_” 来代替单个的参数，实际上你也可以使用“_”来代替整个参数列表，比如说，你可以使用 print 来代替 println ().someNumbers.foreach(println _)。

Scala 编译器自动将上面代码解释成：

someNumbers.foreach( x => println (x))
因此这里的“_” 代表了 Println 的整个参数列表，而不仅仅替代单个参数。当你采用这种方法使用“_”，你就创建了一个部分应用的函数(partially applied function)。 在 Scala 中，当你调用函数，传入所需参数，你就把函数“应用”到参数。 比如：一个加法函数。

scala> def sum = (_:Int) + (_ :Int) + (_ :Int)
sum: (Int, Int, Int) => Int
scala> sum (1,2,3)
res0: Int = 6
一个部分应用的函数指的是你在调用函数时，不指定函数所需的所有参数，这样你就创建了一个新的函数，这个新的函数就称为原始函数的部分应用函数，比如说我们固定 sum 的第一和第三个参数，定义如下的部分应用函数：

scala> val b = sum ( 1 , _ :Int, 3)
b: Int => Int = <function1>
scala> b(2)
res1: Int = 6
变量 b 的类型为一函数，具体类型为 Function1（带一个参数的函数），它是由 sum 应用了第一个和第三个参数，构成的。调用b(2），实际上调用 sum (1,2,3)。

再比如我们定义一个新的部分应用函数，只固定中间参数：

scala> val c = sum (_:Int, 2, _:Int)
c: (Int, Int) => Int = <function2>
scala> c(1,3)
res2: Int = 6
变量 c 的类型为 Function2,调用 c(1,3) 实际上也是调用 sum (1,2,3)。

在 Scala 中，如果你定义一个部分应用函数并且能省去所有参数，比如 println ，你也可以省掉“”本身，比如：

someNumbers.foreach(println _)
可以写成：

someNumbers.foreach(println)

到目前为止我们介绍的函数都只引用到传入的参数，假如我们定义如下的函数：

(x:Int) => x + more
这里我们引入一个自由变量 more。它不是所定义函数的参数，而这个变量定义在函数外面，比如：

var more =1
那么我们有如下的结果：

scala> var more =1
more: Int = 1
scala> val addMore = (x:Int) => x + more
addMore: Int => Int = <function1>
scala> addMore (100)
res1: Int = 101
这样定义的函数变量 addMore 成为一个“闭包”，因为它引用到函数外面定义的变量，定义这个函数的过程是将这个自由变量捕获而构成一个封闭的函数。有意思的是，当这个自由变量发生变化时，Scala 的闭包能够捕获到这个变化，因此 Scala 的闭包捕获的是变量本身而不是当时变量的值。

比如：

scala> more =  9999
more: Int = 9999
scala> addMore ( 10)
res2: Int = 10009
同样的，如果变量在闭包在发生变化，也会反映到函数外面定义的闭包的值。比如：

scala> val someNumbers = List ( -11, -10, -5, 0, 5, 10)
someNumbers: List[Int] = List(-11, -10, -5, 0, 5, 10)
scala> var sum =0
sum: Int = 0
scala> someNumbers.foreach ( sum += _)
scala> sum
res4: Int = -11
可以看到在闭包中修改 sum 的值，其结果还是传递到闭包的外面。

如果一个闭包所访问的变量有几个不同的版本，比如一个闭包使用了一个函数的局部变量（参数），然后这个函数调用很多次，那么所定义的闭包应该使用所引用的局部变量的哪个版本呢？ 简单的说，该闭包定义所引用的变量为定义该闭包时变量的值，也就是定义闭包时相当于保存了当时程序状态的一个快照。比如我们定义下面一个函数闭包：

scala> def makeIncreaser(more:Int) = (x:Int) => x + more
makeIncreaser: (more: Int)Int => Int
scala> val inc1=makeIncreaser(1)
inc1: Int => Int = <function1>
scala> val inc9999=makeIncreaser(9999)
inc9999: Int => Int = <function1>
scala> inc1(10)
res5: Int = 11
scala> inc9999(10)
res6: Int = 10009
当你调用 makeIncreaser(1)时，你创建了一个闭包，该闭包定义时 more的值为 1，而调用 makeIncreaser(9999)所创建的闭包的 more 的值为 9999。此后你也无法修改已经返回的闭包的 more 的值。因此 inc1 始终为加一，而 inc9999 始终为加 9999。

重复参数
Scala 在定义函数时允许指定最后一个参数可以重复（变长参数），从而允许函数调用者使用变长参数列表来调用该函数，Scala 中使用“*”来指明该参数为重复参数。例如：

scala> def echo (args: String *) =
     |   for (arg <- args) println(arg)
echo: (args: String*)Unit
scala> echo()
scala> echo ("One")
One
scala> echo ("Hello","World")
Hello
World
在函数内部，变长参数的类型，实际为一数组，比如上例的 String * 类型实际为 Array[String]。 然而，如今你试图直接传入一个数组类型的参数给这个参数，编译器会报错：

scala> val arr= Array("What's","up","doc?")
arr: Array[String] = Array(What's, up, doc?)
scala> echo (arr)
<console>:10: error: type mismatch;
 found   : Array[String]
 required: String
              echo (arr)
                    ^
为了避免这种情况，你可以通过在变量后面添加 _*来解决，这个符号告诉 Scala 编译器在传递参数时逐个传入数组的每个元素，而不是数组整体。

scala> echo (arr: _*)
What's
up
doc?
命名参数
通常情况下，调用函数时，参数传入和函数定义时参数列表一一对应。

scala> def  speed(distance: Float, time:Float) :Float = distance/time
speed: (distance: Float, time: Float)Float
scala> speed(100,10)
res0: Float = 10.0
使用命名参数允许你使用任意顺序传入参数，比如下面的调用：

scala> speed( time=10,distance=100)
res1: Float = 10.0
scala> speed(distance=100,time=10)
res2: Float = 10.0
缺省参数值
Scala 在定义函数时，允许指定参数的缺省值，从而允许在调用函数时不指明该参数，此时该参数使用缺省值。缺省参数通常配合命名参数使用，例如：

scala> def printTime(out:java.io.PrintStream = Console.out, divisor:Int =1 ) =
     | out.println("time = " + System.currentTimeMillis()/divisor)
printTime: (out: java.io.PrintStream, divisor: Int)Unit
scala> printTime()
time = 1383220409463
scala> printTime(divisor=1000)
time = 1383220422

使用高级函数可以帮助你简化代码，它支持创建一个新的程序控制结构来减低代码重复。比如，你打算写一个文件浏览器，你需要写一个 API 支持搜索给定条件的文件。首先，你添加一个方法，该方法可以通过查询包含给定字符串的文件，比如你可以查所有“.scala”结尾的文件。你可以定义如下的 API：

object FileMatcher {
  private def filesHere = (new java.io.File(".")).listFiles
  def filesEnding(query : String) =
    for (file <-filesHere; if file.getName.endsWith(query))
      yield file
}
filesEnding 方法从本地目录获取所有文件（方法 filesHere)，然后使用过滤条件（文件以给定字符串结尾）输出给定条件的文件。

到目前为止，这代码实现非常好也没有什么重复的代码。后来，你有需要使用新的过滤条件，文件名包含指定字符串，而不仅仅以某个字符串结尾的文件列表。你有实现了下面的 API。

def filesContaining( query:String ) =
    for (file <-filesHere; if file.getName.contains(query))
      yield file
filesContaining 和 filesEnding 的实现非常类似，不同点在于一个使用 endsWith,另一个使用 contains 函数调用。有过了一段时间，你有想支持使用正则表达式来查询文件，你有实现了下面的对象方法：

def filesRegex( query:String) =
   for (file <-filesHere; if file.getName.matches(query))
      yield file
这三个函数的算法非常类似，所不同的是过滤条件稍有不同，在 Scala 中我们可以定义一个高阶函数，将这三个不同过滤条件抽象称一个函数作为参数传给搜索算法，我们可以定义这个高阶函数如下：

def filesMatching( query:String, 
    matcher: (String,String) => Boolean) = {
    for(file <- filesHere; if matcher(file.getName,query))
      yield file
   }
这个函数的第二个参数 matcher 的类型也为函数（如果你熟悉 C#，类似于 delegate)，该函数的类型为 (String,String ) =>Boolean，可以匹配任意使用两个 String 类型参数，返回值类型为 Boolean 的函数。使用这个辅助函数，我们可以重新定义 filesEnding，filesContaining 和 filesRegex。

def filesEnding(query:String) =
   filesMatching(query,_.endsWith(_))
def filesContaining(query:String)=
   filesMatching(query,_.contains(_))
def filesRegex(query:String) =
   filesMatching(query,_.matches(_))
这个新的实现和之前的实现已经简化了不少，实际上代码还可以简化，我们注意到参数 query 在 filesMatching 的作用只是把它传递给 matcher 参数，这种参数传递实际也是无需的，简化后代码如下：

object FileMatcher {
  private def filesHere = (new java.io.File(".")).listFiles
  def filesMatching(
    matcher: (String) => Boolean) = {
    for(file <- filesHere; if matcher(file.getName))
      yield file
   }
  def filesEnding(query:String) =
   filesMatching(_.endsWith(query))
def filesContaining(query:String)= 
   filesMatching(_.contains(query))
def filesRegex(query:String) = 
   filesMatching(_.matches(query))
}
函数类型参数 .endsWith(query)，.contains(query)和_.matches(query)为函数闭包，因为它们绑定了一个自由变量 query，因此我们可以看到闭包也可以用来简化代码。

前面我们说过，Scala 允许程序员自己新创建一些控制结构，并且可以使得这些控制结构在语法看起来和 Scala 内置的控制结构一样，在 Scala 中需要借助于柯里化(Currying)，柯里化是把接受多个参数的函数变换成接受一个单一参数（最初函数的第一个参数）的函数，并且返回接受余下的参数而且返回结果的新函数的技术。

下面先给出一个普通的非柯里化的函数定义，实现一个加法函数：

scala> def plainOldSum(x:Int,y:Int) = x + y
plainOldSum: (x: Int, y: Int)Int
scala> plainOldSum(1,2)
res0: Int = 3
下面在使用“柯里化”技术来定义这个加法函数，原来函数使用一个参数列表，“柯里化”把函数定义为多个参数列表：

scala> def curriedSum(x:Int)(y:Int) = x + y
curriedSum: (x: Int)(y: Int)Int
scala> curriedSum (1)(2)
res0: Int = 3
当你调用 curriedSum (1)(2)时，实际上是依次调用两个普通函数（非柯里化函数），第一次调用使用一个参数 x，返回一个函数类型的值，第二次使用参数y调用这个函数类型的值，我们使用下面两个分开的定义在模拟 curriedSum 柯里化函数：

首先定义第一个函数：

scala> def first(x:Int) = (y:Int) => x + y
first: (x: Int)Int => Int
然后我们使用参数1调用这个函数来生成第二个函数（回忆前面定义的闭包）。

scala> val second=first(1)
second: Int => Int = <function1>
scala> second(2)
res1: Int = 3
first，second的定义演示了柯里化函数的调用过程，它们本身和 curriedSum 没有任何关系，但是我们可以使用 curriedSum 来定义 second，如下：

scala> val onePlus = curriedSum(1)_
onePlus: Int => Int = <function1>
下划线“_” 作为第二参数列表的占位符， 这个定义的返回值为一个函数，当调用时会给调用的参数加一。

scala> onePlus(2)
res2: Int = 3
通过柯里化，你还可以定义多个类似 onePlus 的函数，比如 twoPlus

scala> val twoPlus = curriedSum(2) _
twoPlus: Int => Int = <function1>
scala> twoPlus(2)
res3: Int = 4

你在写代码时，如果发现某些操作需要重复多次，你就可以试着将这个重复操作写成新的控制结构，在前面我们定义过一个 filesMatching 函数

def filesMatching(
    matcher: (String) => Boolean) = {
    for(file <- filesHere; if matcher(file.getName))
      yield file
   }
如果我们把这个函数进一步通用化，可以定义一个通用操作如下：

打开一个资源，然后对资源进行处理，最后释放资源，你可以为这个“模式”定义一个通用的控制结构如下：

def withPrintWriter (file: File, op: PrintWriter => Unit) {
  val writer=new PrintWriter(file)
  try{
    op(writer)
  }finally{
    writer.close()
  }
}
使用上面定义，我们使用如下调用：

withPrintWriter(
   new File("date.txt"),
   writer => writer.println(new java.util.Date)
)
使用这个方法的优点在于 withPrintWriter，而不是用户定义的代码，withPrintWriter 可以保证文件在使用完成后被关闭，也就是不可能发生忘记关闭文件的事件。这种技术成为“租赁模式”，这是因为这种类型的控制结构，比如 withPrintWriter 将一个 PrintWriter 对象“租”给 op 操作，当这个 op 操作完成后，它通知不再需要租用的资源，在 finally 中可以保证资源被释放，而无论 op 是否出现异常。

这里调用语法还是使用函数通常的调用方法，使用（）来列出参数，在 Scala 中如果你调用函数只有一个参数，你可以使用{}来替代().比如下面两种语法是等价的：

scala> println ("Hello,World")
Hello,World
scala> println { "Hello,world" }
Hello,world
上面第二种用法，使用{}替代了()，但这只适用在使用一个参数的调用情况。 前面定义 withPrintWriter 函数使用了两个参数，因此不能使用{}来替代()，但如果我们使用柯里化重新定义下这个函数如下：

import scala.io._
import java.io._
def withPrintWriter (file: File)( op: PrintWriter => Unit) {
  val writer=new PrintWriter(file)
  try{
    op(writer)
  }finally{
    writer.close()
  }
}
将一个参数列表，变成两个参数列表，每个列表含一个参数，这样我们就可以使用如下语法来调用

withPrintWriter 
val file = new File("date.txt")
withPrintWriter(file){
  writer => writer.println(new java.util.Date)
}
第一个参数我们还是使用（）（我们也可以使用{})，第二个参数我们使用{}来替代()，这样修改过的代码使得 withPrintWriter 看起来和 Scala 内置的控制结构语法一样。

传名参数

上篇我们使用柯里化函数定义一个控制机构 withPrintWriter，它使用时语法调用有如 Scala 内置的控制结构：

val file = new File("date.txt")
withPrintWriter(file){
  writer => writer.println(new java.util.Date)
}
不过仔细看一看这段代码，它和 scala 内置的 if 或 while 表达式还是有些区别的，withPrintWrite r的{}中的函数是带参数的含有“writer=>”。 如果你想让它完全和 if 和 while 的语法一致，在 Scala 中可以使用传民参数来解决这个问题。

注：我们知道通常函数参数传递的两种模式，一是传值，一是引用。而这里是第三种按名称传递。

下面我们以一个具体的例子来说明传名参数的用法：

var assertionsEnabled=true
def myAssert(predicate: () => Boolean ) =
  if(assertionsEnabled && !predicate())
    throw new AssertionError
这个 myAssert 函数的参数为一个函数类型，如果标志 assertionsEnabled 为 True 时，mymyAssert 根据 predicate 的真假决定是否抛出异常，如果 assertionsEnabled 为 false，则这个函数什么也不做。

这个定义没什么问题，但调用起来看起来却有些别扭，比如：

myAssert(() => 5 >3 )
还需要 ()=>，你可以希望直接使用 5>3，但此时会报错：

scala> myAssert(5 >3 )
<console>:10: error: type mismatch;
 found   : Boolean(true)
 required: () => Boolean
              myAssert(5 >3 )
此时，我们可以把按值传递（上面使用的是按值传递，传递的是函数类型的值）参数修改为按名称传递的参数，修改方法，是使用=>开始而不是 ()=>来定义函数类型，如下：

def myNameAssert(predicate:  => Boolean ) =
  if(assertionsEnabled && !predicate)
    throw new AssertionError
此时你就可以直接使用下面的语法来调用 myNameAssert：

myNameAssert(5>3)
此时就和 Scala 内置控制结构一样了，看到这里，你可能会想我为什么不直接把参数类型定义为 Boolean，比如：

def boolAssert(predicate: Boolean ) =
  if(assertionsEnabled && !predicate)
    throw new AssertionError
调用也可以使用

boolAssert(5>3)
和 myNameAssert 调用看起来也没什么区别，其实两者有着本质的区别，一个是传值参数，一个是传名参数，在调用 boolAssert(5>3)时，5>3 是已经计算出为 true，然后传递给 boolAssert 方法，而 myNameAssert(5>3)，表达式 5>3 没有事先计算好传递给 myNameAssert，而是先创建一个函数类型的参数值，这个函数的 apply 方法将计算5>3，然后这个函数类型的值作为参数传给 myNameAssert。

因此这两个函数一个明显的区别是，如果设置 assertionsEnabled 为 false，然后试图计算 x/0 ==0，

scala> assertionsEnabled=false
assertionsEnabled: Boolean = false
scala> val x = 5
x: Int = 5
scala> boolAssert ( x /0 ==0)
java.lang.ArithmeticException: / by zero
  ... 32 elided
scala> myNameAssert ( x / 0 ==0)
可以看到 boolAssert 抛出 java.lang.ArithmeticException: / by zero 异常，这是因为这是个传值参数，首先计算 x /0 ，而抛出异常，而 myNameAssert 没有任何显示，这是因为这是个传名参数，传入的是一个函数类型的值，不会先计算 x /0 ==0，而在 myNameAssert 函数体内，由于 assertionsEnabled 为 false，传入的 predicate 没有必要计算(短路计算），因此什么也不会打印。如果我们把 myNameAssert 修改下，把 predicate 放在前面:

scala> def myNameAssert1(predicate:  => Boolean ) =
     |   if( !predicate && assertionsEnabled )
     |     throw new AssertionError
myNameAssert1: (predicate: => Boolean)Unit
scala> myNameAssert1 ( x/0 ==0)
java.lang.ArithmeticException: / by zero
  at $anonfun$1.apply$mcZ$sp(<console>:11)
  at .myNameAssert1(<console>:9)
  ... 32 elided
这个传名参数函数也抛出异常（你可以想想是为什么？）

前面的 withPrintWriter 我们暂时没法使用传名参数，去掉 writer=>，否则就难以实现“租赁模式”，不过我们可以看看下面的例子，设计一个 withHelloWorld 控制结构，这个 withHelloWorld 总打印一个“hello,world”

import scala.io._
import java.io._
def withHelloWorld ( op: => Unit) {
  op   
  println("Hello,world")
}
val file = new File("date.txt")
withHelloWorld{
  val writer=new PrintWriter(file)
  try{
   writer.println(new java.util.Date)
  }finally{
    writer.close()
  }
}
withHelloWorld {
  println ("Hello,Guidebee")
} 
Hello,world 
Hello,Guidebee
Hello,world
可以看到 withHelloWorld 的调用语法和 Scala 内置控制结构非常象了。

作为下一步，我们将向 Element 添加显示宽度和高度的方法，height 方法返回 contents 里的行数。width 方法返回第一行的长度，或如果元素没有行记录，返回零。

abstract class Element { 
  def contents: Array[String] 
  def height: Int = contents.length 
  def width: Int = if (height == 0) 0 else contents(0).length 
}
请注意 Element 的三个方法没一个有参数列表，甚至连个空列表都没有,这种无参数方法在 Scala 里是非常普通的。相对的，带有空括号的方法定义，如 def height(): Int，被称为空括号方法：(empty-paren method)。

Scala的惯例是在方法不需要参数并且只是读取对象状态时使用无参数方法。

此外，我们也可以使用成员变量来定义 width 和 height，例如：

abstract class Element { 
  def contents: Array[String] 
  val height = contents.length 
  val width = if (height == 0) 0 else contents(0).length 
}
从使用这个类的客户代码来说，这两个实现是等价的，唯一的差别是使用成员变量的方法调用速度要快些，因为字段值在类被初始化的时候被预计算，而方法调用在每次调用的时候都要计算。换句话说，字段在每个 Element 对象上需要更多的内存空间。

特别是如果类的字段变成了访问函数，且访问函数是纯函数的，就是说它没有副作用并且不依赖于可变状态，那么类 Element 的客户不需要被重写。这称为统一访问原则： uniform access principle， 就是说客户代码不应受通过字段还是方法实现属性的决定的影响。

Scala 代码可以调用 Java 函数和类，而 Java 没有使用“统一访问原则”，因此 Java 里是 string.length()，不是 string.length。为了解决这个问题，Scala 对于无参数函数和空括号函数的使用上并不是区分得很严格。也就是，你可以用空括号方法重载无参数方法，并且反之亦可。你还可以在调用任何不带参数的方法时省略空的括号。例如，下面两行在 Scala里都是合法的：

Array(1, 2, 3).toString 
"abc".length
原则上 Scala 的函数调用中可以省略所有的空括号。但如果使用的函数不是纯函数，也就是说这个不带参数的函数可能修改对象的状态或是我们需要利用它的一些副作用（比如打印到屏幕，读写 I/o)，一般的建议还是使用空括号，比如：

"hello".length // 没有副作用，所以无须() 
println() // 最好别省略()
总结起来，Scala 里定义不带参数也没有副作用的方法为无参数方法，也就是说，省略空的括号，是鼓励的风格。另一方面，永远不要定义没有括号的带副作用的方法，因为那样的话方法调用看上去会像选择一个字段。

和 Java 稍有不同的一点是，Scala 中成员函数和成员变量地位几乎相同，而且也处在同一个命名空间，也就是 Scala 中不允许定义同名的成员函数和成员变量，但带来的一个好处是，可以使用成员变量来重载一个不带参数的成员函数。比如，接着前面的例子，你可以通过一个成员变量来实现基类中定义的抽象函数 contents。

class ArrayElement(conts: Array[String]) extends Element {
val contents: Array[String] = conts
}
可以看到，这是使用成员变量来实现基类中不带参数的抽象函数的一个非常恰当的例子。 Scala 中的这种实现是 Java 语言所不支持的，一般来说只有两个不同的命名空间来定义类，而 Java 可以有四个，Scala 支持的两个命名空间如下：

值（字段，方法，包还有单例对象）
类型（类和Trait名）
Scala 把字段和方法放进同一个命名空间的理由很清楚，因为这样你就可以使用 val 重载无参数的方法，

我们回到前面定义的类 ArrayElement，它有一个参数 conts，其唯一的目的是用来复制到 contents 成员变量。而参数名称 conts 是为了让它看起来和成员变量 contents 类似，而有不至于和成员变量名冲突。

Scala 支持使用参数化成员变量，也就是把参数和成员变量定义合并到一起来避免上述代码：

class ArrayElement(val contents: Array[String]) 
  extends Element {
}
要注意的是，现在参数 contents 前面加上了 val 关键字，这是前面使用同名参数和同名成员变量的一个缩写形式。使用 val 定义了一个无法重新赋值的成员变量。这个成员变量初始值为参数的值，可以在类的外面访问这个成员变量。它的一个等效的实现如下：

class ArrayElement(val x123: Array[String]) 
  extends Element {
   val contents: Array[String] = x123
}
Scala 也允许你使用 var 关键字来定义参数化成员变量，使用 var 定义的成员变量，可以重新赋值。 此外 Scala 也允许你使用 private，protected，override 来修饰参数化成员变量，和你定义普通的成员变量的用法一样。 比如：

class Cat {
  val dangerous =false
}
class Tiger (
  override val dangerous: Boolean,
  private var age: Int
) extends Cat
这段代码中 Tiger 的定义其实为下面类定义的一个缩写：

class Tiger(param1: Boolean, param2: Int) extends Cat { 
    override val dangerous = param1 
    private var age = param2 
} 
两个成员都初始化自相应的参数。我们任意选择了这些参数名，param1 和 param2。重要的是它们不会与范围内的任何其它名称冲突。

在前面的例子 LineElement 使用了 override 来修饰 width 和 height 成员变量，在 Scala 中需要使用 override 来重载父类的一个非抽象成员，实现抽象成员无需使用 override，如果子类没有重载父类中的成员，不可以使用 override 修饰符。

这个规则可以帮助编译器发现一些难以发现的错误，可以增强系统安全进化。比如，如果你把 height 拼写错误为 hight，使用 override 编译器会报错

root@mail:~/scala# scalac demo.scala 
demo.scala:13: error: method hight overrides nothing
  override def hight = 1
               ^
one error found
这个规则对于系统的演讲尤为重要，假设你定义了一个 2D 图形库。你把它公开，并广泛使用。库的下一个版本里你想在你的基类 Shape 里增加一个新方法：

def hidden(): Boolean
你的新方法将被用在许多画图方法中去决定是否需要把形状画出来，这将可以大大提高系统绘图的性能，但你不可以冒着破坏客户代码的风险做这件事。毕竟客户说不定已经使用不同的 hidde n实现定义了 Shape 的子类。或许客户的方法实际上是让对象消失而不是检测是否对象是隐藏的。因为这两个版本的 hidden 互相重载，你的画图方法将停止对象的消失，这可真不是你想要的！

如果图形库和它的用户是用 Scala 写的，那么客户的 hidden 原始实现就不会有 override 修饰符，因为这时候还没有另外一个使用那个名字的方法。一旦你添加了 hidden 方法到你 Shape 类的第二个版本，客户的重编译将给出像下列这样的错误：

.../Shapes.scala:6: error: error overriding method 
        hidden in class Shape of type ()Boolean; 
method hidden needs 'override' modifier 
def hidden(): Boolean =
也就是说，代之以错误的执行，你的客户将得到一个编译期错误，这常常是更可取的。

我们接着实现类 Element 的其它方法，如 above, beside 和 toString 方法。

above 方法，意味着把一个布局元素放在另外一个布局元素的上方，也就是把这两个元素的 contents 的内容连接起来。我们首先实现 above 函数的第一个版本：

def above(that: Element) :Element =
    new ArrayElement(this.contents ++ that.contents)
Scala 中 Array 使用 Java Array 来实现，但添加很多其它方法，尤其是 Scala 中 Array 可以转换为 scala.Seq 类的实例对象，scala.Seq 为一个序列结构并提供了许多方法来访问和转换这个序列。

实际上上面 above 的实现不是十分有效，因为它不允许你把不同长度的布局元素叠加到另外一个布局元素上面，但就目前来说，我们还是暂时使用这个实现，只使用同样长度的布局元素，后面再提供这个版本的增强版本。

下面我们再实现类 Element 的另外一个 beside 方法，把两个布局元素并排放置，和前面一样，为简单起见，我们暂时只考虑相同高度的两个布局元素：

def beside(that: Element) :Element = {
    val contents = new Array[String](this.contents.length)
    for(i <- 0 until this.contents.length)
      contents(i)=this.contents(i) + that.contents(i)
    new ArrayElement(contents)
  }
尽管上面的实现满足 beside 要求，但采用的还是指令式编程，我们使用函数说编程可以实现下面简化代码：

def beside(that: Element) :Element = {
    new ArrayElement(
      for(
        (line1,line2) <- this.contents zip that.contents
      ) yield line1+line2
    )
  } 
这里我们使用了 Array 的 zip 操作符，可以用来将两个数组转换成二元组的数组，zip 分别取两个数组对应的元素组成一个新的二元祖，比如：

scala> Array( 1,2,3) zip Array("a","b")
res0: Array[(Int, String)] = Array((1,a), (2,b))
如果一个数组长度大于另外一个数组，多余的元素被忽略。 for 的 yield 部分用来构成一个新元素。

最后我们实现 Element 的 toString 方法用来显示布局元素的内容：

override def toString = contents mkString "\n"
这里使用 mkString 函数，这个函数可以应用到任何序列数据结构（包括数组），也就是把 contents 的每个元素调用 toString，然后使用“\n”分隔。

前面我们介绍了 Scala 的类的继承，本篇我们介绍 Scala 语言自身定义的类的层次关系，在 Scala 中，所有的类都有一个公共的基类称为 Any，此外还定义了所有类的子类 Nothing，下面的图给出的 Scala定义的类层次关系的一个概要：



由于所有的类都继承自 Any，因此 Scala 中的对象都可以使用==,!=,或 equals 来比较，使用##或 hashCode 给出 hash 值，使用 toString 转为字符串。Any 的==和!=定位为 fianl，因此不可以被子类重载。==实际上和 equals 等价，!=和 equals 的否定形式等价，因此重载 equals 可以修改==和!=的定义。

根类 Any 有两个子类：AnyVal 和 AnyRef。AnyVal 是 Scala 里每个内建值类型的父类。有九个这样的值类型：Byte，Short，Char，Int，Long，Float，Double，Boolean 和 Unit。其中的前八个对应到 Java 的基本数值类型，它们的值在运行时表示成 Java 的类型。Scala 里这些类的实例都写成字面量。例如，42 是 Int 的实例，’x’是 Char 的实例，false 是 Boolean 的实例。值类型都被定义为即是抽象的又是 final 的，你不能使用 new 创造这些类的实例。

scala> new Int
<console>:8: error: class Int is abstract; cannot be instantiated
              new Int
              ^
scala> 
另一个值类，Unit，大约对应于 Java 的 void 类型；被用作不返回任何有趣结果的方法的结果类型。 Unit 只有一个实例值，被写作().

值类支持作为方法的通用的数学和布尔操作符。例如，Int 有名为+和*的方法，Boolean 有名为||和&&的方法。值类也从类 Any 继承所有的方法。你可以在解释器里测试:

scala> 42 toString
res3: String = 42
scala> 42.hashCode
res6: Int = 42
可以看到 Scala 的值类型之间的关系是扁平的，所有的值类都是 scala.AnyVal 的子类型，但是它们不是互相的子类。代之以它们不同的值类类型之间可以隐式地互相转换。例如，需要的时候，类 scala.Int 的实例可以自动放宽（通过隐式转换）到类 scala.Long 的实例。隐式转换还用来为值类型添加更多的功能。例如，类型 Int 支持以下所有的操作：

scala> 42 max 43
res0: Int = 43
scala> 42 min 43
res1: Int = 42
scala> 1 until 5
res2: scala.collection.immutable.Range = Range(1, 2, 3, 4)
scala> 1 to 5 
res3: scala.collection.immutable.Range.Inclusive = Range(1, 2, 3, 4, 5)
scala> 3.abs
res4: Int = 3
scala> (-3).abs
res5: Int = 3
这里解释其工作原理：方法 min，max，until，to 和 abs 都定义在类 scala.runtime.RichInt 里，并且有一个从类 Int 到 RichInt 的隐式转换。当你在 Int 上调用没有定义在 Int 上但定义在 RichInt 上的方法时，这个转换就被应用了。

类 Any 的另一个子类是类 AnyRef。这个是 Scala 里所有引用类的基类。正如前面提到的，在 Java 平台上 AnyRef 实际就是类 java.lang.Object 的别名。因此 Java 里写的类和 Scala 里写的都继承自 AnyRef。如此说来，你可以认为 java.lang.Object 是 Java 平台上实现 AnyRef 的方式。因此，尽管你可以在 Java 平台上的 Scala 程序里交换使用 Object 和 AnyRef，推荐的风格是在任何地方都只使用 AnyRef。

Scala 类与 Java 类不同在于它们还继承自一个名为 ScalaObject 的特别的 Marker Trait（Trait 我们在后面再进一步解释）。

可以看到 Scala 的值类型之间的关系是扁平的，所有的值类都是 scala.AnyVal 的子类型，但是它们不是互相的子类。代之以它们不同的值类类型之间可以隐式地互相转换。例如，需要的时候，类 scala.Int 的实例可以自动放宽（通过隐式转换）到类 scala.Long 的实例。隐式转换还用来为值类型添加更多的功能。例如，类型 Int 支持以下所有的操作：

scala> 42 max 43
res0: Int = 43
scala> 42 min 43
res1: Int = 42
scala> 1 until 5
res2: scala.collection.immutable.Range = Range(1, 2, 3, 4)
scala> 1 to 5 
res3: scala.collection.immutable.Range.Inclusive = Range(1, 2, 3, 4, 5)
scala> 3.abs
res4: Int = 3
scala> (-3).abs
res5: Int = 3
这里解释其工作原理：方法 min，max，until，to 和 abs 都定义在类 scala.runtime.RichInt 里，并且有一个从类 Int 到 RichInt 的隐式转换。当你在 Int 上调用没有定义在 Int 上但定义在 RichInt 上的方法时，这个转换就被应用了。

类 Any 的另一个子类是类 AnyRef。这个是 Scala 里所有引用类的基类。正如前面提到的，在 Java 平台上 AnyRef 实际就是类 java.lang.Object 的别名。因此 Java 里写的类和 Scala 里写的都继承自 AnyRef。如此说来，你可以认为 java.lang.Object 是 Java 平台上实现 AnyRef 的方式。因此，尽管你可以在 Java 平台上的 Scala 程序里交换使用 Object 和 AnyRef，推荐的风格是在任何地方都只使用 AnyRef。

Scala 类与 Java 类不同在于它们还继承自一个名为 ScalaObject 的特别的 Marker Trait（Trait 我们在后面再进一步解释）。

而，有些情况你需要使用引用相等代替用户定义的相等。例如，某些时候效率是首要因素，你想要把某些类哈希合并： hash cons 然后通过引用相等比较它们的实例，为这种情况，类 AnyRef 定义了附加的 eq 方法，它不能被重载并且实现为引用相等（也就是说，它表现得就像 Java 里对于引用类型的==那样）。同样也有一个 eq 的反义词，被称为 ne。例如：

scala> val x =new String("abc")
x: String = abc
scala> val y = new String("abc")
y: String = abc
scala> x == y
res0: Boolean = true
scala> x eq y
res1: Boolean = false
scala> x ne y
res2: Boolean = true

在 Scala中Trait 为重用代码的一个基本单位。一个 Traits 封装了方法和变量，和 Interface 相比，它的方法可以有实现，这一点有点和抽象类定义类似。但和类继承不同的是，Scala 中类继承为单一继承，也就是说子类只能有一个父类。当一个类可以和多个 Trait 混合，这些 Trait 定义的成员变量和方法也就变成了该类的成员变量和方法，由此可以看出 Trait 集合了 Interface 和抽象类的优点，同时又没有破坏单一继承的原则。

定义一个 Trait 的方法和定义一个类的方法非常类似，除了它使用 trait 而非 class 关键字来定义一个 trait。

trait Philosophical{
  def philosophize() {
    println("I consume memeory, therefor I am!")
  }
}
这个 Trait 名为 Philosophical。它没有声明基类，因此和类一样，有个缺省的基类 AnyRef。它定义了一个方法，叫做 philosophize。这是个简单的 Trait，仅够说明 Trait 如何工作。

一但定义好 Trait，它就可以用来和一个类混合，这可以使用 extends 或 with 来混合一个 trait。例如：

class Frog extends Philosophical{
  override def toString="gree"
}

如果你需要把某个 Trait 添加到一个有基类的子类中，使用 extends 继承基类，而可以通过 with 添加 Trait。比如：

class Animal
class Frog extends Animal with Philosophical{
  override def toString="green"
}
还是和 Interface 类似，可以为某个类添加多个 Trait 属性，此时使用多个 with 即可，比如：

class Animal
trait HasLegs 
class Frog extends Animal with Philosophical with HasLegs{
  override def toString="green"
}

目前为止你看到的例子中，类 Frog 都继承了 Philosophical 的 philosophize 实现。此外 Frog 也可以重载 philosophize 方法。语法与重载基类中定义的方法一样。

class Animal
trait HasLegs 
class Frog extends Animal with Philosophical with HasLegs{
  override def toString="green"
  def philosophize() {
    println("It ain't easy being " + toString + "!")
  }
}
因为 Frog 的这个新定义仍然混入了特质 Philosophize，你仍然可以把它当作这种类型的变量使用。但是由于 Frog 重载了 Philosophical 的 philosophize 实现，当你调用它的时候，你会得到新的回应：

scala> val phrog:Philosophical = new Frog
phrog: Philosophical = green
scala> phrog.philosophize
It ain't easy being green!
这时你或许推导出以下结论：Trait 就像是带有具体方法的 Java 接口，不过其实它能做的更多。Trait 可以，比方说，声明字段和维持状态值。实际上，你可以用 Trait 定义做任何用类定义做的事，并且语法也是一样的，除了两点。第一点，Trait 不能有任何“类”参数，也就是说，传递给类的主构造器的参数。换句话说，尽管你可以定义如下的类：

class Point(x: Int, y: Int)
但下面的 Trait 定义直接报错：

scala> trait NoPoint(x:Int,y:Int)
<console>:1: error: traits or objects may not have parameters
       trait NoPoint(x:Int,y:Int)


下面的例子使用=>重命名类型：

import Fruits.{Apple=>MaIntosh,Orange}
同样重命名也可以重新定义包名称，比如:

import java.{sql => S}
将引入的包 java.sql 该名为 java.S 因此可以使用 S.Date 来代替 sql.Date。

如果需要隐藏某个类型，可以使用 Type => _ ，将某个类型改名为_就达到隐藏某个类型的效果，比如:

import Fruits.{Apple=>_,_}
这个引用，引入 Fruits 中除 Apple 之外的其它类型。

