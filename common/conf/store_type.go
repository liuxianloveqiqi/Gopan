package conf

// 存储类型(表示文件存到哪里)
type StoreType int64

const (
	_ StoreType = iota
	// StoreLocal : 节点本地
	StoreLocal
	// StoreCeph : Ceph集群
	StoreMinio
	// StoreCOS : 腾讯云COS
	StoreCOS
	// StoreMix : 混合(Ceph及OSS)
	StoreMix
	// StoreAll : 所有类型的存储都存一份数据
	StoreAll
)
