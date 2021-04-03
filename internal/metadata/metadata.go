package metadata

import (
	"os"

	"github.com/prologic/bitcask/internal"
)

type MetaData struct {
	// 所以是否是最新的。一般两种情况会写入索引，一种是一个wal文件写满了，关闭时会保存一下索引。此时不是最新，因为还有一个新的数据的索引没有保存
	// 还有一种情况是当数据库关闭时，也会保存索引。此时也是最新的
	IndexUpToDate    bool  `json:"index_up_to_date"`
	// 统计可以重复使用的空间，当删除、或者插入已经存在的一个key时，都会进行统计
	ReclaimableSpace int64 `json:"reclaimable_space"`
}

func (m *MetaData) Save(path string, mode os.FileMode) error {
	return internal.SaveJsonToFile(m, path, mode)
}

func Load(path string) (*MetaData, error) {
	var m MetaData
	err := internal.LoadFromJsonFile(path, &m)
	return &m, err
}
