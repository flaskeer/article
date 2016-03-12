java并发采用的是共享内存模型。java线程之间的通信总在隐式进行。整个通信过程对程序员完全透明。
java中，所有实例域，静态域和数组元素存储在堆内存中，堆内存在线程之间共享。局部变量，方法定义参数和异常处理器参数不会再线程之间共享，他们不会有内存可见性问题，也不受内存模型的影响。
JMM定义了线程和主内存之间的抽象关系，线程之间的共享变量存储在主内存中，每个线程都有一个私有的本地内存，本地内存存储了该线程以读/写共享变量的副本。本地内存是JMM的一个抽象概念。
线程A与线程B要通信的话：
1首先，线程A把本地内存A中更新过的共享变量刷新到主内存中去。
2然后，线程B到主内存中去读取线程A之前已更新过的共享变量。

JMM通过控制主内存与每个线程的本地内存之间的交互。来提供内存可见性保障。

为了提高性能：重排序

对于处理器重排序，JMM的处理器重排序规则会要求java编译器在生成指令序列时，插入特定类型的内存屏障（memory barriers，intel称之为memory fence）指令，通过内存屏障指令来禁止特定类型的处理器重排序（不是所有的处理器重排序都要禁止）。

JMM属于语言级的内存模型，它确保在不同的编译器和不同的处理器平台之上，通过禁止特定类型的编译器重排序和处理器重排序，为程序员提供一致的内存可见性保证。

一个操作执行的结果需要对另一个操作可见，那么这两个操作之间必须存在happens-before关系。这里提到的两个操作既可以是在一个线程之内，也可以是在不同线程之间。 与程序员密切相关的happens-before规则如下：

程序顺序规则：一个线程中的每个操作，happens- before 于该线程中的任意后续操作。
监视器锁规则：对一个监视器锁的解锁，happens- before 于随后对这个监视器锁的加锁。
volatile变量规则：对一个volatile域的写，happens- before 于任意后续对这个volatile域的读。
传递性：如果A happens- before B，且B happens- before C，那么A happens- before C。

如果两个操作访问同一个变量，且这两个操作中有一个为写操作，此时这两个操作之间就存在数据依赖性。

重排序会遵守数据依赖性。

as-if-serial单线程的程度的执行结果不能改变。


数据竞争：
在一个线程中写一个变量。
在另一个线程读同一个变量。
而且写和读没有通过同步排序

如果程序是正确同步的，程序的执行将具有顺序一致性。
同步是广义上的同步（包括lock，volatile和final）

顺序一致性特性：
一个线程中的所有操作必须按照程序的顺序来执行。
不管程序是否同步，所有线程都只能看到一个单一的操作执行，
在顺序一致性模型中，每个操作都必须原子执行且立刻对所有线程可见。

当我们声明共享变量为volatile后，对这个变量的读/写将会很特别。理解volatile的一个好方法：把对volatile变量的单个读/写，看成是使用同一个监视器锁对这些单个读/写操作做了同步。

```java
class VolatileFeaturesExample {
    volatile long vl = 0L;  //使用volatile声明64位的long型变量

    public void set(long l) {
        vl = l;   //单个volatile变量的写
    }

    public void getAndIncrement () {
        vl++;    //复合（多个）volatile变量的读/写
    }


    public long get() {
        return vl;   //单个volatile变量的读
    }
}

假设有多个线程分别调用上面程序的三个方法，这个程序在语意上和下面程序等价：

class VolatileFeaturesExample {
    long vl = 0L;               // 64位的long型普通变量

    public synchronized void set(long l) {     //对单个的普通 变量的写用同一个监视器同步
        vl = l;
    }

    public void getAndIncrement () { //普通方法调用
        long temp = get();           //调用已同步的读方法
        temp += 1L;                  //普通写操作
        set(temp);                   //调用已同步的写方法
    }
    public synchronized long get() { 
    //对单个的普通变量的读用同一个监视器同步
        return vl;
    }
}


```

对一个volatile变量的单个读/写操作，与对一个普通变量的读/写操作使用同一个监视器锁来同步，它们之间的执行效果相同。

对一个volatile变量的读，总是能看到任意线程对这个volatile变量最后的写入。

监视器锁的语义决定了临界区代码的执行具有原子性，

