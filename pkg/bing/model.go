package bing

import "time"

type WallpaperData struct {
	Images   []Image  `json:"images"`
	Tooltips Tooltips `json:"tooltips"`
}

type Image struct {
	StartDate     Date          `json:"startdate"`
	FullStartDate Date          `json:"fullstartdate"`
	EndDate       Date          `json:"enddate"`
	URL           string        `json:"url"`
	UrlBase       string        `json:"urlbase"`
	Copyright     string        `json:"copyright"`
	CopyrightLink string        `json:"copyrightlink"`
	Title         string        `json:"title"`
	Quiz          string        `json:"quiz"`
	Wp            bool          `json:"wp"`
	Hsh           string        `json:"hsh"`
	Drk           int64         `json:"drk"`
	Top           int64         `json:"top"`
	Bot           int64         `json:"bot"`
	Hs            []interface{} `json:"hs"`
}

type Tooltips struct {
	Loading  string `json:"loading"`
	Previous string `json:"previous"`
	Next     string `json:"next"`
	WallErr  string `json:"walle"`
	Walls    string `json:"walls"`
}

type Date string

func (d Date) Time() time.Time {
	t, _ := time.Parse("20060102", string(d))
	return t
}
