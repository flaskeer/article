多个线程访问一个类时，如果不用考虑这些线程在运行时环境下的调度和交替执行，并且不需要额外的同步及在调用方代码不必做其他协调，这个类的行为仍然是正确的，那么这个类是线程安全的。

无状态对象永远是线程安全的。

为了保护状态的一致性，要在单一的原子操作中更新相互关联的状态变量。

synchronized：每个java对象都可以隐式的扮演一个用于同步的锁的角色，这些内置的锁被称为内部锁或监视器锁，获得内部锁的唯一途径是进入这个内部锁保护的同步块或方法。

内部锁在java中扮演了互斥锁的角色，意味着至多只有一个线程可以拥有锁。
重进入的实现是通过为每个锁关联一个请求计数和一个占有它的线程。当计数为0时，认为锁时未被占有的。线程请求一个未被占有的锁时，JVM将记录锁的占有者，并且将请求计数置为1，同一线程再次请求这个锁，计数将递增。每次占用线程退出同步块，计数器值将递减，直到计数器到0，锁被释放。

锁规则：在对象内部封装所有的可变状态，通过对象的内部锁来同步任何访问可变状态的代码路径，保护它在并发访问的安全。

有些耗时的计算或操作，比如网络或控制台IO，难以快速完成，执行这些操作期间不要占有锁。

单个线程中，只要重排序不会对结果产生影响，那么就不能保证其中的操作一定按照程序写定的顺序执行--即使重排序对于其他线程来说会产生明显的影响。

内置锁可以用来确保一个线程以某种可预见的方式看到另一个线程的影响，当A执行一个同步块时，B也随后进入被同一个锁监视的同步块，这可以保证在锁释放之前对A可见的变量的值，B获得锁之后同样可见。
如果没有同步就没有这样的保证。

访问共享变量需要锁同步--- 为了保证一个线程对数值进行的写入，其他线程也都可见。
如果一个线程在没有恰当的使用锁的情况下读取了变量，那么这个变量可能是一个过期数据。

当一个域声明为volatile类型后，编译器与运行时会监视这个变量，它是共享的，而且对它的操作不会与其他内存操作一起被重排序。读取一个volatile变量时，总会返回由某一个线程所写入的最新值。

访问volatile变量的操作不会加锁。也就不会引起执行线程的阻塞。
是轻量级的同步机制。
从内存可见性而言：写入volatile变量就像退出同步块，读取volatile变量就像进入同步块。


加锁可以保证可见性和原子性，volatile只能保证可见性。

ThreadLocal允许你将每个线程与持有数值的对象关联在一起。ThreadLocal提供了get和set访问器。为每个使用它的线程维护一份单独的拷贝。所以get总是返回由当前执行线程通过set设置的最新值。

所有域都是final类型的对象仍然是可变的，因为final域可以获得一个到可变对象的引用。
不可变对象：
它的状态不能再创建后再被修改
所有域都是final类型
它被正确创建。

final域使得确保初始化安全性称为可能。

安全发布对象：
通过静态初始化器初始化对象的引用
将它的引用存储到volatile域或AtomicReference
将它的引用存储到正确创建的对象的final域中
或者将它的引用存储到由锁正确保护的域中


为了让方法正确工作，我们必须保证方法所使用的锁，与List用于客户端加锁与外部加锁时所用的锁是同一个锁。客户端加锁必须保证客户端代码与对象X保护自己状态时使用的是相同的锁。为了正确的执行客户端加锁，必须知道X使用了哪个锁。


及时失败(fail-fast)--当它们察觉容器在迭代开始后被修改，会抛出一个未检查的concurrentModificationException

同步容器类在每个操作的执行期间都持有一个锁。
ConcurrentHashMap使用了分离锁，允许更深层次的共享访问。任意数量的读线程可以并发访问Map,读者和写者可以并发访问Map，并且有限数量的写线程还可以并发修改Map.结果是为并发访问带来更高的吞吐量，同时没有损失单个线程访问的性能。

提供了不会抛出ConcurrentModificationException的迭代器。返回的迭代器具有若一致性，而非及时失败的。、可以允许并发修改，当迭代器被创建时，会遍历已有的元素。并且可以保证感应到在迭代器被创建后对容器的修改。

