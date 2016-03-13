由于数组的长度非固定，可以动态增删，因此数组为引用类型：
```javascript
var array = [1,2,3,4,5];
var arrayRef = array;
array.push(6);
print(arrayRef);
```
引用指向的是地址，也就是说，引用不会指向引用本身，而是指向该引用所对应的实际对象。因此通过修改array指向的数组，则arrayRef指向的是同一个对象，因此运行效果如下：
1,2,3,4,5,6

变量被定义的区域即为其作用域，全局变量具有全局作用域；局部变量，比如声明在函数内部的变量则具有局部作用域，在函数的外部是不能直接访问的

原型(prototype)，是JavaScript特有的一个概念，通过使用原型，JavaScript可以建立其传统OO语言中的继承，从而体现对象的层次关系。JavaScript本身是基于原型的，每个对象都有一个prototype的属性，这个prototype本身也是一个对象，因此它本身也可以有自己的原型，这样就构成了一个链结构。
访问一个属性的时候，解析器需要从下向上的遍历这个链结构，直到遇到该属性，则返回属性对应的值，或者遇到原型为null的对象(JavaScript的基对象Object的构造器的默认prototype有一个null原型)，如果此对象仍没有该属性，则返回undefined.

由于遍历原型链的时候，是有下而上的，所以最先遇到的属性值最先返回，通过这种机制可以完成继承及重载等传统的OO机制。

，而在JavaScript中，this表示当前上下文，即调用者的引用。

，this的值并非函数如何被声明而确定，而是被函数如何被调用而确定，这一点与传统的面向对象语言截然不同
//设置printName的上下文为jack, 此时的this为jack
print(printName.call(jack));
//设置printName的上下文为abruzzi,此时的this为abruzzi
print(printName.call(abruzzi));
运行结果：
jack
Abruzzi

 通过new操作符作用于Object对象，构造一个新的对象，然后动态的添加属性，从无到有的构建一个对象。
 定义对象的“类”：原型，然后使用new操作符来批量的构建新的对象。
 使用JSON，这个在下一节来进行详细说明。
这一节我们详细说明第二种方式，如：
```javascript
//定义一个"类"，Address
function Address(street, xno){
this.street = street || 'Huang Quan Road';
this.xno = xno || 135;
this.toString = function(){
return "street : " + this.street + ", No : " + this.xno;
}
}

var obj = {
name : "abruzzi",
age : 26,
birthday : new Date(1984, 4, 5),
addr : {
street : "Huang Quan Road",
xno : "135"
}
}
```
JSON的另一个应用场景是：当一个函数拥有多个返回值时，在传统的面向对象语言中，我们需要组织一个对象，然后返回，而JavaScript则完全不需要这么麻烦，比如：
```javascript
function point(left, top){
this.left = left;
this.top = top;
//handle the left and top
return {x: this.left, y:this.top};
}
直接动态的构建一个新的匿名对象返回即可：
var pos = point(3, 4);
//pos.x = 3;
//pos.y = 4;
```
而在JavaScript中，函数本身与其他任何的内置对象在地位上是没有任何区别的，也就是说，函数本身也是对象。
总的来说，函数在JavaScript中可以：
 被赋值给一个变量
 被赋值为对象的属性
 作为参数被传入别的函数
 作为函数的结果被返回
 用字面量来创建

function关键字会调用Function来new一个对象，并将参数表和函数体准确的传递给Function的构造器。
通常来说，在全局作用域(作用域将在下一节详细介绍)内声明一个对象，只不过是对一个属性赋值而已，比如上例中的add函数，事实上只是为全局对象添加了一个属性，属性名为add，而属性的值是一个对象，即function(x, y){return x+y;}，理解这一点很重要，这条语句在语法上跟：
var str = "This is a string";
并无二致。都是给全局对象动态的增加一个新的属性，如此而已。

事实上，JavaScript在处理函数的参数时，与其他编译型的语言不一样，解释器传递给函数的是一个类似于数组的内部值，叫arguments，这个在函数对象生成的时候就被初始化了。比如我们传递给adPrint一个参数的情况下，其他两个参数分别为undefined.这样，我们可以才adPrint函数内部处理那些undefined参数，从而可以向外部公开：我们可以处理任意参数。

