一个使用go实现的反向代理服务器

## config
```
# port默认为为1088
port: 1088
# 并发量，默认5000
paralle: 10000

# 本地文件服务器路径
root: "file path"

# 缓存目录，默认如下
cache_path: /var/goproxy/cache
# 缓存过期时间如下
cache_expiration: 3m

# 路由设置
locations:
  - 
    path: "/a/b"
    proxy_pass: "http://server.com"
    # 缓存过期时间如下
    cache_expiration: 3m
```

## archetecture


## TODO
- 缓存设置过期时间
- 缓存设置大小上线，到达上限则开始清理
- 缓存支持层级缓存
- 支持更精细的缓存设置
- 当backend故障时，可以选择返回已过期的缓存数据，提供额外的服务容错能力
- 提供类似nginx proxy_cache_xxxx的缓存控制