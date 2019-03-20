PockerView
====

CUI program which supports modified UPI(Universal Porker Interface) protocol.

## Description
When you making a poker AI, minding an interface part is meaningless. Probably you want to concentrate an engine itself.
So, you need something which absorbs the difference of protocols of interface parts. This CUI program does it.

Now we are thinking PockerView supports three type of uses cases.
+ CUI through standart input
+ Remote server which supports ACPC
+ multiple engines games


## Usage
First, you prepare a binary of an engine.
And then you can build this program after writing some configs in config.toml. 

### BUILD

``make -f docker.Makefile binary`` for linux.

``make -f docker.Makefile binary-osx`` for macOS.

``make -f docker.Makefile binary-windows`` for windows.

## Requirement
Docker

## Contribution
Pull Request is welcome!
You can prepare a development environment by ``make -f docker.Makefile dev``.

## Licence
MIT

## Author

[KazuyaTomita](https://github.com/KazuyaTomita)

