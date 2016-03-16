package com.sung.netty.practice.client;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.concurrent.LinkedBlockingQueue;
import java.util.concurrent.TimeUnit;

/**
 * Created by user on 2016/3/16.
 */
public class NettyClientHandler extends SimpleChannelInboundHandler<String>{

    private static Logger logger = LoggerFactory.getLogger(NettyClientHandler.class);

    private final LinkedBlockingQueue<String> queue;

    public NettyClientHandler() {
        queue = new LinkedBlockingQueue<>();
    }

    @Override
    protected void channelRead0(ChannelHandlerContext ctx, String msg) throws Exception {
        queue.add(msg);
    }

//    @Override
//    public void channelActive(ChannelHandlerContext ctx) throws Exception {
//        ctx.writeAndFlush("send message to server....");
//        TimeUnit.SECONDS.sleep(1000);
//
//    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception {
        logger.error("unexpection happens:{}",cause);
        ctx.close();
    }

    public String getMessage() {
        String message = null;
        try {
            message = queue.poll(1,TimeUnit.SECONDS);
        } catch (InterruptedException e) {
            logger.error(e.getMessage(),e);
        }
        return message;
    }
}
