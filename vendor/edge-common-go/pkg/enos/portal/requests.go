/*
 * Copyright © 2020 江苏远景科创智能科技有限公司
 * All rights reserved.
 *
 * Filename : requests.go
 * Description :
 *
 * Created by Laiyuan YANG at 2020/7/10 18:0:54
 *
 * http://www.envision-group.com/
 */

package portal

import "net/http"

//get http get
func (c *Client) get(uri string, query map[string]string, needAuth bool) ([]byte, error) {
	req, err := c.client.GenReq(http.MethodGet, uri, nil, query)
	if err != nil {
		return nil, err
	}

	if needAuth {
		if c.token == "" {
			c.login()
		}
		c.setAuth(req, c.token)
	}

	return c.client.Request(req)
}

func (c *Client) setAuth(req *http.Request, token string) {

}

func (c *Client) login() {

}
