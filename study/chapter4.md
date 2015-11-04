#闭包

闭包是从函数式编程的Lambda（λ）表达式派生而来的。根据Robert Sebesta的Concepts of
Programming Languages [Seb04]一书，“一个Lambda表达式指定了一个函数的参数与映射”。闭包
是Groovy最强大的特性之一，而且语法上非常优雅。或者如计算机科学家和函数式编程先驱Peter
J. Landin所言：“（闭包是）可以帮你消化λ演算的一点语法糖。”
闭包是轻量级的，短小、简洁，而且将会是我们在Groovy中使用最多的特性之一。过去传递
匿名类实例的地方，现在可以传递闭包。

pickEven()方法是一个高阶函数，即以函数为参数，或返回一个函数作为结果的函数①。该
方法对值进行迭代（和前面一样），但不同的是它将值发送给了一个代码块。在Groovy中，我们
称这种匿名代码块为闭包（Closure），

 def pickEven(n,block){
    for(int i = 2;i <= n;i++){
        block(i)
    }
}
pickEven(10,{print it + " "})
2 3 4 5 6 7 8 9 10

变量block保存了一个指向闭包的引用。可以像传递对象一样传递闭包。变量名没必要一定
命名为block，可以使用任何合法的变量名。当调用pickEven()方法时，现在可以像前面代码中
演示的那样向其发送代码块。代码块（{}内的代码）被传给形参block，就像把值10传给变量n。
在Groovy中，想传递多少闭包就可以传递多少。例如，方法调用的第一个、第三个和最后一个实
参都可以是闭包。如果闭包是最后一个实参，可以用下面这种优雅的语法：

pickEven(10) {println it}
2
3
4
5
6
7
8
9
10


代码块中的it是什么呢？如果只向代码块中传递一个参数，那么可以使用it这个特殊的变量
名来指代该参数。如果你喜欢，也可以像下面的例子这样，用其他名字代替it：
pickEven(10) {evenNumber -> println evenNumber}

total = 0
pickEven(10) {println total += it}
2
5
9
14
20
27
35
44
54

除了语法上的优雅，闭包还为函数将部分实现逻辑委托出去提供了一种简单、方便的方式。
前面示例中的代码块所做的事情，要比我们更早之前看到的代码块多。它将触角伸到了
pickEven()的调用者的作用域之内，使用了变量product。这是闭包的一个有趣特性。闭包是
一个函数，这里变量都绑定到了一个上下文或环境中，这个函数就在其中执行。

闭包能够扩充、优化或增强另一段代码。比如，可以将选择对象的操作通过一个谓词或条件
提炼出来，而闭包对于表达这样的谓词或条件可能很有用。另外，也可以通过闭包来使用协程
（Coroutine），实现诸如迭代器或循环中的控制流转移。
闭包有两个非常擅长的具体领域：一个是辅助资源清理，另一个是辅助创建内
部的领域特定语言。

普通函数在实现某个目标明确的任务时优于闭包。重构的过程是引入闭包的好时机。

闭包应该保持短小，有内聚性。闭包应该设计为附到方法调用上的小段代码，只有几行。当
编写使用闭包的方法时，最好不要滥用闭包的动态属性，比如在运行时确定参数的数目和类型。
在调用方法时实现的闭包一定要非常简单，而且做到显而易见。

```Java
def totalSelectValues(n,closure){
    total = 0
    for(i in 1..n){
        if(closure(i)){total += i}
    }
    total
}
def isOdd = {it % 2 != 0}
print "Total of even numbers from 1 to 10 is "
println totalSelectValues(10,isOdd)
Total of even numbers from 1 to 10 is 25


 def totalSelectValues(n,closure){
    total = 0
    for(i in 1..n){
        if(closure(i)){total += i}
    }
    total
}

print "Total of even numbers from 1 to 10 is "
println totalSelectValues(10) {it % 2 == 0}
Total of even numbers from 1 to 10 is 30
```

totalSelectValues()方法从1迭代到n，它会对每个值调用闭包，以确定该值是否要用于计
算中，它将选择过程委托给了该闭包。
即便在闭包中，return也是可选的。如果没有显式的return，最后一个表达式的值（可能
是null）会自动返回给调用者。
在第一次调用totalSelectValues()时，将闭包内联到了方法调用中，该闭包仅选择偶数。
另一方面，预先定义了要传给第二个调用的闭包。这个通过变量isOdd引用的闭包仅选择奇数。
与调用时直接创建的闭包不同，这种预先定义的闭包可以在多个调用中复用。顺便插一句，不费
吹灰之力，这个例子就实现了策略模式。


> class Equipment{
    def calculator
    Equipment(calc) {calculator = calc}
    def simulate(){
        println "Running simulation"
        calculator()
    }
}
def eq1 = new Equipment({println "calculator 1"})
def aCalculator = {println "calculator 2"}
def eq2 = new Equipment(aCalculator)
eq1.simulate()
eq2.simulate()
Running simulation
calculator 1
Running simulation
calculator 2

