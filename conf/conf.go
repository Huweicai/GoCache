package config

type Conf struct {
	Http HTTPConf `json:"http"`
	Tcp TCPConf   `json:"tcp"`
}
type HTTPConf struct {
	Enable string `json:"enable"`
	Port string `json:"port"`
}

type TCPConf struct {
	Enable string `json:"enable"`
	Port string `json:"port"`
}