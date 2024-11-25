package utils

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type Properties struct {
	propertyMap map[string]string
}

func NewProperties(propertyMap map[string]string) (this *Properties) {
	this = new(Properties)
	this.propertyMap = propertyMap
	return
}
func (this *Properties) GetPropertyString(propertyKey string) (retValue string, err error) {
	if propertyStr, ok := this.propertyMap[strings.ToUpper(propertyKey)]; ok {
		retValue = propertyStr
		return
	}
	err = errors.New("property " + propertyKey + " is not exist")
	return
}

func (this *Properties) GetPropertyInt(propertyKey string) (retValue int, err error) {
	if propertyStr, ok := this.propertyMap[strings.ToUpper(propertyKey)]; ok {
		propertyInt, parseErr := strconv.Atoi(propertyStr)
		if parseErr != nil {
			err = parseErr
			return
		}
		retValue = propertyInt
		return
	}
	err = errors.New("property " + propertyKey + " is not exist")
	return
}

func (this *Properties) GetPropertyInt32(propertyKey string) (retValue int32, err error) {
	if propertyStr, ok := this.propertyMap[strings.ToUpper(propertyKey)]; ok {
		propertyInt, parseErr := strconv.Atoi(propertyStr)
		if parseErr != nil {
			err = parseErr
			return
		}
		retValue = int32(propertyInt)
		return
	}
	err = errors.New("property " + propertyKey + " is not exist")
	return
}

func GetPropertiesPath(propertiesFile string) (propertiesPath string, err error) {
	currentDir, osErr := os.Getwd()
	if osErr != nil {
		err = osErr
		return
	}
	PropertiesDir := currentDir
	if "windows" == runtime.GOOS {
		PropertiesDir = PropertiesDir + "\\properties\\"
	} else {
		PropertiesDir = "/usr/properties/"
	}
	mkErr := os.MkdirAll(PropertiesDir, 0777)
	if mkErr != nil {
		err = mkErr
		return
	}
	propertiesPath = PropertiesDir + propertiesFile

	return
}
func ParseProperties(configFile string) (properties *Properties, err error) {
	if IsFileExist(configFile) {
		f, openFileErr := os.Open(configFile)
		if openFileErr != nil {
			err = openFileErr
			return
		}
		defer f.Close()
		br := bufio.NewReader(f)
		propertyMap := make(map[string]string)
		for {
			line, readLineErr := br.ReadString('\n')
			if readLineErr == nil {
				strArr := strings.Split(strings.Trim(strings.Trim(strings.Trim(line, "\n"), "\r"), " "), "=")
				if len(strArr) == 2 {
					propertyMap[strings.ToUpper(strings.Trim(strArr[0], " "))] = strings.Trim(strArr[1], " ")
				}
			} else if readLineErr == io.EOF {
				strArr := strings.Split(strings.Trim(strings.Trim(strings.Trim(line, "\n"), "\r"), " "), "=")
				if len(strArr) == 2 {
					propertyMap[strings.ToUpper(strings.Trim(strArr[0], " "))] = strings.Trim(strArr[1], " ")
				}
				break
			} else {
				log.Fatalln("line:", line, "readLineErr:", readLineErr)
				err = openFileErr
				return
			}
		}
		properties = NewProperties(propertyMap)
		return
	}
	err = errors.New("file " + configFile + "not exist")
	return
}