CopyOnWriteArrayList是同步List的并发替代器，避免了在迭代期间对容器加锁和复制。每次修改时，他们会创建并重新发布一个新的容器拷贝。写入时复制(copy-on-write)容器的迭代器保留了一个底层基础数组的引用。这个数组作为迭代器的起点永远不会被修改。对它的同步不过是为了确保数组内容的可见性。

阻塞队列(blocking queue)提供了可阻塞的put和take方法，它们与可定时的offer和poll是等价的。如果Queue已经满了，put会被阻塞直到有空间可用。如果queue是空的，那么take方法会被阻塞，直到有元素可用。queue长度可以有限，也可以无限，无限的queue永远不会充满，所以put方法永远会不会阻塞。
阻塞队列支持生产者-消费者模式。
生产者把数据放入队列，并使数据可用，当消费者为适当的行为做好准备时会从队列中获取数据。生产者不需要知道消费者的身份或者数量。甚至根本没有消费者--他们只负责把数据放入队列。消费者也不需要知道 生产者是谁，以及是谁给他们安排的工作，BlockingQueue可以使用任意数量的生产者和消费者。将线程池与工作队列想结合。

take会保持阻塞直到可用数据出现，如果生产者不能足够快的产生工作，让消费者忙碌起来那么消费者只能一直等待。

如果我们使用有界队列，那么当队列充满的时候，生产者就会阻塞，咱不能生成更多的工作，从而给消费者时间来追赶进度。

阻塞队列提供了offer方法，如果条目不能被加入到队列里。它会返回一个失败状态。这使得你能够创建更多灵活的策略来处理超负荷工作。比如减轻负载，序列化剩余工作条目并写入硬盘，减少生产者线程。或者用其他办法遏制生产者线程。

LinkedBlockingQueue和ArrayBlockingQueue是FIFO队列，PriorityBlockingQueue是一个按优先级顺序排序的队列。可以比较元素本身的自然顺序（实现comparable），

生产者和消费者可以并发的执行，如果一个受限于IO，另一个受限于CPU，那么并发执行的全部产出护高于顺序执行的产出，如果生产者和消费者在不同层面并行执行，那么紧密的耦合会减弱并行性，减少并行化的活动。
```java
public class FileCrawler implements Runnable{
	private final BlockingQueue<File> fileQueue;
	private final FileFilter fileFilter;
	private final File root;
	public void run(){
	  try{
	  crawl(root);
	}catch(){}
  }
  private void crawl(File root){
    File[] entries = root.listFile(fileFilter);
    if(entries != null){
	  for(File entry:entries){
		if(entry.isDirectory()){
		   crawl(entry);
		}else{
             fileQueue.put(entry);
		}
	  }
	}
  }
}


public class Indexer implements Runnable{
	private final BlockingQueue<File> queue;
	private Indexer(BlockingQueue<File> queue){
		this.queue = queue;
	}
	public void run(){
	  try{
		while(true){
			indexFile(queue.take());
		}catch(){}
	  }	 
	}

	//开始搜索
	public static void startIndexing(File[] roots){
		BlockingQueue<File> queue = new LinkedBlockingQueue<BOUND>);
		FileFilter filter = new FileFilter(){
			public boolean accept(File file)  {return true;}
		};
		for(File root:roots){
			new Thread(new FileCrawler(queue,filter,root)).start();
		}
		for(int i = 0;i < N_CONSYMERS;i++)
			new Thread(new Indexer(queue)).start();
	}
}

```

Deque是一个双端队列，允许高效的在头和尾分别进行插入和移除。
时他们与自身与一种窃取工作(work stealing)的模式相关联。
C-P设计中，所有的消费者只共享一个工作队列，在窃取工作的设计中，每一个消费者都有一个自己的双端队列，如果一个消费者完成自己双端队列中的全部工作，可以偷取其他消费者的双端队列中的末尾人物。
因为工作者线程并不会竞争一个共享的任务队列。所以伸缩性更大，他会从尾部截取，而不是从头部。从而降低对双端队列的争夺。

阻塞操作与普通操作差别：被阻塞的线程必须等待一个事件的发生才能继续运行，并且这个事件是超越它自己控制的，因此需要花费更长的时间--等待IO操作完成，锁可用。或者是外部计算结束。外部事件发生后，线程被置回RUNNABLE状态，重新获得调度的机会。

interrupt方法，用来中断一个线程，或者查询某线程是否被中断。每一个线程都有一个布尔类型的属性。这个属性代表了线程的中断状态。

