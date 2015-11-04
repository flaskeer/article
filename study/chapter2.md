

那代码块中的$it是什么呢？在这个上下文中，它代表进行循环时的索引值。upto()方法
接受一个闭包作为参数。如果闭包只需要一个参数，在Groovy中则可以使用默认的名字it来表
示该参数。（先记住这一点，第4章将更详细地讨论闭包。）变量it前面的$让print()方法打印
该变量的值，而非打印it这两个字符。利用该特性，我们可以在字符串中嵌入表达式，第5章有
此类用法。


 0.step(10,2) {print "$it "}
0 2 4 6 8

 println "svn help".execute().text

 println "".execute().getClass().name
java.lang.ProcessImpl

 println "cmd /C: dir".execute().text

字符串反转
def foo(str){
    str?.reverse()
}
println foo('eval')
println foo(null)
lave
null

 def openFile(fileName){
    new FileInputStream(fileName)
}
try{
    openFile("nofile")
}catch(FileNotFoundException ex){
    println "Oops:" + ex
}
Oops:java.io.FileNotFoundException: nofile (系统找不到指定的文件。)
</code>
>Groovy还有其他一些使这门语言更为轻量级、更为易用的特性，试举几例如下。
 return语句几乎总是可选的（参见2.11.1节）。
 尽管可以使用分号分隔语句，但它几乎总是可选的（参见2.11.6节）。
 方法和类默认是公开（public）的。
 ?.操作符只有对象引用不为空时才会分派调用。
 可以使用具名参数初始化JavaBean（参见2.2节）。
 Groovy不强迫我们捕获自己不关心的异常，这些异常会被传递给代码的调用者。
 静态方法内可以使用this来引用Class对象。在下面的例子中，learn()方法返回的是
Class对象，所以可以使用链式调用：

javaBean
> class Car{
    def miles = 0
    final year
    Car(theYear){year = theYear}
}
Car car = new Car(2008)
println "Year :$car.year"
println "Miles:$car.miles"
car.miles = 25
println "Miles:$car.miles"
Year :2008
Miles:0
Miles:25

>def在这个上下文中声明了一个属性。我们可以像例子中这样使用def声明属性，还可以像
int miles或int miles = 0这样给出类型（以及可选的值）。Groovy会在背后默默地为其创建
一个访问器和一个更改器（就像在Java中，如果没有编写任何构造器，则Java编译器会创建一个）。
当在代码中调用miles时，其实并非引用一个字段，而是调用该属性的访问器。要把属性设置为
只读的，需要使用final来声明该属性，这和Java中一样。在这种情况下，Groovy会为该属性提
供一个访问器，但不提供更改器。修改final字段的任何尝试都会导致异常。可以根据需要向声
明中加入类型信息。可以把字段标记为private，但是Groovy并不遵守这一点①。因此，如果想
把变量设置为私有的，必须实现一个拒绝任何修改的更改器。

Groovy中可以灵活地初始化一个JavaBean类。在构造对象时，可以简单地以逗号分隔的名
值对来给出属性值。如果类有一个无参构造器，该操作会在构造器之后执行。① 也可以设计自己
的方法，使其接受具名参数。要利用这一特性，需要把第一个形参定义为Map



> class Robot{
    def type,height,width
    def access(location,weight,fragile){
        println "Received fragile? $fragile,weigt:$weight,location:$location"
    }
}
robot = new Robot(type:'arm',width: 10,height: 40)
println "$robot.type,$robot.height,$robot.width"
robot.access(x:30,y:20,z:10,50,true)
robot.access(x:20,50,true)
arm,40,10
Received fragile? true,weigt:50,location:[x:30, y:20, z:10]
Received fragile? true,weigt:50,location:[x:20]

如果发送的实参的个数多于方法的形参的个数，而且多出的实参是名值对，那么Groovy会假
设方法的第一个形参是一个Map，然后将实参列表中的所有名值对组织到一起，作为第一个形参
的值。之后，再将剩下的实参按照给出的顺序赋给其余形参，正如我们在输出中看到的那样。


Groovy中可以把方法和构造器的形参设为可选的。实际上，我们想设置多少就可以设置多少，
但这些形参必须位于形参列表的末尾。利用这一特性，可以在演进式设计中向已有方法添加新的
形参。
要定义可选形参，只需要在形参列表中给它赋上一个值

