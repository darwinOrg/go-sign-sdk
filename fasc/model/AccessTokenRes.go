package model

type AccessTokenResData struct {
	AccessToken string `json:"accessToken"`
	ExpiresIn   string `json:"expiresIn"`
}

type AccessTokenRes struct {
	Code string             `json:"code"`
	Msg  string             `json:"msg"`
	Data AccessTokenResData `json:"data"`
}
