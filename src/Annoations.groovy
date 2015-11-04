import groovy.transform.*

/**
 * Created by hao on 2015/6/20.
 */



//@Canonical(excludes = "lastName,age")
//class Person{
//
//    String firstName
//    String lastName
//    int age
//
//}
//def sara = new Person(firstName: "Sars",lastName: "walker",age: 44)
//println sara

//class Worker{
//    def work() {println 'get work done'}
//    def analyze() {println 'analyze'}
//    def writeReport() {println 'get report written'}
//}
//
//class Expert{
//    def analyze() { println "export"}
//}
//
//class Manager{
//    @Delegate Expert expert = new Expert()
//    @Delegate Worker worker = new Worker()
//
//}
//
//def bernie = new Manager()
//bernie.analyze()
//bernie.work()
//bernie.writeReport()


//class Heavy{
//    def size = 10
//    Heavy() {println "Creating heavy with $size"}
//}
//
//class AsNeeded{
//    def value
//    @Lazy Heavy heavy1 = new Heavy()
//    @Lazy Heavy heavy2 = {new Heavy(size: value)}()
//    AsNeeded() {println "Created AsNeeded"}
//}
//
//def asNeeded = new AsNeeded(value: 100)
//println asNeeded.heavy1.size
//println asNeeded.heavy1.size
//println asNeeded.heavy2.size


@Singleton(lazy = true)
class TheUnique{

    private TheUnique() {println 'Instance created'}

    def hello() {println 'hello'}
}

println "Accessing the unique"
TheUnique.instance.hello()
TheUnique.instance.hello()