> def log(x,base = 10){
    Math.log(x) /Math.log(base)
}
println log(1024)
3.0102999566398116
> def log(x,base = 10){
    Math.log(x) /Math.log(base)
}
println log(1024,2)
10.0

Groovy还会把末尾的数组形参视作可选的。所以在下面的例子中，可以为最后一个形参提供
零个或多个值：

> def task(name,String[] details){
    println "$name - $details"
}
task 'call','124','323465436'
call - [124, 323465436]

#返回多条赋值语句
> def splitName(fullName) {fullName.split(' ')}
def (firstName,lastName) = splitName('bibd sde')
println "$firstName,$lastName"
bibd,sde

还可以使用该特性来交换变量，无需创建中间变量来保存被交换的值，只需将欲交换的变量
放在圆括号内，置于赋值表达式左侧，同时将它们以相反顺序放于方括号内，置于右侧即可。

> def name1 = "Tom"
def name2 = "John"
(name1,name2) = [name2,name1]
println "$name1,$name2"
John,Tom

> def (first,second,third) = ['tome','jerry']
println "$first,$second,$third"
tome,jerry,null

在Groovy中，可以把一个映射或一个代码块转化为接口，因此可以快速实现带有多个方法的
接口

> button.addActionListener(
        {JOPtionPane.showMessageDialog(frame,"you clicked")} as java.awt.event.ActionListener

)

调用了addActionListener方法，同时为该方法提供了一个代码块，借助as操作符，相当
于实现了ActionListener接口。

就是它了！Groovy自会处理剩下的工作。它会拦截对接口中任何方法的调用（这个例子中就
是actionPerformed()），然后将调用路由到我们提供的代码块。要运行这段代码，还需要创建
窗体（Frame）及其组件

display = {positionLabel.setText("$it.x.$it.y")}
frame.addMouseListener(display as MouseListener)
frame.addMouseListener(display as java.awt.event.MouseMotionListener)
前面的例子中又出现了it变量。it表示方法的参数。如果正在实现的一个接口中的方法需要
多个参数，那么可以将其分别定义为独立的参数，也可以定义为一个数组类型的参数，具体情况
将在第4章讨论。

Groovy没有强制实现接口中的所有方法：可以只定义自己关心的，而不考虑其他方法。如果
剩下的方法从来不会被调用，那也就没必要去实现这些方法了。当在单元测试中通过实现接口来
模拟某些行为时，这项技术非常有用。


好了，这挺不错的，但是在大多数实际情况下，接口中的每个方法需要不同的实现。不用担
心，Groovy可以摆平。只需要创建一个映射，以每个方法的名字作为键，以方法对应的代码体作
为键值，同时使用简单的Groovy风格，用冒号（:）分隔方法名和代码块即可。此外，不必实现
所有方法，只需实现真正关心的那些即可。如果未予实现的方法从未被调用过，那么也就没有必
要浪费精力去实现这些伪存根。当然，如果没提供的方法被调用了，则会出现NullPointer
Exception。下面把这些内容放到一个例子里看看：

> handleFocus = [
        focusGained:{msgLable.setText("good to see you")},
        focusLost:{msgLable.setText("come back soon")}
]
button.addFcousListener(handleFocus as java.awt.event.FocusListener)

每当例子中的按钮获得焦点时，与focusGained键关联的第一个代码块就会被调用。当按钮
失去焦点时，与focusLost键关联的代码块则会被调用。在这种情况下，这里的键相当于
FocusListener接口中的方法。

如果知道所实现接口的名字，使用as操作符即可，但如果应用要求的行为是动态的，而且
只有在运行时才能知道接口的名字，又该如何呢？asType()方法可以帮忙。通过将欲实现接口
的Class元对象作为一个参数发送给asType()，可以把代码块或映射转化为接口。我们来看一
个例子。

> events = ['WindosListener','ComponentListener']
handler = { msgLabel.setText("$it")}
for(event in events){
    handlerImpl = handler.asType(Class.forName("java.awt.event.${event}"))
    frame."add${event}"(handlerImpl)
}

想实现的接口（也就是想处理的事件）在列表events中。该列表是动态的，假设它会在代
码执行期间通过输入来填充。事件公共的处理器位于变量handler指向的代码块中。我们对事件
进行循环，对于每个事件，都使用了asType()方法为该接口创建了一个实现。在代码块上调用
asType()方法，同时把使用forName()方法获得的、该接口的Class元对象传给它。一旦手头有
了监听器接口的实现，就可以通过调用相应的add方法（如addWindowListener()）来注册该实
现。调用add方法本身就是动态的。

