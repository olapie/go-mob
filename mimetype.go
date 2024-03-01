package mob

import (
	"mime"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"go.olapie.com/ola/mimetypes"
)

func IsTextFile(f FileInfo) bool {
	return IsMIMEText(f.MIMEType())
}

func IsImageFile(f FileInfo) bool {
	return IsMIMEImage(f.MIMEType())
}

func IsVideoFile(f FileInfo) bool {
	return IsMIMEVideo(f.MIMEType())
}

func IsAudioFile(f FileInfo) bool {
	return IsMIMEAudio(f.MIMEType())
}

func IsMIMEText(t string) bool {
	return mimetypes.IsText(t)
}

func IsMIMEImage(t string) bool {
	return strings.HasPrefix(t, "image") || strings.HasPrefix(t, "photo")
}

func IsMIMEVideo(t string) bool {
	return strings.HasPrefix(t, "video")
}

func IsMIMEAudio(t string) bool {
	return strings.HasPrefix(t, "audio")
}

func ExtensionByMIMEType(mimeType string) string {
	if m := mimetype.Lookup(mimeType); m != nil && m.Extension() != "" {
		return m.Extension()
	}

	exts, err := mime.ExtensionsByType(mimeType)
	if err != nil || len(exts) == 0 {
		return ""
	}
	return exts[0]
}
