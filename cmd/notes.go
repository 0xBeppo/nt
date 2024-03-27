package cmd

type MyNote struct {
	Title string
	Date  string
	Tags  []string
}

var tpl = `
---
title: {{ .Title}}
date: {{ .Date }}
tags: 
---
`
