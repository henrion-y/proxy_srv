#### 目录结构

```bash 
├─lib                               公共库
│  ├─helper                         返回参数格式化辅助工具
│  └─middleware                     中间件
│      └─cors
├─pool_services                     代理站点集合
│  └─colly_kuaidaili_proxy          快代理(demo)
├─proxy_host                        代理池
├─router                            web服务
│  └─proxy_pool
│      ├─api                        api接口
│      └─form                       请求、返回参数
└─set_proxy                         设置代理
```

#### 项目介绍
> 本项目是用gocolly编写的一个爬虫demo。
>用户可仿照 colly_kuaidaili_proxy 自行添加其他代理到站点的代理。
>如果是单个点使用， 可直接调用 set_proxy 中的方法设置代理。
>如果是多个点使用， 则可运行web服务， 通过api的方式调用获取代理。
