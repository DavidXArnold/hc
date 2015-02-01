package common

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func SerialFilename(accessoryName string) string {
	return accessoryName + ".serial"
}

func GetSerialNumberForAccessoryName(name string, storage Storage) string {
	serial_file := SerialFilename(name)
	serial_bytes, _ := storage.Get(serial_file)
	serial := string(serial_bytes)
	if len(serial) == 0 {
		serial = GenerateSerialNumber()
		storage.Set(serial_file, []byte(serial))
	}

	return serial
}

func GenerateSerialNumber() string {
	t := time.Now().Format(time.RFC3339Nano)

	h := md5.New()
	h.Write([]byte(t))
	result := h.Sum(nil)

	return hex.EncodeToString(result)
}
