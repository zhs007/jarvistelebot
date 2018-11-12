protoc -I chatbotdb/proto/ chatbotdb/proto/chatbotdb.proto --go_out=plugins=grpc:chatbotdb/proto
protoc -I assistantdb/proto/ assistantdb/proto/assistant.proto --go_out=plugins=grpc:assistantdb/proto