//恢复中断状态
```
Thread.currentThread.intterupt();
```
Synchronizer都享有类似的结构特性，它们封装状态，而这些状态决定着线程执行到在某一点时是通过还是被迫等待，它们还提供了操控状态的方法，以及高效的等待着Synchronizer进入到期望状态的方法。

latch是一种synchronizer，它可以延迟线程的进度直到线程到达终止状态。一个闭锁工作起来就像一道大门，直到闭锁达到终点之前，门一直是关闭的，没有线程能够通过，在终点状态到来的时候，门开了，允许所有的线程都通过，一旦闭锁到达了终点状态，它就不能够再改变状态了。闭锁可以用来确保特定活动直到其他活动完成时才发生。

确保一个计算不会执行，直到他需要的资源被初始化。一个二元闭锁(两个状态)可以用来表达资源R已经被初始化。，并且所有需要用到R的活动受限都要在闭锁中等待。

CountDownLatch是一个灵活的闭锁实现，允许一个或多个线程等待一个事件集的发生，闭锁的状态包括一个计数器，初始化为一个正数，用来表现需要等待的事件数。countDown方法对计数器做减操作，用来表示一个事件已经发生了，而await方法等待计数器达到零，此时所有需要等待的事件都已经发生，如果计数器入口时值为非零，await会一直阻塞直到计数器为零。或者等待线程中断以及超时。

TestHarness中创建两个线程，并发的执行给定的任务。使用两个闭锁，一个开始阀门，一个结束阀门。开始阀门将计数器初始化为1，结束阀门将计数器初始化为工作线程的数量，每一个工作线程要做的第一件事情等待开始阀门打开，这样就能确保直到所有的线程都做好准备才开始工作，每个线程的最后一个工作是为结束阀门减1，这样做使控制线程有效的等待，直到最后一个工作线程完成任务，这样就能计算整个用时了。

FutureTask实现描述了一个抽象的可以携带结果的计算。是通过Callable实现的。有3个状态。等待，运行，完成。完成包括所有计算以及以任意的方式结束。包括正确结束，取消和异常。一旦futuretask进入完成状态，会永远停止在这个状态。

future.get依赖于任务状态。如果已经完成，get可以立刻得到返回结果，否则会被阻塞直到任务转入完成状态。然后返回结果或者抛出异常。
```
public class Preloader{
	private final FutureTask<ProductInfo> future = 
	   new FutureTask<ProductInfo>(new Callable<ProductInfo>{
			public ProductInfo call(){
			  return loadProductInfo();
		}
	   });

	   private final Thread thread = new Thread(future);
	   public void start(){ thread.start();}
	   ...
	   return future.get();
}
```

计数信号量(semaphore)用来控制能够同时访问某特定资源的活动的数量，或者同时执行某一给定操作的数量，技术信号量可以用来实现资源池或者给一个容器限定边界。
一个semaphore管理一个有效的许可集，许可的初始量通过构造函数传递给semaphore。活动就能获得许可，并在使用之后释放许可，如果已经没有可用的许可，那么acquire就会阻塞，知道有可用的为止。release方法向信号量返回一个许可，计算信号量的一种退化形式是二元信号量，一个计数初始值为1的semaphore，二元信号量可用作互斥锁，有不可重入锁的意思。
可以用来实现资源池。可以使用它把任何容器转化为有界的阻塞容器。信号量被初始化为容器所期望容量的最大值。


闭锁是一次性使用的对象，一旦进入到最终状态，就不能被重置了。

关卡(barrier)类似于闭锁，他们都能够则色一组线程，直到某些事件发生，其中关卡与闭锁关键的不同在于所有线程必须同时到达关卡点，才能继续处理。闭锁等待的是事件，关卡等待的是其他线程。关卡实现的协议，就像是我们6点在KFC见面，不见不散，然后决定接下来做什么。

CyclicBarrier允许一个给定数量的成员多次集中在一个关卡点，这在并行迭代算法中非常有用，这个算法把一个问题拆分成一系列相互独立的子问题，当线层到达关卡点时，调用await会阻塞。直到所有线程都到达关卡点，如果所有线程都到达关卡点，关卡就被成功的突破，这样所有的线程都会被释放，关卡会重置以备下一次使用。如果对await调用超时，或者阻塞中的线程被中断，那么关卡就被认为是失败的。所有对await未完成的调用都通过brokenBarrierException终止。如果成功通过关卡，await为每一个线程返回一个唯一的到达索引号，可以用他选举产生一个领导，在下一次迭代中承担一些特殊工作。
当成功通过关卡的时候，会在一个子任务线程中执行，但是在阻塞线程被释放之前是不能执行的。

