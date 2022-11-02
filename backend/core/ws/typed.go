package ws

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"time"
)

var ErrorJsonUnmarshalFail = errors.New("invalid json string")

type Params map[string]string

func (p Params) Get(key string) (string, bool) {
	v := p[key]
	return v, len(v) > 0
}

func (p Params) Parse(key string, lpObject interface{}) bool {
	v := p[key]
	pointer := reflect.ValueOf(lpObject)
	if pointer.Kind() != reflect.Pointer {
		return false
	}
	ll := pointer.Elem().Interface()
	switch ll.(type) {
	case string:
		{
			vlp := lpObject.(*string)
			*vlp = v
		}
	case bool:
		{
			vlp := lpObject.(*bool)
			*vlp = (v == "true")
		}
	case int8:
		{
			vlp := lpObject.(*int8)
			value, err := strconv.ParseInt(v, 10, 8)
			if err != nil {
				return false
			}
			*vlp = int8(value)
		}
	case int16:
		{
			vlp := lpObject.(*int16)
			value, err := strconv.ParseInt(v, 10, 16)
			if err != nil {
				return false
			}
			*vlp = int16(value)
		}
	case int32:
		{
			vlp := lpObject.(*int32)
			value, err := strconv.ParseInt(v, 10, 32)
			if err != nil {
				return false
			}
			*vlp = int32(value)
		}
	case int64:
		{
			vlp := lpObject.(*int64)
			value, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				return false
			}
			*vlp = int64(value)
		}
	case int:
		{
			vlp := lpObject.(*int)
			value, err := strconv.Atoi(v)
			if err != nil {
				return false
			}
			*vlp = value
		}
	case uint:
		{
			vlp := lpObject.(*uint)
			value, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return false
			}
			*vlp = uint(value)
		}
	case uint8:
		{
			vlp := lpObject.(*uint8)
			value, err := strconv.ParseUint(v, 10, 8)
			if err != nil {
				return false
			}
			*vlp = uint8(value)
		}
	case uint16:
		{
			vlp := lpObject.(*uint16)
			value, err := strconv.ParseUint(v, 10, 16)
			if err != nil {
				return false
			}
			*vlp = uint16(value)
		}
	case uint32:
		{
			vlp := lpObject.(*uint32)
			value, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return false
			}
			*vlp = uint32(value)
		}
	case uint64:
		{
			vlp := lpObject.(*uint64)
			value, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return false
			}
			*vlp = value
		}
	case time.Time:
		{
			var now time.Time
			vlp := lpObject.(*time.Time)
			loc, err := time.LoadLocation("Asia/Shanghai")
			if err == nil {
				now, err = time.ParseInLocation(`"2006-01-02 15:04:05"`, v, loc)
			}
			if err != nil {
				return false
			}
			*vlp = now
		}
	case float32:
		{
			vlp := lpObject.(*float32)
			value, err := strconv.ParseFloat(v, 32)
			if err != nil {
				return false
			}
			*vlp = float32(value)
		}
	case float64:
		{
			vlp := lpObject.(*float32)
			value, err := strconv.ParseFloat(v, 64)
			if err != nil {
				return false
			}
			*vlp = float32(value)
		}
	default:
		{
			return false
		}
	}
	return true
}

type AuthType int

const (
	// 不需要加入频道即可发送 POST SEND 消息
	Auth_Anonymous = AuthType(0)
	// POST消息 需要加入频道才可以发送
	Auth_PostNeedJoin = AuthType(1)
	// SEND消息 需要加入频道才可以发送
	Auth_SendNeedJoin = AuthType(2)
	// POST消息 和 SEND消息 需要加入频道才可以发送
	Auth_PostAndSendNeedJoin = Auth_PostNeedJoin | Auth_SendNeedJoin
)

type MessageType int

const (
	TextMessage = MessageType(1)

	// BinaryMessage denotes a binary data message.
	BinaryMessage = MessageType(2)

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = MessageType(8)

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = MessageType(9)

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = MessageType(10)
)

type RequestAction int8
type ResponseCode int8

const (
	// 操作成功
	Success = ResponseCode(0)
	// 操作失败
	Failure = ResponseCode(-1)
)

const (
	// 加入频道
	JoinChannel = RequestAction(100)
	// 离开频道
	LevelChannel = RequestAction(101)
	// 传输数据 仅发送 不需要考虑响应
	TransferPost = RequestAction(102)
	// 传输数据 发送后需要回应
	TransferSend = RequestAction(103)
)

type PayloadDomain string

type RequestMessage struct {
	// 动作
	Action RequestAction `json:"action"`
	// 消息属于哪个频道
	Channel string `json:"channel"`
	// 请求ID
	TraceId string `json:"traceId"`
	// 方法
	Method string `json:"method"`
	// 数据负载
	Payload PayloadDomain `json:"payload,omitempty"`
	// 是否为托管的， 由对象池生成的对象为托管对象，Free时会放回对象池内，非对象池生成的对象 Free时忽略
	managed bool `json:"-"`
}

// 返回一个回应结构体
func (r *RequestMessage) Response(code ResponseCode, message string) *ResponseMessage {
	if response, err := MallocResponseMessage(); err == nil {
		response.Code = code
		response.Message = message
		response.TraceId = r.TraceId
		return response
	}
	return &ResponseMessage{}
}

// 返回一个成功的回应消息体
func (r *RequestMessage) Success(message string) *ResponseMessage {
	if response, err := MallocResponseMessage(); err == nil {
		response.Code = Success
		response.Message = message
		response.TraceId = r.TraceId
		return response
	}
	return &ResponseMessage{}
}

// 返回一个失败的回应消息体
func (r *RequestMessage) Failure(message string) *ResponseMessage {
	if response, err := MallocResponseMessage(); err == nil {
		response.Code = Failure
		response.Message = message
		response.TraceId = r.TraceId
		return response
	}
	return &ResponseMessage{}
}

func (r *RequestMessage) Unmarshal(data []byte) error {
	if err := json.Unmarshal(data, r); err != nil {
		return err
	}
	// Name不能为空
	if r.Action == 0 {
		return ErrorJsonUnmarshalFail
	}
	if r.Channel == "" {
		return ErrorJsonUnmarshalFail
	}
	if r.TraceId == "" {
		return ErrorJsonUnmarshalFail
	}
	return nil
}

type ResponseMessage struct {
	// 响应代码
	Code ResponseCode `json:"code"`
	// 请求ID
	TraceId string `json:"traceId"`
	// 消息
	Message string `json:"msg"`
	// 响应数据
	Payload PayloadDomain `json:"payload,omitempty"`

	// 是否为托管的， 非托管对象不会被放入对象池中
	managed bool `json:"-"`
}
