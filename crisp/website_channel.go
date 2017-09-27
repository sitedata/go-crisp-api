// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteChannelEmailData mapping
type WebsiteChannelEmailData struct {
  Data  *WebsiteChannelEmail  `json:"data,omitempty"`
}

// WebsiteChannelEmail mapping
type WebsiteChannelEmail struct {
  Domain  *string  `json:"domain,omitempty"`
  Email   *string  `json:"email,omitempty"`
}

// WebsiteChannelEmailRequestChangeData mapping
type WebsiteChannelEmailRequestChangeData struct {
  Data  *WebsiteChannelEmailRequestChange  `json:"data,omitempty"`
}

// WebsiteChannelEmailRequestChange mapping
type WebsiteChannelEmailRequestChange struct {
  Domain  *string                                 `json:"domain,omitempty"`
  Setup   *WebsiteChannelEmailRequestChangeSetup  `json:"setup,omitempty"`
}

// WebsiteChannelEmailRequestChangeSetup mapping
type WebsiteChannelEmailRequestChangeSetup struct {
  Records  *[]WebsiteChannelEmailRequestChangeSetupRecord  `json:"records,omitempty"`
}

// WebsiteChannelEmailRequestChangeSetupRecord mapping
type WebsiteChannelEmailRequestChangeSetupRecord struct {
  Type   *string  `json:"type,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// WebsiteChannelEmailRequest mapping
type WebsiteChannelEmailRequest struct {
  Domain  *string  `json:"domain,omitempty"`
}


// String returns the string representation of WebsiteChannelEmail
func (instance WebsiteChannelEmail) String() string {
  return Stringify(instance)
}


// GetWebsiteEmailChannel resolves the website email channel value.
func (service *WebsiteService) GetWebsiteEmailChannel(websiteID string) (*WebsiteChannelEmail, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  channel := new(WebsiteChannelEmailData)
  resp, err := service.client.Do(req, channel)
  if err != nil {
    return nil, resp, err
  }

  return channel.Data, resp, err
}


// RequestWebsiteEmailChannelChange requests a change in the email channel domain used to send and receive emails.
func (service *WebsiteService) RequestWebsiteEmailChannelChange(websiteID string, domain string) (*WebsiteChannelEmailRequestChange, *Response, error) {
  url := fmt.Sprintf("website/%s/channel/email", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteChannelEmailRequest{Domain: &domain})

  channel := new(WebsiteChannelEmailRequestChangeData)
  resp, err := service.client.Do(req, channel)
  if err != nil {
    return nil, resp, err
  }

  return channel.Data, resp, err
}
