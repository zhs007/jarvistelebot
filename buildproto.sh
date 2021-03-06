protoc -I chatbot/proto/ chatbot/proto/cmdline.proto --go_out=plugins=grpc:chatbot/proto
protoc -I chatbotdb/proto/ chatbotdb/proto/chatbotdb.proto --go_out=plugins=grpc:chatbotdb/proto
protoc -I assistantdb/proto/ assistantdb/proto/assistant.proto --go_out=plugins=grpc:assistantdb/proto
protoc -I plugins/core/proto/ plugins/core/proto/core.proto --go_out=plugins=grpc:plugins/core/proto
protoc -I plugins/jarvisnode/proto/ plugins/jarvisnode/proto/jarvisnode.proto --go_out=plugins=grpc:plugins/jarvisnode/proto
protoc -I plugins/jarvisnodeex/proto/ plugins/jarvisnodeex/proto/jarvisnodeex.proto --go_out=plugins=grpc:plugins/jarvisnodeex/proto
protoc -I plugins/usermgr/proto/ plugins/usermgr/proto/usermgr.proto --go_out=plugins=grpc:plugins/usermgr/proto
protoc -I plugins/userscript/proto/ plugins/userscript/proto/userscript.proto --go_out=plugins=grpc:plugins/userscript/proto
protoc -I plugins/filetemplate/proto/ plugins/filetemplate/proto/filetemplate.proto --go_out=plugins=grpc:plugins/filetemplate/proto
protoc -I plugins/assistant/proto/ plugins/assistant/proto/assistant.proto --go_out=plugins=grpc:plugins/assistant/proto
protoc -I plugins/notekeyword/proto/ plugins/notekeyword/proto/notekeyword.proto --go_out=plugins=grpc:plugins/notekeyword/proto
protoc -I plugins/crawler/proto/ plugins/crawler/proto/crawler.proto --go_out=plugins=grpc:plugins/crawler/proto
protoc -I plugins/translate/proto/ plugins/translate/proto/translate.proto --go_out=plugins=grpc:plugins/translate/proto
protoc -I plugins/dtdata/proto/ plugins/dtdata/proto/dtdata.proto --go_out=plugins=grpc:plugins/dtdata/proto
protoc -I plugins/duckling/proto/ plugins/duckling/proto/duckling.proto --go_out=plugins=grpc:plugins/duckling/proto
protoc -I plugins/generatepwd/proto/ plugins/generatepwd/proto/generatepwd.proto --go_out=plugins=grpc:plugins/generatepwd/proto
protoc -I plugins/dtdata2/proto/ plugins/dtdata2/proto/dtdata2.proto --go_out=plugins=grpc:plugins/dtdata2/proto
protoc -I jarviscrawlercore/ jarviscrawlercore/result.proto --go_out=plugins=grpc:jarviscrawlercore