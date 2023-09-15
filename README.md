# G-Net Keyring

Abstraction layer for turning cryptographic coprocessors into hardware security modules. I'm planning to support the following coprocessors in my initial release:

 - Trusted Computing Group's [TPM 2.0](https://trustedcomputinggroup.org/resource/tpm-library-specification/)
 - Yubico's [Yubikey (PIV)](https://developers.yubico.com/PIV/)

I plan to support the following coprocessors at some point in the future:

 - Microchip Technology's [ATECC608A](https://www.microchip.com/en-us/product/ATECC608A)
 - [Tentative] Apple's [Secure Enclave](https://support.apple.com/guide/security/secure-enclave-sec59b0b31ff/web)

This package defines a common interface for interacting with the hardware listed above.