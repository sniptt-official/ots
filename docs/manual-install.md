# Manual install

## macOS and Linux

See available binaries on the [releases page](https://github.com/sniptt-official/ots/releases).

```sh
$ curl -L https://github.com/sniptt-official/ots/releases/download/v0.0.7/ots_0.0.7_darwin_amd64.tar.gz -o ots.tar.gz
$ sudo mkdir -p /usr/local/ots-cli
$ sudo tar -C /usr/local/ots-cli -xvf ots.tar.gz
$ sudo ln -sf /usr/local/ots-cli/ots /usr/local/bin/ots
$ rm ots.tar.gz
```

Assuming `/usr/local/bin` is on your `PATH`, you can now run:

```sh
$ ots --version
ots version 0.0.7
```

### Uninstall

1.  Find the folder that contains the symlink to the main binary.

```sh
$ which ots
/usr/local/bin/ots
```

2.  Using that information, run the following command to find the installation folder that the symlink points to.

```sh
$ ls -l /usr/local/bin/ots
lrwxr-xr-x  1 root  admin  26 15 Jul 16:00 /usr/local/bin/ots -> /usr/local/ots-cli/ots
```

3.  Delete the symlink in the first folder. If your user account already has write permission to this folder, you don't need to use `sudo`.

```sh
$ sudo rm /usr/local/bin/ots
```

4.  Delete the main installation folder.

```sh
$ rm -rf /usr/local/ots-cli
```