去看Swing.groovy

Java要求if语句的条件部分必须是一个布尔表达式，比如前面例子中的if(obj != null)和
if(val > 0)。

> str = 'hello'
if(str) {print str}
hello

必须承认，前面关于true的说法并不完全正确。如果对象引用不为null，表达式的结果还
与对象的类型有关。例如，如果对象是一个集合（如java.util.ArrayList），那么Groovy会检
查该集合是否为空。因此，在这种情况下，只有当obj不为null，而且该集合至少包含一个元素
时，表达式if(obj)才会被计算为true；

表2-1 类型与布尔求值对它们的特殊处理
类 型 为真的条件
Boolean 值为true
Collection 集合不为空
Character 值不为0
CharSequence 长度大于0
Enumeration Has More Elements()为true
Iterator hasNext()为true
Number Double值不为0
Map 该映射不为空
Matcher 至少有一个匹配
Object[] 长度大于0
其他任何类型 引用不为null

> str = 'hello'
if(str) {print str}
hello> lst0 = null
println lst0 ? 'lst0 true' :'lst0 false'
lst0 false
> lst1 = [1,2,3]
println lst1 ? "lst1 true":"lst1 false"
lst1 true

Groovy支持操作符重载，可以巧妙地应用这一点来创建DSL（领域特定语言，参见第19章）。
Java是不支持操作符重载的，那Groovy又是如何做到的呢？其实很简单：每个操作符都会映射到
一个标准的方法①。在Java中，可以使用那些方法；而在Groovy中，既可以使用操作符，也可以
使用与之对应的方法。
 for(ch = 'a';ch < 'd';ch++){
    println ch
}
a
b
c

> for(ch in 'a'..'c'){
    println ch
}
a
b
c

String类重载了很多操作符，5.4节将予以介绍。类似地，为方便使用，集合类（如ArrayList
和Map）也重载了一些操作符。
要向集合中添加元素，可以使用<<操作符，该操作符会被转换为Groovy在Collection上添
加的leftShift()方法，如下所示：

> lst = ['hello']
lst << 'there'
println lst
[hello, there]


class ComplexNumber{
    def real,imaginary
    def plus(other){
        new ComplexNumber(real: real+other.real,imaginary:imaginary+other.imaginary)
    }
    String toString(){"$real ${imaginary > 0 ? '+' :''} ${imaginary}i"}
}
c1 = new ComplexNumber(real: 1,imaginary: 2)
c2 = new ComplexNumber(real: 3,imaginary: 4)
println c1 + c2

> int val = 4
println val.getClass().name
java.lang.Integer

在2.0版本之前，Groovy中所有基本类型都被看作对象。为了改进性能，也为了能在基本类
型的操作上使用更为直接的字节码，从2.0版本起，Groovy做了一些优化。基本类型只在必要时
才会被看作对象，比如，在其上调用了方法，或者将其传给了对象引用。否则，Groovy会在字节
码级别将其保留为基本类型。

> String[] greetings = ["hello","hi","howdy"]
for(greet in greetings){
    println greet

}
hello
hi
howdy


Groovy提供了对enum的支持，这是Java 5为解决枚举问题而引入的特性。它是类型安全的（比
如，我们可以区分得出用enum表示的衬衫尺寸和一周中的每一天），还具有可打印、可序列化等
特点。

> enum MethodLogies{
    Evo(5),
    XP(23),
    Scrum(30);

    final int daysInIteration
    MethodLogies(days){daysInIteration = days}
    def iterationDetails(){
        println "${this} recommands $daysInIteration days for iteration"
    }
}

for(methodlogy in MethodLogies.values()){
    methodlogy.iterationDetails()
}
Evo recommands 5 days for iteration
XP recommands 23 days for iteration
Scrum recommands 30 days for iteration

Groovy以两种方式支持Java 5的变长参数特性，除了支持使用省略符号标记形参，对于以数
组作为末尾形参的方法，也可以向其传递数目不等的参数。

> def receiveVarArgs(int a,int ...b){
    println "you passed $a and $b"
}
def receiveArray(int a,int[] b){
    println "you passed $a and $b"
}
receiveVarArgs(1,2,3,4,5)
receiveArray(1,2,3,4,5)
you passed 1 and [2, 3, 4, 5]
you passed 1 and [2, 3, 4, 5]