Exchanger是关卡的另一种形式，它是一种两步关卡，在关卡点会交换数据，当两方进行的活动不对称时，Exchanger是非常有用的，比如当一个线程向缓冲写入一个数据，这时另一个线程充当消费者使用这个数据，这些线程可以使用Exchanger进行会面，并用完整的缓冲与空缓冲进行交换，当两个线程通过Exchanger交换对象时，交换为双方的对象建立了一个安全的发布。
交换的时机取决于应用程序的响应需求，最简单的方案是当写入任务的缓冲写满就发生交换，并且当清除任务的缓冲清空后也发生交换，这样做使交换的次数最少，但是如果新数据的到达率不可预测，处理一些数据会发生延迟。另一个方案是缓冲满了就发生交换，但是当缓冲部分充满却已经存在了特定长时间时，也会发生交换。

NCPU/CPUT+1个线程会产生最优吞吐量。
```
public class Memorizer<A,V> implements Commputable<A,V>{
	private final Map<A,Future<V>> cache = new 
	    ConcurrentHashMap<>();
	prvate final Commputable<A,V> c;
	public Memorizer(Commputable<A,V> c) {this.c = c;}
	public V  compute(Final A arg){
		Future<V> f = cache.get(arg);
		if(f == null){
			Callable<V> eval = new Callable<V>(){
            	public V call(){
            	return c.compute(arg);
            }
		};
		FutureTask<V> ft = new FutureTask<V>(eval);
		if(f == null){
			cache.putIfAbsent(arg,f);
			f.run();
	     }
		
		}
		return f.get();
	}
}
```

32位机器上，主要限制因素是线程栈的地址空间。每个线程都维护着两个执行栈，一个用于java代码，另一个用于原生代码，典型的JVM默认会产生一个组合的栈，大概半兆字节左右（-Xss JVM参数），如果为每个线程分配了大小是232字节的栈，那么你的线程数量将被限制在几千到几万间不等。


任务是逻辑上的工作单元，线程是使任务异步执行的机制。
---Question:
任务在什么（what）线程中执行
任务以什么（what）顺序执行（FIFO,LIFO,优先级）
可以由多少个(how many)任务并发执行
可以由多少个(how many)任务进入等待执行队列
如果系统过载，需要放弃一个任务，应该挑选哪一个(which)任务，另外。如何(how)通知应用程序知道这一切呢？
在一个任务的执行前与执行后，应该做什么(what)处理？

newFixedThreadPool创建一个定长的线程池，每当提交一个任务就创建一个线程直到到达池的最大长度，这时线程池会保持长度不在变化。（如果一个线程由于非预期的exception结束，线程池会补充一个新的线程）。
newCachedThreadPool创建一个可缓存的线程池，如果当前线程池的长度超过了处理的需要时，它可以灵活的回收空闲的线程，当需求增加时，它可以灵活的添加新的线程。不对池的长度做任何限制。
newSingleThreadExecutor创建一个单线程化的executor，它只会创建唯一的工作者线程来执行任务，如果这个线程异常结束，会有另一个取代它，executor会保证任务依照任务队列所规定的顺序执行。
newScheduledThreadPool创建一个定长的线程池，而且支持定时的以及周期性的任务执行，类似于Timer

executorService暗示了生命周期3种状态，运行，关闭，终止。

Timer只创建唯一的线程来执行所有timer任务，如果一个timer任务的执行很耗时，会导致其他timerTask的时效准确性出问题。

TimerTask抛出未检查的异常，Timer会产生无法预料的行为。
调度线程池可以提供多个线程来执行延迟。并具有周期性的任务。

调度服务使用DelayQueue