volatile的特性：
可见性：对一个vola变量的读，总是能看到任意线程对这个volatile变量最后的写入。
原子性：对任意单个volatile变量的读/写具有原子性，但类似于volatile++这种复合操作不具有原子性。

volatile变量的写-读可以实现线程之间的通信。

从内存语义的角度来说，volatile与监视器锁有相同的效果：volatile写和监视器的释放有相同的内存语义；volatile读与监视器的获取有相同的内存语义。

这里A线程写一个volatile变量后，B线程读同一个volatile变量。A线程在写volatile变量之前所有可见的共享变量，在B线程读同一个volatile变量后，将立即变得对B线程可见。

当写一个volatile变量时，JMM会把该线程对应的本地内存中的共享变量刷新到主内存。

当读取一个volatile变量时，JMM会把该线程对应的本地内存置为无效。线程接下来将从主内存中读取共享变量。

如果我们把volatile写和volatile读这两个步骤综合起来看的话，在读线程B读一个volatile变量后，写线程A在写这个volatile变量之前所有可见的共享变量的值都将立即变得对读线程B可见。

下面对volatile写和volatile读的内存语义做个总结：

线程A写一个volatile变量，实质上是线程A向接下来将要读这个volatile变量的某个线程发出了（其对共享变量所在修改的）消息。
线程B读一个volatile变量，实质上是线程B接收了之前某个线程发出的（在写这个volatile变量之前对共享变量所做修改的）消息。
线程A写一个volatile变量，随后线程B读这个volatile变量，这个过程实质上是线程A通过主内存向线程B发送消息。

当第二个操作是volatile写时，不管第一个操作是什么，都不能重排序。这个规则确保volatile写之前的操作不会被编译器重排序到volatile写之后。
当第一个操作是volatile读时，不管第二个操作是什么，都不能重排序。这个规则确保volatile读之后的操作不会被编译器重排序到volatile读之前。
当第一个操作是volatile写，第二个操作是volatile读时，不能重排序。

因此在旧的内存模型中 ，volatile的写-读没有监视器的释放-获所具有的内存语义。为了提供一种比监视器锁更轻量级的线程之间通信的机制，JSR-133专家组决定增强volatile的内存语义：严格限制编译器和处理器对volatile变量与普通变量的重排序，确保volatile的写-读和监视器的释放-获取一样，具有相同的内存语义。从编译器重排序规则和处理器内存屏障插入策略来看，只要volatile变量与普通变量之间的重排序可能会破坏volatile的内存语意，这种重排序就会被编译器重排序规则和处理器内存屏障插入策略禁止。

由于volatile仅仅保证对单个volatile变量的读/写具有原子性，而监视器锁的互斥执行的特性可以确保对整个临界区代码的执行具有原子性。在功能上，监视器锁比volatile更强大；在可伸缩性和执行性能上，volatile更有优势。如果读者想在程序中用volatile代替监视器锁，请一定谨慎。

锁时java并发编程中最重要的同步机制。锁除了让临界区互斥执行外，还可以让释放锁的线程向获取同一个锁的线程发送代码。

```java
class MonitorExample {
    int a = 0;

    public synchronized void writer() {  //1
        a++;                             //2
    }                                    //3

    public synchronized void reader() {  //4
        int i = a;                       //5
        ……
    }                                    //6
}

```

假设线程A执行writer()方法，随后线程B执行reader()方法。根据happens before规则，这个过程包含的happens before 关系可以分为两类：

根据程序次序规则，1 happens before 2, 2 happens before 3; 4 happens before 5, 5 happens before 6。
根据监视器锁规则，3 happens before 4。
根据happens before 的传递性，2 happens before 5。

线程A在释放锁之前所有可见的共享变量，在线程B获取同一个锁之后，将立刻变得对B线程可见。

当线程释放锁时，JMM会把该线程对应的本地内存中的共享变量刷新到主内存中。

当线程获取锁时，JMM会把该线程对应的本地内存置为无效。从而使得被监视器保护的临界区代码必须要从主内存中去读取共享变量。

锁释放与volatile写有着相同的内存语义，锁获取与volatile读有着相同的内存语义。

