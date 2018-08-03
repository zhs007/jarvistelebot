package jarviscore

import (
	pb "github.com/zhs007/jarvistelebot/proto"
)

type NodeInfo struct {
	Name     string
	ServAddr string
	Token    string
	NodeType pb.NODETYPE
}
