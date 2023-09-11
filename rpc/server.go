package rpc

import (
	"fmt"
	"github.com/Allen9012/AllenGame/log"
	"github.com/Allen9012/AllenGame/network"
	"math"
	"runtime"
	"strings"
)

/**
  Copyright © 2023 github.com/Allen9012 All rights reserved.
  @author: Allen
  @since: 2023/9/8
  @desc:
  @modified by:
**/

type RpcProcessorType uint8

const (
	RpcProcessorJson   RpcProcessorType = 0
	RpcProcessorGoGoPB RpcProcessorType = 1
)

var arrayProcessor = []IRpcProcessor{&JsonProcessor{}, &PBProcessor{}}
var arrayProcessorLen uint8 = 2
var LittleEndian bool

type Server struct {
	functions       map[interface{}]interface{}
	rpcHandleFinder RpcHandleFinder
	rpcServer       *network.TCPServer

	compressBytesLen int
}

/*	Implement Agent  */
type RpcAgent struct {
	conn      network.Conn
	rpcServer *Server
	userData  interface{}
}

func (server *Server) NewAgent(c *network.TCPConn) network.Agent {
	agent := &RpcAgent{conn: c, rpcServer: server}

	return agent
}

func (server *Server) Init(rpcHandleFinder RpcHandleFinder) {
	server.rpcHandleFinder = rpcHandleFinder
	server.rpcServer = &network.TCPServer{}
}

func (server *Server) Start(listenAddr string, maxRpcParamLen uint32, compressBytesLen int) {
	splitAddr := strings.Split(listenAddr, ":")
	if len(splitAddr) != 2 {
		log.Fatal("listen addr is failed", log.String("listenAddr", listenAddr))
	}

	server.rpcServer.Addr = ":" + splitAddr[1]
	server.rpcServer.MinMsgLen = 2
	server.compressBytesLen = compressBytesLen
	if maxRpcParamLen > 0 {
		server.rpcServer.MaxMsgLen = maxRpcParamLen
	} else {
		server.rpcServer.MaxMsgLen = math.MaxUint32
	}

	server.rpcServer.MaxConnNum = 100000
	server.rpcServer.PendingWriteNum = 2000000
	server.rpcServer.NewAgent = server.NewAgent
	server.rpcServer.LittleEndian = LittleEndian
	server.rpcServer.WriteDeadline = Default_ReadWriteDeadline
	server.rpcServer.ReadDeadline = Default_ReadWriteDeadline
	server.rpcServer.LenMsgLen = DefaultRpcLenMsgLen

	server.rpcServer.Start()
}

/*	Implement Agent  */

