package com.sung.netty.practice.client;

import io.netty.channel.Channel;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.socket.SocketChannel;
import io.netty.handler.codec.DelimiterBasedFrameDecoder;
import io.netty.handler.codec.Delimiters;
import io.netty.handler.codec.string.StringDecoder;
import io.netty.handler.codec.string.StringEncoder;
import io.netty.util.CharsetUtil;

/**
 * Created by user on 2016/3/16.
 */
public class ClientChannelInitializer extends ChannelInitializer<SocketChannel> {

    private static final StringDecoder DECODER = new StringDecoder(CharsetUtil.UTF_8);
    private static final StringEncoder ENCODER = new StringEncoder(CharsetUtil.UTF_8);
    private static final NettyClientHandler HANDLER = new NettyClientHandler();

    @Override
    protected void initChannel(SocketChannel ch) throws Exception {
        ch.pipeline().addLast("framer",new DelimiterBasedFrameDecoder(1024,Delimiters.lineDelimiter()))
                     .addLast("decoder",DECODER)
                     .addLast("encoder",ENCODER)
                     .addLast("handler",HANDLER);
    }

    public NettyClientHandler getClientHandler() {
        return HANDLER;
    }
}
