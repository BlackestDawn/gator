package config

import "io/fs"

const configFileName string = ".gatorconfig.json"
const configFileMode = fs.FileMode(0644)
