package com.sung.netty.practice;

import io.netty.channel.ChannelInitializer;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.DelimiterBasedFrameDecoder;
import io.netty.handler.codec.Delimiters;
import io.netty.handler.codec.string.StringDecoder;
import io.netty.handler.codec.string.StringEncoder;

/**
 * Created by user on 2016/3/15.
 */
public class ClientInitialize extends ChannelInitializer<SocketChannel>{

    private static final StringEncoder ENCODER = new StringEncoder();
    private static final StringDecoder DECODER = new StringDecoder();
    private static final ClientHandler HANDLER = new ClientHandler();


    @Override
    protected void initChannel(SocketChannel ch) throws Exception {
        ch.pipeline().addLast("framer",new DelimiterBasedFrameDecoder(8192, Delimiters.lineDelimiter()))
                     .addLast("encoder",ENCODER)
                     .addLast("decoder",DECODER)
                     .addLast("handler",HANDLER);

    }
}
