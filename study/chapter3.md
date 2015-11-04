动态类型语言中的类型是在运行时推断的，方法及其实参也是在运行时检查的。通过这种能
力，可以在运行时向类中注入行为，从而使代码比严格的静态类型具有更好的可扩展性。

真正的多态并不关注类型——把一个消息发送给一个对象，在运行时，它自会确定所要使
用的相应实现。因此，动态类型语言可以实现比传统的静态类型语言更高程度的多态。

#动态类型不等于弱类型

有些语言，比如C++，会维护一个方法分派表，其中保存了多态方法的地址，参见Margaret A. Ellis和Bjarne Stroustrup
合著的The Annotated C++ Reference Manual [ES90]一书。

作为Java程序员，我们严重依赖接口。我们推崇“契约式设计”（Design By Contract），在这
种设计中，接口定义了交流的契约，类负责实现并遵守这些契约——参见Bertrand Meyer的Object-
Oriented Software Construction[Mey97]一书。
```groovy
def takeHelper(helper){
    helper.helpMoveThing()
}
```
takeHelp()接受一个helper，但是没有指定其类型，这样类型默认为Object。此外，这里
在它上面调用了helpMoveThings()方法。这就是能力式设计（Design By Capability）①。不同于
让helper遵守某些显式的接口，我们利用了对象的能力——依赖一个隐式的接口。这被称作鸭
子类型，它基于这一观点：“如果它走路像鸭子，叫起来也像鸭子，那它就是一只鸭子。”②
```groovy
class Man{
    void helpMoveThing(){
        println "Man's helping"
    }
}

takeHelper(new Man())   // Man's helping
```
我们不需要我们不需要扩展任何公共类，也不需要实现任何公共接口，不过借助动态特性 我们就可以实现

例如，在一个订单处理系统中，可以使用一个模拟（Mock）对象毫不
费力地替换掉一个信用卡处理对象，以便进行快速的自动化测试，而不必提前做出优雅的设计决
策。这也意味着，可以比较方便地加入一些设计反思，这为创建容易扩展的代码提供了更多的灵
活性和动力。

 在创建一个helper时，可能会敲错方法的名字。
 没有类型信息，怎么知道给方法发什么呢？
 如果把方法发给一个不能提供帮助的事物（一个不能搬动物体的对象），又会怎么样呢？

Groovy是动态类型的，同时也是可选类型的。这意味着，既可以将类型的转盘拨到一个极
端——不指定任何类型，让Groovy确定；也可以将它拨到另一个极端，精确地指定所要使用的变
量或引用的类型。
请记住，Groovy是一门运行在JVM上的语言。可选类型有助于集成Groovy代码和Java的库、
框架以及工具。有时候，Groovy的动态类型映射与当前使用的库、框架或工具并不匹配。这种
情况在Groovy中并不突出——开发者可以轻松地切换类型模式，指明类型信息。可选类型在其
他情况下也是有用的，比如有时候需要类型信息来生成数据库模式，或者创建GORM/Grails中
的验证器。

这里的“过早的优化”指的是提前确定调用的版本，在GiveRaiseJava.java这个例子中，raise(java.math.
BigDecimal amount) 方法在重载解析时就被排除掉了，而Groovy则直到最后调用时才根据目标对象和所提供的
参数确定实际要调用的方法版本。

如果一个类中有重载的方法，Groovy会聪明地选择正确的实现——不仅基于目标对象（调用
方法的对象），还基于所提供的参数。因为方法分派基于多个实体——目标加参数，所以这被称
作多分派或多方法（Multimethods）。

lst引用的是一个ArrayList<String>实例，而
Collection<String>类型的col引用的是同一实例。我们向lst中加入3个元素，然后移除1个。
移除操作去掉了列表中的第一个元素。现在我们想调用col.remove(0)来移除另一个元素。然而，
Collection接口的remove()方法想接收的是一个Object，所以Java把0装箱成一个Integer。因
为这个Integer实例不是列表中的元素，所以这个方法调用没有移除掉任何东西。
```groovy
ArrayList<String> lst = new ArrayList<>();
Collection<String> col = lst;
lst.add("one")
lst.add("two")
lst.add("three")
lst.remove(0)
col.remove(0)
//java
lst.size() => 2
col.size() => 2

//groovy
lst.size() => 1
col.size() => 1
```
groovy不会招惹装箱这种麻烦

其实，完全可以让Groovy自己识别正确的类型，并确保调用的方法和访问的属性在该类型上
是合法的。可以使用特殊的注解@TypeChecked，让Groovy去检查这些种错误，这个注解可以用
于类或单个方法上。如果用于一个类，则类型检查会在该类中所有的方法、闭包和内部类上执行。
如果用于一个方法，则类型检查仅在目标方法的成员上执行。
```groovy
 def shout(String str){
    println "Print in uppercase"
    println str.toUpperCase()
    println "again"
    println str.toUppercase()
}
try{
    shout("hello")
}catch(ex){
    println "failed..."
}
Print in uppercase
HELLO
again
failed...
类型检查
@groovy.transform.TypeChecked
def shout(String str){
    println "Print in uppercase"
    println str.toUpperCase()
    println "again"
    println str.toUppercase()
}
try{
    shout("hello")
}catch(ex){
    println "failed..."
}
org.codehaus.groovy.control.MultipleCompilationErrorsException: startup failed:
```

>动态向string的实例添加方法
```groovy
 def shoutString(String str){
    println str.shout()
}
str = 'hello'
str.metaClass.shout = {-> toUpperCase()}
shoutString(str)
HELLO

> def printInReverse(String str){
    println str.reverse()
}
printInReverse 'hello'
olleh
```
要利用静态类型检查，必须要指明方法和闭包的形参类型。能够在形参上调用的方法，被限
制为该类型在编译时已知支持的方法。Groovy会推断闭包的返回类型，并相应地执行类型检查，
所以不必担心此类细节。

与Java相比，Groovy的类型检查有一个优势。如果使用instanceOf检查类型，在使用该类
型特定的方法或属性时，并不需要执行强制转换，
```groovy
 def use(Object instance){
    if(instance instanceof String){
        println instance.length()
    }else{
        println instance
    }
}
use 'hello'
use 2
5
2
```
Groovy元编程和动态类型的优点显而易见，但是这些优点需要以性能为代价。性能的下降与
代码、所调用方法的个数等因素相关。当不需要元编程和动态能力时，与等价的Java代码相比，
性能损失可能高达10%。Java 7的InvokeDynamic特性就旨在缓解这种痛苦，但是对于使用老版本
Java的人而言，静态编译可能是个有用的特性。
我们可以关闭动态类型，阻止元编程，放弃多方法，并让Groovy生成性能足以与Java媲美的、
高效的字节码。
可以使用@CompileStatic注解让Groovy执行静态编译。这样为目标代码生成的字节码会和
javac生成的字节码很像



