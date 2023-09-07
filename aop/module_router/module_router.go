package module_router

import (
	"github.com/Allen9012/AllenServer/network/protocol/gen/messageID"
	"github.com/Allen9012/AllenServer/protos/gen/module"
)

type ModuleMessageHandler func(messageId uint64, data []byte)

var (
	Module2MessageId2Handler map[module.Module]map[messageID.MessageId]ModuleMessageHandler
)

func RegisterModuleMessageHandler(moduleId module.Module, msgId messageID.MessageId, handler ModuleMessageHandler) {
	if Module2MessageId2Handler == nil {
		Module2MessageId2Handler = make(map[module.Module]map[messageID.MessageId]ModuleMessageHandler)
	}
	if Module2MessageId2Handler[moduleId] == nil {
		Module2MessageId2Handler[moduleId] = make(map[messageID.MessageId]ModuleMessageHandler)
	}
	if Module2MessageId2Handler[moduleId][msgId] != nil {
		panic("[RegisterModuleMessageHandler] repeated register")
	}
	Module2MessageId2Handler[moduleId][msgId] = handler
}

func GetModuleHandler(moduleId module.Module, messageId messageID.MessageId) ModuleMessageHandler {
	message2Handler, ok := Module2MessageId2Handler[moduleId]
	if !ok {
		return nil
	}
	handler, exist := message2Handler[messageId]
	if !exist {
		return nil
	}
	return handler
}
