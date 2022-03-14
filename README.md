# Go Ethereum

## CHANGELOG

### v1.0.1 2021-12-28
* [FIX] 修复以太坊事件解码异常问题 （ref: https://github.com/ethereum/go-ethereum/pull/23573）

### v1.0.0 2021-12-24

* [FEATURE] ethclient: 支持 `ethclient.WithConsistentHashKey` 方法，配合网关实现负载均衡的一致性哈希。
* [FEATURE] ethclient: 支持 `ethclient.WithHTTPHeader` 方法，在调用 JSON RPC 方法时传入自定义HTTP请求头。