Future:任务的状态决定了get方法的行为，如果任务已经完成，get会立即返回或者抛出一个Exception，如果任务没有完成，get会阻塞直到它完成，如果任务抛出了异常，get会把该异常封装成ExecutionException，然后重新抛出。
FutureTask实现了Runnable。所以既可以提交给Executor执行，也可以调用run方法运行。
```
void renderPage(){
	Callable<List<ImageData>> task = new Callable<>(){
	//...
		return result;
      };
      FutureTask<List<ImageData>> future = executor.submit(task);
      try{
       List<ImageData> imageData = future.get();

  		}catch(){
  		 Thread.currentThread.intterupt();
  		 //我们不需要结果，取消任务
  		 future.cancel(true);
  		}
}
```

CompletionService整合了Executor和BlockingQueue的功能，可以将callable任务提交给它去执行，使用类似队列的take和poll方法，结果完整时获得这个结果吗，像个打包的future，ExecutorCompletionService是一个实现类，将计算任务委托给一个executor
```
 void  renderPage(){
   CompletionService = new ExecutorCompletionService(executor);
   completionService.submit(new Callable(){
		return imageInfo.download();
   });
 }

 Future<ImageData> f = completionService.take();
 ImageData imageData = f.get();
 renderImage(imageData);
```
超时的Future
```
Future<Ad> f = exec.submit(new FetchAdTask());
long timeLeft = endNanos - System.nanoTime();
ad = f.get(timeLeft,NANOSECONDS);
```
使用限时的get方法通过future顺序的获取每一个结果
invokeAll将多个任务提交到一个ExecutorService，并且获得其结果，invokeAll处理一个任务的容器，并且返回一个future的容器，两个容器具有相同的结构，invokeAll将future添加到返回的容器中，这样可以使用任务容器的迭代器，所有任务完成时，调用线程被中断时或者超过时限时，限时版本的invokeAll都会返回结果，超过时限后，任何尚未完成的任务都会被取消，作为invokeAll的返回值，每个任务要么正常的完成，要么被取消。

线程中断是一个协作机制，一个线程给另一个线程发送信号，通知它在方便或者可能的情况下停止正在做的工作，去做其他事情。

静态的interrupted应该小心使用，他会清除并发线程的中断状态。如果你用了interrupted，并且它返回了true，必须对其进行处理，除非你想掩盖这个中断，可以抛出InterruptedException，或者通过再次调用interrupt保存中断状态。
中断通常是实现取消最明智的选择。
```
public static void timeRun(Runnable r,long timeout,TimeUnit unit){
	Future<?> task = taskExec.submit(r);
	try{
		task.get(timeout,unit);
	}finally{
		task.cancel(true);
	}
}



public void start(){
	Runtime.getRuntime().addShutdownHook(new Thread(){
		public void run(){
			try{LogService.this.stop();}
			catch(){}
		}
	});
}
```
线程池中不应该用ThreadLocal传递任务间的数值

线程池中如果一个任务依赖于其他任务的执行，就可能产生死锁，
对于一个单线程化的executor，一个任务将另一个任务提交到相同的executor,并等待新提交的任务的结果，这总会引发死锁。如果所有线程执行的任务都阻塞在线程池中，等待着仍然处于同一工作队列的其他任务，那么会发生线程饥饿死锁。
```
int N_CPU = Runtime.getRuntime().availableProcessors();
```
newFixedThreadPool为请求的池设置了核心池的大小和最大池的大小，而且池永远不会超时，newCachedThreadPool工厂将最大池的大小设置Integer.MAX_VALUE，核心池的大小设置为零，超时设置为一分钟。

对于庞大或者无限的池，可以使用synchronousQueue，完全绕开队列。
```
void processInParallel(Executor exec,List<Element> elements){
	for(final Element e:elements){
		exec.execute(new Runnable(){
			public void run(){
				process(e);
			}
		});
	}
}
```
所有的下载任务都进入到了executor的队列，就会立刻返回。而不用等到这些任务全部完成。如果需要提交一个任务集并等待他们，那么可以使用executorservice.invokeAll,当所有结果都可用后，可以使用completionService获取结果。

数据库检测到一个事务集发生了死锁（通过在表示正在等待关系的有向图上搜索循环），他会选择一个牺牲者，使它退出事务，这个牺牲者释放的资源，使得其他事务能够继续进行。应用程序可以重新执行那个被强行退出的事务，现在这个事务可能就能成功完成了。

如果所有线程以通过固定的秩序获得锁，程序就不会出现锁顺序死锁的问题了。

如果两个线程同时调用transferMoney，一个从X向Y转账，另一个从Y向X转账，那么就会发生死锁。
transferMoney(myAccount,yourAccount,10)
transferMoney(yourAccount,myAccount,10)


