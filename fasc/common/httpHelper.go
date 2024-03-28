package common

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

// SendPost httpPost 请求
func SendPost(urlStr string, headMaps map[string]string, data []byte) ([]byte, string) {
	client := &http.Client{}
	values := url.Values{}
	values.Add("bizContent", string(data))
	encodedData := values.Encode()
	r, _ := http.NewRequest("POST", urlStr, bytes.NewReader([]byte(encodedData))) // URL-encoded payload
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headMaps {
		r.Header.Add(k, v)
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil, ""
	}
	defer resp.Body.Close()
	requestId := resp.Header.Values("X-FASC-Request-Id")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	if requestId == nil {
		return body, ""
	}
	return body, requestId[0]
}

// UploadFilePost 上传即时文件post
func UploadFilePost(urlStr string, headMaps map[string]string, file []byte, bizContent string, fileName string) ([]byte, string, error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)
	//boundary :=  bodyWriter.Boundary()
	bodyWriter.WriteField("bizContent", bizContent)
	// 2. 内存中的文件1，FormFile1
	_, err := bodyWriter.CreateFormFile("fileContent", fileName)
	if err != nil {
		return nil, "", err
	}
	bodyBuf.Write(file)
	// 结束整个消息body
	bodyWriter.Close()
	reqReader := io.MultiReader(bodyBuf)
	req, err := http.NewRequest("POST", urlStr, reqReader)
	if err != nil {
		return nil, "", err
	}
	for k, v := range headMaps {
		req.Header.Add(k, v)
	}
	// 添加Post头
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.ContentLength = int64(bodyBuf.Len())
	// 发送消息
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	requestId := resp.Header.Values("X-FASC-Request-Id")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("读取回应消息异常:%s", err)
	}
	if requestId == nil {
		return body, "", err
	}
	return body, requestId[0], err
}

// PutFileUpload 本地PUT请求上传文件
func PutFileUpload(uploadUrl string, filePath string) error {
	//读取文件
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyBuf.Write(file)
	// 结束整个消息body
	bodyWriter.Close()
	reqReader := io.MultiReader(bodyBuf)
	req, err := http.NewRequest("PUT", uploadUrl, reqReader)
	if err != nil {
		return err
	}
	req.Header.Set("Connection", "close")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	req.ContentLength = int64(bodyBuf.Len())
	// 发送消息
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		//return errors.New("wrong http status code")
		return fmt.Errorf("wrong http status code %d", resp.StatusCode)
	}
	return nil
}

// DownLoadFilesPost 下载文件HttpPost
func DownLoadFilesPost(urlStr string, headMaps map[string]string, data []byte) ([]byte, string, string) {
	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, bytes.NewReader(data)) // URL-encoded payload
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	for k, v := range headMaps {
		r.Header.Add(k, v)
	}
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "", ""
	}
	defer resp.Body.Close()
	requestId := resp.Header.Values("X-FASC-Request-Id")
	contentType := resp.Header.Values("Content-Type")
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, int64(5<<20)))
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	if requestId == nil {
		return body, "", contentType[0]
	}
	if contentType == nil {
		return body, requestId[0], ""
	}
	return body, requestId[0], contentType[0]
}

func PostWithFormData(method, url string, postData *map[string]string) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range *postData {
		w.WriteField(k, v)
	}
	w.Close()
	req, _ := http.NewRequest(method, url, body)
	req.Header.Set("Content-Type", w.FormDataContentType())
	resp, _ := http.DefaultClient.Do(req)
	data, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Printf("%s", data)
}
