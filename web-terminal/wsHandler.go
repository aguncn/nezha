package main

import (
	"encoding/json"
	"fmt"

	//"io/ioutil"
	"log"
	"net/http"

	//"time"

	"github.com/gorilla/websocket"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

const (
	// EndOfTransmission end
	EndOfTransmission = "\u0004"
)

// token认证json返回值
type tokenResponse struct {
	Code uint16 `json:"code"` // 非200表示认证失败
	Msg  string `json:"msg"`  // 一点小说明
}

// web终端发来的包
type xtermMessage struct {
	MsgType string `json:"type"`  // 类型:resize客户端调整终端, input客户端输入
	Input   string `json:"input"` // msgtype=input情况下使用
	Rows    uint16 `json:"rows"`  // msgtype=resize情况下使用
	Cols    uint16 `json:"cols"`  // msgtype=resize情况下使用
}

// ssh流式处理器
type streamHandler struct {
	wsConn     *websocket.Conn
	resizeChan chan remotecommand.TerminalSize
	doneChan   chan struct{}
}

// executor回调获取web是否resize
func (handler *streamHandler) Next() (size *remotecommand.TerminalSize) {
	select {
	case size := <-handler.resizeChan:
		return &size
	case <-handler.doneChan:
		return nil
	}
}

func (handler *streamHandler) Close() error {
	return handler.wsConn.Close()
}

// Done done, must call Done() before connection close, or Next() would not exits.
func (handler *streamHandler) Done() {
	close(handler.doneChan)
}

// executor回调读取web端的输入
func (handler *streamHandler) Read(p []byte) (size int, err error) {

	_, message, err := handler.wsConn.ReadMessage()
	if err != nil {
		log.Printf("read message err: %v", err)
		return copy(p, EndOfTransmission), err
	}

	xmsg := xtermMessage{}
	if err := json.Unmarshal(message, &xmsg); err != nil {
		log.Println("json.Unmarshal err: ", err)
		return copy(p, EndOfTransmission), err
	}

	switch xmsg.MsgType {
	case "input":
		return copy(p, xmsg.Input), nil
	case "resize":
		handler.resizeChan <- remotecommand.TerminalSize{Width: xmsg.Cols, Height: xmsg.Rows}
		return 0, nil
	default:
		log.Printf("unknown message type '%s'", xmsg.MsgType)
		// return 0, nil
		return copy(p, EndOfTransmission), fmt.Errorf("unknown message type '%s'", xmsg.MsgType)
	}
}

// executor回调向web端输出
func (handler *streamHandler) Write(p []byte) (size int, err error) {
	var copyData []byte

	// 产生副本
	copyData = make([]byte, len(p))
	copy(copyData, p)
	size = len(p)
	if err := handler.wsConn.WriteMessage(websocket.TextMessage, copyData); err != nil {
		log.Printf("write message err: %v", err)
		return 0, err
	}
	return
}

func wsHandler(resp http.ResponseWriter, req *http.Request) {
	var (
		wsConn        *websocket.Conn
		sshReq        *rest.Request
		podName       string
		podNs         string
		containerName string
		executor      remotecommand.Executor
		handler       *streamHandler
		err           error
	)

	// 解析GET参数
	if err = req.ParseForm(); err != nil {
		return
	}

	//userToken = req.Form.Get("userToken")
	podNs = req.Form.Get("namespace")
	podName = req.Form.Get("pod")
	containerName = req.Form.Get("container")

	//accessUrl := fmt.Sprintf("%s?token=%s", *apiUrl, userToken)
	//timeout := time.Duration(3 * time.Second)
	//client := http.Client{
	//	Timeout: timeout,
	//}

	//respToken, err := client.Get(accessUrl)
	//if respToken != nil {
	//	defer respToken.Body.Close()
	//}

	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//body, err := ioutil.ReadAll(respToken.Body)
	//if err != nil {
	//	panic(err.Error())
	//}

	//var tokenJson = tokenResponse{}
	//json.Unmarshal(body, &tokenJson)

	//if tokenJson.Code != 200 {
	//h.conn.WriteMessage(websocket.TextMessage, []byte("亲，你无权进入此容器。bye~bye~"))
	//h.conn.Close()
	//	return
	//}

	// 得到websocket长连接
	upgrader := websocket.Upgrader{
		// 允许跨域,这个回调函数需要特别指定下。
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	if wsConn, err = upgrader.Upgrade(resp, req, nil); err != nil {
		log.Fatalf("error creating ws conn: %v", err)
	}

	// URL长相:
	// https://172.18.11.25:6443/api/v1/namespaces/default/pods/nginx-deployment-5cbd8757f-d5qvx/exec?command=sh&container=nginx&stderr=true&stdin=true&stdout=true&tty=true
	sshReq = clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(podNs).
		SubResource("exec").
		VersionedParams(&v1.PodExecOptions{
			Container: containerName,
			Command:   []string{"sh"},
			Stdin:     true,
			Stdout:    true,
			Stderr:    true,
			TTY:       true,
		}, scheme.ParameterCodec)

	// 创建到容器的连接
	if executor, err = remotecommand.NewSPDYExecutor(kconfig, "POST", sshReq.URL()); err != nil {
		log.Fatalf("error creating spdy executor: %v", err)
	}

	// 配置与容器之间的数据流处理回调
	handler = &streamHandler{
		wsConn:     wsConn,
		resizeChan: make(chan remotecommand.TerminalSize),
		doneChan:   make(chan struct{})}

	err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             handler,
		Stdout:            handler,
		Stderr:            handler,
		TerminalSizeQueue: handler,
		Tty:               true,
	})

	if err != nil {
		log.Println("error executor: ", err)
		handler.wsConn.WriteMessage(websocket.TextMessage, []byte("亲，你选择的容器不存在！bye~bye~"))
		handler.wsConn.Close()
	}

	return
}
