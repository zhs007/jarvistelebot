package jarvisnode

import (
	"github.com/zhs007/jarviscore"
	pb "github.com/zhs007/jarviscore/proto"
)

// Init - init node module
func Init(filename string) (*jarviscore.BaseInfo, error) {
	cfg, err := jarviscore.LoadConfig(filename)
	if err != nil {
		return nil, err
	}

	jarviscore.InitJarvisCore(*cfg)

	bi := &jarviscore.BaseInfo{
		Name:     cfg.BaseNodeInfo.NodeName,
		BindAddr: cfg.BaseNodeInfo.BindAddr,
		ServAddr: cfg.BaseNodeInfo.ServAddr,
		NodeType: pb.NODETYPE_SH,
	}

	return bi, nil
}

// Release - release node module
func Release() {
	jarviscore.ReleaseJarvisCore()
}

// NewNode - new node
func NewNode(myinfo *jarviscore.BaseInfo) jarviscore.JarvisNode {
	return jarviscore.NewNode(*myinfo)
}
