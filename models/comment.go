package models

import (
	"NeteaseCloudMusic/utils"
	"net/http"
)

const (
	GetMusicComment    = `/weapi/v1/resource/comments/R_SO_4_`
	GetAlbumComment    = `/weapi/v1/resource/comments/R_AL_3_`
	GetPlaylistComment = `/weapi/v1/resource/comments/A_PL_0_`
	GetMVComment       = `/weapi/v1/resource/comments/R_MV_5_`
	GetDJComment       = `/weapi/v1/resource/comments/A_DJ_1_`
	LikeComment        = `/weapi/v1/comment/`
	PublishComment     = `/weapi/resource/comments/`
)

type ICommentOperator interface {
	GetMusicComment(rid string) interface{}
	GetAlbumComment(rid string) interface{}
	GetPlaylistComment(rid string) interface{}
	GetMVComment(rid string) interface{}
	GetDJComment(rid string) interface{}
	LikeComment(action string) interface{}
	Publish(threadId string, action string) interface{}
}

type BasicParams struct {
	ThreadId  string `json:"threadId"`
	CommentId string `json:"commentId"`
}

type Comment struct {
	Cookies []*http.Cookie
	PageLimitParams
	BasicParams
}

func (comment *Comment) GetMusicComment(rid string) interface{} {
	pageLimitParams, _ := utils.TransformStructToJSONMap(comment.PageLimitParams)
	reqParams := make(map[string]interface{})
	reqParams["rid"] = rid
	_reqParams := utils.MergeMaps(pageLimitParams, reqParams)
	response, _, _ := utils.NeteaseCloudRequest(GetMusicComment+rid+`/?csrf_token=`, _reqParams, comment.Cookies, http.MethodPost)
	return response
}
func (comment *Comment) GetAlbumComment(rid string) interface{} {
	pageLimitParams, _ := utils.TransformStructToJSONMap(comment.PageLimitParams)
	reqParams := make(map[string]interface{})
	reqParams["rid"] = rid
	_reqParams := utils.MergeMaps(pageLimitParams, reqParams)
	response, _, _ := utils.NeteaseCloudRequest(GetAlbumComment+rid+`/?csrf_token=`, _reqParams, comment.Cookies, http.MethodPost)
	return response
}
func (comment *Comment) GetPlaylistComment(rid string) interface{} {
	pageLimitParams, _ := utils.TransformStructToJSONMap(comment.PageLimitParams)
	reqParams := make(map[string]interface{})
	reqParams["rid"] = rid
	_reqParams := utils.MergeMaps(pageLimitParams, reqParams)
	response, _, _ := utils.NeteaseCloudRequest(GetPlaylistComment+rid+`/?csrf_token=`, _reqParams, comment.Cookies, http.MethodPost)
	return response
}
func (comment *Comment) GetMVComment(rid string) interface{} {
	pageLimitParams, _ := utils.TransformStructToJSONMap(comment.PageLimitParams)
	reqParams := make(map[string]interface{})
	reqParams["rid"] = rid
	_reqParams := utils.MergeMaps(pageLimitParams, reqParams)
	response, _, _ := utils.NeteaseCloudRequest(GetMVComment+rid+`/?csrf_token=`, _reqParams, comment.Cookies, http.MethodPost)
	return response
}
func (comment *Comment) GetDJComment(rid string) interface{} {
	pageLimitParams, _ := utils.TransformStructToJSONMap(comment.PageLimitParams)
	reqParams := make(map[string]interface{})
	reqParams["rid"] = rid
	_reqParams := utils.MergeMaps(pageLimitParams, reqParams)
	response, _, _ := utils.NeteaseCloudRequest(GetDJComment+rid+`/?csrf_token=`, _reqParams, comment.Cookies, http.MethodPost)
	return response
}
func (comment *Comment) LikeComment(action string) interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(comment.BasicParams)
	response, _, _ := utils.NeteaseCloudRequest(LikeComment+action, reqParams, comment.Cookies, http.MethodPost)
	return response
}
func (comment *Comment) Publish(content string, action string) interface{} {
	reqParams, _ := utils.TransformStructToJSONMap(comment.BasicParams)
	if content != "" {
		reqParams["content"] = content
	}
	response, _, _ := utils.NeteaseCloudRequest(PublishComment+action, reqParams, comment.Cookies, http.MethodPost)
	return response
}