总结：
线程A释放一个锁，实质上是线程A向接下来要获取这个锁的某个线程发出了（线程A对共享变量所做修改的）消息。
线程B获取一个锁，实质上是线程B接收了之前某个线程发出的（在释放这个锁之前对共享变量所做修改的）消息。
线程A释放锁，随后线程B获取这个锁，这个过程实质上是线程A通过主内存向线程B发送消息。
```java
class ReentrantLockExample {
int a = 0;
ReentrantLock lock = new ReentrantLock();

public void writer() {
    lock.lock();         //获取锁
    try {
        a++;
    } finally {
        lock.unlock();  //释放锁
    }
}

public void reader () {
    lock.lock();        //获取锁
    try {
        int i = a;
        ……
    } finally {
        lock.unlock();  //释放锁
    }
}
}
```
ReentrantLock中，调用lock()方法获取锁，调用unlock()方法释放锁。

ReentrantLock的实现依赖于java同步器框架AbstractQueueSynchronizer(AQS),AQS使用一个整型的volatile变量（命名为state）来维护同步状态。这个volatile变量是ReentrantLock内存语义实现的关键。

ReentrantLock分为公平锁和非公平锁。
使用公平锁加锁的调用轨迹：
ReentrantLock:lock()
FairSync:lock()
AbstractQueuedSynchronizer:acquire(int arg)
ReentrantLock:tryAcquire(int acquires)
第4步真正开始加锁：源码：
```java
protected final boolean tryAcquire(int acquires) {
    final Thread current = Thread.currentThread();
    int c = getState();   //获取锁的开始，首先读volatile变量state
    if (c == 0) {
        if (isFirst(current) &&
            compareAndSetState(0, acquires)) {
            setExclusiveOwnerThread(current);
            return true;
        }
    }
    else if (current == getExclusiveOwnerThread()) {
        int nextc = c + acquires;
        if (nextc < 0)  
            throw new Error("Maximum lock count exceeded");
        setState(nextc);
        return true;
    }
    return false;
}
```

加锁首先读volatile变量state

解锁Unlock()调用轨迹：
ReentrantLock:unlock()
AbstractQueuedSynchronizer:release(int arg)
Sync:tryRelease(int release)

第3步开始释放锁：
```java
protected final boolean tryRelease(int releases) {
    int c = getState() - releases;
    if (Thread.currentThread() != getExclusiveOwnerThread())
        throw new IllegalMonitorStateException();
    boolean free = false;
    if (c == 0) {
        free = true;
        setExclusiveOwnerThread(null);
    }
    setState(c);           //释放锁的最后，写volatile变量state
    return free;
}
```
在释放锁的最后写volatile变量state

公平锁在释放锁的最后写volatile变量state，在获得锁时首先读取这个volatile变量，根据volatile的happens-before规则，释放锁的线程在写volatile变量之前可见的共享变量，在获取锁的线程读取同一个volatile变量后将立即变的对获取锁的线程可见。

非公平锁的释放和公平锁完全一样，所以这里仅仅分析非公平锁的获取。

使用非公平锁时，加锁方法lock()的方法调用轨迹如下：

ReentrantLock : lock()
NonfairSync : lock()
AbstractQueuedSynchronizer : compareAndSetState(int expect, int update)
在第3步真正开始加锁，下面是该方法的源代码：

protected final boolean compareAndSetState(int expect, int update) {
    return unsafe.compareAndSwapInt(this, stateOffset, expect, update);
}

该方法以原子操作的方式更新state变量，把java的compareAndSet简称为CAS。该方法说明如下：如果当前状态值等于预期值，则以原子方式将同步状态设置为给定的更新值，此操作具有volatile读和写的内存语义。

编译器不会对volatile读与volatile读后面的任意内存操作重排序，编译器不会对volatile写和volatile写前面的任意内存操作重排序。组合这两个条件，意味着为了同时实现volatile读和volatile写的内存语义，编译器不能对CAS和CAS前面和后面的任意内存操作重排序。
公平锁和非公平锁内存语义总结：
公平锁和非公平锁释放时，最后都要写一个volatile变量state。
公平锁获取时，首先会去读这个volatile变量。
非公平锁获取时，首先会用CAS更新这个volatile变量，这个操作同时具有volatile读和volatile写的内存语义。

