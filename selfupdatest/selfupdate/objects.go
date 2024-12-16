package selfupdate

type updateVersionInfo struct {
	CurrentVersion           string `json:"current_version"`
	WindowsAmd64UpdateBinary string `json:"windows_amd64_update_binary"`
	LinuxAmd64UpdateBinary   string `json:"linux_amd64_update_binary"`
	DarwinAmd64UpdateBinary  string `json:"darwin_amd64_update_binary"`
	DarwinArm64UpdateBinary  string `json:"darwin_arm64_update_binary"`
}

type osInfo struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}
