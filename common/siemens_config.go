package common

/*
*
* 西门子PLC的一些配置
*
 */
type S1200Config struct {
	Host          string       `json:"host" validate:"required" title:"IP地址" info:""`          // 127.0.0.1
	Port          *int         `json:"port" validate:"required" title:"端口号" info:""`           // 0
	Rack          *int         `json:"rack" validate:"required" title:"架号" info:""`            // 0
	Slot          *int         `json:"slot" validate:"required" title:"槽号" info:""`            // 1
	Model         string       `json:"model" validate:"required" title:"型号" info:""`           // s7-200 s7 1500
	Timeout       *int         `json:"timeout" validate:"required" title:"连接超时时间" info:""`     // 5s
	IdleTimeout   *int         `json:"idleTimeout" validate:"required" title:"心跳超时时间" info:""` // 5s
	ReadFrequency *int         `json:"readFrequency" validate:"required" title:"采集频率" info:""` // 5s
	Blocks        []S1200Block `json:"blocks" validate:"required" title:"采集配置" info:""`        // Db
}
type S1200Block struct {
	Tag     string `json:"tag"`     // 数据tag
	Address int    `json:"address"` // 地址
	Start   int    `json:"start"`   // 起始地址
	Size    int    `json:"size"`    // 数据长度
}
type S1200BlockValue struct {
	Tag     string `json:"tag"`     // 数据tag
	Address int    `json:"address"` // 地址
	Start   int    `json:"start"`   // 起始地址
	Size    int    `json:"size"`    // 数据长度
	Value   []byte `json:"value"`
}
