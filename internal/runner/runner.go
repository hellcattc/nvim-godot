package runner

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"time"
)

type Runner struct {
	Options Options
}

type Options struct {
	Executable string
	Addr       string
    RemoteSend string
    ProjectDir string
    DebugOn bool
}

func NewRunner(opts Options) *Runner {
   return &Runner{Options: opts}
}

func (r *Runner) Run() error {
    err := r.checkAddr()
    if err != nil {
        if err = r.startVim(); err != nil {
            log.Fatal("Error starting Vim")
        }
        if err = r.checkVim(); err != nil {
            log.Fatal(err)
        }
    }
    cmd := exec.Command(r.Options.Executable, "--", "--server", r.Options.Addr, "--remote-send", r.Options.RemoteSend)
    fmt.Printf("cmd.String(): %v\n", cmd.String())
    if err = cmd.Start(); err != nil {
        log.Fatal("Error with remote-send arg")
    }
    return nil
} 

func (r *Runner) getConn() (net.Conn, error) {
    return net.Dial("tcp", r.Options.Addr)
}

func (r *Runner) checkAddr() error {
	conn, err := r.getConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Printf("Listener found on addr %s\n", r.Options.Addr)
	return nil
}

func (r *Runner) startVim() error {
    var args []string
	if r.Options.Executable == "nvim-qt" {
		args = append(args, "--") 
	}
    if r.Options.DebugOn {
        args = append(args, "-V16vimlog.log", "--startuptime", "vimlog.txt")
    }
	args = append(args, "vimlog.txt", "--listen", r.Options.Addr, r.Options.ProjectDir) 
	cmd := exec.Command(r.Options.Executable, args...)
    cmd.Dir = "."
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (r *Runner) checkVim() error {
    retryInterval := 500 * time.Millisecond
    
    maxTries := 20

    for i := 1; i < maxTries; i++ {
        if err := r.checkAddr(); err == nil {
            return nil
        }
        time.Sleep(retryInterval)
    }

    return fmt.Errorf("Neovim didn't become ready")
}