加时赛锁 --保证一次只有一个线程执行这个有风险的操作以未知的顺序获得锁。

在持有锁的时候调用外部方法是在挑战活跃度问题，外部方法可能会获得其他锁，或者遭遇严重超时的阻塞，当你持有锁的时候会延迟其他试图获得该锁的线程。

持有锁的时候调用了一个外部方法很难进行分析，因此是危险的。

检测死锁和从死锁从恢复的技术，是使用每个显示Lock类中定时tryLock特性，来替代使用内部锁机制。在内部锁机制中，只要没有获得锁，就会永远保持等待。而显示的锁使你能够定义超时的时间。在规定时间之后，tryLock还没有获得锁就返回失败，通过使用超时，尽管这段时间比你预期能够获得锁的时间长很多，你仍然可以再意外发生后重新获得控制权。使用可轮询检查的tryLock基本上避免了死锁发生的可能性。

kill -3出发线程触发线程转储

如果可运行的线程数大于CPU的数量，OS最终会强行换出正在执行的线程。从而使其他线程能够使用CPU，这会引起上下文切换，他会保存当前运行线程的执行上下文，并重建新调入线程的执行上下文。

切换上下文是有代价的，线程的调度需要操控OS和JVM共享的数据机构。你的程序与OS，JVM使用相同的CPU,CPU在JVM和OS的代码花费越多时间，意味着用于你的程序的时间就越少。但是JVM和OS活动花费并不是切换上下文开销唯一来源。当一个新的线程被换入后，它所需要的数据可能不在当前处理器本地的缓存中，所以切换上下文会引起缓存缺失的小恐慌，因此线程在第一次调度的时候会运行的稍慢一些。

当线程因为竞争一个锁而阻塞时，JVM通常会将这个线程挂起，允许它被换出，如果线程频繁发生阻塞，那线程就不能完整的使用它的调度限额了。一个程序发生越多的阻塞（阻塞IO，等待竞争锁，或者等待条件变量）,与受限于CPU的程序相比，就会造成越多的上下文切换，这增加了调度的开销，并减少了吞吐量。

vmstat报告上下文切换次数和内核占用的时间等信息，高内核占用率通常象征繁重的调度活动，这很可能是IO阻塞。或竞争锁引起的。

synchronized和volatile提供的可见性保证要求使用一个特殊的，存储关卡(memory barrier)的指令，来刷新缓存，使缓存无效，刷新硬件的写缓冲，并延迟执行的传递，从存储关卡可能同样会对性能产生影响。因为它们抑制了其他编译器的优化，在存储关卡中，大多数操作是不能被重排序的。

>逸出分析
>锁的粗化

自旋等待（spin-waiting不断尝试获取锁，直到成功）

分拆锁和分离锁：
也就是采用相互独立的锁，守卫多个独立的状态变量，在改变之前，它们都是由一个锁守护的。

ConcurrentHashMap实现使用了一个包含16个锁的Array，每一个锁都守护Hash Bucket的1/16，Bucket N由第N mod 16个锁来守护，架设哈希提供合理的拓展特性，并且关键字能够以统一的方式访问，这将会把对于所的请求减少到约为原来的1/16,这使得ConcurrentHashMap能够支持16个并发的Writer，为了对多处理器系统的大负荷访问提供更好的并发性，锁的数量还可以增加。
分离锁的一个负面作用是：对容器加锁，进行独占访问更加困难，并且更加昂贵了。
ConcurrentHashMap通过枚举每个条目获得size，并把这个值加入到每个条目。而不是维护一个全局计数，为了避免列举所有元素，ConcurrentHashMap为每个条目维护一个独立的计数域，同样由分离的锁守护。

iostat检测是否受限于磁盘

Object.wait和Condition.await允许发出假唤醒，让一个线程从WAITING或TIMED_WAITING临时的转换为RUNNABLE，即使它等待的条件尚未成真。
thread.interrupt();
thread.join(LOCKUP_DETECT_TIMEOUT);定时的join能确保测试完成。


