package assistantdb

import (
	pb "github.com/zhs007/jarvistelebot/assistantdb/proto"
)

// ResultAssistantData -
type ResultAssistantData struct {
	AssistantData struct {
		MaxMsgID int64    `json:"maxMsgID"`
		Keys     []string `json:"keys"`
	} `json:"assistantData"`
}

// ResultUpdMsg -
type ResultUpdMsg struct {
	UpdMsg struct {
		MsgID      int64    `json:"msgID"`
		Data       string   `json:"data"`
		Keys       []string `json:"keys"`
		CreateTime int64    `json:"createTime"`
		UpdateTime int64    `json:"updateTime"`
	} `json:"updMsg"`
}

// ResultUpdAssistantData -
type ResultUpdAssistantData struct {
	UpdAssistantData struct {
		MaxMsgID int64    `json:"maxMsgID"`
		Keys     []string `json:"keys"`
	} `json:"updAssistantData"`
}

// ResultMsg -
type ResultMsg struct {
	Msg struct {
		MsgID      int64    `json:"msgID"`
		Data       string   `json:"data"`
		Keys       []string `json:"keys"`
		CreateTime int64    `json:"createTime"`
		UpdateTime int64    `json:"updateTime"`
	} `json:"msg"`
}

// ResultAssistantData2AssistantData - ResultAssistantData -> AssistantData
func ResultAssistantData2AssistantData(result *ResultAssistantData) *pb.AssistantData {
	dat := &pb.AssistantData{
		MaxMsgID: result.AssistantData.MaxMsgID,
		Keys:     result.AssistantData.Keys,
	}

	return dat
}

// ResultUpdAssistantData2AssistantData - ResultUpdAssistantData -> AssistantData
func ResultUpdAssistantData2AssistantData(result *ResultUpdAssistantData) *pb.AssistantData {
	dat := &pb.AssistantData{
		MaxMsgID: result.UpdAssistantData.MaxMsgID,
		Keys:     result.UpdAssistantData.Keys,
	}

	return dat
}

// ResultUpdMsg2Msg - ResultUpdMsg -> Message
func ResultUpdMsg2Msg(result *ResultUpdMsg) *pb.Message {
	msg := &pb.Message{
		MsgID:      result.UpdMsg.MsgID,
		Data:       result.UpdMsg.Data,
		Keys:       result.UpdMsg.Keys,
		CreateTime: result.UpdMsg.CreateTime,
		UpdateTime: result.UpdMsg.UpdateTime,
	}

	return msg
}

// ResultMsg2Msg - ResultMsg -> Message
func ResultMsg2Msg(result *ResultMsg) *pb.Message {
	msg := &pb.Message{
		MsgID:      result.Msg.MsgID,
		Data:       result.Msg.Data,
		Keys:       result.Msg.Keys,
		CreateTime: result.Msg.CreateTime,
		UpdateTime: result.Msg.UpdateTime,
	}

	return msg
}
