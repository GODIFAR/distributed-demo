package registry

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// /注册服务
func RegisterService(r Registration) error {

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	err := enc.Encode(r)
	if err != nil {
		return err
	}
	res, err := http.Post(ServicesURL, "application/json", buf)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		//注册服务失败。注册服务“+”响应了代码
		return fmt.Errorf("Failed to register service. Registry service"+"responded with code %v", res.StatusCode)
	}
	return nil
}

// 取消服务
func ShutdownService(url string) error {
	req, err := http.NewRequest(http.MethodDelete, ServicesURL, bytes.NewBuffer([]byte(url)))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Tyep", "text/plain")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to deregister service.Registry"+" service responded with code %v", res.StatusCode)
	}
	return nil
}
