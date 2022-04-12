package argument

import "os/exec"

const Version = "v.1.0.0"
const git = "https://github.com/yoehwan/youtube-dl-server.git"

type Upgrade struct {
}

func (u *Upgrade) Do() {

}

func checkVersion() ([]byte, bool) {
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

func showLogo() {

}

func printVersion() {

}

func pullNewVersion() {

}

func runMain() {

}

func command(name string, cmd ...string) ([]byte, error) {
	exe := exec.Command(name, cmd...)
	return exe.Output()
}
