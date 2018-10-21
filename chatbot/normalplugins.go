package chatbot

// normalPlugins - normal plugins
type normalPlugins struct {
}

// RegNormalPlugins - reg normal plugins
func RegNormalPlugins(mgr PluginsMgr) {
	Info("RegPlugins - normalPlugins")

	mgr.RegPlugins(&normalPlugins{})
}

// OnMessage - get message
func (p *normalPlugins) OnMessage(chatbot ChatBot, msg Message) (bool, error) {
	from := msg.GetFrom()
	if from == nil {
		return false, ErrMsgNoFrom
	}

	if from.IsMaster() {
		chatbot.SendMsg(from, "Yes, master.")
	} else {
		chatbot.SendMsg(from, "sorry, you are not my master.")
	}

	return true, nil
}
