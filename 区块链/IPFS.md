# IPFS与HTTP的区别

```
IPFS协议栈分：
身份层；（描述加密数据）kad算法生成节点身份信息，节点身份信息用DHT分布式哈希表的结构，互相关联。

网络层；路由层；交换层；（数据传输）libp2p支持多种传输协议；bt做数据分发交换。

对象层；命名层；文件层；（数据结构）文克尔树列数据结构，Git控制版本迭代IPFS网络找文件方式被称作“文件在哪里”，直接输入文件hash地址，对应节点服务器提供带宽资源。

目前互联网：TCP/IP
1、应用层：HTTP 响应请求的

2、传输层：TCP,UDP...运输车

3、网络层：IP 目的地

4、物理链路层 大马路HTTP网络找文件方式被称作“哪里有文件”，会出现输入一个地址，对应的服务器没有我们要的文件，就要重新输入新网址，

重新找总结：IPFS协议替代TCP/IP应用层的HTTP，底下的传输层，网络层，物理链路层可以不变，只是用户发起数据请求，做响应用的不在是HTTP，而是IPFS，两种方式：HTTP输入的是地址，IPFS输入的是文件数据加密之后的hash地址，一个是文件在哪，另一个是哪有文件。
```



# 使用IPFS是免费的？为什么还需要Filecoin？

```
当前，我们使用IPFS网络是免费的。但是请注意，添加到IPFS网络的数据不会自动复制到其他节点，换句话说，IPFS不提供可用性保证。而Filecoin建立在IPFS之上，是一个持久层，为用户数据提供可用性保证。

另外，IPFS不像区块链那样工作，也不使用区块链。IPFS是一个去中心化的P2P和内容可寻址的存储和交付网络。而Filecoin是一个区块链项目，可以与IPFS共同成为存储区块链数据的理想选择。

总的来说，IPFS不包括鼓励为其他人存储数据的内置机制。这是Filecoin希望解决的挑战。Filecoin网络为长期存储创建了一个分布式存储市场。存储容量大的节点可以向用户出租存储空间，并从中获得收益。

Filecoin可以被视为一个冷存储层，非常适合安全地存储大量数据。IPFS将是热存储层，旨在快速检索和分发内容。
```

