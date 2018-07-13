package c2s_proto

import (
	"fmt"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"

	// 使用binary协议，因此匿名引用这个包，底层会自动注册
	_ "github.com/davyxu/cellnet/codec/binary"
	_ "github.com/davyxu/cellnet/util"
	"reflect"
)

///////cs通信
type HttpEchoREQ struct {
	UserName string
}

type HttpEchoACK struct {
	Token  string
	Status int32
}


// 用于消息日志打印消息内容
//client2login/login2client/client2game/game2client
func (self *HttpEchoREQ) String() string { return fmt.Sprintf("%+v", *self) }
func (self *HttpEchoACK) String() string { return fmt.Sprintf("%+v", *self) }

// 引用消息时，自动注册消息，这个文件可以由代码生成自动生成
func init() {

	cellnet.RegisterHttpMeta(&cellnet.HttpMeta{
		Path:         "/",
		Method:       "POST",
		RequestCodec: codec.MustGetCodec("httpjson"),
		RequestType:  reflect.TypeOf((*HttpEchoREQ)(nil)).Elem(),

		// 请求方约束
		ResponseCodec: codec.MustGetCodec("httpjson"),
		ResponseType:  reflect.TypeOf((*HttpEchoACK)(nil)).Elem(),
	})
}
