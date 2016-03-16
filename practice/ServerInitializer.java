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
public class ServerInitializer extends ChannelInitializer<SocketChannel> {


    private static final StringEncoder ENCODER = new StringEncoder();
    private static final StringDecoder DECODER = new StringDecoder();
    private static final ServerHandler HANDLER = new ServerHandler();
    @Override
    protected void initChannel(SocketChannel socketChannel) throws Exception {
        socketChannel.pipeline().addLast("framer",new DelimiterBasedFrameDecoder(9182, Delimiters.lineDelimiter()))
                .addLast("decoder",DECODER)
                .addLast("encoder",ENCODER)
                .addLast("handler",HANDLER);
    }
}
