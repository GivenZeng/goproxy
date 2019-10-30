一个使用go实现的反向代理服务器.

尚在开放中，还未提供cache control功能...

## config
见etc_sample

## build & run
```
make run
```

## archetecture


## TODO
- 缓存设置大小上线，到达上限则开始清理
- 缓存支持层级缓存
- 支持更精细的缓存设置
- 当backend故障时，可以选择返回已过期的缓存数据，提供额外的服务容错能力
- 提供类似nginx proxy_cache_xxxx的缓存控制