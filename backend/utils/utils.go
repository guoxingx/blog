package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// ReadConfig read config params from configPath into given instance with json decode.
func ReadConfig(configPath string, config interface{}) error {
	absPath, err := filepath.Abs(configPath)
	if err != nil {
		return err
	}

	file, err := os.Open(absPath)
	if err != nil {
		return err
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	if err := jsonParser.Decode(config); err != nil {
		return err
	}
	return nil
}

// MustParseDuration cause panic if parse duration failed
func MustParseDuration(s string) time.Duration {
	value, err := time.ParseDuration(s)
	if err != nil {
		panic("util: Can't parse duration `" + s + "`: " + err.Error())
	}
	return value
}

// ToInt parse string to int
func ToInt(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

// ParseTimeInToday ...
func ParseTimeInToday(s string) (time.Time, error) {
	dateS := time.Now().Format("2006-01-02")
	timeS := fmt.Sprintf("%s %s", dateS, s)
	t, err := time.ParseInLocation("2006-01-02 15:04", timeS, time.Local)
	return t, err
}

// ParseTimeInTodayWithDefault ...
func ParseTimeInTodayWithDefault(s, d string) time.Time {
	if s == "" {
		t, _ := ParseTimeInToday(d)
		return t
	}
	t, err := ParseTimeInToday(s)
	if err != nil {
		t, _ = ParseTimeInToday(d)
		return t
	}
	return t
}

// CheckFileExisted ...
func CheckFileExisted(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

// CreateFileReplace ...
func CreateFileReplace(absPath string) (*os.File, error) {
	if CheckFileExisted(absPath) {
		err := os.Remove(absPath)
		if err != nil {
			return nil, err
		}
	}
	return os.Create(absPath)
}

// DownloadFile ...
func DownloadFile(url, filename string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	file, err := CreateFileReplace(filename)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(file, res.Body)
	file.Close()
	if err != nil {
		return "", err
	}

	return filename, nil
}

// DownloadFileInRaw ...
func DownloadFileInRaw(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// DownloadFileInRawWithClient ...
func DownloadFileInRawWithClient(url string, c *http.Client) ([]byte, error) {
	request, _ := http.NewRequest("GET", url, nil)
	resp, err := c.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return data, fmt.Errorf("status code is %d instead of 200", resp.StatusCode)
	}

	return data, nil
}
