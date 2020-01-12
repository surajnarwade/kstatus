# kstatus

* Build the binary

```
make bin
```

* move it to PATH

```
mv kubectl-status ~/.local/bin
```

note: I prefer `~/.local/bin`, you can use different PATH as well


* your subcommand is ready to use with kubectl

```
kubectl status
```