package util

import "strings"

const (
	base64Keys   = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="
	uuidTemplate = "........-....-....-....-............"
	hexChars     = "0123456789abcdef"
)

func Decompress(base64 string) string {
	length := len(base64)
	if length == 22 {
		return decompress(base64, 2)
	} else if length == 23 {
		return decompress(base64, 5)
	} else {
		return base64
	}
}

func decompress(base64 string, start int) string {
	indices := make([]byte, 0)
	for i, v := range uuidTemplate {
		if v != 45 {
			indices = append(indices, byte(i))
		}
	}

	uuid := []byte(uuidTemplate)
	uuid[0] = base64[0]
	uuid[1] = base64[1]

	var base64Values [123]byte
	for i := 0; i < 123; i++ {
		base64Values[i] = 64
	}
	for i := 0; i < 64; i++ {
		base64Values[base64Keys[i]] = byte(i)
	}

	for i, j := start, 2; i < 22; i = i + 2 {
		lhs := base64Values[base64[i]]
		rhs := base64Values[base64[i+1]]
		uuid[indices[j]] = hexChars[lhs>>2]
		uuid[indices[j+1]] = hexChars[(lhs&3)<<2|rhs>>4]
		uuid[indices[j+2]] = hexChars[rhs&0xf]
		j += 3
	}
	return string(uuid)
}

func Compress(uuid string) string {
	if len(uuid) != 36 {
		return uuid
	}

	uuid = strings.ReplaceAll(uuid, "-", "")

	zipUuid := make([]byte, 22)
	zipUuid[0] = uuid[0]
	zipUuid[1] = uuid[1]

	hexMap := make(map[byte]byte, 16)
	for i, v := range hexChars {
		hexMap[byte(v)] = byte(i)
	}

	for i, j := 2, 2; i < 32; i += 3 {
		left := hexMap[uuid[i]]
		mid := hexMap[uuid[i+1]]
		right := hexMap[uuid[i+2]]

		zipUuid[j] = base64Keys[(left<<2)+(mid>>2)]
		zipUuid[j+1] = base64Keys[((mid&3)<<4)|right]

		j += 2
	}
	return string(zipUuid)
}
