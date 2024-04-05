

### test function을 만들고 디버깅하면 안된다. 
```
Starting: /home/jhyunlee/go/bin/dlv dap --listen=127.0.0.1:36465 --log-dest=3 from /home/jhyunlee/go/src/devgo/sutest
DAP server listening at: 127.0.0.1:36465
Type 'dlv help' for list of commands.
exit status 1
```

### main.go 컴파일 하고  F5를 통해서 디버기하면 된다.  
```
root@Good:/home/jhyunlee/go/src/devgo#  cd /home/jhyunlee/go/src/devgo ; /usr/bin/env GOPATH=/home/jhyunlee/go /usr/bin/sudo /home/jhyunlee/go/bin/dlv dap --client-addr=:38045 
```

### vscode의 extention에서 제공하는 dlv
* F5 눌러서 디버깅하는 것과 차이가 있기 때문에 
* vscode 

### delveConfig 
* go.delveConfig
   * debugAdapter: Controls which debug adapter to use (default: legacy). Select ‘dlv-dap’.
   * showGlobalVariables: Show global variables in the Debug view (default: false).
   * substitutePath: Path mappings to apply to get from a path in the editor to a path in the compiled program (default: []).


   