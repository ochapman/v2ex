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
	stats, err := v2ex.Stats()
	if err != nil {
		t.Errorf("TestStats failed: %s\n", err)
		return
	}
	t.Logf("topic_max: %d, member_max: %d\n", stats.Topic_max, stats.Member_max)
}

func TestInfo(t *testing.T) {
	info, err := v2ex.Info()
	if err != nil {
		t.Errorf("TestInfo failed: %s\n", err)
		return
	}
	t.Logf("title: %s\nslogan: %s\ndescription: %s\ndomain: %s",
		info.Title, info.Slogan, info.Description, info.Domain)
}

func TestNodeByID(t *testing.T) {
	node, err := v2ex.NodeByID(334)
	if err != nil {
		t.Errorf("TestInfo failed: %s\n", err)
		return
	}
	t.Logf("id: %d\nname: %s\nurl: %s\ntitle: %s\ntitle_alternative: %s"+
		"\ntopics: %d\nheader: %s\nfooter: %s\ncreated: %d\navatar_mini: %s"+
		"\navatar_normal: %s\navatar_large: %s",
		node.ID, node.Name, node.URL, node.Title, node.Title_alternative,
		node.Topics, node.Header, node.Footer, node.Created, node.Avatar_mini,
		node.Avatar_normal, node.Avatar_large)

}

func TestNodeByName(t *testing.T) {
	node, err := v2ex.NodeByName("linux")
	if err != nil {
		t.Errorf("TestInfo failed: %s\n", err)
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
	nodes, err := v2ex.Nodes()
	if err != nil {
		t.Errorf("TestNodes failed: %s\n", err)
		return
	}
	for _, node := range nodes {
		t.Logf("%#v\n", node)
	}
}
