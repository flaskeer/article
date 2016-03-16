package com.sung.netty.practice;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.Channel;
import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelOption;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

/**
 * Created by user on 2016/3/15.
 */
public class TelnetClient {

    private  final String host;
    private  final int port;

    public TelnetClient(String host, int port) {
        this.host = host;
        this.port = port;
    }

    public void run() throws InterruptedException, IOException {
        EventLoopGroup group = new NioEventLoopGroup();
        try{
            Bootstrap b = new Bootstrap();
            b.group(group)
             .channel(NioSocketChannel.class)
             .handler(new ClientInitialize())
             .option(ChannelOption.SO_BACKLOG,128);
            Channel channel = b.connect(host, port).sync().channel();
            ChannelFuture lastFuture = null;
            BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
            for (;;) {
                String line = reader.readLine();
                if(line == null) {
                    break;
                }
                lastFuture = channel.writeAndFlush(line + "\r\n");
                if("bye".equals(line.toLowerCase())){
                    channel.closeFuture().sync();
                    break;
                }

                if(lastFuture != null) {
                    lastFuture.sync();
                }
            }
        } finally {
            group.shutdownGracefully();
        }

    }

    public static void main(String[] args) throws Exception {
        new TelnetClient("localhost",8080).run();
    }
}
