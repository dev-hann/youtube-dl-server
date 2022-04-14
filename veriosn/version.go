package veriosn

import (
	"os/exec"
)

const _Version = "v.1.0.0"
const git = "https://github.com/yoehwan/youtube-dl-server.git"

type Version struct{}

func InitVersion() *Version {
	return &Version{}
}

func (v *Version) CurrentVersion() string {
	return _Version
}
func (v *Version) CheckVersion() ([]byte, bool) {
	data, err := command("git", "fetch")
	if err != nil {
		return data, false
	}
	var remote, local []byte
	remote, err = command("git", "rev-parse", "HEAD")
	if err != nil {
		return data, false
	}
	local, err = command("git", "rev-parse", "@{u}")
	if err != nil {
		return data, false
	}
	return nil, string(remote) == string(local)
}

func (v *Version) PullNewVersion() ([]byte, error) {
	return command("git", "pull")

}

func (v *Version) Build() ([]byte, error) {
	return command("go", "build", ".")
}

func command(name string, cmd ...string) ([]byte, error) {
	exe := exec.Command(name, cmd...)
	return exe.Output()
}
