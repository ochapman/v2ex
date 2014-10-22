/*
* File Name:	v2ex.go
* Description:  www.v2ex.com API
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-10-19
 */

// https://www.v2ex.com/p/7v9TEc53
// http://www.v2ex.com/t/85402

package v2ex

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var (
	ErrInvalidUsername = errors.New("v2ex: invalid user name")
	ErrInvalidNodename = errors.New("v2ex: invalid node name")
)

const v2exAPI = "https://www.v2ex.com/api"

func getAPIData(url string, v interface{}) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(v)
	return
}

type Stats struct {
	Topic_max  uint32 `json:topic_max` //当前社区话题数量
	Member_max uint32 `json:topic_max` //当前社区用户数量
}

//获取社区统计信息
func GetStats() (stats Stats, err error) {
	url := v2exAPI + "/site/stats.json"
	err = getAPIData(url, &stats)
	return
}

type Info struct {
	Title       string `json:title`       //当前社区站名
	Slogan      string `json:slogan`      //当前社区口号
	Description string `json:description` //当前社区描述
	Domain      string `json:domain`      //社区网址
}

//获取社区介绍
func GetInfo() (info Info, err error) {
	url := v2exAPI + "/site/info.json"
	err = getAPIData(url, &info)
	return
}

type Node struct {
	ID                uint32 `json:id`                //节点 ID
	Name              string `json:name`              //节点缩略名
	URL               string `json:url`               //节点地址
	Title             string `json:title`             //节点名称
	Title_alternative string `json:title_alternative` //备选节点名称
	Topics            uint32 `json:topics`            //节点主题总数
	Header            string `json:header`            //节点头部信息
	Footer            string `json:footer`            //节点脚部信息
	Created           int64  `json:created`           //节点创建时间
	Avatar
}

//通过节点ID获取单个节点信息
func GetNodeByID(id uint32) (node Node, err error) {
	url := v2exAPI + "/nodes/show.json?id=" + strconv.Itoa(int(id))
	err = getAPIData(url, &node)
	return
}

//通过节点名字获取单个节点信息
func GetNodeByName(name string) (node Node, err error) {
	if name == "" {
		return node, ErrInvalidUsername
	}
	url := v2exAPI + "/nodes/show.json?name=" + name
	err = getAPIData(url, &node)
	return
}

type Nodes []Node

//获取全部节点
func GetNodes() (nodes Nodes, err error) {
	url := v2exAPI + "/nodes/all.json"
	err = getAPIData(url, &nodes)
	return
}

//
type Avatar struct {
	Avatar_mini   string `json:avatar_mini`
	Avatar_normal string `json:avatar_normal`
	Avatar_large  string `json:avatar_large`
}

//用户的自我介绍，及其登记的社交网络信息
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

//通过用户的ID获取
func GetMemberByID(id uint32) (member Member, err error) {
	url := v2exAPI + "/members/show.json?id=" + strconv.Itoa(int(id))
	err = getAPIData(url, &member)
	return
}

//通过用户的名字获取
func GetMemberByUsername(name string) (member Member, err error) {
	if name == "" {
		return member, ErrInvalidUsername
	}
	url := v2exAPI + "/members/show.json?username=" + name
	err = getAPIData(url, &member)
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

func GetLatest() (topics Topics, err error) {
	url := v2exAPI + "/topics/latest.json"
	err = getAPIData(url, &topics)
	return
}

//相当于首页右侧的 10 大每天的内容
func GetHot() (topics Topics, err error) {
	url := "https://www.v2ex.com/api/topics/hot.json"
	err = getAPIData(url, &topics)
	return
}

//通过id获取对应主题
func GetTopicByID(id uint32) (topic Topic, err error) {
	var topics Topics
	url := v2exAPI + "/topics/show.json?id=" + strconv.Itoa(int(id))
	err = getAPIData(url, &topics)
	topic = topics[0]
	return
}

//获取指定用户的所有主题
func GetTopicsByUsername(name string) (topics Topics, err error) {
	if name == "" {
		err = ErrInvalidUsername
		return
	}
	url := v2exAPI + "/topics/show.json?username=" + name
	err = getAPIData(url, &topics)
	return
}

//获取节点下所有主题
func GetTopicsByNodename(name string) (topics Topics, err error) {
	if name == "" {
		err = ErrInvalidNodename
		return
	}
	url := v2exAPI + "/topics/show.json?node_name=" + name
	err = getAPIData(url, &topics)
	return
}

//获取节点ID下所有主题
func GetTopicsByNodeID(nodeid uint32) (topics Topics, err error) {
	url := v2exAPI + "/topics/show.json?node_id=" + strconv.Itoa(int(nodeid))
	err = getAPIData(url, &topics)
	return
}

//回复
type Reply struct {
	ID               uint32 `json:id`     //Reply ID
	Thanks           uint32 `json:thanks` //感谢数量
	Content          string `json:content`
	Content_rendered string `json:content_rendered`
	Member           Member `json:member`
	Created          int64  `json:created`
	Last_modified    int64  `json:last_modified`
}

type Replies []Reply

// id: topic ID
func GetRepliesByTopicID(id uint32) (replies Replies, err error) {
	url := v2exAPI + "/replies/show.json?topic_id=" + strconv.Itoa(int(id))
	err = getAPIData(url, &replies)
	return
}
