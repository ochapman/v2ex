/*
* File Name:	v2ex_test.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2014-10-19
 */
package v2ex_test

import (
	"github.com/ochapman/v2ex"
	"testing"
)

func TestStats(t *testing.T) {
	stats, err := v2ex.GetStats()
	if err != nil {
		t.Errorf("TestGetStats failed: %s\n", err)
		return
	}
	t.Logf("topic_max: %d, member_max: %d\n", stats.Topic_max, stats.Member_max)
}

func TestGetInfo(t *testing.T) {
	info, err := v2ex.GetInfo()
	if err != nil {
		t.Errorf("TestGetInfo failed: %s\n", err)
		return
	}
	t.Logf("title: %s\nslogan: %s\ndescription: %s\ndomain: %s",
		info.Title, info.Slogan, info.Description, info.Domain)
}

func TestGetNodeByID(t *testing.T) {
	node, err := v2ex.GetNodeByID(334)
	if err != nil {
		t.Errorf("TestGetInfo failed: %s\n", err)
		return
	}
	t.Logf("id: %d\nname: %s\nurl: %s\ntitle: %s\ntitle_alternative: %s"+
		"\ntopics: %d\nheader: %s\nfooter: %s\ncreated: %d\navatar_mini: %s"+
		"\navatar_normal: %s\navatar_large: %s",
		node.ID, node.Name, node.URL, node.Title, node.Title_alternative,
		node.Topics, node.Header, node.Footer, node.Created, node.Avatar_mini,
		node.Avatar_normal, node.Avatar_large)

}

func TestGetNodeByName(t *testing.T) {
	node, err := v2ex.GetNodeByName("linux")
	if err != nil {
		t.Errorf("TestGetInfo failed: %s\n", err)
		return
	}
	t.Logf("id: %d\nname: %s\nurl: %s\ntitle: %s\ntitle_alternative: %s"+
		"\ntopics: %d\nheader: %s\nfooter: %s\ncreated: %d\navatar_mini: %s"+
		"\navatar_normal: %s\navatar_large: %s",
		node.ID, node.Name, node.URL, node.Title, node.Title_alternative,
		node.Topics, node.Header, node.Footer, node.Created, node.Avatar_mini,
		node.Avatar_normal, node.Avatar_large)

}

func TestNodes(t *testing.T) {
	nodes, err := v2ex.GetNodes()
	if err != nil {
		t.Errorf("TestGetNodes failed: %s\n", err)
		return
	}
	for _, node := range nodes {
		t.Logf("%#v\n", node)
	}
}

func TestGetLatest(t *testing.T) {
	topics, err := v2ex.GetLatest()
	if err != nil {
		t.Errorf("TestGetLatest failed: %s\n", err)
		return
	}
	for _, topic := range topics {
		t.Logf("%#v\n", topic)
	}
}

func TestGetHot(t *testing.T) {
	topics, err := v2ex.GetHot()
	if err != nil {
		t.Errorf("TestGetHot failed: %s\n", err)
		return
	}
	for _, topic := range topics {
		t.Logf("%#v\n", topic)
	}
}

