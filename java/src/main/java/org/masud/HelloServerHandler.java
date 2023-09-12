package org.masud;

/**
 * @author Mahfuzur Rahman
 */

import io.netty.channel.ChannelHandler.Sharable;
import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;
import io.netty.handler.codec.http.DefaultFullHttpResponse;
import io.netty.handler.codec.http.FullHttpRequest;
import io.netty.handler.codec.http.FullHttpResponse;
import io.netty.handler.codec.http.HttpResponseStatus;
import io.netty.handler.codec.http.HttpVersion;

@Sharable
public class HelloServerHandler extends SimpleChannelInboundHandler<FullHttpRequest> {


  @Override
  protected void channelRead0(ChannelHandlerContext ctx, FullHttpRequest req) throws Exception {
    String uri = req.uri();

    if ("/".equals(uri)) {
      FullHttpResponse response = new DefaultFullHttpResponse(HttpVersion.HTTP_1_1, HttpResponseStatus.OK);
      response.content().writeBytes("Hello, Java Netty!".getBytes());
      ctx.write(response);
      ctx.flush();
      ctx.close();
    }
  }
//  @Override
//  public void channelRead(ChannelHandlerContext ctx, Object msg) {
//    ByteBuf buffer = (ByteBuf) msg;
//    String message = buffer.toString(io.netty.util.CharsetUtil.US_ASCII);
//    buffer.release();
//
//    if ("/".equals(message)) {
//      ByteBuf response = ctx.alloc().buffer();
//      response.writeBytes("Hello, Java Netty!".getBytes());
//      ctx.write(response);
//      ctx.flush();
//    }
//  }

  @Override
  public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
    cause.printStackTrace();
    ctx.close();
  }


}

