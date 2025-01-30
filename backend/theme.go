package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
)

type ThemeSummary struct {
	ID       string `json:"id"`
	Author   string `json:"author"`
	Name     string `json:"name"`
	FileName string `json:"fileName"`
	Hash     string `json:"hash"`
}

type Theme struct {
	ID              string       `json:"id"`
	Name            string       `json:"name"`
	Author          string       `json:"author"`
	BackgroundColor string       `json:"backgroundColor"`
	TextColor       string       `json:"textColor"`
	Grid            Grid         `json:"grid"`
	Field           Field        `json:"field"`
	Button          ButtonColors `json:"button"`
	Sidebar         Sidebar      `json:"sidebar"`
	Input           Input        `json:"input"`
}

func (th *Theme) GetHash() string {
	data, _ := json.Marshal(th)
	hash := sha1.Sum(data)
	return hex.EncodeToString(hash[:])
}

type Field struct {
	Default ButtonColors `json:"default"`
	Checked ButtonColors `json:"checked"`
	Gay     ButtonColors `json:"gay"`
}

type ButtonColors struct {
	Color            string `json:"color"`
	HoverColor       string `json:"hoverColor"`
	TextOutlineColor string `json:"textOutlineColor,omitempty"`
	TextColor        string `json:"textColor,omitempty"`
}

type Grid struct {
	Color string `json:"color"`
}

type Input struct {
	Color          string `json:"color"`
	ActiveColor    string `json:"activeColor"`
	IndicatorColor string `json:"indicatorColor"`
}

type Sidebar struct {
	Color          string `json:"color"`
	BackgroundBlur string `json:"backgroundBlur"`
}
