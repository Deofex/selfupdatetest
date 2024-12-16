package selfupdate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

const versionIndexUrl = "https://deofex.github.io/selfupdatetest/version.json"

func SelfUpdate(currentVersion string) (err error) {
	upstreamVersion, err := getUpstreamVersion()
	if err != nil {
		return
	}
	if upstreamVersion.CurrentVersion == currentVersion {
		return
	}
	path, err := os.Executable()
	if err != nil {
		return
	}
	fmt.Printf("New version available (%s), current version (%s) updating...\n", upstreamVersion.CurrentVersion, currentVersion)
	err = updateExecutable(upstreamVersion)
	if err != nil {
		return
	}
	fmt.Println("Update successfully installed, restarting executable...")
	err = restartProgram(path)
	if err != nil {
		fmt.Println("Unable to restart program")
	}
	return
}

func getUpstreamVersion() (versionInfo updateVersionInfo, err error) {
	resp, err := http.Get(versionIndexUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return updateVersionInfo{}, fmt.Errorf("unable to get version: %s", resp.Status)
	}
	versionBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(versionBytes, &versionInfo)
	if err != nil {
		return
	}
	return versionInfo, nil
}

func updateExecutable(versionInfo updateVersionInfo) (err error) {
	updateUrl, err := getBinaryDownloadPath(versionInfo)
	if err != nil {
		return
	}
	resp, err := http.Get(updateUrl)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unable to download update: %s", resp.Status)
	}
	path, err := os.Executable()
	if err != nil {
		return
	}
	backupFile := path + ".bak"
	err = os.Rename(path, backupFile)
	if err != nil {
		return
	}
	newExec, err := os.Create(path)
	if err != nil {
		_ = os.Rename(backupFile, path)
		return
	}
	defer newExec.Close()
	_, err = io.Copy(newExec, resp.Body)
	if err != nil {
		_ = os.Rename(backupFile, path)
		return
	}
	err = os.Chmod(path, 0755)
	if err != nil {
		return
	}
	os.Remove(backupFile)
	return
}

func getBinaryDownloadPath(versionInfo updateVersionInfo) (url string, err error) {
	osi := getOsInfo()
	var updateUrl string
	switch osi.OS {
	case "windows":
		updateUrl = versionInfo.WindowsAmd64Binary
	case "linux":
		updateUrl = versionInfo.LinuxAmd64Binary
	case "darwin":
		switch osi.Arch {
		case "amd64":
			updateUrl = versionInfo.DarwinAmd64Binary
		case "arm64":
			updateUrl = versionInfo.DarwinArm64Binary
		default:
			return "", fmt.Errorf("unsupported Architecture: %s (OS: %s)", osi.Arch, osi.OS)
		}
	default:
		return "", fmt.Errorf("unsupported OS: %s, Architecture: %s", osi.OS, osi.Arch)
	}
	if updateUrl == "" {
		return "", fmt.Errorf("no update binary found for %s %s", osi.OS, osi.Arch)
	}
	return updateUrl, nil
}

func getOsInfo() (os osInfo) {
	return osInfo{
		OS:   runtime.GOOS,
		Arch: runtime.GOARCH,
	}
}

func restartProgram(execPath string) error {
	args := os.Args
	env := os.Environ()

	cmd := exec.Command(execPath, args[1:]...)
	cmd.Env = env
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to start new process: %w", err)
	}

	os.Exit(0)

	return nil
}
