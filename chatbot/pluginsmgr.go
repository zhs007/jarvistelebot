package chatbot

import (
	"context"
	"strings"
)

// PluginsMgr - chat bot plugins interface
type PluginsMgr interface {
	// RegPlugin - reg plugin
	RegPlugin(plugin Plugin) error
	// OnMessage
	OnMessage(ctx context.Context, bot ChatBot, msg Message) error
	// GetComeInPlugin
	GetComeInPlugin(code string) Plugin
	// ComeInPlugin
	ComeInPlugin(plugin Plugin)
	// GetCurPlugin
	GetCurPlugin() Plugin
	// OnStart - on start
	OnStart(ctx context.Context) error
}

// NewPluginsMgr - new default plugins mgr
func NewPluginsMgr() PluginsMgr {
	return &pluginsMgr{
		plugins:   make([]Plugin, 0, 16),
		lstComeIn: make([]Plugin, 0, 16),
	}
}

// PluginsMgr - chat bot plugins
type pluginsMgr struct {
	plugins   []Plugin
	lstComeIn []Plugin
	curPlugin Plugin
}

func (mgr *pluginsMgr) GetComeInPlugin(code string) Plugin {
	for _, v := range mgr.lstComeIn {
		if v.GetComeInCode() == code {
			return v
		}
	}

	return nil
}

func (mgr *pluginsMgr) RegPlugin(plugin Plugin) error {
	mgr.plugins = append(mgr.plugins, plugin)

	if plugin.GetComeInCode() != "" {
		mgr.lstComeIn = append(mgr.lstComeIn, plugin)
	}

	return nil
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

	for _, v := range mgr.plugins {
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
	for _, v := range mgr.plugins {
		go v.OnStart(ctx)
	}

	return nil
}
