package chatbot

import (
	"context"
	"strings"

	"github.com/golang/protobuf/proto"
)

// FuncNewPlugin - func newPlugin(cfgPath string) (Plugin, error)
type FuncNewPlugin func(cfgPath string) (Plugin, error)

const (
	// PluginTypeNormal - normal plugin
	PluginTypeNormal = 0
	// PluginTypeCommand - command plugin
	PluginTypeCommand = 1
	// PluginTypeWritableCommand - writable command plugin
	PluginTypeWritableCommand = 2
)

// PluginsMgr - chat bot plugins interface
type PluginsMgr interface {
	// SetDefaultPlugin - set default plugin
	SetDefaultPlugin(pluginName string) error
	// NewPlugin - New a plugin
	NewPlugin(pluginName string) error

	// OnMessage
	OnMessage(ctx context.Context, bot ChatBot, msg Message) error
	// SetCurPlugin
	SetCurPlugin(plugin Plugin)
	// GetCurPlugin
	GetCurPlugin() Plugin
	// OnStart - on start
	OnStart(ctx context.Context) error

	// RegPlugin - Registered a new plugin
	RegPlugin(pluginName string, funcNewPlugin FuncNewPlugin) error
	// CanNewPlugin - can new this plugin
	CanNewPlugin(pluginName string) bool

	// HasPlugin - has a plugin
	HasPlugin(pluginName string) bool
	// FindPlugin - find a plugin
	FindPlugin(pluginName string) Plugin

	// GetPlugins - get plugins
	GetPlugins() []string
}

// NewPluginsMgr - new default plugins mgr
func NewPluginsMgr(cfgPath string) PluginsMgr {
	return &pluginsMgr{
		lstNormal:          make([]Plugin, 0, 16),
		lstCommand:         make([]Plugin, 0, 16),
		lstWritableCommand: make([]Plugin, 0, 16),
		mapFuncNewPlugin:   make(map[string]FuncNewPlugin),
		mapPlugin:          make(map[string]Plugin),
		cfgPath:            cfgPath,
	}
}

type pluginParsingInfo struct {
	plugin  Plugin
	cmdline proto.Message
}

// PluginsMgr - chat bot plugins
type pluginsMgr struct {
	lstNormal          []Plugin
	lstCommand         []Plugin
	lstWritableCommand []Plugin
	curPlugin          Plugin
	mapFuncNewPlugin   map[string]FuncNewPlugin
	mapPlugin          map[string]Plugin
	cfgPath            string
	defaultPlugin      Plugin
}

// NewPlugin - New a plugin
func (mgr *pluginsMgr) NewPlugin(pluginName string) error {
	if mgr.HasPlugin(pluginName) {
		return ErrRepeatPlugins
	}

	funcNewPlugin, ok := mgr.mapFuncNewPlugin[pluginName]
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
	} else if pt == PluginTypeWritableCommand {
		mgr.lstWritableCommand = append(mgr.lstWritableCommand, plugin)

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
		cmdline, _ := mgr.curPlugin.ParseMessage(params)

		params.CurPlugin = mgr.curPlugin
		params.CommandLine = cmdline

		r, err := mgr.curPlugin.OnMessage(ctx, params)
		if err != nil {
			mgr.curPlugin = nil

			return err
		}

		// if this plugins process current message
		// then break
		if r {
			return nil
		}
	}

	var lstppi []*pluginParsingInfo

	if mgr.isWritableCommand(params) {
		for _, v := range mgr.lstWritableCommand {
			cmdline, err := v.ParseMessage(params)
			if err == nil {
				lstppi = append(lstppi, &pluginParsingInfo{
					plugin:  v,
					cmdline: cmdline,
				})
			}
		}
	}

	if mgr.isCommand(params) {
		for _, v := range mgr.lstCommand {
			cmdline, err := v.ParseMessage(params)
			if err == nil {
				lstppi = append(lstppi, &pluginParsingInfo{
					plugin:  v,
					cmdline: cmdline,
				})
			}
		}
	}

	for _, v := range mgr.lstNormal {
		cmdline, err := v.ParseMessage(params)
		if err == nil {
			lstppi = append(lstppi, &pluginParsingInfo{
				plugin:  v,
				cmdline: cmdline,
			})
		}
	}

	if len(lstppi) > 0 {
		params.CurPlugin = lstppi[0].plugin
		params.CommandLine = lstppi[0].cmdline
		r, err := lstppi[0].plugin.OnMessage(ctx, params)
		if err != nil {
			return err
		}

		if r {
			return nil
		}
	}

	params.CurPlugin = mgr.defaultPlugin
	r, err := mgr.defaultPlugin.OnMessage(ctx, params)
	if err != nil {
		return err
	}

	if r {
		return nil
	}

	return ErrPluginsEmpty
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

// isCommand - is command
func (mgr *pluginsMgr) isCommand(params *MessageParams) bool {
	if params.Msg.GetFile() != nil {
		return true
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">" {
		return true
	}

	return false
}

// isWritableCommand - is writable command
func (mgr *pluginsMgr) isWritableCommand(params *MessageParams) bool {
	if params.Msg.GetFile() != nil {
		return true
	}

	if len(params.LstStr) > 1 && params.LstStr[0] == ">>" {
		return true
	}

	return false
}

// RegPlugin - Registered a new plugin
func (mgr *pluginsMgr) RegPlugin(pluginName string, funcNewPlugin FuncNewPlugin) error {
	_, ok := mgr.mapFuncNewPlugin[pluginName]
	if ok {
		return ErrSamePluginName
	}

	mgr.mapFuncNewPlugin[pluginName] = funcNewPlugin

	return nil
}

// SetDefaultPlugin - set default plugin
func (mgr *pluginsMgr) SetDefaultPlugin(pluginName string) error {
	if mgr.HasPlugin(pluginName) {
		return ErrRepeatPlugins
	}

	funcNewPlugin, ok := mgr.mapFuncNewPlugin[pluginName]
	if !ok {
		return ErrNoPluginName
	}

	plugin, err := funcNewPlugin(mgr.cfgPath)
	if err != nil {
		return err
	}

	mgr.defaultPlugin = plugin
	mgr.mapPlugin[pluginName] = plugin

	return nil
}

// CanNewPlugin - can new this plugin
func (mgr *pluginsMgr) CanNewPlugin(pluginName string) bool {
	_, ok := mgr.mapFuncNewPlugin[pluginName]
	return ok
}

// HasPlugin - has a plugin
func (mgr *pluginsMgr) HasPlugin(pluginName string) bool {
	_, ok := mgr.mapPlugin[pluginName]
	return ok
}

// FindPlugin - find a plugin
func (mgr *pluginsMgr) FindPlugin(pluginName string) Plugin {
	plugin, ok := mgr.mapPlugin[pluginName]
	if ok {
		return plugin
	}

	return nil
}

// GetPlugins - get plugins
func (mgr *pluginsMgr) GetPlugins() []string {
	var lst []string

	for _, v := range mgr.lstWritableCommand {
		lst = append(lst, v.GetPluginName())
	}

	for _, v := range mgr.lstCommand {
		lst = append(lst, v.GetPluginName())
	}

	for _, v := range mgr.lstNormal {
		lst = append(lst, v.GetPluginName())
	}

	return lst
}

// SetCurPlugin
func (mgr *pluginsMgr) SetCurPlugin(plugin Plugin) {
	mgr.curPlugin = plugin
}