对于单参数的闭包，it是该参数的默认名称。只要知道只传入一个参数，就可以使用it。如
果传入的参数多于一个，就需要通过名字一一列出了


在调用闭包closure时，tellFortune()方法提供了两个参数：一个Date实例，一个表示运
势信息的String。该闭包分别用name和fortune引用它们。符号->将闭包的参数声明与闭包主
体分隔开来
> def tellFortune(closure){
    closure  new Date("09/22/1223"),"your day is filled with ceremony"
}
tellFortune(){date,fortune ->
    println "Fortune for ${date} is  '${fortune}'"
}
Fortune for Fri Sep 22 00:00:00 CST 1223 is  'your day is filled with ceremony'


如果为参数选择了表现力好的名字，通常可以避免定义类型。后面会看到，在元编程中，我
们可以使用闭包来覆盖或替代方法，而在那种情况下，类型信息对于确保实现的正确性非常重要。

> new FileWriter('out.txt').withWriter{writer ->
    writer.write('a')
}
使用Groovy添加的withWriter()方法重写这段代码。当从闭包返回时，withWriter()会自
动刷新（flush）并关闭这个流。

Execute Around Method模式
如果有一对必须连续执行的动作，比如打开和关闭，我们就可以使用Execute Around
Method模式，这是一个Smalltalk模式，Kent Beck的Smalltalk Best Practice Patterns [Bec96]一
书中曾经讨论过。编写一个Execute Around方法，它接收一个块作为参数。在这个方法中，
把对该块的调用夹到对那对方法的调用之间。即先调用第一个方法，然后调用该块，最后调
用第二个方法。方法的使用者不必担心这对动作，它们会自动被调用。我们甚至可以在
Execute Around方法内处理异常。
def static use(closure){
    def r = new Resource()
    try{
        r.open()
        closure(r)
    }finally{
        r.close()
      }
}

Resource.use{ res ->
    res.read()
    res.write()

}

调用一个函数或方法会在程序的执行序列中创建一个新的作用域。我们会在一个入口点（方
法最上面的语句）进入函数。在方法完成之后，回到调用者的作用域。
协程（Coroutine）则支持多个入口点，每个入口点都是上次挂起调用的位置。我们可以进入
一个函数，执行部分代码，挂起，再回到调用者的上下文或作用域内执行一些代码。之后我们可
以在挂起的地方恢复该函数的执行。正如Donald E. Knuth所言，“与主例程和子例程之间的不对
称关系不同，协程之间是完全对称的，可以互相调用。① ”
协程对于实现某些特殊的逻辑或算法非常方便，比如用在生产者－消费者问题中。生产者会
接收一些输入，对输入做一些初始处理，通知消费者拿走处理过的值做进一步计算，并输出或存
储结果。消费者处理它的那部分工作，完成之后通知生产者以获取更多输入。

> def iterate(n,closure){
    1.upto(n){
        println "In iterate with value ${it}"
        closure(it)
    }
}
println "calling iterate"
total = 0
iterate(4){
    total += it
    println "In closure total so far is ${total}"
}
println "Done"
calling iterate
In iterate with value 1
In closure total so far is 1
In iterate with value 2
In closure total so far is 3
In iterate with value 3
In closure total so far is 6
In iterate with value 4
In closure total so far is 10
Done

闭包可能不接受任何形参，也可能接受多个形参。每次调用一个闭包时，它会期望我们为其
每一个形参传入相应的实参。然而，如果在多次调用之间，有一个或多个实参是相同的，传参就
会变得枯燥乏味。预先绑定一些闭包形参可以缓解这种痛苦。
带有预绑定形参的闭包叫做科里化闭包（Curried Closure）。虽然英文单词curry有“咖喱”之
意，但科里化闭包与我最喜爱的印度菜并没有什么关系。（术语“科里化”源自对Lambda演算作
出重要贡献的著名数学家Haskell B. Curry的名字，Christopher Strachey、Moses Schönfinkel和
Friedrich Ludwig创造了这一术语。具体概念则是由Gottlob Frege发明的。）当对一个闭包调用
curry()时，就是要求预先绑定某些形参。在预先绑定了一个形参之后，调用闭包时就不必重复
为这个形参提供实参了。如图4-2所示，方法调用现在可以接受较少的参数。这有助于去掉方法
调用中的冗余或重复，


> def tellFortunes(closure){
    Date date = new Date("09/22/1223")
    postFortune = closure.curry(date)
    postFortune "your day is filled with ceremony"
    postFortune "They are features.not bugs"
}
tellFortunes(){date,fortune ->
    println "Fortune for ${date} is '${fortune}'"

}
Fortune for Fri Sep 22 00:00:00 CST 1223 is 'your day is filled with ceremony'
Fortune for Fri Sep 22 00:00:00 CST 1223 is 'They are features.not bugs'

