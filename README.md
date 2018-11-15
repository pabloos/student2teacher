# std2tch

[![Build Status](https://travis-ci.org/pabloos/student2teacher.svg?branch=master)](https://travis-ci.org/pabloos/student2teacher)   [![Coverage Status](https://coveralls.io/repos/github/pabloos/student2teacher/badge.svg?branch=master)](https://coveralls.io/github/pabloos/student2teacher?branch=master)

std2tch is a web app for send and process the student source code to the teacher ("std" is in revenge for stdio.h).

The main idea is to run the server in the teacher's computer and wait to receive the files from the students via webpage.

If required, you can run the server program indicating the dir where you want the codefiles to be stored by passing the dirname as a flag:

`./std2tch -path=mydir`

Also you can run the server without any flag and use the default dir `./home`.

## docker deploy
There's a Makefile for automate task. One of them is to build and run the app  (both of them in docker containers). If you you want to build a linux binary:

```shell
make docker-build
```

And for run the app:

```shell
make docker-run
```