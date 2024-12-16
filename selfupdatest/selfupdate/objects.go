package selfupdate

type updateVersionInfo struct {
	CurrentVersion     string `json:"current_version"`
	WindowsAmd64Binary string `json:"windows_amd64_binary"`
	LinuxAmd64Binary   string `json:"linux_amd64_binary"`
	DarwinAmd64Binary  string `json:"darwin_amd64_binary"`
	DarwinArm64Binary  string `json:"darwin_arm64_binary"`
}

type osInfo struct {
	OS   string `json:"os"`
	Arch string `json:"arch"`
}
