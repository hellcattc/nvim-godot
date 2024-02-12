# nvim-godot
Script to launch nvim, if it hasn't been, on editing .gdscript files in Godot engine

I'm using Windows, not sure how it would work with other setup.

I've setup my Neovim following [this guide](https://mb-izzo.github.io/nvim-godot-solution/)

# Usage

Set the binary as the executable in the Godot.

### Set the flags:

Server you'll be using for communication

```code
--server [your server arg, for example "127.0.0.1:33422", or "godot.pipe", this is redirected to nvim]
```

The executable for nvim (currently only nvim-qt works)

```code
--exec [your terminal command for nvim, I use nvim-qt]
```

The remote-send options for listener

```code
--remote-send [your remote-send arg, this is redirected to nvim]
```

Godot project directory so that neovim sets it as the working directory

```code
//Just plug the {project} here
--project {project} 
```

Create debug logs on starting neovim

```code
--debug
```
