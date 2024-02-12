# nvim-godot
Script to launch nvim, if it hasn't been, on editing .gdscript files in Godot engine

I'm using Windows, not sure how it would work with other setup.

I've setup my Neovim following [this guide](https://mb-izzo.github.io/nvim-godot-solution/)

# Usage

Set the binary as the executable in the Godot.

Set the flags:
```code
--server [your server arg, for example "127.0.0.1:33422", or "godot.pipe", this is redirected to nvim]
```

```code
--exec [your terminal command for nvim, I use nvim-qt]
//Currently works only with nvim-qt
```

```code
--remote-send [your remote-send arg, this is redirected to nvim]
```

```code
//Just plug the {project} here
--project {project} 
```
