package com.sung.netty.test;

import io.netty.channel.ChannelInitializer;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.DelimiterBasedFrameDecoder;
import io.netty.handler.codec.Delimiters;
import io.netty.handler.codec.string.StringDecoder;
import io.netty.handler.codec.string.StringEncoder;

/**
 * Created by user on 2016/3/16.
 */
public class MyServerInitializer extends ChannelInitializer<SocketChannel>{

    private static final StringEncoder ENCODER = new StringEncoder();
    private static final StringDecoder DECODER = new StringDecoder();
    private static final MyServerHandler HANDLER = new MyServerHandler();

    @Override
    protected void initChannel(SocketChannel ch) throws Exception {
        ch.pipeline().addLast("framer",new DelimiterBasedFrameDecoder(1024, Delimiters.lineDelimiter()))
                     .addLast("decoder",DECODER)
                     .addLast("encoder",ENCODER)
                     .addLast("handler",HANDLER);
    }


}