JavaScript中的变量作用域为函数体内有效，而无块作用域，

JavaScript的函数是在局部作用域内运行的，在局部作用域内运行的函数体可以访问其外层的(可能是全局作用域)的变量和函数。JavaScript的作用域为词法作用域，所谓词法作用域是说，其作用域为在定义时(词法分析时)就确定下来的，而并非在执行时确定，
```javascript
var str = "global";
function scopeTest(){
print(str);
var str = "local";
print(str);
}
scopeTest();
```
运行结果是什么呢？初学者很可能得出这样的答案：
global
local
而正确的结果应该是：
undefined
local

那么，局部变量又隶属于什么对象呢？就是我们要讨论的调用对象。在执行一个函数时，函数的参数和其局部变量会作为调用对象的属性进行存储。同时，解释器会为函数创建一个执行器上下文(context)，与上下文对应起来的是一个作用域链。顾名思义，作用域链是关于作用域的链，通常实现为一个链表，链表的每个项都是一个对象，在全局作用域中，该链中有且只有一个对象，即全局对象。对应的，在一个函数中，作用域链上会有两个对象，第一个(首先被访问到的)为调用对象，第二个为全局对象。
应该注意的是，作用域链随着嵌套函数的层次会变的很长，但是查找变量的过程依旧是遍历作用域链(链表)，一直向上查找，直到找出该值，如果遍历完作用域链仍然没有找到对应的属性，则返回undefined

函数的上下文是可以变化的，因此，函数内的this也是可以变化的，函数可以作为一个对象的方法，也可以同时作为另一个对象的方法，总之，函数本身是独立的。可以通过Function对象上的call或者apply函数来修改函数的上下文：

call和apply通常用来修改函数的上下文，函数中的this指针将被替换为call或者apply的第一个参数

apply的第二个参数为一个函数需要的参数组成的一个数组，而call则需要跟若干个参数，参数之间以逗号(,)隔开即可。

只有一个参数的时候call和apply的使用方式是一样的，如果有多个参数：
```javascript
setName.apply(jack, ["Jack Sept."]);
print(printName.apply(jack));
setName.call(abruzzi, "John Abruzzi");
print(printName.call(abruzzi));
得到的结果为：
Jack Sept.
John Abruzzi
```
//声明一个函数，接受两个参数，返回其和
```javascript
function add(x, y){
return x + y;
}
var a = 0;
a = add;//将函数赋值给一个变量
var b = a(2, 3);//调用这个新的函数a
print(b);
```
这段代码会打印”5”，因为赋值之后，变量a引用函数add，也就是说，a的值是一个函数对象(一个可执行代码块)，因此可以使用a(2, 3)这样的语句来进行求和操作。
//高级打印函数的第二个版本
```javascript
function adPrint2(str, handler){
print(handler(str));
}
//将字符串转换为大写形式，并返回
function up(str){
return str.toUpperCase();
}
//将字符串转换为小写形式，并返回
function low(str){
return str.toLowerCase();
}
adPrint2("Hello, world", up);
adPrint2("Hello, world", low);
```
应该注意到，函数adPrint2的第二个参数，事实上是一个函数，将这个处理函数作为参数传入，在adPrint2的内部，仍然可以调用这个函数，这个特点在很多地方都是有用的，特别是，当我们想要处理一些对象，但是又不确定以何种形式来处理，则完全可以将“处理方式”作为一个抽象的粒度来进行包装(即函数)。

另一个与其他语言的数组不同的是，字符串也可以作为数组的下标，事实上，在JavaScript的数组中，字符串型下标和数字型的下标会被作为两个截然不同的方式来处理，一方面，如果是数字作为下标，则与其他程序设计语言中的数组一样，可以通过index来进行访问，而使用字符串作为下标，就会采用访问JavaScript对象的属性的方式进行，毕竟JavaScript内置的Array也是从Object上继承下来的。

