# 简介
gap-proxy 是一个加速网络的 SOCKS5 安全代理工具。

# mTCP
[mTCP](./mtcp/protocol.txt) 是专门为 gap-proxy 设计的快速、可靠、加密安全、基于 UDP 的传输控制协议，借鉴了 KCP、TCP、Quic 协议的优点。技术特性：

* 支持连接迁移，在 IP 地址变化的情况下保持连接不断
* 从协议头到应用数据都使用了 AES-256-CFB 加密
* 类似 HMAC 机制，可防止数据被篡改
* 选择确认、快速重传、快速超时选择重传
* KeepAlive，及时释放意外终止的连接
* 快速建立、释放连接

mTCP 没有拥塞控制，只有流量控制，因此在高丢包率网络环境中，比使用了拥塞控制的 TCP 更快。

# 支持平台
gap-proxy 仅支持 macOS, Linux, 其他类 Unix 理论上支持，但并未测试过。

# 安装

#### 下载
根据所使用的操作系统从 [releases](https://github.com/fanpei91/gap-proxy/releases/) 下载相应已编译好的 `gap-local` 和 `gap-server` 二进制文件压缩包。

# 基本使用

#### gap-local

```
$ gap-local --local-addr "127.0.0.1:1086" --server-addr "8.8.8.8:1086" --key "key"
```


#### gap-server

```
$ gap-server --server-addr "8.8.8.8:1086" --key "key"
```

完毕！

#### 代理
gap-local 只是个简单的 SOCKS5 代理，如果日常上网需要智能代理，可自己在浏览器安装代理管理插件，插件的 SOCKS5 地址填写为配置文件的 `--local-addr` 参数的地址。


# 感谢及参考
* [netstack](https://github.com/google/netstack)
* [kcp-go](https://github.com/xtaci/kcp-go)
* [shadowsockts-go](https://github.com/shadowsocks/shadowsocks-go)
* [go-daemon](https://github.com/sevlyar/go-daemon)
* [v2ray-core](https://github.com/v2ray/v2ray-core)
* [quic-go](https://github.com/lucas-clemente/quic-go)

# 许可证
MIT