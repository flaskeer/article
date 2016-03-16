package com.sung.netty.practice.client;

import com.alibaba.dubbo.common.utils.NetUtils;
import io.netty.bootstrap.Bootstrap;
import io.netty.channel.*;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetSocketAddress;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.ScheduledThreadPoolExecutor;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * Created by user on 2016/3/16.
 */
public class NettyClient {

    private static final Logger logger = LoggerFactory.getLogger(NettyClient.class);
    private String host;
    private int port;
    private int timeout = 1000;
    private int connectionTimeout = 3000;

    private final EventLoopGroup group = new NioEventLoopGroup();
    private ClientChannelInitializer initializer;
    private Bootstrap bootstrap;
    private volatile Channel channel;
    private volatile ChannelFuture channelFuture;
    private volatile ScheduledFuture<?> reconnectExecutorFuture = null;
    private long lastConnectionTime = System.currentTimeMillis();
    private final AtomicInteger reconnect_count = new AtomicInteger(0);
    private final AtomicBoolean reconnect_error_log_flag = new AtomicBoolean(false);
    private final int reconnect_warning_period = 1800;
    private final long shutdown_timeout = 1000 * 60 * 15;
    private static final ScheduledThreadPoolExecutor reconnectService = new ScheduledThreadPoolExecutor(2,new NamedThreadFactory("ClientReconnectTimer",true));

    public NettyClient(String host, int port, ClientChannelInitializer initializer) throws Exception {
        this.host = host;
        this.port = port;
        this.initializer = initializer;
        try{
            doOpen();
        }catch (Throwable t) {
            close();
            throw new Exception("Failed to start " + getClass().getSimpleName() + " " + NetUtils.getLocalAddress()  + " connect to the server " + host + ", cause: " + t.getMessage(), t);
        }
        try{
            connect();
            logger.info("Start " + getClass().getSimpleName() + " " + NetUtils.getLocalAddress() + " connect to the server " + host);
        }catch (Throwable t){
            throw new Exception("Failed to start " + getClass().getSimpleName() + " " + NetUtils.getLocalAddress()  + " connect to the server " + host + ", cause: " + t.getMessage(), t);
        }
    }

    private void doOpen() throws Throwable{
        bootstrap = new Bootstrap();
        bootstrap.option(ChannelOption.TCP_NODELAY,true)
                 .option(ChannelOption.SO_KEEPALIVE,true);
        bootstrap.group(group)
                 .channel(NioSocketChannel.class)
                 .handler(initializer);
    }

    private void connect() throws Throwable{
        try{
            if(isConnect()) {
                return;
            }
            initConnectStatusCheckCommand();
            doConnection();
            if(!isConnect()) {
                throw new Exception("Failed connect to server " + getRemoteAddress() + " from " + getClass().getSimpleName() + " "
                        + NetUtils.getLocalHost() + ", cause: Connect wait timeout: " + getTimeout() + "ms.");
            }else{
                logger.info("Successed connect to server " + getRemoteAddress() + " from " + getClass().getSimpleName() + " "
                        + NetUtils.getLocalHost() + ", channel is " + this.channel);
            }
            reconnect_count.set(0);
            reconnect_error_log_flag.set(false);
        }catch (Throwable t) {
            logger.error("fail to connect ");
        }

    }

    private void doConnection() throws Exception {
        long start = System.currentTimeMillis();
        channelFuture = bootstrap.connect(getConnectAddress());
        try{
            boolean ret = channelFuture.awaitUninterruptibly(getConnectionTimeout(), TimeUnit.MILLISECONDS);
            if(ret && channelFuture.isSuccess()) {
                Channel newChannel = channelFuture.channel();
                try{
                    Channel oldChannel = this.channel;
                    if(oldChannel != null){
                        logger.info("closed old channel " + oldChannel + " on create new netty channel " + newChannel);
                        oldChannel.close();
                    }
                }finally {
                    this.channel = newChannel;
                }
            }else if(channelFuture.cause() != null){
                throw new Exception("client failed to connect server " + getRemoteAddress() + ",error message is " + channelFuture.cause().getMessage(),channelFuture.cause());
            }else{
                throw new Exception("client fail to connect to server " + getRemoteAddress() + " client side timeout " + getTimeout() +"ms(elapsed:" + (System.currentTimeMillis() - start) + "ms) from netty client" + NetUtils.getLocalHost());
            }
        }finally {
            if(!isConnect()) {
                channelFuture.cancel(true);
            }
        }
    }

    private void initConnectStatusCheckCommand() {
        if(reconnectExecutorFuture == null || reconnectExecutorFuture.isCancelled()){
            Runnable connectStatusCheckCommand = new Runnable() {
                @Override
                public void run() {
                    try {

                        if (!isConnect()) {
                            connect();
                        }else{
                            lastConnectionTime = System.currentTimeMillis();
                        }
                    }catch (Throwable e) {
                        String errorMsg = "client reconnect to " + getRemoteAddress() + " find error. ";
                        if(System.currentTimeMillis() - lastConnectionTime > shutdown_timeout) {
                            if(!reconnect_error_log_flag.get()){
                                reconnect_error_log_flag.set(true);
                                logger.error(errorMsg,e);
                                return;
                            }
                        }
                        if(reconnect_count.getAndIncrement() % reconnect_warning_period == 0) {
                            logger.warn(errorMsg,e);
                        }
                    }
                }
            };
            reconnectExecutorFuture = reconnectService.scheduleWithFixedDelay(connectStatusCheckCommand,2*1000,2*1000, TimeUnit.MILLISECONDS);
        }
    }


    public String receiveMessage() {
        return initializer.getClientHandler().getMessage();
    }

    public void sendMessage(String message) {
        channel.writeAndFlush(message);
    }

    public boolean isConnect() {
        if(channel == null) {
            return false;
        }
        return channel.isActive();
    }

    public void close() {
        destoryConnectStatusCheckCommand();
        try {
            if(channel != null) {
                channel.close();
            }
        }catch (Throwable t) {
            logger.warn(t.getMessage(),t);
        }
        try {
            group.shutdownGracefully();
        }catch (Throwable t) {
            logger.warn(t.getMessage());
        }
    }

    private synchronized void destoryConnectStatusCheckCommand() {
        try {
            if(reconnectExecutorFuture != null || !reconnectExecutorFuture.isDone()) {
                reconnectExecutorFuture.cancel(true);
                reconnectService.purge();
            }
        }catch (Throwable t) {
            logger.error(t.getMessage(),t);
        }
    }

    private InetSocketAddress getConnectAddress() {
        return new InetSocketAddress(host,port);
    }

    private String getRemoteAddress() {
        return host + ":" + port;
    }

    public int getTimeout() {
        return timeout;
    }

    public int getConnectionTimeout() {
        return connectionTimeout;
    }
}