使用Thread.yield()激发更多的上下文转换
```
public class BarrierTimer implements Runnable{
	private boolean started;
	private long startTime,endTime;
	public synchronized void run(){
		long t = System.nanoTime();
		if(!started){
			started = true;
			startTime = t;
		}else{
			endTime = t;
		}
	}
	public synchronized void clear(){
		started = false;
	}
	public synchronized long getTime(){
		return endTime - startTime;
	}
}
```
连接队列的put和take操作允许有比基于数组的队列更好的并发访问，这是因为最佳的链接队列算法允许队列的头和尾彼此独立的更新，由于分配操作通常是线程本地的，算法通过多做一些分配操作可以降低竞争。

如果需要等待一个状态转换的发生，闭锁或者条件等待通常是更好的技术，而不是自旋循环。

内部锁的局限：不能中断那些正在等待获取锁的线程，并且在请求锁失败的情况下必须无限等待，内部锁必须在获取他们的代码块中被释放，这很好的简化了代码，与异常处理机制能够良好的互动。

可定时的与可轮询的锁获取模式，是由tryLock方法实现，与无条件的锁获取相比，它具有更完善的错误恢复机制。内部锁中，死锁是致命的--唯一的恢复方法是重新启动程序，唯一的预防方法是在构建程序时不要出错。所以不可能允许不一致的锁顺序。

使用tryLock试图获得两个锁，如果不能同时获得两个，就回退，重新尝试，休眠时间由一个特定的组件管理，并由一个随机组件减小活锁发生的可能性。如果一定时间没有能获得所需要的锁，就返回一个失败状态。

定时锁能够在时间预算内设定相应的超时，如果活动在期待的时间内没能获得结果，这个机制使得程序能够提前返回，使用内部锁一旦开始请求，锁就不能停止了。
```
if(!lock.tryLock(nanosToLock,NANOSECONDS)){
	return false;
}
```
lock.lockInterruptibly() 可中断的锁

ReentrantLock创建公平锁和非公平锁--非公平锁允许闯入，但不是有意鼓励闯入，倘若遇到闯入的发生，他们不会有意避开，公平锁中，如果锁已经被其他线程占有，新的请求线程会加入到等待队列，或者已经有一些线程在等待锁了，在非公平锁中，线程只有当锁正在被占用时才会等待。

公平会因为挂起和重新开始线程的代价带来巨大的性能开销，实践中，统计上的公平性保证--承诺一个阻塞的线程最终能够获得锁

闯入锁比公平锁性能好的原因之一是--挂起的线程重新开始，与它真正开始运行，两者之间会产生严重的延迟。
当持有锁的时间相对较长，或者请求锁的平均时间间隔比较长，那么使用公平锁是比较好的。

多处理器系统中，频繁的访问主要是为读取数据结构的时候，读-写能够改进性能，在其他情况下运行的情况比独占的锁要稍差一些。
ReadWriteLock实现选择：
释放优先
读者闯入
重进入
降级
升级


ReentrantReadWriteLock为两个锁提供了可重进入的加锁语义，可以被构造为非公平或者公平的。公平锁中选择权交给等待时间最长的线程，如果锁由读者获得，而一个线程请求写入锁，那么不在允许读者获得读取锁。直到写者被受理。并且已经释放了写入锁。

锁持有的时间相对较长，并且大部分操作都不会改变锁守护的资源，那么读写锁能够改进并发性。

生产者--消费者的的设计经常使用ArrayBlockingQueue这种有限缓存，一个有限缓存提供的put和take操作，都有一个先验条件，你不能从空缓存中获取元素，也不能把元素置入已满的缓存中，如果依赖于状态的操作在处理先验条件时失败，可以抛出异常或者返回错误状态。也可以保持阻塞直到对象转入正确的状态。

调用者选择休眠(sleep)以避免消耗过多的CPU时间。
Threa.yield可以给调度器一个提示，我现在可以让出一定的时间让另外的线程运行。

条件队列可以让一组线程--称为等待集--以某种方式等待相关条件变成真。不同于传统的队列，他们的元素是数据项，条件队列的元素是等待相关条件的线程。
Object的wait,notify,notifyAll构成了内部条件队列的API
一个对象的内部锁与它的内部条件队列是相关的。为了调用对象中任一个条件队列方法，必须持有对象锁，因为等待基于状态的条件机制必须和维护状态一致性紧密绑定在一起。除非你能检查状态，否则你不能等待条件。同时除非你能改变状态。否则你不能从条件等待队列中释放其他的线程。
Object.wait会自动释放锁，请求OS挂起当前线程，让其他线程获得该锁进而修改对象的状态，当它被唤醒时，它会在返回前重新获得锁。
```
public synchronized void put(V v){
	while(isFull())
		wait();
	doPut(v);
	notifyAll();
}

public synchronized V take(){
	while(isEmpty())
		wait();
	V v = doTake();
	notifyAll();
	return v;
}
```
wait会释放锁，并阻塞当前线程，然后等待，直到特定时间超时过后，线程被中断或被通知唤醒，线程被唤醒后。wait会在返回运行前重新请求锁，一个从wait方法中唤醒的线程，在重新请求锁的过程没有任何特殊的优先级。

