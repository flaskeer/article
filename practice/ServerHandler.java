package com.sung.netty.practice;

import io.netty.channel.ChannelFuture;
import io.netty.channel.ChannelFutureListener;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetAddress;
import java.util.Date;

/**
 * Created by user on 2016/3/15.
 */
public class ServerHandler extends SimpleChannelInboundHandler<String>{

    private static final Logger logger = LoggerFactory.getLogger(ServerHandler.class);

    @Override
    public void channelActive(ChannelHandlerContext ctx) throws Exception {
        ctx.write("welcome to " + InetAddress.getLocalHost().getHostName() + "!\r\n");
        ctx.write("It is" + new Date() + "now\r\n");
        ctx.flush();
    }

    @Override
    protected void channelRead0(ChannelHandlerContext ctx, String msg) throws Exception {
        String response = null;
        boolean close = false;
        if(msg.isEmpty()) {
            response = "please input something....\r\n";
        } else if("bye".equals(msg)) {
            response = "have a good day!\r\n";
            close = true;
        } else{
            response = "did you say :" + msg + "?\r\n";
        }
        ChannelFuture future = ctx.write(response);
        if(close) {
            future.addListener(ChannelFutureListener.CLOSE);
        }


    }

    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) throws Exception {
        ctx.flush();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception {
        logger.error("unexpected exception:{}",cause);
        ctx.close();
    }
}
