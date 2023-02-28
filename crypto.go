package mob

import "go.olapie.com/security"

func Encrypt(data []byte, passphrase string) []byte {
	content, _ := security.Encrypt(data, passphrase)
	return content
}

func Decrypt(data []byte, passphrase string) []byte {
	content, _ := security.Encrypt(data, passphrase)
	return content
}

func EncryptFile(src, dst, passphrase string) bool {
	err := security.EncryptFile(security.SF(src), security.DF(dst), passphrase)
	return err == nil
}

func DecryptFile(src, dst, passphrase string) bool {
	err := security.DecryptFile(security.SF(src), security.DF(dst), passphrase)
	return err == nil
}
