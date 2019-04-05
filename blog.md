# JarvisTeleBot Development Log

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
translate -s zh-CN -d en -p google -r true
# 退出翻译模式
translate -s zh-CN -d en -p google -r false
```

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

