package models

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type ResourceClient struct {
	BaseURL    *url.URL
	UserAgent  string
	HttpClient *http.Client
}

func (s *ResourceClient) Post(path string, form map[string]string, v interface{}) (err error) {
	baseURL, err := url.Parse(MyConfig.WxSessionAPIURL)
	if err != nil {
		return
	}
	rel := &url.URL{Path: path}
	u := baseURL.ResolveReference(rel)
	f := url.Values{}
	for k, v := range form {
		f.Add(k, v)
	}
	resp, err := http.PostForm(u.String(), f)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(v)
	} else {
		err = errors.New("status:" + resp.Status)
	}
	return
}

func (s *ResourceClient) PostForm(url string, form map[string]io.Reader, v interface{}) (err error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range form {
		var fw io.Writer
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return err
		}

	}

	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()

	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest("POST", url, &b)

	if err != nil {
		return
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	// Submit the request
	resp, err := httpClient.Do(req)

	if err != nil {
		return
	}

	// Check the response
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(&v)
	} else {
		err = errors.New("status:" + resp.Status)
	}
	return
}

func (s *ResourceClient) PostFile(path string, file io.Reader, v interface{}) (err error) {
	baseURL, err := url.Parse(MyConfig.WxSessionAPIURL)
	if err != nil {
		return
	}
	rel := &url.URL{Path: path}
	u := baseURL.ResolveReference(rel)

	resp, err := http.Post(u.String(), "application/x-www-form-urlencoded", file)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		err = json.NewDecoder(resp.Body).Decode(v)
	} else {
		err = errors.New("status:" + resp.Status)
	}
	return
}

func (s *ResourceClient) Get(path string, query map[string]string, v interface{}) (err error) {
	baseURL, err := url.Parse(MyConfig.WxSessionAPIURL)
	if err != nil {
		return
	}
	rel := &url.URL{Path: path}
	u := baseURL.ResolveReference(rel)
	q := u.Query()

	if query != nil {
		for k, v := range query {
			q.Add(k, v)
		}
	}
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	log.Println("url:", u.String())
	if err != nil {
		return
	}
	req.Header.Set("Accept", "application/json")

	resp, err := httpClient.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		bodyByte, _ := ioutil.ReadAll(resp.Body)
		log.Println("bodyByte:", strings.NewReader(string(bodyByte)))
		dec := json.NewDecoder(strings.NewReader(string(bodyByte)))
		dec.UseNumber()
		dec.Decode(&v)

	} else {
		err = errors.New("status:" + resp.Status)
	}
	return
}

var (
	wxClient   *ResourceClient
	httpClient *http.Client
)

func init() {
	baseUrl, err := url.Parse(MyConfig.WxSessionAPIURL)
	if err != nil {
		panic(err)
	}
	httpClient = &http.Client{
		Transport: &http.Transport{
			//			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression:    true,
			ResponseHeaderTimeout: time.Second * 5,
		},
	}
	wxClient = &ResourceClient{
		BaseURL:    baseUrl,
		HttpClient: httpClient,
	}
}
