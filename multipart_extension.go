package json_api

import "mime/multipart"

type FormFile struct {
	multipart.File
	multipart.FileHeader
	error
}