对于接受变长参数或者以数组作为末尾形参的方法，可以向其发送数组或离散的值，Groovy
知道该做什么。

在发送数组而非离散值时，请务必谨慎。Groovy会将包围在方括号中的值看作ArrayList的
一个实例， 而不是纯数组。所以如果简单地发送如[2, 3, 4, 5] 这样的值， 将出现
MethodMissingException。要发送数组，可以定义一个指向该数组的引用，或使用as操作符。

对于Java中与编译相关的注解，Groovy的处理方式有所不同。例如，groovyc会忽略@Override
。

首先，它实现了静态导入。我们可以像在Java中那样使用。
当然可以随意丢掉分号，它们在Groovy中是可选的。其次，在Groovy中可以为静态方法和类名定
义别名。要定义别名，需要在import语句中使用as操作符：
> import static Math.random as rand
import groovy.lang.ExpandoMetaClass as EMC
double value = rand()
def metaClass = new EMC(Integer)
assert metaClass.getClass().name == 'groovy.lang.ExpandoMetaClass'


在调用add()方法的过程中，Groovy更大程度上是将类型信息看作一个建议。当对集合进行
循环时，Groovy会尝试将其中的元素强制转换为int。如果无法转换，则会导致运行时错误。

Groovy在支持动态行为的同时支持泛型。前面的代码示例也说明了这两种概念有趣的相互作
用。对于Groovy的这种双重性，我们一开始可能会感到惊讶，但是当学到Groovy元编程（参见第
三部分）的好处时，你会看到其意义。


只有当派生类是真正可替换的，而且可代替基类使用时，继承才显示出其优势。从纯粹的代
码复用角度看，对于其他大部分用途，委托要优于继承。然而在Java中我们不太愿意使用委托，
因为会导致代码冗余，而且需要更多工作。Groovy使委托变得非常容易，所以我们可以做出正确
的设计选择。


@Canonical(excludes = "lastName,age")
class Person{

    String firstName
    String lastName
    int age

}
def sara = new Person(firstName: "Sars",lastName: "walker",age: 44)
println sara

Person(Sars)


class Worker{
    def work() {println 'get work done'}
    def analyze() {println 'analyze'}
    def writeReport() {println 'get report written'}
}

class Expert{
    def analyze() { println "export"}
}

class Manager{
    @Delegate Expert expert = new Expert()
    @Delegate Worker worker = new Worker()

}

def bernie = new Manager()
bernie.analyze()
bernie.work()
bernie.writeReport()

在编译时，Groovy会检查Manager类，如果该类中没有被委托类中的方法，就把这些方法从
被委托类中引入进来。因此，首先它会引入Expert类中的analyze()方法。而从Worker类中，
只会把work()和writeReport()方法因为进来。这时候，因为从Expert类带来的analyze()方
法已经出现在Manager类中，所以Worker类中的analyze()方法会被忽略。



不可变对象天生是线程安全的，将其字段标记为final是很好的实践选择。如果用
@Immutable注解标记一个类，Groovy会将其字段标记为final的，并且额外为我们创建一些便
捷方法，从而使得“做正确的事情”变得更容易了。

@Immutable
class CreditCard{
 String cardNumber
 int creditLimit

}

作为反馈，Groovy给我们提供了一个构造器，其参数以类中字段定义的顺序依次列出。在构
造时间过后，字段就无法修改了。此外，Groovy还添加了hashCode()、equals()和toString()
方法。运行所提供的构造器和toString()方法，看一下输出：
可以使用@Immutable注解轻松地创建轻量级的不可变值对象。在基于Actor模型的并发应用
中，线程安全是个大问题，而这些不可变值对象是作为消息传递的理想实例。


我们想把耗时对象的构建推迟到真正需要时。完全可以懒惰与高效并得，编写更少的代码，
同时又能获得惰性初始化的所有好处。
下面的例子将推迟创建Heavy实例，直到真正需要它时。既可以在声明的地方直接初始化实
例，也可以将创建逻辑包在一个闭包中。
class Heavy{
    def size = 10
    Heavy() {println "Creating heavy with $size"}
}