tellFortunes()方法多次调用了一个闭包。该闭包接受两个形参。因此，每次调用
tellFortunes()时都要提供第一个参数date。作为一种选择，可以以date作为一个参数来调用
curry()方法，实现形参date的科里化。postFortune保存着科里化之后的闭包的引用，它已经
预先绑定了date的值。
现在可以调用科里化闭包了，只需要传入原来闭包的第二个形参（fortune）。科里化闭包
负责把fortune和预先绑定的形参date发送给原来的闭包：
可以使用curry()方法科里化任意多个形参，但这些形参必须是从前面开始的连续若干个。
也就是说，如果有n个形参，我们可以任意科里化前k个，其中0 <= k <= n。
如果想科里化后面的形参，可以使用rcurry()方法。如果想科里化位于形参列表中间的形
参，可以使用ncurry()①方法，传入要进行科里化的形参的位置，同时提供相应的值。
科里化是一种变换，将一个接受多个形参的函数变成了一个接受较少（通常是一个）形参的
函数。函数f(X,Y) -> Z上的科里化函数被定义为curry(f): X -> (Y -> Z)。科里化有助于
简化数学证明方法。就我们的目的而言，在Groovy中，科里化可以减少代码中的噪音。

可以确定一个闭包是否已经提供。如果尚未提供，比如说是一个算法，我们可以决定使用该
算法的一个默认实现来代替调用者未能提供的特殊实现

> def doSomething(closure){
    if(closure){
        closure()
    }else{
        println "Using default implemention"
    }
}
doSomething() {println "Use specialized implementation"}

doSomething()
Use specialized implementation
Using default implemention

this、owner和delegate是闭包的三个属性，用于确定由哪个对象处理该闭包内的方法调用。
一般而言，delegate会设置为owner，但是对其加以修改，可以挖掘出Groovy的一些非常好的元
编程能力。我们来观察一下闭包的这三个属性：

使用递归会遇到一些较为常见的问题，而借助Groovy中的闭包，我们可以在获得递归之优势
的同时避免这些问题。
递归可以通过子问题的解决方案来解决主干问题。递归解决方案的魅力在于非常简洁，而且
只需利用输入规模较小的相同问题的解决方案来组合出最终解决方案，这点很酷。尽管存在这些
优势，但是程序员往往对递归解决方案敬而远之。在输入规模较大的情况下，潜在的
StackOverflowError威胁，使得最优秀的程序员都有可能望而却步。

> def factorial(BigInteger number){
    if(number == 1) 1 else number * factorial(number - 1)
}
try{
    println "factorial of 4 is ${factorial(5)}"
    println "Number of bits in the result is ${factorial(4000).bitCount()}"
}catch(Throwable e){
    println "caught ${e.class.name}"
}
factorial of 4 is 120
caught java.lang.StackOverflowError

这里定义了一个名为factorial的变量，并将一个闭包赋给它。该闭包接受两个参数：一个
是number，要计算的就是它的阶乘；一个是theFactorial，它表示通过这个递归计算出的部分
结果。在闭包中，如果给定的number是1，就返回theFactorial的值作为结果。 否则，就通过
调用trampoline()方法递归地调用该闭包。将number - 1作为第一个参数传给该方法，以缩减
计算范围。第二个参数是到目前为止计算出的部分阶乘结果。
factorial变量本身被赋的就是在闭包上调用trampoline()方法的结果。
Groovy中的尾递归实现非常出彩，没有对语言本身做任何修改就实现了。当我们调用
trampoline()方法时，该闭包会直接返回一个特殊类TrampolineClosure的一个实例。当我们
向该实例传递参数时，比如像factorial(5, 1)中这样，其实是调用了该实例的call()方法。
该方法使用了一个简单的for循环来调用闭包上的call方法，直到不再产生TrampolineClosure
的实例。这种简单的技术在背后将递归调用转换成了一个简单的迭代。
这种递归之所以叫作尾递归，是因为方法中最后的表达式或者是结束递归，或者是调用自身。
相反，在直接递归计算阶乘时，最后的表达式调用的是*，即乘法操作符。
> def factorial
factorial = {int number,BigInteger theFactorial ->
    number == 1? theFactorial:
        factorial.trampoline(number-1,number*theFactorial)
}.trampoline()
println "factorial of 5 is ${factorial(5,1)}"
println "Number results is ${factorial(6000,1).bitCount()}"
factorial of 5 is 120

上一节介绍了一种可以使递归调用更高效地使用内存的技巧。递归本质上是一种使用子问题
的解决方案来解决问题本身的方式。这种技巧有一个变种（被奇怪地命名为动态规划）①，将问
题分解为可以多次重复解决的若干部分。在执行期间，我们将子问题的结果保存下来，当调用到
重复的计算时，只需要简单地使用保存下来的结果。这就避免了重复运行，因此极大地减少了计
算时间。记忆化可以将一些算法的计算时间复杂度从输入规模（n）的指数级（O(k^n)）降低到
只有线性级（O(n)）。


