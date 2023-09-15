package errors

const (
	ErrInvalidSigningKey    KeyError = "invalid signing key"
	ErrInvalidDecryptionKey KeyError = "invalid decryption key"
)

const (
	ErrUnsupportedAlgorithmForKeySlot KeystoreError = "unsupported algorithm for key slot"
	ErrExportNotAllowed               KeystoreError = "export not allowed"
	ErrImportNotAllowed               KeystoreError = "import not allowed"
	ErrKeyAlreadyExists               KeystoreError = "key already exists"
	ErrFileAlreadyExists              KeystoreError = "file already exists"
	ErrCertAlreadyExists              KeystoreError = "certificate already exists"
	ErrUnableToOpenKeystore           KeystoreError = "unable to open keystore"
	ErrKeystoreLocked                 KeystoreError = "keystore is locked"
	ErrKeystoreHandleClosed           KeystoreError = "keystore handle is closed"
	ErrResetNotAllowed                KeystoreError = "reset not allowed"
	ErrKeyGenFailed                   KeystoreError = "key generation failed"
)
