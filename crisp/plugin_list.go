// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// PluginListData mapping
type PluginListData struct {
  Data  *[]PluginInformation  `json:"data,omitempty"`
}


// ListAllPlugins lists all available plugins on the marketplace.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-list-get
func (service *WebsiteService) ListAllPlugins(pageNumber uint) (*[]PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugins/list/all/%d", pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// ListFeaturedPlugins lists featured plugins on the marketplace, from the newer to older.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-list-get-1
func (service *WebsiteService) ListFeaturedPlugins(pageNumber uint) (*[]PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugins/list/featured/%d", pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}


// SearchPlugins searches for plugins in the marketplace, given a search term.
// Reference: https://docs.crisp.im/api/v1/#plugin-plugins-list-get-2
func (service *WebsiteService) SearchPlugins(query string, pageNumber uint) (*[]PluginInformation, *Response, error) {
  url := fmt.Sprintf("plugins/list/search/%d?query=%s", pageNumber, url.QueryEscape(query))
  req, _ := service.client.NewRequest("GET", url, nil)

  plugins := new(PluginListData)
  resp, err := service.client.Do(req, plugins)
  if err != nil {
    return nil, resp, err
  }

  return plugins.Data, resp, err
}