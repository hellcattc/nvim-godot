package main

import (
	"flag"
	"log"

	"github.com/hellcattc/godot-nvim-script/internal/runner"
)

func main() {
    exec_arg := flag.String("exec", "nvim", "The executable for the editor, for example 'nvim' or 'nvim-qt'")
    server_arg := flag.String("server", "", "The server parameters, can be anything the --server argument gets")
    remote_send_arg := flag.String("remote-send", "", "The remote send parameters, can be anything the --remote-send argument gets")
    project_dir_arg := flag.String("project", "", "The project dir, do the {project} here")
    debug_info_arg := flag.Bool("debug", false, "Enable args '-V16vimlog.log' and '-startuptime vimlog.txt'")

    flag.Parse()

    if *server_arg == "" {
        log.Fatalln("Please specify the server arg")
    }
    if *remote_send_arg == "" {
        log.Fatalln("Please specify the remote-send argument")
    }
    if *project_dir_arg == "" {
        log.Fatalln("Please specify the project dir")
    }

    runner := runner.NewRunner(runner.Options{Executable: *exec_arg, Addr: *server_arg, RemoteSend: *remote_send_arg, ProjectDir: *project_dir_arg, DebugOn: *debug_info_arg})

    err := runner.Run()

    if err != nil {
        log.Fatalln(err)
    }
}
