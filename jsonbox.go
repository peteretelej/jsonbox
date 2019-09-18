package jsonbox

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

// Errors returned when working with jsonbox package
var (
	ErrTODO = errors.New("functionality not implemented")
	ErrName = errors.New("invalid name for either BOX_ID or COLLECTION")
)

// New returns a new jsonbox client
func NewClient(baseURL string) (*Client, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, errors.New("invalid host url")
	}
	return &Client{
		baseURL: baseURL,
		URL:     url,
		Timeout: time.Second * 10,
	}, nil
}

// NameRegExp defines valid BOX_ID or COLLECTION names reges
var NameRegExp = regexp.MustCompile("^[a-zA-Z0-9_]*$")

// Client defines the structure of a jsonbox client
type Client struct {
	baseURL string

	URL     *url.URL
	Timeout time.Duration
}

// Create creates or adds a record to a boxId
func (c Client) Create(boxID string, val []byte) ([]byte, error) {
	if ok := NameRegExp.MatchString(boxID); !ok {
		return nil, ErrName
	}
	return c.Request(http.MethodPost, boxID, val)
}

// Read reads records for a boxId or boxId with query
func (c Client) Read(boxID string) ([]byte, error) {
	return c.Request(http.MethodGet, boxID, nil)
}

// Update updates a boxID record given the record key and new value
func (c Client) Update(boxID string, recordID string, val []byte) ([]byte, error) {
	p := boxID + "/" + recordID
	return c.Request(http.MethodPut, p, val)
}

// Delete deletes a boxID record given the record Id
func (c Client) Delete(boxID string, recordID string) error {
	p := boxID + "/" + recordID
	_, err := c.Request(http.MethodDelete, p, nil)
	return err
}

// IDs returns all IDs in a BoxID collection
func (c Client) IDs(boxID string) ([]string, error) {
	out, err := c.Read(boxID)
	if err != nil {
		return nil, err
	}
	var ids []string
	metas, err := GetRecordMetas(out)
	if err != nil {
		return nil, err
	}
	for _, m := range metas {
		ids = append(ids, m.ID)
	}
	return ids, nil
}

// DeleteAll deletes all records for boxID
func (c Client) DeleteAll(boxID string) error {
	out, err := c.Read(boxID)
	if err != nil {
		log.Fatalf("failed to READ record(s) from boxID %s: %v", boxID, err)
	}
	metas, err := GetRecordMetas(out)
	for _, m := range metas {
		err := c.Delete(boxID, m.ID)
		if err != nil {
			return fmt.Errorf("Failed to delete record %s/%s: %v", boxID, m.ID, err)
		}
	}
	return nil
}

// Request makes a request to json box url
func (c Client) Request(method, urlPath string, dat []byte) ([]byte, error) {
	if c.baseURL == "" {
		return nil, errors.New("invalid client, create using NewClient")
	}

	cl := &http.Client{Timeout: c.Timeout}
	urlPath = strings.TrimPrefix(urlPath, "/")
	urlString := c.URL.String() + "/" + urlPath

	var req *http.Request
	var err error

	switch method {
	case http.MethodPost, http.MethodPut:
		if dat == nil {
			return nil, errors.New("missing request payload")
		}
		req, err = http.NewRequest(method, urlString, bytes.NewBuffer(dat))
		req.Header.Set("content-type", "application/json")
	default:
		req, err = http.NewRequest(method, urlString, nil)
	}
	if err != nil {
		return nil, err
	}
	resp, err := cl.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("%s - %s", resp.Status, out)
	}
	return out, nil

}

// Meta defines metadata for json stored
type Meta struct {
	ID        string `json:"_id"`
	CreatedOn string `json:"_createdOn"`
}

// GetRecordID returns the id of a record, or the first ID if multiple records
func GetRecordID(dat []byte) (string, error) {
	fmt.Printf("Dat : %s", dat)
	multipleRecords := bytes.HasPrefix(dat, []byte("["))
	if multipleRecords {
		mm, err := GetRecordMetas(dat)
		if err != nil {
			return "", err
		}
		return mm[0].ID, nil
	}
	var m Meta
	if err := json.Unmarshal(dat, &m); err != nil {
		return "", err
	}
	fmt.Printf("Meta : %s", m)
	return m.ID, nil
}

// GetRecordMetas returns a list of record Metas, fails for single record
func GetRecordMetas(dat []byte) ([]Meta, error) {
	var mm []Meta
	if err := json.Unmarshal(dat, &mm); err != nil {
		return nil, err
	}
	if len(mm) == 0 {
		return nil, errors.New("no records found")
	}
	return mm, nil
}
