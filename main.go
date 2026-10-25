package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var flint = []byte(`name: flint

on: [push, pull_request]

jobs:
  flint:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: luanti-org/flint@master
        with:
          command: lint
      - uses: luanti-org/flint@master
        with:
          command: format
      - uses: luanti-org/flint@master
        with:
          command: compress
`)

var luacheck = []byte(`name: luacheck
on:
  push:
    paths:
      - "**.lua"
  pull_request:
    paths:
      - "**.lua"
jobs:
  luacheck:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@main
      - name: Luacheck
        uses: lunarmodules/luacheck@master
`)
var reuse = []byte(`name: REUSE compliance check
on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@main
      - name: REUSE Compliance Check
        uses: fsfe/reuse-action@main
`)

func main() {
	dirPath := filepath.Join(".github", "workflows")
	err := os.MkdirAll(dirPath, 0775)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filePath := filepath.Join(dirPath, "flint.yml")
	err = os.WriteFile(filePath, flint, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filePath = filepath.Join(dirPath, "luacheck.yml")
	err = os.WriteFile(filePath, luacheck, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filePath = filepath.Join(dirPath, "reuse.yml")
	err = os.WriteFile(filePath, reuse, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
