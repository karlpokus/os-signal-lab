# os-signal-lab
Experiments with os signals and process trees

There's a lot to read about os signals on linux but I want some hard data. Let's investigate:
- What signals are sent (if any) to processes on system reboot?
- How does a process send a reboot command to a parent and still exit before being restarted? Assuming we have a restart directive defined.

# signals
Here's a mapping between signals and strings that `signal.Notify` sends on a channel
- SIGINT: interrupt
- SIGTERM: terminated
- SIGHUP: hangup
- SIGQUIT: quit

We know we can't catch SIGKILL but what about SIGSTOP and companion SIGCONT? Well, I was unable to catch SIGSTOP and it looks like sending SIGSTOP to a process that's connected to a tty in another shell - makes it at least stop reading stdin. Then resumes after SIGCONT. And if I pass SIGINT on stdin while the process is being STOPPED it gets processed after SIGCONT. cool.

Both `reboot` and `systemctl reboot [--no-block]` sends SIGHUP which we can catch. The latter is much faster though. Both allow graceful exits.

# test
Use the binary to gather data on a linux device. Output includes pid for manual signaling from another shell. After we catch the signal we wait 3s to see if we're allowed a graceful exit.

````bash
$ go build
# run and optionally keep a record
$ ./signal-catcher 2>> catcher.log &
````

# license
MIT
