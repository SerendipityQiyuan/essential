package utils

import "mime/multipart"

func ImageTypeValid(headers *multipart.FileHeader) bool {
	ImageContentTypeArr := []string{"image/jpeg", "image/jpg", "image/png", "image/gif"}
	for _, item := range ImageContentTypeArr {
		if headers.Header.Get("Content-Type") == item {
			//如果是以上ImageContentTypeArr中的一种则格式正确
			return true
		}
	}
	return false
}
