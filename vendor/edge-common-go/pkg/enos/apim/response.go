/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : response.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 16:56:21
 *
 * http://www.envision-group.com/
 */

package apim

type RespBase struct {
	Code      int    `json:"code"`
	Message   string `json:"msg"`
	RequestId string `json:"requestId"`
}
