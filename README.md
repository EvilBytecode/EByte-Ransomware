# EByte-Ransomware
Go ransomware leveraging ChaCha20 and ECIES encryption with a web-based control panel.

---
- Contact me at : https://t.me/codepulze
- <a href="https://t.me/+TRYMuOVDiWA4MjM0"><img src="https://img.shields.io/badge/Join%20my%20Telegram%20group-2CA5E0?style=for-the-badge&logo=telegram&labelColor=db44ad&color=5e2775"></a>
- [Join our Discord server!](https://discord.gg/NRTdwYUtdQ)

---

## Disclaimer
**This project is strictly for educational purposes only.**  
It is intended to demonstrate the implementation of cryptographic techniques and must not be used for malicious, illegal, or unethical purposes. The misuse of this code for unauthorized activities, including data encryption or extortion, may result in severe legal consequences.  
The author assumes no liability or responsibility for any damage, loss, or legal action caused directly or indirectly by the use or misuse of this project. Always adhere to ethical guidelines, cybersecurity laws, and regulations.

## Brief Overview
EByte is a ransomware written from scratch in Go. It uses a mixture of ChaCha20 and ECIES cryptography to encrypt files securely so that they cannot be recovered by traditional recovery tools. Files encrypted by EByte can only be decrypted using the corresponding decryptor.

## Installation & Setup
### Pre-requisites:
- [The Go Programming Language](https://go.dev)

## Running:
- ```go run server.go```

## Encryption Process
- The encryptor enumerates all drives on the system and proceeds to iterate through each directory recursively.
- It ignores blacklisted files, directories, and extensions.
- It generates a unique ChaCha20 key and nonce for each file and encrypts the file using a pattern of 1 byte encrypted, 2 bytes unencrypted.
- It encrypts the ChaCha20 key and nonce using the ECIES public key and prepends them to the start of the file.

## Benefits of ChaCha20 and ECIES
I chose this unique combination of encryption methods for several reasons:
- ChaCha20's stream-based approach allows for byte-by-byte encryption, enabling the pattern of 1 byte encrypted, 2 bytes unencrypted.
- ECIES offers similar security to RSA with shorter key lengths, making it a more efficient choice.

## Showcase:
![First](https://github.com/user-attachments/assets/7c742129-81c1-45c4-9044-6da7583091e7)
![Second](https://github.com/user-attachments/assets/4e227eae-7a61-4a05-9914-4276ad68027e)
![In Action](https://github.com/user-attachments/assets/6b50e00d-9160-462e-b16a-5876536248ee)

# Credits:
- https://github.com/SecDbg/Prince-Ransomware
