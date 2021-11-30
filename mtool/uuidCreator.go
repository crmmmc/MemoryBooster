package mtool

import (
	uuid "github.com/satori/go.uuid"
	"hash/crc32"
	"strconv"
	"sync"
	"time"
)

var mutex sync.Mutex

//改进方法
func GetMyUUID() string {
	mutex.Lock()
	defer mutex.Unlock()
	currentTime := time.Now()
	timestamp := currentTime.UnixNano()
	uuid := uuid.NewV4()
	uuidHash := int(crc32.ChecksumIEEE([]byte(uuid.String())))

	if -uuidHash >= 0 {
		uuidHash = -uuidHash
	}
	uuidStr := strconv.FormatInt(timestamp, 10) + strconv.Itoa(uuidHash)
	return uuidStr
}
