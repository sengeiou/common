// Copyright 2019 gf Author(https://github.com/qnsoft/common). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/qnsoft/common.

// go test *.go -bench=".*"

package qn_page_test

import (
	"testing"

	"github.com/qnsoft/common/util/qn_page"

	"github.com/qnsoft/common/test/qn_test"
)

func Test_New(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(9, 2, 1, `/user/list?page={.page}`)
		t.Assert(page.TotalSize, 9)
		t.Assert(page.TotalPage, 5)
		t.Assert(page.CurrentPage, 1)
	})
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(9, 2, 0, `/user/list?page={.page}`)
		t.Assert(page.TotalSize, 9)
		t.Assert(page.TotalPage, 5)
		t.Assert(page.CurrentPage, 1)
	})
}

func Test_Basic(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(9, 2, 1, `/user/list?page={.page}`)
		t.Assert(page.NextPage(), `<a class="qn_pageLink" href="/user/list?page=2" title="">></a>`)
		t.Assert(page.PrevPage(), `<span class="qn_pageSpan"><</span>`)
		t.Assert(page.FirstPage(), `<span class="qn_pageSpan">|<</span>`)
		t.Assert(page.LastPage(), `<a class="qn_pageLink" href="/user/list?page=5" title="">>|</a>`)
		t.Assert(page.PageBar(), `<span class="qn_pageSpan">1</span><a class="qn_pageLink" href="/user/list?page=2" title="2">2</a><a class="qn_pageLink" href="/user/list?page=3" title="3">3</a><a class="qn_pageLink" href="/user/list?page=4" title="4">4</a><a class="qn_pageLink" href="/user/list?page=5" title="5">5</a>`)
	})

	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(9, 2, 3, `/user/list?page={.page}`)
		t.Assert(page.NextPage(), `<a class="qn_pageLink" href="/user/list?page=4" title="">></a>`)
		t.Assert(page.PrevPage(), `<a class="qn_pageLink" href="/user/list?page=2" title=""><</a>`)
		t.Assert(page.FirstPage(), `<a class="qn_pageLink" href="/user/list?page=1" title="">|<</a>`)
		t.Assert(page.LastPage(), `<a class="qn_pageLink" href="/user/list?page=5" title="">>|</a>`)
		t.Assert(page.PageBar(), `<a class="qn_pageLink" href="/user/list?page=1" title="1">1</a><a class="qn_pageLink" href="/user/list?page=2" title="2">2</a><span class="qn_pageSpan">3</span><a class="qn_pageLink" href="/user/list?page=4" title="4">4</a><a class="qn_pageLink" href="/user/list?page=5" title="5">5</a>`)
	})

	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(9, 2, 5, `/user/list?page={.page}`)
		t.Assert(page.NextPage(), `<span class="qn_pageSpan">></span>`)
		t.Assert(page.PrevPage(), `<a class="qn_pageLink" href="/user/list?page=4" title=""><</a>`)
		t.Assert(page.FirstPage(), `<a class="qn_pageLink" href="/user/list?page=1" title="">|<</a>`)
		t.Assert(page.LastPage(), `<span class="qn_pageSpan">>|</span>`)
		t.Assert(page.PageBar(), `<a class="qn_pageLink" href="/user/list?page=1" title="1">1</a><a class="qn_pageLink" href="/user/list?page=2" title="2">2</a><a class="qn_pageLink" href="/user/list?page=3" title="3">3</a><a class="qn_pageLink" href="/user/list?page=4" title="4">4</a><span class="qn_pageSpan">5</span>`)
	})
}

