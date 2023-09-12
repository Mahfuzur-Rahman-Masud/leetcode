package org.masud;

/**
 * @author Mahfuzur Rahman
 */
import io.netty.bootstrap.ServerBootstrap;
import io.netty.channel.*;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.SocketChannel;
import io.netty.channel.socket.nio.NioServerSocketChannel;
import io.netty.handler.codec.http.HttpObjectAggregator;
import io.netty.handler.codec.http.HttpServerCodec;

public class NettyServer {
  private static final HelloServerHandler helloServerHandler = new HelloServerHandler();

  public static void main(String[] args) throws Exception {
    EventLoopGroup selectors = new NioEventLoopGroup();
    EventLoopGroup workers = new NioEventLoopGroup();

    try {
      ServerBootstrap bootstrap = new ServerBootstrap();
      bootstrap.group(selectors, workers)
        .channel(NioServerSocketChannel.class)
        .childHandler(new ChannelInitializer<SocketChannel>() {
          @Override
          public void initChannel(SocketChannel ch) throws Exception {
            ch.pipeline()
              .addLast(new HttpServerCodec())
              .addLast(new HttpObjectAggregator(1024*256))
              .addLast(helloServerHandler);
          }
        });

      ChannelFuture future = bootstrap.bind(5557).sync();
      System.out.println("Server started and listening on port 5557");

      future.channel().closeFuture().sync();
    } finally {
      selectors.shutdownGracefully();
      workers.shutdownGracefully();
    }
  }
}
