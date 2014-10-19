/*
* File Name:	v2ex.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-10-19
 */

// http://www.v2ex.com/t/85402
package v2ex

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func get(url string, v interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(v)
	return
}

type stats struct {
	Topic_max  uint32 `json:topic_max`
	Member_max uint32 `json:topic_max`
}

func Stats() (stats stats, err error) {
	url := "http://www.v2ex.com/api/site/stats.json"
	err = get(url, &stats)
	return
}

type info struct {
	Title       string `json:title`
	Slogan      string `json:slogan`
	Description string `json:description`
	Domain      string `json:domain`
}

func Info() (info info, err error) {
	url := "http://www.v2ex.com/api/site/info.json"
	err = get(url, &info)
	return
}

type node struct {
	ID                uint32 `json:id`
	Name              string `json:name`
	URL               string `json:url`
	Title             string `json:title`
	Title_alternative string `json:title_alternative`
	Topics            uint32 `json:topics`
	Header            string `json:header`
	Footer            string `json:footer`
	Created           int64  `json:created`
	Avatar_mini       string `json:avatar_mini`
	Avatar_normal     string `json:avatar_normal`
	Avatar_large      string `json:avatar_large`
}

func NodeByID(id uint32) (node node, err error) {
	url := "http://www.v2ex.com/api/nodes/show.json?id=" + strconv.Itoa(int(id))
	err = get(url, &node)
	return
}

func NodeByName(name string) (node node, err error) {
	url := "http://www.v2ex.com/api/nodes/show.json?name=" + name
	err = get(url, &node)
	return
}

type nodes []node

func Nodes() (nodes nodes, err error) {
	url := "http://www.v2ex.com/api/nodes/all.json"
	err = get(url, &nodes)
	return
}
