# JarvisTeleBot Development Log

### 2019-04-23

增加了随机密码插件。  
一样需要在config.yaml里开启。

```
generatepassword -m normal -l 16
```

### 2019-04-22

前几天基本上都在尝试rasa，也做了一些训练，基本的基本上搞清楚了。  
现在最大的问题其实是我觉得准确率堪忧，首先就是ner的识别率不够高，前期缺少时间的识别，后来发现duckling就是干这事的，然后看到duckling上还有人问rasa和duckling结果不一样，我看rasa维护的duckling是4个月以前的，所以就自己做了duckling的服务。  
haskell编译很慢，不知道是不是服务器机器配置的问题，折腾了几次都失败了，后来干脆直接丢给Jarvis处理了，比我ssh连上去要顺利多了。

接下来，想用bert来做ner，缺数据。  

估计还是会分几步走，一方面训练现在的rasa，一方面用自己的想法来优化吧。

今天加入了duckling的plugin，可以先测一下duckling解析。  
还是加了duckling的配置文件，以及需要修改config.yaml，开启duckling插件。

```
duckling -l zh_CN -t "这个星期四要去看复联"
duckling -l en_GB -t "I am going to see the movie this Thursday."
```

### 2019-04-18

今天增加了``dtdata``，还是有单独的配置文件，并需要配置插件开关。  
这样就可以查数据了。  
目前有些bug，有些查询本地测试没问题，但放服务器会gateway504错误。  
换了台内存更大些的机器，另外一个就可以用了，不确定是否和机器内存有关。

```
getdtdata -m gamedatareport -s 2019-04-17 -e 2019-04-17
```

这个是目前最全的订阅新闻配置。

```
subscribearticles -t 300 -w=huxiu -w=baijingapp -w=tmtpost -w=36kr -w=geekpark -w=lieyunwang -w=techcrunch -w=techinasia -w=iheima -w=smzdm.post -w=smzdm.news
```

### 2019-04-17

昨天开始实际的部署rasa了，今天把最简单那个joke的项目试了一遍，觉得那个太简单了。  
然后和官网的rasa实际的聊了一下，觉得这个太基础了，于是又找了一圈，发现google、microsoft、facebook其实都有类似的服务，dialogflow、Luis 和 wit.ai，这些目前看来都有试用的版本，但看了下文档，其实用法都和rasa的差不多，都是要提供很多句子，然后分析intents。  
rasa甚至有他们的移植方案。  
目前还是就用rasa吧。  
rasa由于底层库的原因，其实对中文支持得并不好，所以前期也不打算直接用中文，先把rasa官方维护得比较好的英文支持掉吧，后面再来考虑多语言的问题。

### 2019-04-14

今天将``crawler``插件从命令行方式调整为service方式了，配置文件有修改。

```yaml
# crawler service addr
crawlerservaddr: 127.0.0.1:7051

# ankadb config
ankadb:
  # dbpath
  dbpath: './dat'
  # engine
  engine: 'leveldb'
  # httpaddr
  httpaddr: ''
```

然后``Message``有点小修改，支持了``markdown``。

前几天才知道其实我们现在做的``Jarvis``，其实有个名字叫``ChatOps``，``github``最早开始就做了，他们那个叫``Hubot``。  
现在``Jarvis``在``devops``这块功能其实已经差不多了，对于我们内部使用是完全没问题的，前后端各种工作流都可以很方便的挂接到``Jarvis``上，甚至国外算法同事的调试更新，去年年底也迁移到``Jarvis``上了。  
安全性方面，我们核心加密算法其实就是``BTC``的，现在有信任机制，所以应该不会有什么问题，对外发布的话，需要把文件系统隔离出来。  
现在除了``ChatOps``外，我们主要在做一组特别的bot，譬如confluence、oa、后台这些，希望是可以将各种环境通过bot彻底打通，而且不需要特殊的环境部署，这个概念其实前段时间，也看到阿里云提过，但他们更多的是对接现有文件流程，我们是希望通过``bot``在后端来实现，而且尽可能是通过常规人操作的方式来实现，譬如``headless chrome``、``ADB``这样的。  

很喜欢``chatbot``这种交互方式。

### 2019-04-05

今天加了``translate``插件，当然，首先需要在``config.yaml``里配置加载插件。

``` yaml
# plugins - enable plugins
plugins:
  - 'core'
  - 'assistant'
  - 'jarvisnode'
  - 'jarvisnodeex'
  - 'timestamp'
  - 'xlsx2json'
  - 'filetransfer'
  - 'usermgr'
  - 'userscript'
  - 'filetemplate'
  - 'crawler'
  - 'translate'
```

翻译插件直接用``jarviscrawlercore``的服务，也有自己的配置文件``translate.yaml``。

``` yaml
# jarvis telegram bot plugin translate config file

# translate service addr
translateservaddr: 127.0.0.1:7051
```


有几个命令

``` sh
# 进入翻译模式
translate -s zh-CN -d en -p google -r=true
# 退出翻译模式
translate -s zh-CN -d en -p google -r=false
```

这个翻译其实最主要是在聊天群组里用的，我觉得现在短句的聊天，其实google翻译已经可以很方便的让对话双方能看懂了，这个就是在聊天群组里，可以将国外同事的说话自动翻译成中文，而把自己的自动翻译成英文，还可以强制双向翻译（也就是中->英->中），这样你可以发现表述方式可能不太对，需要调整一下。  
对我来说，这个功能有了以后，能省掉不少时间。

### 2019-03-24

``exparticle``已经可以主动发文件给你了。

### 2019-03-20

今天把crawler控制插件完成了。  

crawler插件有自己的配置文件，例子如下：

``` yaml
# jarvis telegram bot plugin crawler config file

# crawler node addr
crawlernodeaddr: 12LyThj17Dj88EsHgVonn1eJffMSwjsXf4
# crawler path
crawlerpath: /mnt/jarviscrawlercore

# updatescript - the script will run when the update crawler was starting.  
#     There are some variables that can be used in this script.
#       - CrawlerPath: crawler path
updatescript: |
  cd {{.CrawlerPath}}
  git pull
  sh builddocker.sh

# exparticlescript - the script will run when the export article was starting.  
#     There are some variables that can be used in this script.
#       - CrawlerPath: crawler path
#       - URL: article url
#       - PDF: pdf filename
exparticlescript: |
  cd {{.CrawlerPath}}
  docker run --rm -v $PWD/output:/usr/src/app/output jarviscrawlercore node ./bin/jarviscrawler.js exparticle {{.URL}} -p ./output/{{.PDF}} -f A4
```

然后就是需要在``config.yaml``里配置加载插件。

``` yaml
# plugins - enable plugins
plugins:
  - 'core'
  - 'assistant'
  - 'jarvisnode'
  - 'jarvisnodeex'
  - 'timestamp'
  - 'xlsx2json'
  - 'filetransfer'
  - 'usermgr'
  - 'userscript'
  - 'filetemplate'
  - 'crawler'
```

新增了2个指令：

- ``updcrawler``: 更新crawler的，目前不需要任何参数。
- ``exparticle``: 导出article的，参数类似  
``exparticle -u http://www.baijingapp.com/article/22001 -p baijingapp.22001.pdf``  
现在在指令完成后，还需要手动拉取pdf文件，可以requestfile  
``requestfile -n blackbughk001 -f /mnt/jarviscrawlercore/output/baijingapp.22001.pdf``。

