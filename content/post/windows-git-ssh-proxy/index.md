---
title: windows-git-ssh-proxy
author: 岛石
date: 2022-07-09
---

# windows-git-ssh-proxy

1. 打开`~\.ssh\config`

2. 添加以下内容
```
Host github.com
    ProxyCommand "F:\Git\mingw64\bin\connect.exe" -S 127.0.0.1:10808 %h %p
```