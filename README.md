# yae - YAML aliases expander

[![Build Status](https://travis-ci.org/shin1x1/yae.svg?branch=master)](https://travis-ci.org/shin1x1/yae)


## Installation

<https://github.com/shin1x1/yae/releases>

```
go get -u github.com/shin1x1/yae
```

## Usage

```
$ cat src.yaml
---
anchors:
  items:
    - &book
      name: book
      price: 100

stocks:
  - *book
  - <<: *book
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