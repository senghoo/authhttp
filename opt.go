package main

import (
	"flag"
	"fmt"
	"os"
)

const usage string = "Usage: %s [OPTIONS]\n"

type Options struct {
	User   string
	Pass   string
	Path   string
	Listen string
	Realm  string
}

func ParseArgs() (opts *Options, err error) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage, os.Args[0])
		flag.PrintDefaults()
	}

	user := flag.String(
		"user",
		"user",
		"Username for auth")

	pass := flag.String(
		"pass",
		"",
		"Password for user")

	path := flag.String(
		"path",
		"",
		"Path for serve")
	listen := flag.String(
		"listen",
		"0.0.0.0:8080",
		"Listen address")
	realm := flag.String(
		"realm",
		"Auth",
		"realm for server")

	flag.Parse()

	if *path == "" {
		current, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = &current
	}

	return &Options{
		User:   *user,
		Pass:   *pass,
		Path:   *path,
		Listen: *listen,
		Realm:  *realm,
	}, nil
}