slice方法的第一个参数为起始位置，第二个参数为终止位置，操作不影响数组本身。下面我们来看splice方法，虽然这两个方法的拼写非常相似，但是功用则完全不同，事实上，splice是一个相当难用的方法：
```javascript
bigArray.splice(5, 2);
bigArray.splice(5, 0, "very", "new", "item", "here");
```
第一行代码表示，从bigArray数组中，从第5个元素起，删除2个元素；而第二行代码表示，从第5个元素起，删除0个元素，并把随后的所有参数插入到从第5个开始的位置，则操作结果为：
one,two,three,four,five,very,new,item,here,another,array,yet,another,array
```javascript
//Array Remove - By John Resig (MIT Licensed)
Array.prototype.remove = function(from, to) {
var rest = this.slice((to || from) + 1 || this.length);
this.length = from < 0 ? this.length + from : from;
return this.push.apply(this, rest);
};
```
这个函数扩展了JavaScript的内置对象Array，这样，我们以后的所有声明的数组都会自动的拥有remove能力

[\w-]表示所有的字符，数字，下划线及减号，[\w-]+表示这个集合最少重复一次，然后紧接着的这个括号表示一个分组
(*)，表示重复零或多次。这样就可以匹配任意字母，数字，下划线及中划线的集合，且至少重复一次。
而@符号之后的部分与前半部分唯一不同的是，后边的一个分组的修饰符为(+)，表示至少重复一次，那就意味着后半部分至少会有一个点号(.)，而且点号之后至少有一个字符。这个修饰主要是用来限制输入串中必须包含域名。
最后，脱字符(^)和美元符号($)限制，以„„开始，且以„„结束。这样，整个表达式的意义就很明显了。

