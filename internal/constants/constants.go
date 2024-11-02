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

	// group commands
	"usegrp",
	"showgrp",
	"dropgrp",
	"truncategrp",

	//task commands
	"ls",
	"add",
	// "delete",
	"done",
	// "amend",
}
