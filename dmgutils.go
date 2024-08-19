package dmgutils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func MountDMG(dmgPath string, mountPoint string) error {
	cmd := exec.Command("hdiutil", "attach", dmgPath, "-mountpoint", mountPoint)
	return cmd.Run()
}

func UnmountDMG(mountPoint string) {
	exec.Command("hdiutil", "detach", mountPoint).Run()
}

func FindApplication(mountPoint string) (string, error) {
	// Implement your logic to find the application path
	// Here we assume the application is directly under the mount point
	// For example purposes, we'll just look for a `.app` bundle
	files, err := os.ReadDir(mountPoint)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.IsDir() && strings.HasSuffix(file.Name(), ".app") {
			return filepath.Join(mountPoint, file.Name()), nil
		}
	}

	return "", fmt.Errorf("no application found in the DMG")
}

func VerifyCodeSigning(appPath string) (bool, error) {
	cmd := exec.Command("codesign", "--verify", "--deep", "--strict", appPath)
	err := cmd.Run()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 0 {
			return true, nil
		}
		return false, err
	}
	return true, nil
}
