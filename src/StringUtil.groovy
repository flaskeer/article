/**
 * Created by hao on 2015/6/12.
 */
class StringUtil {
    static String sayHello(String name){
        if(name == "Martin" || name == "Ben"){
            "Hello author" + name + "!"
        }else{
            "Hello reader " + name + "!"
        }
    }

    def sayGoodbye = {
        name ->
        if(name == "Martin" || name == "Ben"){
            "Hello author" + name + "!"
        }else{
            "Hello reader " + name + "!"
        }
    }

}