class AsNeeded{
    def value
    @Lazy Heavy heavy1 = new Heavy()
    @Lazy Heavy heavy2 = {new Heavy(size: value)}()
    AsNeeded() {println "Created AsNeeded"}
}

def asNeeded = new AsNeeded(value: 100)
println asNeeded.heavy1.size
println asNeeded.heavy1.size
println asNeeded.heavy2.size
Groovy不仅推迟了创建，还将字段标记为volatile，并确保创建期间是线程安全的。实例
会在第一次访问这些字段的时候被创建，在输出中可以看到：
另一个好处是，@Lazy注解提供了一种轻松实现线程安全的虚拟代理模式（virtual proxy
pattern） 的方式。

在Groovy中，经常会按照传统的Java语法，使用new来创建实例。然而，在创建DSL时，去
掉这个关键字，表达会更流畅。@Newify注解可以帮助创建类似Ruby的构造器，在这里，new是
该类的一个方法。该注解还可以用来创建类似Python的构造器（也类似Scala的applicator），这里
可以完全去掉new。要创建类似Python的构造器，必须向@Newify注解指明类型列表。只有将
auto=false这个值作为一个参数设置给@Newify，Groovy才会创建Ruby风格的构造器。

@Newify([Person,CreditCard])
def fluentCreate(){
    println Person.new(firstName:"John",lastName:"Doe",age:20)
    println Person(firstName:"John",lastName:"Doe",age:20)
    println CreditCard("1334",2000)

}

fluentCreate()


@Singleton(lazy = true)
class TheUnique{

    private TheUnique() {println 'Instance created'}

    def hello() {println 'hello'}
}

println "Accessing the unique"
TheUnique.instance.hello()
TheUnique.instance.hello()

警告 使用@Singleton注解，会使目标类的构造器成为私有的，这在我们意料之中，不过因为
Groovy实现并不区分公开还是私有，所以在Groovy内仍可使用new关键字来创建实例。但
是，必须谨慎恰当地使用这个类，并留心代码分析工具和集成开发环境给出的警告。

在Java中，==和equals()是一个混乱之源，而Groovy加剧了这种混乱。Groovy将==操作符
映射到了Java中的equals()方法。假如我们想比较引用是否相等（也就是原始的==的语义），该
怎么办呢？必须使用Groovy中的is()。

观察发现，Groovy的==映射到equals()，这个结论并不总是成立，当且仅当该类没有实现
Comparable接口时，才会这样映射。如果实现了Comparable接口，则==会被映射到该类的
compareTo()方法。

注意 在比较对象时，请首先问一下自己，要比较的是引用还是值。然后再问一下，是不是使
用了正确的操作符。
> str1 = 'hello'
str2 = str1
str3 = new String('hello')
str4 = 'Hello'
println "str1==str2: ${str1==str2}"
println "str1==str3:${str1==str3}"
println "str1==str4:${str1==str4}"
str1==str2: true
str1==str3:true
str1==str4:false


通过输出可以看到，在实现了Comparable接口的类上，==操作符选择了compareTo()，而
不是equals()。

Groovy的类型是可选的。然而，Groovy编译器groovyc大多数情况下不会执行完整的类型检
查，而是在遇到类型定义时执行强制类型转换。



def和in都是Groovy中的新关键字。def用于定义方法、属性和局部变量。in用于在for循环
中指定循环的区间，比如for(i in 1..10)。
将这些关键字用作变量名或方法名可能会带来问题，尤其是当把现有的Java代码当作Groovy
代码时。
定义名为it的变量也是不明智的。尽管Groovy不会抱怨什么，但是如果在闭包内使用了这样
的变量，它引用的是闭包的参数，而不是类中的一个字段——隐藏变量可无助于偿还技术债②。

Groovy的闭包是使用花括号（{...}）定义的，而定义匿名内部类也是使用花括号。如下面
例子所示，当构造器接收一个闭包作为参数时，就出现问题了：
这只是个小麻烦；与传递内联的闭包相比，传递引用给闭包噪音会小一些。
#闭包在我看来就是个代码块

> int[] arr = [1,2,3,43,5]
println arr
println "class is " + arr.getClass().name
[1, 2, 3, 43, 5]
class is [I

输出表明，所创建实例的类型为[I，JVM用它表示int[]。
> def arr = [1,2,3,43,5] as int[]
println arr
println "class is " + arr.getClass().name
[1, 2, 3, 43, 5]
class is [I



