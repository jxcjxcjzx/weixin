package weixin

import "encoding/xml"

// MsgType 消息类型
type MsgType string

// 微信支持的消息类型
const (
	MsgTypeText       MsgType = "text"       // 文本消息
	MsgTypeImage      MsgType = "image"      // 图片消息
	MsgTypeVoice      MsgType = "voice"      // 语音消息
	MsgTypeVideo      MsgType = "video"      // 视频消息
	MsgTypeShortVideo MsgType = "shortvideo" // 小视频消息
	MsgTypeLocation   MsgType = "location"   // 地理位置消息
	MsgTypeLink       MsgType = "link"       // 链接消息
	MsgTypeMusic      MsgType = "music"      // 音乐消息
	MsgTypeNews       MsgType = "news"       // 图文消息
	MsgTypeWXCard     MsgType = "wxcard"     // 卡券，客服消息时使用
)

// EncMessage 加密消息体
type EncMessage struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"-"` // 开发者微信号
	Encrypt      string   // 加密的消息报文
	MsgSignature string   // 报文签名
	TimeStamp    string   // 时间戳
	Nonce        string   // 随机数
}

// 接收普通消息

// Message 通用类，接收微信各类消息
type Message struct {
	Encrypt string

	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   // 开发者微信号
	FromUserName string   // 发送方帐号（一个OpenID）
	CreateTime   string   // 消息创建时间（整型）
	MsgId        int      // 消息id，64位整型

	// text-文本消息，image-图片消息，voice-语音消息，
	// video-视频消息，shortvideo-小视频消息，
	// location-地理位置消息，link-链接消息，
	// music-音乐，news-图文消息
	MsgType MsgType

	// text-文本消息
	Content string `xml:",omitempty"` // 文本消息内容

	// image-图片消息
	PicUrl  string `xml:",omitempty"` // 图片链接
	MediaId string `xml:",omitempty"` // 图片消息媒体id，可以调用多媒体文件下载接口拉取数据

	// voice-语音消息
	// MediaId string  `xml:",omitempty"`// 语音消息媒体id，可以调用多媒体文件下载接口拉取数据
	Format       string `xml:",omitempty"` // 语音格式，如amr，speex等
	Recongnition string `xml:",omitempty"` // 语音识别结果，使用UTF8编码

	// video-视频消息，shortvideo-小视频消息
	// MediaId string `xml:",omitempty"` // 视频消息媒体id，可以调用多媒体文件下载接口拉取数据
	ThumbMediaId string `xml:",omitempty"` // 视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据

	// location-地理位置消息
	LocationX float64 `xml:"Location_X,omitempty"` // 地理位置维度
	LocationY float64 `xml:"Location_Y,omitempty"` // 地理位置经度
	Scale     int     `xml:",omitempty"`           // 地图缩放大小
	Label     string  `xml:",omitempty"`           // 地理位置信息

	// link-链接消息
	Title       string `xml:",omitempty"` // 消息标题
	Description string `xml:",omitempty"` // 消息描述
	Url         string `xml:",omitempty"` // 消息链接
}
