// package config

// import (
// 	"log"
// 	"os"

// 	"gopkg.in/ini.v1"
// )

// type Configer interface {
// 	String(section, key string) string
// }

// type Confile struct {
// 	Path string
// 	File *ini.File
// }

// func (c *Confile) NewConfig() {

// 	Cfg, err := ini.Load(c.Path)
// 	if err != nil {
// 		log.Println(c.Path, ":", err)
// 		os.Exit(1)
// 	}
// 	c.File = Cfg
// }

// //section可为空时，则没有section
// func (c *Confile) String(section, key string) (value string) {
// 	return c.File.Section(section).Key(key).String()
// }

// //返回数组
// func (c *Confile) StringArray(section string) (keys []string) {
// 	return c.File.Section(section).KeyStrings()
// }
// func (c.Confile)
