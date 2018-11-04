package jarvisnode

import (
	"github.com/zhs007/jarviscore"
)

// Init - init node module
func Init(cfg jarviscore.Config) {
	jarviscore.InitJarvisCore(cfg)
}

// Release - release node module
func Release() {
	jarviscore.ReleaseJarvisCore()
}

// NewNode - new node
func NewNode(myinfo jarviscore.BaseInfo) jarviscore.JarvisNode {
	return jarviscore.NewNode(myinfo)
}
