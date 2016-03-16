package com.sung.netty.practice.client;

import java.util.concurrent.ThreadFactory;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * Created by user on 2016/3/16.
 */
public class NamedThreadFactory implements ThreadFactory{

    private static final AtomicInteger POOL_SEQ = new AtomicInteger(1);
    private final AtomicInteger mThreadNum = new AtomicInteger(1);
    private final String mPrefix;
    private final boolean mDaemon;
    private final ThreadGroup mGroup;

    public NamedThreadFactory(String mPrefix, boolean mDaemon) {
        this.mPrefix = mPrefix + "-thread-" ;
        this.mDaemon = mDaemon;
        SecurityManager s = System.getSecurityManager();
        mGroup = (s == null) ? Thread.currentThread().getThreadGroup() : s.getThreadGroup();
    }

    public NamedThreadFactory() {
        this("pool-" + POOL_SEQ.getAndIncrement(),false);
    }

    public NamedThreadFactory(String mPrefix) {
        this(mPrefix,false);
    }

    @Override
    public Thread newThread(Runnable r) {
        String name = mPrefix + mThreadNum.getAndIncrement();
        Thread ret = new Thread(mGroup,r,name,0);
        ret.setDaemon(true);
        return ret;
    }

    public ThreadGroup getThreadGroup() {
        return mGroup;
    }
}
