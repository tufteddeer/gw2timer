# gw2timer

GW2Timer gives you desktop notifications about upcoming world bosses in Guild Wars 2

# Dependencies

Works only on linux!

- Gtk+3
- notify-send/libnotify

Depending on your desktop environment or window manager, notifications my not be displayed over full screen windows like GW2!

# Build

Stuff i used:
- Go 1.11.4
- gopkg.in/robfig/cron.v2 
- github.com/gotk3/gotk3/gtk 
- github.com/go-bindata/go-bindata
- github.com/go-task/task for the build process

Assuming task is installed and in your path:

```
$ task install # embed assets and run go install
$ taks run # install and start gw2timer
```
gw2timer should now be in your $GOBIN

See Taskfile.yml for details.