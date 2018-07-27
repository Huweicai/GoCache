package config

type Conf struct {
	http HTTPConf `json:"http"`
	tcp TCPConf   `json:"tcp"`
}
type HTTPConf struct {
	Enable string `json:"enable"`
	port string `json:"port"`
}

type TCPConf struct {
	enable string `json:"enable"`
	port string `json:"port"`
}