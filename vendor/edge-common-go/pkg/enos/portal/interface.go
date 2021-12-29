/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : interface.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 17:0:55
 *
 * http://www.envision-group.com/
 */

package portal

type API interface {
	GetRootCertificate(caFileName string) error
}
