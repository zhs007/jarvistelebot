package chatbot

// PluginsMgr - chat bot plugins interface
type PluginsMgr interface {
	// RegPlugins - reg plugins
	RegPlugins(plugins Plugins) error
	// OnMessage
	OnMessage(chatbot ChatBot, msg Message) error
}

// NewPluginsMgr - new default plugins mgr
func NewPluginsMgr() PluginsMgr {
	return &pluginsMgr{
		plugins: make([]Plugins, 0, 16),
	}
}

// PluginsMgr - chat bot plugins
type pluginsMgr struct {
	plugins []Plugins
}

func (mgr *pluginsMgr) RegPlugins(plugins Plugins) error {
	return nil
}

func (mgr *pluginsMgr) OnMessage(chatbot ChatBot, msg Message) error {
	for _, v := range mgr.plugins {
		r, err := v.OnMessage(chatbot, msg)
		if err != nil {
			return err
		}

		// if this plugins process current message
		// then break
		if r {
			return nil
		}
	}

	return ErrPluginsEmpty
}