锁释放--获取的内存语义的实现方式：
利用volatile变量的写--读所具有的内存语义。
利用CAS所附带的volatile读和volatile写的内存语义。

java线程之间通信的四种方式：
1.A线程写volatile变量。随后B读这个volatile变量。
2.A线程写volatile变量，随后B线程用CAS更新这个volatile变量。
3.A线程用CAS更新一个volatile变量。随后B线程用CAS更新这个volatile变量。
4.A线程用CAS更新一个volatile变量，随后B线程读这个volatile变量。

Java的CAS会使用现代处理器上提供的高效机器级别原子指令，这些原子指令以原子方式对内存执行读-改-写操作，这是在多处理器中实现同步的关键（从本质上来说，能够支持原子性读-改-写指令的计算机器，是顺序计算图灵机的异步等价机器，因此任何现代的多处理器都会去支持某种能对内存执行原子性读-改-写操作的原子指令）。同时，volatile变量的读/写和CAS可以实现线程之间的通信。把这些特性整合在一起，就形成了整个concurrent包得以实现的基石。如果我们仔细分析concurrent包的源代码实现，会发现一个通用化的实现模式：

1.首先声明共享变量为volatile
2.使用CAS的原子条件更新来实现线程之间的同步
3.同时，配合以volatile的读/写和CAS所具有的volatile读和写的内存语义来实现线程之间的通信。
AQS，非阻塞数据结构和原子变量类（java.util.concurrent.atomic包中的类），这些concurrent包中的基础类都是使用这种模式来实现的，而concurrent包中的高层类又是依赖于这些基础类来实现的。

对final域的读和写更像是普通的变量访问。对于final域，编译器和处理器遵守两个重排序规则：
在构造函数内对一个final域写入，与随后把这个被构造对象的引用赋值给一个引用变量，这两个操作之间不能重排序。
初次读一个包含final域的对象的引用，与随后初次读这个final域，这两个操作之间不能重排序。

```java
public class FinalExample {
    int i;                            //普通变量
    final int j;                      //final变量
    static FinalExample obj;

    public void FinalExample () {     //构造函数
        i = 1;                        //写普通域
        j = 2;                        //写final域
    }

    public static void writer () {    //写线程A执行
        obj = new FinalExample ();
    }

    public static void reader () {       //读线程B执行
        FinalExample object = obj;       //读对象引用
        int a = object.i;                //读普通域
        int b = object.j;                //读final域
    }
}
```

写final域的重排序规则禁止把final域的写重排序到构造函数之外，包含下面两个方面：
JMM禁止编译器把final域的写重排序到构造函数之外。
编译器会在final域的写之后，构造函数return之前，插入一个StoreStore屏障，这个屏障禁止处理器把final域的写重排序到构造函数之外。

现在让我们分析writer ()方法。writer ()方法只包含一行代码：finalExample = new FinalExample ()。这行代码包含两个步骤：

构造一个FinalExample类型的对象；
把这个对象的引用赋值给引用变量obj。

写final域的重排序规则可以确保，在对象引用为任意线程可见之前，对象的final域已经被正确初始化过了，而普通域不具有这个保障。

读final域的重排序规则如下：
在一个线程中，初次读对象引用与初次读该对象包含的final域，JMM禁止处理器重排序这个两个操作，（这个规则仅仅针对处理器）。编译器会在读final域操作的前面插入一个LoadLoad屏障。
初七读对象引用与初次读该对象包含的final域，这两个操作之间存在间接依赖关系。由于编译器遵守间接依赖关系，因此编译器不会重排序这两个操作。
大多数处理器也会遵守间接依赖，大多数处理器也不会重排序这两个操作。但有少数处理器允许对存在间接依赖关系的操作做重排序（比如alpha处理器），这个规则就是专门用来针对这种处理器。

reader()方法包含三个操作：

初次读引用变量obj;
初次读引用变量obj指向对象的普通域j。
初次读引用变量obj指向对象的final域i。

读final域的重排序规则可以确保，在读一个对象的final域之前，一定会先读包含这个final域的对象的引用，在这个示例程序中，如果该引用不为null，那么引用对象的final域一定已经被A线程初始化过了。

