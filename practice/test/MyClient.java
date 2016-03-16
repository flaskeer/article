package com.sung.netty.test;

import com.sung.netty.practice.client.ClientChannelInitializer;
import com.sung.netty.practice.client.NettyClient;

/**
 * Created by user on 2016/3/16.
 */
public class MyClient {

    public static void main(String[] args) throws Exception {
        NettyClient client = new NettyClient("localhost",8080,new ClientChannelInitializer());
        while(client.isConnect()){
            System.out.println("+++++");
            System.out.println(client.receiveMessage());
            System.out.println("------");
        }
        System.out.println("==========================");
        client.close();
    }
}
