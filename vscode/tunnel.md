### 1. download 

```
curl -Lk 'https://code.visualstudio.com/sha/download?build=stable&os=cli-alpine-x64' --output vscode_cli.tar.gz
tar -xf vscode_cli.tar.gz
```

### 2. create secure tunnel  
```sh
code tunnel

# It will ask you to login with github account and setup device login.
# Press Ctrl+C to exit the tunnel
code tunnel --accept-server-license-terms
```
### 3. connect vscode
```sh
https://vscode.dev/tunnel/<machine_name>/<folder_name>.
```

### 4. system daemon
```sh
sudo vi /etc/systemd/system/code-tunnel.service


[Unit]
Description=VSCode Tunnel as Daemon

[Service]
Type=simple
ExecStart=/snap/bin/code tunnel
User=jhyunlee
Group=jhyunlee
Restart=on-failure

[Install]
WantedBy=multi-user.target



sudo systemctl daemon-reload
sudo systemctl enable code-tunnel
sudo systemctl start code-tunnel
sudo systemctl status code-tunnel
```

```
$ sudo journalctl  -u code-tunnel
```