func Test_CustomTag(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(5, 1, 2, `/user/list/{.page}`)
		page.PrevPageTag = "《"
		page.NextPageTag = "》"
		page.FirstPageTag = "|《"
		page.LastPageTag = "》|"
		page.PrevBarTag = "《《"
		page.NextBarTag = "》》"
		t.Assert(page.NextPage(), `<a class="qn_pageLink" href="/user/list/3" title="">》</a>`)
		t.Assert(page.PrevPage(), `<a class="qn_pageLink" href="/user/list/1" title="">《</a>`)
		t.Assert(page.FirstPage(), `<a class="qn_pageLink" href="/user/list/1" title="">|《</a>`)
		t.Assert(page.LastPage(), `<a class="qn_pageLink" href="/user/list/5" title="">》|</a>`)
		t.Assert(page.PageBar(), `<a class="qn_pageLink" href="/user/list/1" title="1">1</a><span class="qn_pageSpan">2</span><a class="qn_pageLink" href="/user/list/3" title="3">3</a><a class="qn_pageLink" href="/user/list/4" title="4">4</a><a class="qn_pageLink" href="/user/list/5" title="5">5</a>`)
	})
}

func Test_CustomStyle(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(5, 1, 2, `/user/list/{.page}`)
		page.LinkStyle = "MyPageLink"
		page.SpanStyle = "MyPageSpan"
		page.SelectStyle = "MyPageSelect"
		t.Assert(page.NextPage(), `<a class="MyPageLink" href="/user/list/3" title="">></a>`)
		t.Assert(page.PrevPage(), `<a class="MyPageLink" href="/user/list/1" title=""><</a>`)
		t.Assert(page.FirstPage(), `<a class="MyPageLink" href="/user/list/1" title="">|<</a>`)
		t.Assert(page.LastPage(), `<a class="MyPageLink" href="/user/list/5" title="">>|</a>`)
		t.Assert(page.PageBar(), `<a class="MyPageLink" href="/user/list/1" title="1">1</a><span class="MyPageSpan">2</span><a class="MyPageLink" href="/user/list/3" title="3">3</a><a class="MyPageLink" href="/user/list/4" title="4">4</a><a class="MyPageLink" href="/user/list/5" title="5">5</a>`)
		t.Assert(page.SelectBar(), `<select name="MyPageSelect" onchange="window.location.href=this.value"><option value="/user/list/1">1</option><option value="/user/list/2" selected>2</option><option value="/user/list/3">3</option><option value="/user/list/4">4</option><option value="/user/list/5">5</option></select>`)
	})
}

func Test_Ajax(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(5, 1, 2, `/user/list/{.page}`)
		page.AjaxActionName = "LoadPage"
		t.Assert(page.NextPage(), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="">></a>`)
		t.Assert(page.PrevPage(), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title=""><</a>`)
		t.Assert(page.FirstPage(), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">|<</a>`)
		t.Assert(page.LastPage(), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="">>|</a>`)
		t.Assert(page.PageBar(), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="1">1</a><span class="qn_pageSpan">2</span><a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="3">3</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/4')" title="4">4</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="5">5</a>`)
	})
}

func Test_PredefinedContent(t *testing.T) {
	qn_test.C(t, func(t *qn_test.T) {
		page := qn_page.New(5, 1, 2, `/user/list/{.page}`)
		page.AjaxActionName = "LoadPage"
		t.Assert(page.GetContent(1), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">上一页</a> <span class="current">2</span> <a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页</a>`)
		t.Assert(page.GetContent(2), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">首页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title=""><<上一页</a><span class="current">[第2页]</span><a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页>></a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="">尾页</a>第<select name="qn_pageSelect" onchange="window.location.href=this.value"><option value="/user/list/1">1</option><option value="/user/list/2" selected>2</option><option value="/user/list/3">3</option><option value="/user/list/4">4</option><option value="/user/list/5">5</option></select>页`)
		t.Assert(page.GetContent(3), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">首页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">上一页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="1">1</a><span class="qn_pageSpan">2</span><a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="3">3</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/4')" title="4">4</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="5">5</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="">尾页</a><span>当前页2/5</span> <span>共5条</span>`)
		t.Assert(page.GetContent(4), `<a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">首页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="">上一页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/1')" title="1">1</a><span class="qn_pageSpan">2</span><a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="3">3</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/4')" title="4">4</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="5">5</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/3')" title="">下一页</a><a class="qn_pageLink" href="javascript:LoadPage('/user/list/5')" title="">尾页</a>`)
		t.Assert(page.GetContent(5), ``)
	})
}
