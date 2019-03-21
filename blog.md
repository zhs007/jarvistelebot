# JarvisTeleBot Development Log

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

