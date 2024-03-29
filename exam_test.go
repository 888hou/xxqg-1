package main

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/chromedp/cdproto/cdp"
)

func TestCalcAnswers(t *testing.T) {
	testcase := []struct {
		answers []*cdp.Node
		tips    []*cdp.Node
		res     []string
	}{
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "大城市"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "农村"},
			},
			[]*cdp.Node{{NodeValue: "大城市"}, {NodeValue: "农村"}},
			[]string{"大城市", "农村"},
		},
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "大城市"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "农村"},
			},
			[]*cdp.Node{{NodeValue: "大"}, {NodeValue: "城"}, {NodeValue: "市"}, {NodeValue: "农村"}},
			[]string{"大城市", "农村"},
		},
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "大城市 农村"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "农村 大城市"},
			},
			[]*cdp.Node{{NodeValue: "大城市"}, {NodeValue: "农村"}},
			[]string{"大城市 农村"},
		},
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "大城市 农村"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "农村 大城市"},
			},
			[]*cdp.Node{{NodeValue: "大"}, {NodeValue: "城"}, {NodeValue: "市"}, {NodeValue: "农村"}},
			[]string{"大城市 农村"},
		},
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "线上"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "线下"},
				{NodeValue: "C"},
				{NodeValue: "."},
				{NodeValue: "线上线下同步"},
			},
			[]*cdp.Node{{NodeValue: "线上线下同步"}},
			[]string{"线上线下同步"},
		},
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "正确"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "错误"},
			},
			[]*cdp.Node{{NodeValue: "线上线下同步"}},
			nil,
		},
		{
			[]*cdp.Node{
				{NodeValue: "A"},
				{NodeValue: "."},
				{NodeValue: "正确说"},
				{NodeValue: "B"},
				{NodeValue: "."},
				{NodeValue: "错误说"},
			},
			[]*cdp.Node{{NodeValue: "正确"}},
			[]string{"正确说"},
		},
	}

	for _, tc := range testcase {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		res, err := calcAnswers(ctx, tc.answers, tc.tips)
		if err != nil {
			t.Error(err)
		}
		if !reflect.DeepEqual(tc.res, res) {
			t.Errorf("expected %q; got %q", tc.res, res)
		}
		cancel()
	}
}