func TestGetTopicByID(t *testing.T) {
	topic, err := v2ex.GetTopicByID(123)
	if err != nil {
		t.Errorf("TestGetTopicByID failed: %s\n", err)
		return
	}
	topic_target := v2ex.Topic{
		ID:               123,
		Title:            "NoSQL 区",
		URL:              "http://www.v2ex.com/t/123",
		Content:          "欢迎在这里讨论所有的 NoSQL 技术。",
		Content_rendered: "欢迎在这里讨论所有的 NoSQL 技术。",
		Replies:          3,
		Member: v2ex.Member{
			ID:       1,
			Username: "Livid",
			Tagline:  "Beautifully Advance",
			Avatar: v2ex.Avatar{
				Avatar_mini:   "//cdn.v2ex.com/avatar/c4ca/4238/1_mini.png?m=1401650222",
				Avatar_normal: "//cdn.v2ex.com/avatar/c4ca/4238/1_normal.png?m=1401650222",
				Avatar_large:  "//cdn.v2ex.com/avatar/c4ca/4238/1_large.png?m=1401650222",
			},
		},
		Node: v2ex.Node{
			ID:     51,
			Name:   "nosql",
			Title:  "NoSQL",
			URL:    "http://www.v2ex.com/go/nosql",
			Topics: 13,
			Avatar: v2ex.Avatar{
				Avatar_mini:   "/static/img/node_mini.png",
				Avatar_normal: "/static/img/node_normal.png",
				Avatar_large:  "/static/img/node_large.png",
			},
		},
		Created:       1272356978,
		Last_modified: 1335108284,
		Last_touched:  1314771414,
	}
	if topic_target != topic {
		t.Error("TestGetTopicByID failed: unmatch\n")
		return
	}
}

func TestGetTopicsByUsername(t *testing.T) {
	topics, err := v2ex.GetTopicsByUsername("ochapman")
	if err != nil {
		t.Error("TestGetTopicsByUsername failed: %s\n", err)
		return
	}
	for _, topic := range topics {
		t.Logf("%#v\n", topic)
	}
}

func TestGetTopicsByNodename(t *testing.T) {
	topics, err := v2ex.GetTopicsByNodename("go")
	if err != nil {
		t.Error("TestGetTopicsByNodeName failed: %s\n", err)
		return
	}
	for _, topic := range topics {
		t.Logf("%#v\n", topic)
	}
}

func TestGetTopicsByNodeID(t *testing.T) {
	topics, err := v2ex.GetTopicsByNodeID(375)
	if err != nil {
		t.Error("TestGetTopicsByNodeID failed: %s\n", err)
		return
	}
	for _, topic := range topics {
		t.Logf("%#v\n", topic)
	}
}

func TestGetRepliesByTopicID(t *testing.T) {
	replies, err := v2ex.GetRepliesByTopicID(123)
	if err != nil {
		t.Error("TestGetRepliesByTopicID failed: %s\n", err)
		return
	}

	replies_target := v2ex.Replies{
		{
			ID:               587,
			Thanks:           0,
			Content:          "我觉得应该加上板块说明/简介功能",
			Content_rendered: "我觉得应该加上板块说明/简介功能",
			Member: v2ex.Member{
				ID:       7,
				Username: "Sai",
				Tagline:  "いのち短し 恋せよ乙女",
				Avatar: v2ex.Avatar{
					Avatar_mini:   "//cdn.v2ex.com/avatar/8f14/e45f/7_mini.png?m=1334913596",
					Avatar_normal: "//cdn.v2ex.com/avatar/8f14/e45f/7_normal.png?m=1334913596",
					Avatar_large:  "//cdn.v2ex.com/avatar/8f14/e45f/7_large.png?m=1334913596",
				},
			},
			Created:       1272368938,
			Last_modified: 1335092184,
		},

		{
			ID:               590,
			Thanks:           0,
			Content:          "嗯。\u000D\u000A\u000D\u000Anode.header 和 node.footer 目前在后台里有，只是前台还没有加上。",
			Content_rendered: "嗯。\u003Cbr /\u003E\u003Cbr /\u003Enode.header 和 node.footer 目前在后台里有，只是前台还没有加上。",
			Member: v2ex.Member{
				ID:       1,
				Username: "Livid",
				Tagline:  "Beautifully Advance",
				Avatar: v2ex.Avatar{
					Avatar_mini:   "//cdn.v2ex.com/avatar/c4ca/4238/1_mini.png?m=1401650222",
					Avatar_normal: "//cdn.v2ex.com/avatar/c4ca/4238/1_normal.png?m=1401650222",
					Avatar_large:  "//cdn.v2ex.com/avatar/c4ca/4238/1_large.png?m=1401650222",
				},
			},
			Created:       1272369274,
			Last_modified: 1335092197,
		},

		{
			ID:               137613,
			Thanks:           0,
			Content:          "开始研究下mangoDB",
			Content_rendered: "开始研究下mangoDB",
			Member: v2ex.Member{
				ID:       10911,
				Username: "vven",
				Tagline:  "举头望明月，低头敲代码。",
				Avatar: v2ex.Avatar{
					Avatar_mini:   "//cdn.v2ex.com/avatar/6bd2/ab1d/10911_mini.png?m=1335099920",
					Avatar_normal: "//cdn.v2ex.com/avatar/6bd2/ab1d/10911_normal.png?m=1335099920",
					Avatar_large:  "//cdn.v2ex.com/avatar/6bd2/ab1d/10911_large.png?m=1335099920",
				},
			},
			Created:       1314685014,
			Last_modified: 1335092195,
		},
	}

	if len(replies) != len(replies_target) {
		t.Error("TestGetRepliesByTopicID failed: len is not equal")
		return
	}

	for i := 0; i < len(replies); i++ {
		if replies[i] != replies_target[i] {
			t.Error("TestGetRepliesByTopicID failed: unmatch\n")
			return
		}
	}
}