notifyAll唤醒后，控制流重新进入调用wait代码，重新请求与条件队列相关联的锁，在重新请求锁的时刻又再次变为假
```
while(!conditionPredicate()){
	lock.wait();
}
```
优先使用notifyAll();

二元闭锁：只有初始状态和终止状态。闭锁会阻止线程通过开始阀门，直到阀门被打开。此时所有的线程都可以通过。，但是闭锁阀门一旦被打开就不能再重新关闭。

每个Lock都有任意数量的Condition对象，使用的是await() --signal() --  aingall()

ReentrantLock和Semaphore这两个接口特点：都扮演了类似阀门角色。每次只允许有限数目的线程通过它，线程到达阀门后，可以允许通过(lock或acquire成功返回),可以等待(lock或acquire阻塞),也可以被取消(tryLock或tryAcquire返回false， 指明在允许的时间内，锁或者许可不可用)，他们都允许可中断，不可中断的，可限时的请求尝试。

AbstractQueuedSynchronizer(AQS)是一个用来构建锁和synchronizer的框架。

一个基于AQS的Synchronizer执行的基本操作是不同形式的获取和释放

ReentrantLock只支持独占的获取操作

当一个线程尝试去获取锁时，tryAcquire会首先请求锁的状态，如果锁未被占有，它就会尝试更新锁的状态，表明锁已被占有。
tryAcquire使用compareAndSetState尝试原子的更新状态。

FutureTask使用AQS类型同步状态来持有任务的状态，运行，完成或取消。

volatile不会引起上下文切换和线程调度。但是不能用于构建原子化的复合操作。

独占锁是一项悲观的技术--假设最坏情况(如果不锁门，捣蛋鬼就会闯入破坏物品秩序)，并且会通过正确的锁来避免其他线程的打扰。直到做出保证才能继续进行。

细粒度的操作有乐观的解决办法--凭借新的方法，我们可以指望不受打扰的完成更新，这个方法依赖于冲突检测。从而判定更新过程是否存在其他成员的干涉。

CAS--比较并交换，有3个操作数，内存位置V，旧的预期值A和新值B。当且仅当V符合旧预期值A时，CAS用新值B原子化的更新V的值，否则他什么都不做，都会返回V的知识值。（compare-and-set）
是一项乐观技术。
```
public class ConcurrentStack<E> {
	AtomicReference<Node<E>> top = new AtomicReference<>();
	public void push(E item){
		Node<E> newHead = new Node<>(item);
		Node<E> oldHead;
		do{
			oldHead = top.get();
			newHead.next = oldHead;
		}while(!top.compareAndSet(oldHead,newHead));
	}
	public E pop(){
		Node<E> oldHead;
		Node<E> newHead;
		do{
			oldHead = top.get();
			if(oldHead == null)
				return null;
			newHead = oldHead.next;
			oldHead = null;
		}while(!top.compareAndSet(oldHead,newHead));
		return oldHead.item;
	}

	public static class Node<E>{
		public final E item;
		public Node<E> next;
		public Node(E item){
			this.item = item;
		}
	}
}
```

ABA问题是因为在算法中误用比较并交换而引起的反常现象，节点被循环使用（主要存在于不能被垃圾回收的环境中）。CAS的有效请求仍为A么，并且如果成立就继续处理更新，算法中如果进行自身链接节点对象的内存管理，那么就可能出现ABA问题。

简单方案--更新一对值，包括引用和版本号，而不是仅更新该值的引用。

JMM为所有程序内部的动作定义了一个偏序关系叫happens-before，要想保证执行动作B的线程看到动作A的结果（无论A和B是否发生在同一个线程中），A和B之间就必须满足happens-before关系。未按照这个关系排序，JVM可以对他们随意的重排序