第二种情况，括号用来分组，当正则表达式执行完成之后，与之匹配的文本将会按照规则填入各个分组，比如，某个数据库的主键是这样的格式：四个字符表示省份，然后是四个数字表示区号，然后是两位字符表示区县，如yunn0871cg表示云南省昆明市呈贡县(当然，看起来的确很怪，只是举个例子)，我们关心的是区号和区县的两位字符代码，怎么分离出来呢？
```javascript
var pattern = /\w{4}(\d{4})(\w{2})/;
var result = pattern.exec("yunn0871cg");
print("city code = "+result[1]+", county code = "+result[2]);
result = pattern.exec("shax0917cc");
print("city code = "+result[1]+", county code = "+result[2]);
```
正则表达式的exec方法会返回一个数组(如果匹配成功的话)，数组的第一个元素(下标为0)表示整个串，第一个元素为第一个分组，第二个元素为第二个分组，以此类推。因此上例的执行结果即为：
```javascript
city code = 0871, county code = cg
city code = 0917, county code = cc
```
我们的正则表达式还是可以匹配，注意这两个字符串两侧的引号不匹配！我们需要的是，前边是单引号，则后边同样是单引号，反之亦然。因此，我们需要知道前边匹配的到底是“单”还是“双”。这里就需要用到引用，JavaScript中的引用使用斜杠加数字来表示，如\1表示第一个分组(括号中的规则匹配的文本)，\2表示第二个分组，以此类推。因此我们就设计出了这样的表达式：
var pattern = /(['"])[^'"]*\1/;
在我们新设计的这个语言中，为了某种原因，在单引号中我们不允许出现双引号，同样，在双引号中也不允许出现单引号，我们可以稍作修改即可完成：
var pattern = /(['"])[^\1]*\1/;

RegExp对象的方法：
方法名
描述 test() 测试串中是否有合乎模式的匹配
exec()
对串进行匹配 compile() 编译正则表达式

match 匹配正则表达式，返回匹配数组
replace
替换 split 分割
search
查找，返回首次发现的位置


：由于JavaScript中，函数是对象，对象是属性的集合，而属性的值又可以是对象，则在函数内定义函数成为理所当然，如果在函数func内部声明函数inner，然后在函数外部调用inner，这个过程即产生了一个闭包。

。闭包允许你引用存在于外部函数中的变量。然而，它并不是使用该变量创建时的值，相反，它使用外部函数中该变量最后的值。

，所有的变量，如果不加上var关键字，则默认的会添加到全局对象的属性上去，这样的临时变量加入全局对象有很多坏处，比如：别的函数可能误用这些变量；造成全局对象过于庞大，影响访问速度(因为变量的取值是需要从原型链上遍历的)。除了每次使用变量都是用var关键字外，我们在实际情况下经常遇到这样一种情况，即有的函数只需要执行一次，其内部变量无需维护，比如UI的初始化，那么我们可以使用闭包：
```javascript
var datamodel = {
table : [],
tree : {}
};
(function(dm){
for(var i = 0; i < dm.table.rows; i++){
var row = dm.table.rows[i];
for(var j = 0; j < row.cells; i++){
drawCell(i, j);
}
}
//build dm.tree
})(datamodel);
```
我们创建了一个匿名的函数，并立即执行它，由于外部无法引用它内部的变量，执行完后很快就会被释放，最主要的是这种机制不会污染全局对象。

闭包的另一个重要用途是实现面向对象中的对象，传统的对象语言都提供类的模板机制，这样不同的对象(类的实例)拥有独立的成员及状态，互不干涉。虽然JavaScript中没有类这样的机制，但是通过使用闭包，我们可以模拟出这样的机制
```javascript
function Person(){
var name = "default";
return {
getName : function(){
return name;
},
setName : function(newName){
name = newName;
}
}
};
var john = Person();
print(john.getName());
john.setName("john");
print(john.getName());
var jack = Person();
print(jack.getName());
jack.setName("jack");
print(jack.getName());
运行结果如下：
default
john
default
jack
```
JavaScript的解释器都具备垃圾回收机制，一般采用的是引用计数的形式，如果一个对象的引用计数为零，则垃圾回收机制会将其回收，这个过程是自动的。但是，有了闭包的概念之后，这个过程就变得复杂起来了，在闭包中，因为局部的变量可能在将来的某些时刻需要被使用，因此垃圾回收机制不会处理这些被外部引用到的局部变量，而如果出现循环引用，即对象A引用B，B引用C，而C又引用到A，这样的情况使得垃圾回收机制得出其引用计数不为零的结论，从而造成内存泄漏。
```javascript
$(function(){
var con = $("div#panel");
this.id = "content";
con.click(function(){
alert(this.id);//panel
});
});
```
此处的alert(this.id)到底引用着什么值呢？很多开发者可能会根据闭包的概念，做出错误的判断：
content
理由是，this.id显示的被赋值为content,而在click回调中，形成的闭包会引用到this.id，因此返回值为content。然而事实上，这个alert会弹出”panel”，究其原因，就是此处的this,虽然闭包可以引用局部变量，但是涉及到this的时候，情况就有些微妙了，因为调用对象的存在，使得当闭包被调用时(当这个panel的click事件发生时)，此处的this引用的是con这个jQuery对象。而匿名函数中的this.id = “content”是对匿名函数本身做的
操作。两个this引用的并非同一个对象。
如果想要在事件处理函数中访问这个值，我们必须做一些改变：
```javascript
$(function(){
var con = $("div#panel");
this.id = "content";
var self = this;
con.click(function(){
alert(self.id);//content
});
});
```
这样，我们在事件处理函数中保存的是外部的一个局部变量self的引用，而并非this。这种技巧在实际应用中多有应用，我们在后边的章节里进行详细讨论。关于闭包的更多内容，我们将在第九章详细讨论，包括讨论其他命令式语言中的“闭包”，闭包在实际项目中的应用等等。

。在JavaScript中，通过new操作符来作用与一个函数，实质上会发生这样的动作：
首先，创建一个空对象，然后用函数的apply方法，将这个空对象传入作为apply的第一个参数，及上下文参数。这样函数内部的this将会被这个空的对象所替代：
```javascript
var triangle = new Shape("triangle");
//上一句相当于下面的代码
var triangle = {};
Shape.apply(triangle, ["triangle"]);

var outter = function(){
var x = 0;
return function(){
return x++;
}
}
var a = outter();
print(a());
print(a());
var b = outter();
print(b());
print(b());
运行结果为：
0
1
0
1
```
变量a通过闭包引用outter的一个内部变量，每次调用a()就会改变此内部变量，应该注意的是，当调用a时，函数outter已经返回了，但是内部变量x的值仍然被保持。而变量b也引用了outter，但是是一个不同的闭包，所以b开始引用的x值不会随着a()被调用而改变，两者有不同的实例，这就相当于面向对象中的不同实例拥有不同的私有属性，互不干涉。
```javascript
function func(){
//do something
}
var func = function(){
//do something
}
```
这两个语句的意义是一样的，它们都表示，为全局对象添加一个属性func，属性func的值为一个函数对象，而这个函数对象是匿名的。匿名函数的用途非常广泛，在JavaScript代码中，我们经常可以看到这样的代码：
```javascript
var mapped = [1, 2, 3, 4, 5].map(function(x){return x * 2});
print(mapped);
```

//update会返回一个函数，这个函数可以设置id属性为item的web元素的内容
```javascript
function update(item){
return function(text){
$("div#"+item).html(text);
}
}
//Ajax请求，当成功是调用参数callback
function refresh(url, callback){
var params = {
type : "echo",
data : ""
};
$.ajax({
type:"post",
url:url,
cache:false,
async:true,
dataType:"json",
data:params,
//当异步请求成功时调用
success: function(data, status){
callback(data);     //update("newsPanel")(data =>text)
},
//当请求出现错误时调用
error: function(err){
alert("error : "+err);
}
});
}
refresh("action.do?target=news", update("newsPanel")); refresh("action.do?target=articles", update("articlePanel")); refresh("action.do?target=pictures", update("picturePanel"));
```
其中，update函数即为柯里化的一个实例，它会返回一个函数，即：
```javascript
update("newsPanel") = function(text){
$("div#newsPanel").html(text);
}
```
由于update(“newsPanel”)的返回值为一个函数，需要的参数为一个字符串，因此在refresh的Ajax调用中，当success时，会给callback传入服务器端返回的数据信息，从而实现newsPanel面板的刷新，其他的文章面板articlePanel,图片面板picturePanel的刷新均采取这种方式，这样，代码的可读性，可维护性均得到了提高。

lambda演算的先驱们，天才的发明了一个神奇的函数，成为Y-结合子。使用Y-结合子，可以做到对匿名函数使用递归。关于Y-结合子的发现及推导过程的讨论已经超出了本部分的范围，有兴趣的读者可以参考附录中的资料。我们来看看这个神奇的Y-结合子：
```javascript
var Y = function(f) {
return (function(g) {
return g(g);
})(function(h) {
return function() {
return f(h(h)).apply(null, arguments);
};
});
};
```
事实上，在JavaScript中，我们有一种简单的方法来实现Y-结合子：
```javascript
var fact = function(x){
return x == 0 : 1 : x * arguments.callee(x-1);
}
fact(10);
```
或者：
```javascript
(function(x){
return x == 0 ? 1 : x * arguments.callee(x-1);
})(10);//3628800
```
其中，arguments.callee表示函数自身，而arguments.caller表示函数调用者，因此省去了很多复杂的步骤。

能有多个。
引擎在调用一个函数时，进入该函数上下文，并执行函数体，与其他程序设计语言类似，函数体内可以有递归，也可以调用其他函数(进入另外一个上下文，此时调用者被阻塞，直至返回)。调用eval会有类似的情况。

所谓词法作用域(静态作用域)是指，在函数对象的创建时，作用域”[[scope]]”就已经建立，而并非到执行时，因为函数创建后可能永远不会被执行，但是作用域是始终存在的。

我们在基础部分提到过，JavaScript代码都有一个执行环境，所有的JavaScript代码均被包含在一个全局对象中，比如在所有函数之外声明：
```javascript
var x = 3;
var str = "global";
这两个变量声明语句事实上是为全局对象添加了两个属性x和str。可以将全局对象想象成一个大的匿名自执行函数：
(function(){
var x = 3;
var str = "global";
//...
})();
```
在浏览器端，这个全局的对象称为window，window是所有JavaScript对象的根，我们可以通过window对象的属性document来访问页面本身，也可以调用window的一些方法来与用户交互