func TestGetMemberByID(t *testing.T) {
	member, err := v2ex.GetMemberByID(123)
	if err != nil {
		t.Errorf("TestGetMemberByID failed: %s\n", err)
		return
	}

	member_target := v2ex.Member{
		Status:   "found",
		ID:       123,
		URL:      "http://www.v2ex.com/member/romoo",
		Username: "romoo",
		Website:  "twitter.com/romoo",
		Twitter:  "romoo",
		Psn:      "",
		Github:   "romoo",
		Btc:      "1NKcbDqhCrK7WXTww6ovrma4AB5h36JJPQ",
		Location: "Beijing, China",
		Tagline:  "Aha",
		Bio:      "",
		Avatar: v2ex.Avatar{
			Avatar_mini:   "//cdn.v2ex.com/avatar/202c/b962/123_mini.png?m=1335117969",
			Avatar_normal: "//cdn.v2ex.com/avatar/202c/b962/123_normal.png?m=1335117969",
			Avatar_large:  "//cdn.v2ex.com/avatar/202c/b962/123_large.png?m=1335117969",
		},
		Created: 1272264436,
	}

	if member_target != member {
		t.Error("TestGetMemberByID failed: not match")
		return
	}
}

func TestGetMemberByUsername(t *testing.T) {
	member, err := v2ex.GetMemberByUsername("ochapman")
	if err != nil {
		t.Errorf("TestGetMemberByUsername failed: %s\n", err)
		return
	}

	member_target := v2ex.Member{
		Status:   "found",
		ID:       59852,
		URL:      "http://www.v2ex.com/member/ochapman",
		Username: "ochapman",
		Website:  "",
		Twitter:  "",
		Psn:      "",
		Github:   "",
		Btc:      "",
		Location: "",
		Tagline:  "",
		Bio:      "",
		Avatar: v2ex.Avatar{
			Avatar_mini:   "https://secure.gravatar.com/avatar/ddfc45c85ffa86cc0afff074b57df297?s=24&d=https%3A%2F%2Fcdn.v2ex.com%2Fstatic%2Fimg%2Favatar_mini.png",
			Avatar_normal: "https://secure.gravatar.com/avatar/ddfc45c85ffa86cc0afff074b57df297?s=48&d=https%3A%2F%2Fcdn.v2ex.com%2Fstatic%2Fimg%2Favatar_normal.png",
			Avatar_large:  "https://secure.gravatar.com/avatar/ddfc45c85ffa86cc0afff074b57df297?s=73&d=https%3A%2F%2Fcdn.v2ex.com%2Fstatic%2Fimg%2Favatar_large.png",
		},
		Created: 1396860133,
	}

	if member != member_target {
		t.Error("TestGetMemberByUsername failed, not match\n")
		return
	}
}