```java
public class FinalReferenceExample {
final int[] intArray;                     //final是引用类型
static FinalReferenceExample obj;

public FinalReferenceExample () {        //构造函数
    intArray = new int[1];              //1
    intArray[0] = 1;                   //2
}

public static void writerOne () {          //写线程A执行
    obj = new FinalReferenceExample ();  //3
}

public static void writerTwo () {          //写线程B执行
    obj.intArray[0] = 2;                 //4
}

public static void reader () {              //读线程C执行
    if (obj != null) {                    //5
        int temp1 = obj.intArray[0];       //6
    }
}
}
```
这里final域为一个引用类型，它引用一个int型的数组对象。对于引用类型，写final域的重排序规则对编译器和处理器增加了如下约束：

在构造函数内对一个final引用的对象的成员域的写入，与随后在构造函数外把这个被构造对象的引用赋值给一个引用变量，这两个操作之间不能重排序。

JMM可以确保读线程C至少能看到写线程A在构造函数中对final引用对象的成员域的写入。即C至少能看到数组下标0的值为1。而写线程B对数组元素的写入，读线程C可能看的到，也可能看不到。JMM不保证线程B的写入对读线程C可见，因为写线程B和读线程C之间存在数据竞争，此时的执行结果不可预知。

如果想要确保读线程C看到写线程B对数组元素的写入，写线程B和读线程C之间需要使用同步原语（lock或volatile）来确保内存可见性。

为什么final引用不能从构造函数内“逸出”

前面我们提到过，写final域的重排序规则可以确保：在引用变量为任意线程可见之前，该引用变量指向的对象的final域已经在构造函数中被正确初始化过了。其实要得到这个效果，还需要一个保证：在构造函数内部，不能让这个被构造对象的引用为其他线程可见，也就是对象引用不能在构造函数中“逸出”。为了说明问题，让我们来看下面示例代码：
```java
public class FinalReferenceEscapeExample {
final int i;
static FinalReferenceEscapeExample obj;

public FinalReferenceEscapeExample () {
    i = 1;                              //1写final域
    obj = this;                          //2 this引用在此“逸出”
}

public static void writer() {
    new FinalReferenceEscapeExample ();
}

public static void reader {
    if (obj != null) {                     //3
        int temp = obj.i;                 //4
    }
}
}
```
假设一个线程A执行writer()方法，另一个线程B执行reader()方法。这里的操作2使得对象还未完成构造前就为线程B可见。即使这里的操作2是构造函数的最后一步，且即使在程序中操作2排在操作1后面，执行read()方法的线程仍然可能无法看到final域被初始化后的值，因为这里的操作1和操作2之间可能被重排序。


JSR-133专家组增强了final的语义。通过为final域增加写和读重排序规则，可以为java程序员提供初始化安全保证：只要对象是正确构造的（被构造对象的引用在构造函数中没有“逸出”），那么不需要使用同步（指lock和volatile的使用），就可以保证任意线程都能看到这个final域在构造函数中被初始化之后的值。

顺序一致性内存模型是一个理论参考模型，JMM和处理器内存模型在设计时通常会把顺序一致性内存模型作为参照。JMM和处理器内存模型在设计时会对顺序一致性模型做一些放松，因为如果完全按照顺序一致性模型来实现处理器和JMM，那么很多的处理器和编译器优化都要被禁止，这对执行性能将会有很大的影响。

根据对不同类型读/写操作组合的执行顺序的放松，可以把常见处理器的内存模型划分为下面几种类型：

放松程序中写-读操作的顺序，由此产生了total store ordering内存模型（简称为TSO）。
在前面1的基础上，继续放松程序中写-写操作的顺序，由此产生了partial store order 内存模型（简称为PSO）。
在前面1和2的基础上，继续放松程序中读-写和读-读操作的顺序，由此产生了relaxed memory order内存模型（简称为RMO）和PowerPC内存模型。
注意，这里处理器对读/写操作的放松，是以两个操作之间不存在数据依赖性为前提的（因为处理器要遵守as-if-serial语义，处理器不会对存在数据依赖性的两个内存操作做重排序）


JMM把happens- before要求禁止的重排序分为了下面两类：

会改变程序执行结果的重排序。
不会改变程序执行结果的重排序。
JMM对这两种不同性质的重排序，采取了不同的策略：

