package errors

const (
	ErrInvalidManagementKey KeyError = "invalid management key"
	ErrInvalidSigningKey    KeyError = "invalid signing key"
	ErrInvalidDecryptionKey KeyError = "invalid decryption key"
	ErrUnsupportedAlgorithm KeyError = "unsupported algorithm"
)

const (
	ErrInvalidKeySlot                 KeystoreError = "invalid key slot"
	ErrUnsupportedAlgorithmForKeySlot KeystoreError = "unsupported algorithm for key slot"
	ErrExportNotAllowed               KeystoreError = "export not allowed"
	ErrImportNotAllowed               KeystoreError = "import not allowed"
	ErrKeyAlreadyExists               KeystoreError = "key already exists"
	ErrKeyNotFound                    KeystoreError = "key not found"
	ErrFileAlreadyExists              KeystoreError = "file already exists"
	ErrCertAlreadyExists              KeystoreError = "certificate already exists"
	ErrCertNotFound                   KeystoreError = "cert not found"
	ErrUnableToOpenKeystore           KeystoreError = "unable to open keystore"
	ErrKeystoreLocked                 KeystoreError = "keystore is locked"
	ErrKeystoreHandleClosed           KeystoreError = "keystore handle is closed"
	ErrKeystoreNotfound               KeystoreError = "keystore not found"
	ErrResetNotAllowed                KeystoreError = "reset not allowed"
	ErrKeyGenFailed                   KeystoreError = "key generation failed"
)
