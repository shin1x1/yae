# yae - YAML aliases expander

[![Build Status](https://travis-ci.org/shin1x1/yae.svg?branch=master)](https://travis-ci.org/shin1x1/yae)

`yae` is a YAML aliases expander, written by Go.

## Motivation

When refactoring a long YAML using anchor/alias, I created it to make sure that the original YAML and the refactored YAML are identical.

By comparing the original YAML with the refactored YAML output from the `yae` command with `diff`, you can verify that the contents of YAML have not changed.

```
$ diff <(yae original.yaml) <(yae refactored.yaml)
```

Since the `yae` command can change the indentation of the list, the sequence of keys, and other parts that are not relevant to the meaning of YAML, the original YAML uses this command to ensure that the diff command does not cause differences in formatting.

## Installation

Download binary from below link.

<https://github.com/shin1x1/yae/releases>

You can install using `go get`.

```
go get -u github.com/shin1x1/yae
```

## Usage

Run the `yae` command with the YAML file path that you want to expand alias as an argument.

```
$ cat original.yaml
---
anchors:
  items:
    - &book # anchor
      name: book
      price: 100

stocks:
  - *book # alias
  - <<: *book # alias with overwrite key
    price: 200

$ yae src.yaml
anchors:
  items:
  - name: book
    price: 100
stocks:
- name: book
  price: 100
- name: book
  price: 200
```