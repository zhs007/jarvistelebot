package jarvisnode

import (
	"github.com/zhs007/jarviscore"
	"github.com/zhs007/jarvistelebot/basedef"
)

// Init - init node module
func Init(filename string) (*jarviscore.Config, error) {
	cfg, err := jarviscore.LoadConfig(filename)
	if err != nil {
		return nil, err
	}

	jarviscore.InitJarvisCore(cfg)

	// bi := &jarviscore.BaseInfo{
	// 	Name:     cfg.BaseNodeInfo.NodeName,
	// 	BindAddr: cfg.BaseNodeInfo.BindAddr,
	// 	ServAddr: cfg.BaseNodeInfo.ServAddr,
	// }

	return cfg, nil
}

// Release - release node module
func Release() {
	jarviscore.ReleaseJarvisCore()
}

// NewNode - new node
func NewNode(cfg *jarviscore.Config) (jarviscore.JarvisNode, error) {
	node, err := jarviscore.NewNode(cfg)
	if err != nil {
		return nil, err
	}

	node.SetNodeTypeInfo(basedef.JARVISNODETYPE, basedef.VERSION)

	return node, nil
}
