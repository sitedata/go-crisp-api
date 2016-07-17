// Copyright 2016 Crisp IM. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
)


// WebsiteOperatorListData mapping
type WebsiteOperatorListData struct {
  Data  *[]WebsiteOperatorListOne  `json:"data,omitempty"`
}

// WebsiteOperatorListOne mapping
type WebsiteOperatorListOne struct {
  Type     *string           `json:"type,omitempty"`
  Details  *WebsiteOperator  `json:"details,omitempty"`
}

// WebsiteOperatorData mapping
type WebsiteOperatorData struct {
  Data  *WebsiteOperator  `json:"data,omitempty"`
}

// WebsiteOperator mapping
type WebsiteOperator struct {
  UserID        *string  `json:"user_id,omitempty"`
  Email         *string  `json:"email,omitempty"`
  FirstName     *string  `json:"first_name,omitempty"`
  LastName      *string  `json:"last_name,omitempty"`
  Role          *string  `json:"role,omitempty"`
  Availability  *string  `json:"availability,omitempty"`
}

// WebsiteOperatorInvite mapping
type WebsiteOperatorInvite struct {
  Email  *string  `json:"email,omitempty"`
  Role   *string  `json:"role,omitempty"`
}

// WebsiteOperatorEdit mapping
type WebsiteOperatorEdit struct {
  Role  *string  `json:"role,omitempty"`
}


// ListWebsiteOperators lists all operator members of website.
// Reference: https://docs.crisp.im/api/v1/#website-website-operator-get
func (service *WebsiteService) ListWebsiteOperators(websiteID string) (*[]WebsiteOperatorListOne, *Response, error) {
  url := fmt.Sprintf("website/%s/operator", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  operators := new(WebsiteOperatorListData)
  resp, err := service.client.Do(req, operators)
  if err != nil {
    return nil, resp, err
  }

  return operators.Data, resp, err
}


// GetWebsiteOperator resolves a given website operator.
// Reference: https://docs.crisp.im/api/v1/#website-website-operator-get-1
func (service *WebsiteService) GetWebsiteOperator(websiteID string, userID string) (*WebsiteOperator, *Response, error) {
  url := fmt.Sprintf("website/%s/operator/%s", websiteID, userID)
  req, _ := service.client.NewRequest("GET", url, nil)

  operator := new(WebsiteOperatorData)
  resp, err := service.client.Do(req, operator)
  if err != nil {
    return nil, resp, err
  }

  return operator.Data, resp, err
}


// InviteWebsiteOperator invites an email to join website as operator.
// Reference: https://docs.crisp.im/api/v1/#website-website-operator-post
func (service *WebsiteService) InviteWebsiteOperator(websiteID string, email string, role string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operator", websiteID)
  req, _ := service.client.NewRequest("POST", url, WebsiteOperatorInvite{Email: &email, Role: &role})

  return service.client.Do(req, nil)
}


// ChangeOperatorRole changes the role of an existing operator. Useful to downgrade or upgrade an operator from/to owner role.
// Reference: https://docs.crisp.im/api/v1/#website-website-operator-patch
func (service *WebsiteService) ChangeOperatorRole(websiteID string, userID string, role string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operator/%s", websiteID, userID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteOperatorEdit{Role: &role})

  return service.client.Do(req, nil)
}


// UnlinkOperatorFromWebsite unlinks given operator from website. Note that the last operator in the website cannot be unlinked.
// Reference: https://docs.crisp.im/api/v1/#website-website-operator-delete
func (service *WebsiteService) UnlinkOperatorFromWebsite(websiteID string, userID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/operator/%s", websiteID, userID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}