对于会改变程序执行结果的重排序，JMM要求编译器和处理器必须禁止这种重排序。
对于不会改变程序执行结果的重排序，JMM对编译器和处理器不作要求（JMM允许这种重排序）。

java程序的内存可见性保证按程序类型分为三类：
1.单线程程序。单线程程序不会出现内存可见性问题。编译器，runtime和处理器会共同确保单线程程序的执行结果与该程序在顺序一致性模型中的执行结果相同。
2.正确同步的多线程程序。正确同步的多线程程序的执行将具有顺序一致性（程序的执行结果与该程序在顺序一致性内存模型中的执行结果相同），这是JMM关注的重点，JMM通过限制编译器和处理器的重排序来为程序员提供内存可见性保证。
3.未同步/未正确同步的多线程程序。JMM为他们提供了最小安全性保障。线程执行时读取到的值，要么是之前某个线程写入的值，要么是默认值(0,null,false).

JSR-133对JDK5之前的旧内存模型的修补主要有两个：

增强volatile的内存语义。旧内存模型允许volatile变量与普通变量重排序。JSR-133严格限制volatile变量与普通变量的重排序，使volatile的写-读和锁的释放-获取具有相同的内存语义。
增强final的内存语义。在旧内存模型中，多次读取同一个final变量的值可能会不相同。为此，JSR-133为final增加了两个重排序规则。现在，final具有了初始化安全性。

每次仅有一个线程可以持有管程上的锁，其他视图试图锁定该管程的线程会一直阻塞，直到能从该管程上获得锁为止。

synchronized语句需要一个对象的引用，随后会尝试在该对象的管程上执行lock动作，如果lock动作未能成功完成，将会一直等待。当lock动作执行成功，就会运行synchronized语句块中的代码，一旦语句块中的代码执行结束，不管是正常还是异常结束，都会在之前执行lock动作的那个管程上自动执行一个unlock动作。

synchronized方法在调用时会自动执行一个lock动作。在lock动作成功完成之前，都不会执行方法体。如果是实例方法，锁的是调用该方法的实例（即，方法体执行期间的this）相关联的管程。如果是静态方法，锁的是定义该方法的类所对应的Class对象。一旦方法体执行结束，不管是正常还是异常结束，都会在之前执行lock动作的那个管程上自动执行一个unlock动作。

冲突访问（Conflicting Accesses） 对同一个共享字段或数组元素存在两个访问（读或写），且至少有一个访问是写操作，就称作有冲突。
Happens-Before关系 两个动作（action）可以被happens-before关系排序。如果一个动作happens-before另一个动作，则第一个对第二个可见，且第一个排在第二个之前。必须强调的是，两个动作之间存在happens-before关系并不意味着这些动作在Java中必须以这种顺序发生。happens-before关系主要用于强调两个有冲突的动作之间的顺序，以及定义数据争用的发生时机。可以通过多种方式包含一个happens-before顺序，包括：
 某个线程中的每个动作都happens-before该线程中该动作后面的动作。
 某个管程上的unlock动作happens-before同一个管程上后续的lock动作。
 对某个volatile字段的写操作happens-before每个后续对该volatile字段的读操作。
 在某个线程对象上调用start()方法happens-before该启动了的线程中的任意动作。

某个线程中的所有动作happens-before任意其它线程成功从该线程对象上的join()中返回。
 如果某个动作a happens-before动作b，且b happens-before动作c，则有a happens-before c.

final字段也允许编程人员在不需要同步的情况下实现线程安全的不可变对象。一个线程安全的不可变对象被所有线程都视为不可变的

共享变量/堆内存（Shared variables/Heap memory） 能够在线程间共享的内存称作共享内存或堆内存。所有的实例字段，静态字段以及数组元素都存储在堆内存中。我们使用变量这个词来表示字段和数组元素。方法中的局部变量永远不会在线程间共享且不会被内存模型影响。

同步动作（Synchronization Actions） 同步动作包括锁、解锁、读写volatile变量，用于启动线程的动作以及用于探测线程是否结束的动作。任何动作，只要是synchronizes-with边缘（edge）的起始或结束点，都是同步动作。这些动作会在后面讲到happens-before边缘的地方详讲。
同步顺序（Synchronization Order） 每个执行过程都有一个同步顺序。同步顺序是一次执行过程中的所有同步动作上的全序关系。


