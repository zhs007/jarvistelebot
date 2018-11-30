protoc -I chatbotdb/proto/ chatbotdb/proto/chatbotdb.proto --go_out=plugins=grpc:chatbotdb/proto
protoc -I assistantdb/proto/ assistantdb/proto/assistant.proto --go_out=plugins=grpc:assistantdb/proto
protoc -I plugins/core/proto/ plugins/core/proto/core.proto --go_out=plugins=grpc:plugins/core/proto
protoc -I plugins/jarvisnode/proto/ plugins/jarvisnode/proto/jarvisnode.proto --go_out=plugins=grpc:plugins/jarvisnode/proto