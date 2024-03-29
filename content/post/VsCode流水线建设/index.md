---
title: VSCode全键盘方案
date: 2022-08-30
author: 岛石  
categories: 
- "程序"
tag: 
- "VSCode"
- "环境"
---
# VsCode 流水线建设

## 插件

- vscode.Vim
- yzhang.markdown-all-in-one
- monokai.theme-monokai-pro-vscode

## 全键盘-命令&快捷键

### Command

```shell
code --install-extension <extension-identifier>
code <file>|<dir>
```

### VsCode Default Shortcut

|                      |              |
| -------------------- | ------------ |
| File Explorer        | ctrl+shift+e |
| Extension            | ctrl+shift+x |
| Excute Commands      | ctrl+shift+p |
| Command Palette      | ctrl+p       |
| Document Format      | alt+shift+f  |
| Focus search results | F4/shift+F4  |

### Custome Config

```json
    //settings.json
    //vim剪切板
    "vim.useSystemClipboard": true,
    //jj退出InsertMode
    "vim.insertModeKeyBindings": [
        {
            "before" :["j","j"],
            "after":["<ESC>"]
        }
    ],
    "[markdown]": {
        "editor.defaultFormatter": "yzhang.markdown-all-in-one"
    },
    "workbench.colorTheme": "Monokai Pro"
```

```json
    {
        "key": "ctrl+shift+n",
        "command": "explorer.newFolder",
        "when": "explorerViewletFocus"
    },
    {
        "key": "ctrl+n",
        "command": "explorer.newFile",
        "when": "explorerViewletFocus"
    }
```
