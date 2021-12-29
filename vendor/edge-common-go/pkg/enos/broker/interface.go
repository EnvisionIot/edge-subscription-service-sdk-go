/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : interface.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 20:59:55
 *
 * http://www.envision-group.com/
 */

package broker

type Api interface {
	GetOtaPkg(fileUri string) ([]byte, error)
	DownloadFile(fileUri, category string) ([]byte, error)
}
