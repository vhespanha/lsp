# LSP

A minimalist Language Server Protocol implementation for educational purposes.

## Overview

This project implements basic LSP features to demonstrate protocol fundamentals. It supports markdown files and includes hover information, basic navigation, code actions, and simple completions.

## Installation

Clone and install:

```bash
git clone git.sr.ht/~vhespanha/lsp && cd lsp
make build
make install
```

## Client Configuration

### Neovim

```lua
local client = vim.lsp.start_client {
  name = 'educationalsp',
  cmd = { '$HOME/.local/bin/lsp' },
}

if not client then
  vim.notify 'bad client'
  return
end

vim.api.nvim_create_autocmd('Filetype', {
  pattern = 'markdown',
  callback = function()
    vim.lsp.buf_attach_client(0, client)
  end,
})
```

### VSCode

Requires creating a custom extension. See a reference implementation [here](https://github.com/microsoft/vscode-extension-samples/tree/main/lsp-sample).

## Features

- Document synchronization (open/change)
- Hover: Shows file path and character count
- Definition: Basic line navigation
- Code actions: Foo → Bar/Baz conversion
- Completions: Handles "Melange" with description
- Diagnostics: Basic error detection for "Foo"

## License

[MIT](./LICENSE)
