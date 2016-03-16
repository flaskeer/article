package com.sung.netty.practice;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Created by user on 2016/3/15.
 */
public class ClientHandler extends SimpleChannelInboundHandler<String>{

    private static final Logger logger = LoggerFactory.getLogger(ClientHandler.class);


    @Override
    protected void channelRead0(ChannelHandlerContext ctx, String msg) throws Exception {
        System.out.println(msg);
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) throws Exception {
        logger.error("expection happen:{}",cause);
        ctx.close();
    }
}
