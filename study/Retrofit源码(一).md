
```java

/**
* 如何使用
*/
public interface GitHub {
    @GET("/repos/{owner}/{repo}/contributors")
    Call<List<Contributor>> contributors(
        @Path("owner") String owner,
        @Path("repo") String repo);
  }

// 创建一个实例  是以代理的形式创建  github="retrofit2.Retrofit$1@5e25a92e"
//这个过程会把定义的converterFactory添加进来eg:{serializeNulls:falsefactories:[Factory[typeHierarchy=com.google.gson.JsonElement,adapter=com.google.gson.internal.bind.TypeAdapters$29@4df828d7], com.google.gson.internal.bind.ObjectTypeAdapter$1@b59d31, com.google.gson.internal.Excluder@62fdb4a6, Factory[type=java.lang.String,adapter=com.google.gson.internal.bind.TypeAdapters$17@11e21d0e], Factory[type=java.lang.Integer+int,adapter=com.google.gson.internal.bind.TypeAdapters$11@1dd02175], Factory[type=java.lang.Boolean+boolean,adapter=com.google.gson.internal.bind.TypeAdapters$7@31206beb], Factory[type=java.lang.Byte+byte,adapter=com.google.gson.internal.bind.TypeAdapters$9@3e77a1ed], Factory[type=java.lang.Short+short,adapter=com.google.gson.internal.bind.TypeAdapters$10@3ffcd140], Factory[type=java.lang.Long+long,adapter=com.google.gson.internal.bind.TypeAdapters$12@23bb8443], Factory[type=java.lang.Double+double,adapter=com.google.gson.Gson$3@1176dcec], Factory[type=java.lang.Float+float,adapter=com.google.gson.Gson$4@120d6fe6], Factory[type=java.lang.Number,adapter=com.google.gson.internal.bind.TypeAdapters$15@4ba2ca36], Factory[type=java.util.concurrent.atomic.AtomicInteger,adapter=com.google.gson.TypeAdapter$1@3444d69d], Factory[type=java.util.concurrent.atomic.AtomicBoolean,adapter=com.google.gson.TypeAdapter$1@1372ed45], Factory[type=java.util.concurrent.atomic.AtomicLong,adapter=com.google.gson.TypeAdapter$1@6a79c292], Factory[type=java.util.concurrent.atomic.AtomicLongArray,adapter=com.google.gson.TypeAdapter$1@37574691], Factory[type=java.util.concurrent.atomic.AtomicIntegerArray,adapter=com.google.gson.TypeAdapter$1@25359ed8], Factory[type=java.lang.Character+char,adapter=com.google.gson.internal.bind.TypeAdapters$16@21a947fe], Factory[type=java.lang.StringBuilder,adapter=com.google.gson.internal.bind.TypeAdapters$20@5606c0b], Factory[type=java.lang.StringBuffer,adapter=com.google.gson.internal.bind.TypeAdapters$21@80ec1f8], Factory[type=java.math.BigDecimal,adapter=com.google.gson.internal.bind.TypeAdapters$18@1445d7f], Factory[type=java.math.BigInteger,adapter=com.google.gson.internal.bind.TypeAdapters$19@6a396c1e], Factory[type=java.net.URL,adapter=com.google.gson.internal.bind.TypeAdapters$22@6c3f5566], Factory[type=java.net.URI,adapter=com.google.gson.internal.bind.TypeAdapters$23@12405818], Factory[type=java.util.UUID,adapter=com.google.gson.internal.bind.TypeAdapters$25@314c508a], Factory[type=java.util.Currency,adapter=com.google.gson.TypeAdapter$1@10b48321], Factory[type=java.util.Locale,adapter=com.google.gson.internal.bind.TypeAdapters$28@6b67034], Factory[typeHierarchy=java.net.InetAddress,adapter=com.google.gson.internal.bind.TypeAdapters$24@16267862], Factory[type=java.util.BitSet,adapter=com.google.gson.internal.bind.TypeAdapters$6@453da22c], com.google.gson.internal.bind.DateTypeAdapter$1@71248c21, Factory[type=java.util.Calendar+java.util.GregorianCalendar,adapter=com.google.gson.internal.bind.TypeAdapters$27@442675e1], com.google.gson.internal.bind.TimeTypeAdapter$1@6166e06f, com.google.gson.internal.bind.SqlDateTypeAdapter$1@49e202ad, com.google.gson.internal.bind.TypeAdapters$26@1c72da34, com.google.gson.internal.bind.ArrayTypeAdapter$1@6b0c2d26, Factory[type=java.lang.Class,adapter=com.google.gson.internal.bind.TypeAdapters$5@3d3fcdb0], com.google.gson.internal.bind.CollectionTypeAdapterFactory@641147d0, com.google.gson.internal.bind.MapTypeAdapterFactory@6e38921c, com.google.gson.internal.bind.JsonAdapterAnnotationTypeAdapterFactory@64d7f7e0, com.google.gson.internal.bind.TypeAdapters$30@27c6e487, com.google.gson.internal.bind.ReflectiveTypeAdapterFactory@49070868],instanceCreators:{}}
service="interface com.example.retrofit.SimpleService$GitHub"


GitHub github = retrofit.create(GitHub.class);

// 创建一个call实例
//响应类型为responseType="java.util.List<com.example.retrofit.SimpleService$Contributor>"
Call<List<Contributor>> call = github.contributors("square", "retrofit");

// 抓取并且得到数据 类型为Call的泛型类型
List<Contributor> contributors = call.execute().body();



@Override
void apply(RequestBuilder builder, T value) throws IOException {
      if (value == null) {
        throw new IllegalArgumentException(
            "Path parameter \"" + name + "\" value must not be null.");
      }
      builder.addPathParam(name, valueConverter.convert(value), encoded);
    }

@Override
public String convert(String value) throws IOException {
      return value;
    }
//eg: name="repo" value="retrofit"
void addPathParam(String name, String value, boolean encoded) {
    if (relativeUrl == null) {
      // The relative URL is cleared when the first query parameter is set.
      throw new AssertionError();
    }
    //在这里将relativeUrl中的{} 值转换为传入的值。此时relativeUrl: /repos/square/{repo}/contributors
    relativeUrl = relativeUrl.replace("{" + name + "}", canonicalizeForPath(value, encoded));
  }
/**
* 在 调用call.execute()   过程中 创建一个rawCall
*newCall(request) 会生成一个realCall的实例
* request类型为Request{method=GET, url=https://api.github.com/repos/square/retrofit/contributors, tag=null}
*/
private okhttp3.Call createRawCall() throws IOException {
    Request request = serviceMethod.toRequest(args);
    okhttp3.Call call = serviceMethod.callFactory.newCall(request);
    if (call == null) {
      throw new NullPointerException("Call.Factory returned null.");
    }
    return call;
  }

/**
* RealCall  调用  execute() 的过程
*/
  @Override public Response execute() throws IOException {
    synchronized (this) {
      if (executed) throw new IllegalStateException("Already Executed");
      executed = true;
    }
    try {
      client.dispatcher().executed(this);
      Response result = getResponseWithInterceptorChain(false);
      if (result == null) throw new IOException("Canceled");
      return result;
    } finally {
      client.dispatcher().finished(this);
    }
  }
 /**
   * Figures out what the response source will be, and opens a socket to that source if necessary.
   * Prepares the request headers and gets ready to start writing the request body if it exists.
   *
   * @throws RequestException if there was a problem with request setup. Unrecoverable.
   * @throws RouteException if the was a problem during connection via a specific route. Sometimes
   * recoverable. See {@link #recover}.
   * @throws IOException if there was a problem while making a request. Sometimes recoverable. See
   * {@link #recover (IOException)}.
   *发送请求
   */
  public void sendRequest() throws RequestException, RouteException, IOException { ...}
//最终走的是流
public HttpStream newStream(int connectTimeout, int readTimeout, int writeTimeout,
      boolean connectionRetryEnabled, boolean doExtensiveHealthChecks)
      throws RouteException, IOException {
    try {
      RealConnection resultConnection = findHealthyConnection(connectTimeout, readTimeout,
          writeTimeout, connectionRetryEnabled, doExtensiveHealthChecks);

      HttpStream resultStream;
      if (resultConnection.framedConnection != null) {
        resultStream = new Http2xStream(this, resultConnection.framedConnection);
      } else {
        resultConnection.socket().setSoTimeout(readTimeout);
        resultConnection.source.timeout().timeout(readTimeout, MILLISECONDS);
        resultConnection.sink.timeout().timeout(writeTimeout, MILLISECONDS);
        resultStream = new Http1xStream(this, resultConnection.source, resultConnection.sink);
      }

      synchronized (connectionPool) {
        stream = resultStream;
        return resultStream;
      }
    } catch (IOException e) {
      throw new RouteException(e) ;
    }
  }
/**
  * 读取响应
*/
private Response readNetworkResponse() throws IOException {
    httpStream.finishRequest();

    Response networkResponse = httpStream.readResponseHeaders()
        .request(networkRequest)
        .handshake(streamAllocation.connection().handshake())
        .header(OkHeaders.SENT_MILLIS, Long.toString(sentRequestMillis))
        .header(OkHeaders.RECEIVED_MILLIS, Long.toString(System.currentTimeMillis()))
        .build();

    if (!forWebSocket) {
      networkResponse = networkResponse.newBuilder()
          .body(httpStream.openResponseBody(networkResponse))
          .build();
    }

    if ("close".equalsIgnoreCase(networkResponse.request().header("Connection"))
        || "close".equalsIgnoreCase(networkResponse.header("Connection"))) {
      streamAllocation.noNewStreams();
    }

    return networkResponse;
  }
```

>  最后总结下：生成接口句柄，生成Call调用句柄，按照执行方式分别请求http，获取response响应报文，调用body自动调用json工具转换实体