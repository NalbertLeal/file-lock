# File Lock

![GitHub](https://img.shields.io/github/license/NalbertLeal/file-lock)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/NalbertLeal/file-lock)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/NalbertLeal/file-lock)

## Description

_File lock_ is a Go tool developed to encrypt files using the Advanced Encryption Standard (AES) algorithm. Don't use this as your main encryption tool, just to fast encypt a file that should not be open to anyone.

## How to Use

### Prerequisites

Ensure you have Go installed on your machine. For instructions on Go installation, visit [https://golang.org/doc/install](https://golang.org/doc/install).

### Installation

```bash
go get -u github.com/NalbertLeal/file-lock
```

### Basic Usage

Encrypt a file called _file\_name.txt_ using AES with key _key123_:

```bash
file-lock file_name.txt key123
```

## Contribution

If you wish to contribute improvements, bug fixes, or new features, please feel free to open an issue or submit a pull request. Your contributions are highly appreciated!

## License

This project is licensed under the [GNU LESSER GENERAL PUBLIC LICENSE](https://choosealicense.com/licenses/lgpl-2.1/).