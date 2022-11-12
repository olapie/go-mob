package mobile

import "code.olapie.com/cryptokit"

func Encrypt(data []byte, passphrase string) []byte {
	content, _ := cryptokit.Encrypt(data, passphrase)
	return content
}

func Decrypt(data []byte, passphrase string) []byte {
	content, _ := cryptokit.Decrypt(data, passphrase)
	return content
}

func EncryptFile(dst, src, passphrase string) bool {
	err := cryptokit.EncryptFile(cryptokit.Destination(dst), cryptokit.Source(src), passphrase)
	return err == nil
}

func DecryptFile(dst, src, passphrase string) bool {
	err := cryptokit.DecryptFile(cryptokit.Destination(dst), cryptokit.Source(src), passphrase)
	return err == nil
}
