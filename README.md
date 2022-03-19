# kasa-stocks
_use of this Software is subject to the MIT license as specified here: [LICENSE.txt](./LICENSE.txt)_

This project allows changing the color of a Kasa smart-bulb based on the stock market. 

Its designed to be customizable to change the condition such as crypto market, the weather etc.

## Pre-requisites

- Install [python-kasa](https://github.com/python-kasa/python-kasa). You can install it by running the command: `pip3 install python-kasa`
- Download your corresponding binary of kasa-stocks. If you want to build your own binary, please see the building your own section below. 

## Usage

After you have installed the binary, update config.json with the path to binary and other parameters as needed.

## Build your own (optional)

Pre-requisites:
- Golang dev environment
- Run the command `make build`