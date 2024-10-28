package constants

var Folders = map[string]string{
	"StoreFolder":  "./store",
	"ConfigFolder": "./store/configurations",
	"DataFolder":   "./store/data",
}

var Files = map[string]string{
	"GroupFile": "./store/configurations/group.txt",
}

var ValidCommands = []string{
	"ls",
	"usegrp",
	"add",
	"showgrp",
	"dropgrp",
}
