/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : certificate.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 17:59:32
 *
 * http://www.envision-group.com/
 */

package portal

import "os"

const (
	RootCertificateURI = "/enos/CA/cacert"
)

func (c *Client) GetRootCertificate(caFileName string) error {
	resp, err := c.get(RootCertificateURI, nil, false)
	if err != nil {
		return err
	}

	fd, err := os.OpenFile(caFileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	_, err = fd.Write(resp)
	if err != nil {
		return err
	}
	fd.Close()

	return nil
}