func (agent *RpcAgent) Run() {
	defer func() {
		if r := recover(); r != nil {
			buf := make([]byte, 4096)
			l := runtime.Stack(buf, false)
			errString := fmt.Sprint(r)
			log.Dump(string(buf[:l]), log.String("error", errString))
		}
	}()

	for {
		data, err := agent.conn.ReadMsg()
		if err != nil {
			log.Error("read message is error", log.String("remoteAddress", agent.conn.RemoteAddr().String()), log.ErrorAttr("error", err))
			//will close tcpconn
			break
		}

		bCompress := (data[0] >> 7) > 0
		processor := GetProcessor(data[0] & 0x7f)
		if processor == nil {
			agent.conn.ReleaseReadMsg(data)
			log.Warning("cannot find processor", log.String("RemoteAddr", agent.conn.RemoteAddr().String()))
			return
		}

		//解析head
		var compressBuff []byte
		byteData := data[1:]
		if bCompress == true {
			var unCompressErr error

			compressBuff, unCompressErr = compressor.UncompressBlock(byteData)
			if unCompressErr != nil {
				agent.conn.ReleaseReadMsg(data)
				log.Error("UncompressBlock failed", log.String("RemoteAddr", agent.conn.RemoteAddr().String()), log.ErrorAttr("error", unCompressErr))
				return
			}
			byteData = compressBuff
		}

		req := MakeRpcRequest(processor, 0, 0, "", false, nil)
		err = processor.Unmarshal(byteData, req.RpcRequestData)
		if cap(compressBuff) > 0 {
			compressor.UnCompressBufferCollection(compressBuff)
		}
		agent.conn.ReleaseReadMsg(data)
		if err != nil {
			log.Error("Unmarshal failed", log.String("RemoteAddr", agent.conn.RemoteAddr().String()), log.ErrorAttr("error", err))
			if req.RpcRequestData.GetSeq() > 0 {
				rpcError := RpcError(err.Error())
				if req.RpcRequestData.IsNoReply() == false {
					agent.WriteResponse(processor, req.RpcRequestData.GetServiceMethod(), req.RpcRequestData.GetSeq(), nil, rpcError)
				}
				ReleaseRpcRequest(req)
				continue
			} else {
				ReleaseRpcRequest(req)
				break
			}
		}

		//交给程序处理
		serviceMethod := strings.Split(req.RpcRequestData.GetServiceMethod(), ".")
		if len(serviceMethod) < 1 {
			rpcError := RpcError("rpc request req.ServiceMethod is error")
			if req.RpcRequestData.IsNoReply() == false {
				agent.WriteResponse(processor, req.RpcRequestData.GetServiceMethod(), req.RpcRequestData.GetSeq(), nil, rpcError)
			}
			ReleaseRpcRequest(req)
			log.Error("rpc request req.ServiceMethod is error")
			continue
		}

		rpcHandler := agent.rpcServer.rpcHandleFinder.FindRpcHandler(serviceMethod[0])
		if rpcHandler == nil {
			rpcError := RpcError(fmt.Sprintf("service method %s not config!", req.RpcRequestData.GetServiceMethod()))
			if req.RpcRequestData.IsNoReply() == false {
				agent.WriteResponse(processor, req.RpcRequestData.GetServiceMethod(), req.RpcRequestData.GetSeq(), nil, rpcError)
			}
			log.Error("serviceMethod not config", log.String("serviceMethod", req.RpcRequestData.GetServiceMethod()))
			ReleaseRpcRequest(req)
			continue
		}

		if req.RpcRequestData.IsNoReply() == false {
			req.requestHandle = func(Returns interface{}, Err RpcError) {
				agent.WriteResponse(processor, req.RpcRequestData.GetServiceMethod(), req.RpcRequestData.GetSeq(), Returns, Err)
				ReleaseRpcRequest(req)
			}
		}

		req.inParam, err = rpcHandler.UnmarshalInParam(req.rpcProcessor, req.RpcRequestData.GetServiceMethod(), req.RpcRequestData.GetRpcMethodId(), req.RpcRequestData.GetInParam())
		if err != nil {
			rErr := "Call Rpc " + req.RpcRequestData.GetServiceMethod() + " Param error " + err.Error()
			log.Error("call rpc param error", log.String("serviceMethod", req.RpcRequestData.GetServiceMethod()), log.ErrorAttr("error", err))
			if req.requestHandle != nil {
				req.requestHandle(nil, RpcError(rErr))
			} else {
				ReleaseRpcRequest(req)
			}

			continue
		}

		err = rpcHandler.PushRpcRequest(req)
		if err != nil {
			rpcError := RpcError(err.Error())

			if req.RpcRequestData.IsNoReply() {
				agent.WriteResponse(processor, req.RpcRequestData.GetServiceMethod(), req.RpcRequestData.GetSeq(), nil, rpcError)
			}

			ReleaseRpcRequest(req)
		}
	}
}

func (agent *RpcAgent) OnClose() {
}

func GetProcessorType(param interface{}) (RpcProcessorType, IRpcProcessor) {
	for i := uint8(1); i < arrayProcessorLen; i++ {
		if arrayProcessor[i].IsParse(param) == true {
			return RpcProcessorType(i), arrayProcessor[i]
		}
	}

	return RpcProcessorJson, arrayProcessor[RpcProcessorJson]
}

func GetProcessor(processorType uint8) IRpcProcessor {
	if processorType >= arrayProcessorLen {
		return nil
	}
	return arrayProcessor[processorType]
}
