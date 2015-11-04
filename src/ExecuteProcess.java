import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * Created by hao on 2015/6/20.
 */
public class ExecuteProcess {

    public static void main(String[] args) {
        try{
            Process proc = Runtime.getRuntime().exec("svn help");
            BufferedReader br = new BufferedReader(new InputStreamReader(proc.getInputStream()));
            String line;
            while((line = br.readLine()) != null){
                System.out.println(line);
            }
        }catch(IOException e){
            e.printStackTrace();
        }
    }
}
