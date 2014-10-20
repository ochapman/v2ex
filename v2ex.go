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
	"errors"
	"net/http"
	"strconv"
)

var (
	ErrInvalidUsername = errors.New("v2ex: invalid username")
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

type Node struct {
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

func NodeByID(id uint32) (node Node, err error) {
	url := "http://www.v2ex.com/api/nodes/show.json?id=" + strconv.Itoa(int(id))
	err = get(url, &node)
	return
}

func NodeByName(name string) (node Node, err error) {
	url := "http://www.v2ex.com/api/nodes/show.json?name=" + name
	err = get(url, &node)
	return
}

type nodes []Node

func Nodes() (nodes nodes, err error) {
	url := "http://www.v2ex.com/api/nodes/all.json"
	err = get(url, &nodes)
	return
}

type Avatar struct {
	Avatar_mini   string `json:avatar_mini`
	Avatar_normal string `json:avatar_normal`
	Avatar_large  string `json:avatar_large`
}

type Member struct {
	Status   string `json:status`
	ID       uint32 `json:id`
	URL      string `json:url`
	Username string `json:username`
	Website  string `json:website`
	Twitter  string `json:twitter`
	Psn      string `json:psn`
	Github   string `json:github`
	Btc      string `json:btc`
	Location string `json:location`
	Tagline  string `json:tagline`
	Bio      string `json:bio`
	Avatar
	Created int64 `json:created`
}

func MemberByID(id uint32) (member Member, err error) {
	url := "http://www.v2ex.com/api/members/show.json?id=" + strconv.Itoa(int(id))
	err = get(url, &member)
	return
}

func MemberByUsername(name string) (member Member, err error) {
	if name == "" {
		return Member{}, ErrInvalidUsername
	}
	url := "http://www.v2ex.com/api/members/show.json?username=" + name
	err = get(url, &member)
	return
}

type Topic struct {
	ID               uint32 `json:id`
	Title            string `json:title`
	URL              string `json:url`
	Content          string `json:content`
	Content_rendered string `json:content_rendered`
	Replies          uint32 `json:replies`
	Member           Member `json:member`
	Node             Node   `json:node`
	Created          uint64 `json:created`
	Last_modified    uint64 `json:last_modified`
	Last_touched     uint64 `json:last_touched`
}

type Topics []Topic

func Latest() (topics Topics, err error) {
	url := "http://www.v2ex.com/api/topics/latest.json"
	err = get(url, &topics)
	return
}

//相当于首页右侧的 10 大每天的内容
func Hot() (topics Topics, err error) {
	url := "https://www.v2ex.com/api/topics/hot.json"
	err = get(url, &topics)
	return
}

func TopicByID(id uint32) (topic Topic, err error) {
	var topics Topics
	url := "http://www.v2ex.com/api/topics/show.json?id=" + strconv.Itoa(int(id))
	err = get(url, &topics)
	topic = topics[0]
	return
}

type Reply struct {
	ID               uint32 `json:id` //Reply ID
	Thanks           uint32 `json:thanks`
	Content          string `json:content`
	Content_rendered string `json:content_rendered`
	Member           Member `json:member`
	Created          int64  `json:created`
	Last_modified    int64  `json:last_modified`
}

type Replies []Reply

// id: topic ID
func RepliesByTopicID(id uint32) (replies Replies, err error) {
	url := "http://www.v2ex.com/api/replies/show.json?topic_id=" + strconv.Itoa(int(id))
	err = get(url, &replies)
	return
}
