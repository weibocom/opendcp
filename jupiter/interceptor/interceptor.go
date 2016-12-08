package interceptor

import (
	"github.com/astaxie/beego/context"
	"regexp"
)

//
//func apiInterceptor(ctx *context.Context) {
//	ctx.Output.Header("Content-Type", "application/json; charset=utf-8")
//	user := ctx.Input.CruSession.Get("User")
//	appId, appKey := GetAppIdAndKey(ctx)
//	url := ctx.Request.URL.Path
//	// check appkey
//	if !MatchUrl(url, noAuthUrlArr) {
//		userAppkey, err := settingService.GetAppkey(appId, appKey)
//		if err != nil {
//			beego.Error("check appkey err: ", err)
//			ctx.WriteString(AppkeyFaildRespStr)
//			return
//		}
//		// banned key
//		if !userAppkey.InUse {
//			ctx.WriteString(AppkeyBannedRespStr)
//			return
//		}
//		// call api without login
//		if user == nil {
//			// can not user system key
//			if userAppkey.IsSystem || userAppkey.User.Deleted == 1 {
//				ctx.WriteString(AppkeyFaildRespStr)
//				return
//			} else {
//				// set key user into session
//				if userName := ctx.Input.Header("UserName"); userName == "" {
//					ctx.Input.CruSession.Set("User", userAppkey.User)
//					user = userAppkey.User
//				} else if userAppkey.User.Admin == 1 {
//					// admin can do a trick
//					u, err := userService.GetUserByName(userName)
//					if err != nil {
//						ctx.WriteString(UserHeaderErrRespStr)
//						return
//					}
//					ctx.Input.CruSession.Set("User", &u)
//					user = &u
//				}
//			}
//		}
//	}
//	// check admin auth
//	if util.MatchUrl(url, adminUrlArr) && (user == nil || user.(*User).Admin != 1) {
//		ctx.WriteString(NoAuthRespStr)
//		return
//	}
//	// update custom resp content-type
//	for k, v := range customContentTypeMap {
//		if strings.HasPrefix(url, k) {
//			ctx.Output.Header("Content-Type", v)
//			break
//		}
//	}
//}

func MatchUrl(url string, regexArr []*regexp.Regexp) bool {
	for _, r := range regexArr {
		if r.MatchString(url) {
			return true
		}
	}
	return false
}

func GetAppIdAndKey(ctx *context.Context) (string, string) {
	appId := ""
	appIdSess := ctx.Input.CruSession.Get("AppId")
	if appIdSess == nil {
		appId = ctx.Input.Header("App-Id")
	} else {
		appId = appIdSess.(string)
	}
	appKey := ""
	appKeySess := ctx.Input.CruSession.Get("AppKey")
	if appKeySess == nil {
		appKey = ctx.Input.Header("App-Key")
	} else {
		appKey = appKeySess.(string)
	}
	return appId, appKey
}
