package argument

import (
	"fmt"
	"os"
	"os/exec"
)

const Version = "v.1.0.0"
const git = "https://github.com/yoehwan/youtube-dl-server.git"
const Logo = "YOUTUBEDLSERVER"

type Upgrade struct{}

func (u *Upgrade) Do() {
	res, needUpgrade := checkVersion()
	fmt.Fprintln(os.Stdout, string(res))
	if needUpgrade {
		showLogo()
		res, err := pullNewVersion()
		if err != nil {
			fmt.Fprintln(os.Stdout, string(res))
			return
		}
		res, err = build()
		if err != nil {
			fmt.Fprintln(os.Stdout, string(res))
			return
		}
		u.Do()
	} else {
		fmt.Fprintln(os.Stdout, "Current Version is already Newest.")
		printVersion()
	}

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
	fmt.Fprintln(os.Stdout, Logo)

}

func printVersion() {
	fmt.Fprintln(os.Stdout, "Version : "+Version)

}

func pullNewVersion() ([]byte, error) {
	return command("git", "pull")

}

func build() ([]byte, error) {
	return command("go", "build", ".")
}

func command(name string, cmd ...string) ([]byte, error) {
	exe := exec.Command(name, cmd...)
	return exe.Output()
}
