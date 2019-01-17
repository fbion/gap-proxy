# 简介
gap-proxy 是一个加速网络的 SOCKS5 安全代理工具。

# mTCP
mTCP 是专门为 gap-proxy 设计的快速、可靠，加密安全、基于 UDP 的传输控制协议，借鉴了 KCP，TCP，Quic 协议的优点，技术特性：

* 支持连接迁移，在 IP 地址变化的情况下保持连接不断。
* 从协议头到应用数据都使用了 AES-256-CFB 加密。
* 类似 HMAC 机制，可防止数据被篡改。
* 选择确认，快速重传，快速超时选择重传。
* KeepAlive，及时释放意外终止的连接。
* 快速连接建立、连接释放。

mTCP 没有拥塞控制，只有流量控制。正因为如此，在高丢包率网络环境中，它才比使用了拥塞控制的 TCP 更快，即便是用了 BBR。

[mTCP 协议](./mtcp/protocol.txt)

# 支持平台
gap-proxy 仅支持 macOS, Linux, 其他类 Unix 理论上支持，但并未测试过。

# 安装

**提示：为了简单，如下操作在个人电脑，服务器上可都一样。**

#### 下载
根据所使用的操作系统从 [releases](https://github.com/fanpei91/gap-proxy/releases/) 下载相应已编译好的二进制文件压缩包。

#### 安装
首先解压程序包，进入该目录，把 `.gap-proxy` 目录移动到用户主目录下：

```
$ mv .gap-proxy ~/.gap-proxy
```

接着把 `gap-proxy` 二进制文件移动到 `/usr/local/bin` 或其它在 `$PATH` 环境变量的目录里：

```
$ mv gap-proxy /usr/local/bin/
```

最后编辑 `~/.gap-proxy/config.json` 文件，配置相关参数，举例：

```
{
	"local":  "127.0.0.1:1186", # 客户端监听地址 (SOCKS5)
	"server": "8.8.8.8:2286",   # 服务器端监听地址
	"key":    "gap-proxy"       # 密钥
}
```

# 基本使用
#### 启动
一旦安装并配置好后，在个人电脑系统上执行：

```
$ gap-proxy local start
```

然后在服务器上执行：

```
$ gap-proxy server start
```

这样便分别启动了 `gap-local` 和 `gap-server` 后台进程。

#### 代理
gap-proxy 只是个简单的 SOCKS5 代理，如果日常上网需要智能代理，可自己在浏览器安装代理管理插件，插件的 SOCKS5 地址填写为配置文件的 `local` 字段的地址。

# 高级使用
gap-proxy 的默认接收窗口为 256 个包，每个包最多能装 1420 字节数据，假设数据包传输每轮次需要 300 毫秒（RTT），一秒就有 3.3 个轮次，那么每秒最多能传输 256 * 1420 * 3.3 / 1024 = 1171 KB 有效数据。

默认值 256 基本上能流畅观看 720p 的 youtube 视频。如果看更高清的视频会卡顿，只要带宽允许，可把接收窗口值逐渐调大：

```
$ gap-proxy wnd 512
```

这样便把当前所有连接、新连接的接收窗口都设为 512 个包，跟 TCP 协议一样，`gap-server` 就能通过 `ack` 包的 `Window` 字段知道新发送窗口值。

现在，每秒最多能传输 2343 KB 有效数据。

# 待办事项
* FEC（前向纠错）
* [Proxifier](https://www.proxifier.com/) / [ProxyCap](http://www.proxycap.com/) 的核心功能，强制代理不支持 SOCKS5 协议的程序

# 感谢及参考
* [netstack](https://github.com/google/netstack)
* [kcp-go](https://github.com/xtaci/kcp-go)
* [shadowsockts-go](https://github.com/shadowsocks/shadowsocks-go)
* [go-daemon](https://github.com/sevlyar/go-daemon)
* [v2ray-core](https://github.com/v2ray/v2ray-core)
* [quic-go](https://github.com/lucas-clemente/quic-go)

# 许可
MIT