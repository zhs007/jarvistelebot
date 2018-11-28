package chatbot

import (
	"context"
	"strings"
)

// FuncNewPlugin - func newPlugin(cfgPath string) (Plugin, error)
type FuncNewPlugin func(cfgPath string) (Plugin, error)

const (
	// PluginTypeNormal - normal plugin
	PluginTypeNormal = 0
	// PluginTypeCommand - command plugin
	PluginTypeCommand = 1
)

// PluginsMgr - chat bot plugins interface
type PluginsMgr interface {
	// SetDefaultPlugin - set default plugin
	SetDefaultPlugin(pluginName string) error
	// NewPlugin - New a plugin
	NewPlugin(pluginName string) error
	// OnMessage
	OnMessage(ctx context.Context, bot ChatBot, msg Message) error
	// GetCurPlugin
	GetCurPlugin() Plugin
	// OnStart - on start
	OnStart(ctx context.Context) error
	// RegPlugin - Registered a new plugin
	RegPlugin(pluginName string, funcNewPlugin FuncNewPlugin) error
}

// NewPluginsMgr - new default plugins mgr
func NewPluginsMgr(cfgPath string) PluginsMgr {
	return &pluginsMgr{
		lstNormal:  make([]Plugin, 0, 16),
		lstCommand: make([]Plugin, 0, 16),
		lstComeIn:  make([]Plugin, 0, 16),
		mapPlugin:  make(map[string]FuncNewPlugin),
		cfgPath:    cfgPath,
	}
}

// PluginsMgr - chat bot plugins
type pluginsMgr struct {
	lstNormal     []Plugin
	lstCommand    []Plugin
	lstComeIn     []Plugin
	curPlugin     Plugin
	mapPlugin     map[string]FuncNewPlugin
	cfgPath       string
	defaultPlugin Plugin
}

// NewPlugin - New a plugin
func (mgr *pluginsMgr) NewPlugin(pluginName string) error {
	funcNewPlugin, ok := mgr.mapPlugin[pluginName]
	if !ok {
		return ErrNoPluginName
	}

	plugin, err := funcNewPlugin(mgr.cfgPath)
	if err != nil {
		return err
	}

	pt := plugin.GetPluginType()
	if pt == PluginTypeNormal {
		mgr.lstNormal = append(mgr.lstNormal, plugin)

		return nil
	} else if pt == PluginTypeCommand {
		mgr.lstCommand = append(mgr.lstCommand, plugin)

		return nil
	}

	return ErrInvalidPluginType
}

func (mgr *pluginsMgr) OnMessage(ctx context.Context, bot ChatBot, msg Message) error {
	params := &MessageParams{
		ChatBot:    bot,
		MgrPlugins: mgr,
		Msg:        msg,
		LstStr:     strings.Fields(msg.GetText()),
	}

	if mgr.curPlugin != nil {
		r, err := mgr.curPlugin.OnMessage(ctx, params)
		if err != nil {
			return err
		}

		// if this plugins process current message
		// then break
		if r {
			return nil
		}
	}

	if mgr.isCommand(params) {
		var cp []Plugin
		for _, v := range mgr.lstCommand {
			if v.IsMyMessage(params) {
				cp = append(cp, v)
			}
		}

		if len(cp) > 0 {
			r, err := cp[0].OnMessage(ctx, params)
			if err != nil {
				return err
			}

			if r {
				return nil
			}
		}
	}

	for _, v := range mgr.lstNormal {
		r, err := v.OnMessage(ctx, params)
		if err != nil {
			return err
		}

		// if this plugins process current message
		// then break
		if r {
			return nil
		}
	}

	r, err := mgr.defaultPlugin.OnMessage(ctx, params)
	if err != nil {
		return err
	}

	if r {
		return nil
	}

	return ErrPluginsEmpty
}

// ComeInPlugin
func (mgr *pluginsMgr) ComeInPlugin(plugin Plugin) {
	mgr.curPlugin = plugin
}

// GetCurPlugin
func (mgr *pluginsMgr) GetCurPlugin() Plugin {
	return mgr.curPlugin
}

// OnStart - on start
func (mgr *pluginsMgr) OnStart(ctx context.Context) error {
	for _, v := range mgr.lstCommand {
		go v.OnStart(ctx)
	}

	for _, v := range mgr.lstNormal {
		go v.OnStart(ctx)
	}

	return nil
}

// OnStart - on start
func (mgr *pluginsMgr) isCommand(params *MessageParams) bool {
	if params.Msg.GetFile() != nil {
		return true
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		return true
	}

	return false
}

// RegPlugin - Registered a new plugin
func (mgr *pluginsMgr) RegPlugin(pluginName string, funcNewPlugin FuncNewPlugin) error {
	_, ok := mgr.mapPlugin[pluginName]
	if ok {
		return ErrSamePluginName
	}

	mgr.mapPlugin[pluginName] = funcNewPlugin

	return nil
}

// SetDefaultPlugin - set default plugin
func (mgr *pluginsMgr) SetDefaultPlugin(pluginName string) error {
	funcNewPlugin, ok := mgr.mapPlugin[pluginName]
	if !ok {
		return ErrNoPluginName
	}

	plugin, err := funcNewPlugin(mgr.cfgPath)
	if err != nil {
		return err
	}

	mgr.defaultPlugin = plugin

	return nil
}
