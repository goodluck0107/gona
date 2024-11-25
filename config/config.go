package config

import (
	"bytes"
	"fmt"
	"path/filepath"
	"regexp"

	"gitee.com/andyxt/gona/cast"

	"github.com/spf13/viper"
)

func New() *Config {
	c := &Config{
		Viper: viper.New(),
	}
	return c
}

type Reader interface {
	Names() []string
	Bytes(name string) ([]byte, error)
}

type Config struct {
	*viper.Viper
	runtimeEnv string
}

func (c *Config) SetRuntimeEnv(s string) *Config {
	c.runtimeEnv = s
	return c
}

// regexp to match ${} or $()
var re = regexp.MustCompile(`\$\{([^}]+)\}|\$\(([^\)]+)\)`)

func (c *Config) Parse(s string) string {
	return re.ReplaceAllStringFunc(s, func(v string) string {
		k := v[2 : len(v)-1]
		return cast.ToString(c.Get(k))
	})
}

func (c *Config) ParseDepth(depth int, str string) string {
	switch {
	case depth > 0:
		for i := 0; i < depth; i++ {
			newStr := c.Parse(str)
			if newStr == str {
				return str
			}
			str = newStr
		}
		return str
	case depth < 0:
		for {
			newStr := c.Parse(str)
			if newStr == str {
				return str
			}
			str = newStr
		}
	default:
		return str
	}
}

func (c *Config) Read(b Reader) *Config {
	v := viper.New()
	for _, name := range b.Names() {
		extWithPoint := filepath.Ext(name)
		if extWithPoint == "" {
			continue
		}
		ext := extWithPoint[1:]
		if !c.isExtValid(ext) {
			continue
		}
		v.SetConfigType(ext)
		data, err := b.Bytes(name)
		if err != nil {
			panic(fmt.Errorf("name:%v,error:%v", name, err))
		}
		reader := bytes.NewReader(data)
		if err := v.MergeConfig(reader); err != nil {
			panic(fmt.Errorf("name:%v,error:%v", name, err))
		}
	}

	// merge runtime config
	if c.runtimeEnv != "" {
		runtimeViper := v.Sub(fmt.Sprintf("<%s>", c.runtimeEnv))
		if runtimeViper != nil {
			if err := v.MergeConfigMap(runtimeViper.AllSettings()); err != nil {
				panic(err)
			}
		}
	}

	settings := v.AllSettings()
	for k := range settings {
		if k[0] == '<' && k[len(k)-1] == '>' {
			delete(settings, k)
		}
	}
	if err := c.MergeConfigMap(settings); err != nil {
		panic(err)
	}

	return c
}

func (c *Config) isExtValid(ext string) bool {
	for _, e := range viper.SupportedExts {
		if ext == e {
			return true
		}
	}
	return false
}

func (c *Config) ReadBinary(names func() []string, bytes func(string) ([]byte, error)) *Config {
	c.Read(NewBinaryReader(names, bytes))
	return c
}
