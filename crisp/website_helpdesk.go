// Copyright 2017 Crisp IM, Inc. All rights reserved.
//
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crisp


import (
  "fmt"
  "net/url"
)


// WebsiteHelpdeskLocalesData mapping
type WebsiteHelpdeskLocalesData struct {
  Data  *[]WebsiteHelpdeskLocale  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocale mapping
type WebsiteHelpdeskLocale struct {
  Locale      *string  `json:"locale,omitempty"`
  URL         *string  `json:"url,omitempty"`
  Articles    *uint32  `json:"articles,omitempty"`
  Categories  *uint32  `json:"categories,omitempty"`
}

// WebsiteHelpdeskLocaleData mapping
type WebsiteHelpdeskLocaleData struct {
  Data  *WebsiteHelpdeskLocale  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticlesData mapping
type WebsiteHelpdeskLocaleArticlesData struct {
  Data  *[]WebsiteHelpdeskLocaleArticle  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticle mapping
type WebsiteHelpdeskLocaleArticle struct {
  ArticleID    *string  `json:"article_id,omitempty"`
  Title        *string  `json:"title,omitempty"`
  Content      *string  `json:"content,omitempty"`
  Status       *string  `json:"status,omitempty"`
  Visibility   *string  `json:"visibility,omitempty"`
  URL          *string  `json:"url,omitempty"`
  CreatedAt    *uint    `json:"created_at,omitempty"`
  UpdatedAt    *uint    `json:"updated_at,omitempty"`
  PublishedAt  *uint    `json:"published_at,omitempty"`
}

// WebsiteHelpdeskLocaleArticleData mapping
type WebsiteHelpdeskLocaleArticleData struct {
  Data  *WebsiteHelpdeskLocaleArticle  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticleNewData mapping
type WebsiteHelpdeskLocaleArticleNewData struct {
  Data  *WebsiteHelpdeskLocaleArticleNew  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticleNew mapping
type WebsiteHelpdeskLocaleArticleNew struct {
  ArticleID  *string  `json:"article_id,omitempty"`
}

// WebsiteHelpdeskLocaleArticleCategoryData mapping
type WebsiteHelpdeskLocaleArticleCategoryData struct {
  Data  *WebsiteHelpdeskLocaleArticleCategory  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticleCategory mapping
type WebsiteHelpdeskLocaleArticleCategory struct {
  CategoryID  *string  `json:"category_id,omitempty"`
}

// WebsiteHelpdeskLocaleCategoriesData mapping
type WebsiteHelpdeskLocaleCategoriesData struct {
  Data  *[]WebsiteHelpdeskLocaleCategory  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleCategoryNewData mapping
type WebsiteHelpdeskLocaleCategoryNewData struct {
  Data  *WebsiteHelpdeskLocaleCategoryNew  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticlePublishData mapping
type WebsiteHelpdeskLocaleArticlePublishData struct {
  Data  *WebsiteHelpdeskLocaleArticlePublish  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleArticlePublish mapping
type WebsiteHelpdeskLocaleArticlePublish struct {
  URL  *string  `json:"url,omitempty"`
}

// WebsiteHelpdeskLocaleCategoryNew mapping
type WebsiteHelpdeskLocaleCategoryNew struct {
  CategoryID  *string  `json:"category_id,omitempty"`
}

// WebsiteHelpdeskLocaleCategoryData mapping
type WebsiteHelpdeskLocaleCategoryData struct {
  Data  *WebsiteHelpdeskLocaleCategory  `json:"data,omitempty"`
}

// WebsiteHelpdeskLocaleCategory mapping
type WebsiteHelpdeskLocaleCategory struct {
  CategoryID   *string  `json:"category_id,omitempty"`
  Name         *string  `json:"name,omitempty"`
  Description  *string  `json:"description,omitempty"`
  Color        *string  `json:"color,omitempty"`
  Image        *string  `json:"image,omitempty"`
  URL          *string  `json:"url,omitempty"`
  CreatedAt    *uint    `json:"created_at,omitempty"`
  UpdatedAt    *uint    `json:"updated_at,omitempty"`
}

// WebsiteHelpdeskSettingsData mapping
type WebsiteHelpdeskSettingsData struct {
  Data  *WebsiteHelpdeskSettings  `json:"data,omitempty"`
}

// WebsiteHelpdeskSettings mapping
type WebsiteHelpdeskSettings struct {
  Name        *string                             `json:"name,omitempty"`
  title       *string                             `json:"title,omitempty"`
  Appearance  *WebsiteHelpdeskSettingsAppearance  `json:"appearance,omitempty"`
  Behavior    *WebsiteHelpdeskSettingsBehavior    `json:"behavior,omitempty"`
}

// WebsiteHelpdeskSettingsAppearance mapping
type WebsiteHelpdeskSettingsAppearance struct {
  Logos   *WebsiteHelpdeskSettingsAppearanceLogos  `json:"logos,omitempty"`
  Banner  *string                                  `json:"banner,omitempty"`
}

// WebsiteHelpdeskSettingsAppearanceLogos mapping
type WebsiteHelpdeskSettingsAppearanceLogos struct {
  Header  *string  `json:"header,omitempty"`
  Footer  *string  `json:"footer,omitempty"`
}

// WebsiteHelpdeskSettingsBehavior mapping
type WebsiteHelpdeskSettingsBehavior struct {
  FrequentlyRead  *bool  `json:"frequently_read,omitempty"`
  ShowChatbox     *bool  `json:"show_chatbox,omitempty"`
  AskFeedback     *bool  `json:"ask_feedback,omitempty"`
}

// WebsiteHelpdeskDomainData mapping
type WebsiteHelpdeskDomainData struct {
  Data  *WebsiteHelpdeskDomain  `json:"data,omitempty"`
}

// WebsiteHelpdeskDomain mapping
type WebsiteHelpdeskDomain struct {
  Root    *string  `json:"root,omitempty"`
  Basic   *string  `json:"basic,omitempty"`
  Custom  *string  `json:"custom,omitempty"`
}

// WebsiteHelpdeskDomainSetupFlowData mapping
type WebsiteHelpdeskDomainSetupFlowData struct {
  Data  *WebsiteHelpdeskDomainSetupFlow  `json:"data,omitempty"`
}

// WebsiteHelpdeskDomainSetupFlow mapping
type WebsiteHelpdeskDomainSetupFlow struct {
  Custom  *string                               `json:"custom,omitempty"`
  Setup   *WebsiteHelpdeskDomainSetupFlowSetup  `json:"setup,omitempty"`
}

// WebsiteHelpdeskDomainSetupFlowSetup mapping
type WebsiteHelpdeskDomainSetupFlowSetup struct {
  Records  *[]WebsiteHelpdeskDomainSetupFlowSetupRecord  `json:"records,omitempty"`
}

// WebsiteHelpdeskDomainSetupFlowSetupRecord mapping
type WebsiteHelpdeskDomainSetupFlowSetupRecord struct {
  Type   *string  `json:"type,omitempty"`
  Name   *string  `json:"name,omitempty"`
  Value  *string  `json:"value,omitempty"`
}

// WebsiteHelpdeskLocaleItem mapping
type WebsiteHelpdeskLocaleItem struct {
  Locale  *string  `json:"locale,omitempty"`
}

// WebsiteHelpdeskLocaleArticleItem mapping
type WebsiteHelpdeskLocaleArticleItem struct {
  Title  *string  `json:"title,omitempty"`
}

// WebsiteHelpdeskLocaleArticleCategoryItem mapping
type WebsiteHelpdeskLocaleArticleCategoryItem struct {
  CategoryID  *string  `json:"category_id,omitempty"`
}

// WebsiteHelpdeskLocaleCategoryItem mapping
type WebsiteHelpdeskLocaleCategoryItem struct {
  Name  *string  `json:"name,omitempty"`
}

// WebsiteHelpdeskSettingsUpdate mapping
type WebsiteHelpdeskSettingsUpdate struct {
  Name        string                                   `json:"name,omitempty"`
  title       string                                   `json:"title,omitempty"`
  Appearance  WebsiteHelpdeskSettingsUpdateAppearance  `json:"appearance,omitempty"`
  Behavior    WebsiteHelpdeskSettingsUpdateBehavior    `json:"behavior,omitempty"`
}

// WebsiteHelpdeskSettingsUpdateAppearance mapping
type WebsiteHelpdeskSettingsUpdateAppearance struct {
  Logos   WebsiteHelpdeskSettingsUpdateAppearanceLogos  `json:"logos,omitempty"`
  Banner  string                                        `json:"banner,omitempty"`
}

// WebsiteHelpdeskSettingsUpdateAppearanceLogos mapping
type WebsiteHelpdeskSettingsUpdateAppearanceLogos struct {
  Header  string  `json:"header,omitempty"`
  Footer  string  `json:"footer,omitempty"`
}

// WebsiteHelpdeskSettingsUpdateBehavior mapping
type WebsiteHelpdeskSettingsUpdateBehavior struct {
  FrequentlyRead  bool  `json:"frequently_read,omitempty"`
  ShowChatbox     bool  `json:"show_chatbox,omitempty"`
  AskFeedback     bool  `json:"ask_feedback,omitempty"`
}

// WebsiteHelpdeskDomainChange mapping
type WebsiteHelpdeskDomainChange struct {
  Basic   string  `json:"basic,omitempty"`
  Custom  string  `json:"custom,omitempty"`
}


// String returns the string representation of WebsiteHelpdeskLocale
func (instance WebsiteHelpdeskLocale) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskLocaleArticle
func (instance WebsiteHelpdeskLocaleArticle) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskLocaleArticleNew
func (instance WebsiteHelpdeskLocaleArticleNew) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskLocaleArticleCategory
func (instance WebsiteHelpdeskLocaleArticleCategory) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskLocaleCategoryNew
func (instance WebsiteHelpdeskLocaleCategoryNew) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskLocaleCategory
func (instance WebsiteHelpdeskLocaleCategory) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskSettings
func (instance WebsiteHelpdeskSettings) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskDomain
func (instance WebsiteHelpdeskDomain) String() string {
  return Stringify(instance)
}

// String returns the string representation of WebsiteHelpdeskDomainSetupFlow
func (instance WebsiteHelpdeskDomainSetupFlow) String() string {
  return Stringify(instance)
}


// ListHelpdeskLocales lists locales for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocales(websiteID string, pageNumber uint) (*[]WebsiteHelpdeskLocale, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locales/%d", websiteID, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  locales := new(WebsiteHelpdeskLocalesData)
  resp, err := service.client.Do(req, locales)
  if err != nil {
    return nil, resp, err
  }

  return locales.Data, resp, err
}


// AddHelpdeskLocale adds a locale for helpdesk in website.
func (service *WebsiteService) AddHelpdeskLocale(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale", websiteID)
  req, _ := service.client.NewRequest("POST", url, WebsiteHelpdeskLocaleItem{Locale: &locale})

  return service.client.Do(req, nil)
}


// CheckIfHelpdeskLocaleExists checks if a helpdesk locale exists for helpdesk in website.
func (service *WebsiteService) CheckIfHelpdeskLocaleExists(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocale resolves a locale for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocale(websiteID string, locale string) (*WebsiteHelpdeskLocale, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("GET", url, nil)

  localeData := new(WebsiteHelpdeskLocaleData)
  resp, err := service.client.Do(req, localeData)
  if err != nil {
    return nil, resp, err
  }

  return localeData.Data, resp, err
}


// DeleteHelpdeskLocale deletes a locale for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocale(websiteID string, locale string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s", websiteID, locale)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskLocaleArticles lists articles for a helpdesk locale in website.
func (service *WebsiteService) ListHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint) (*[]WebsiteHelpdeskLocaleArticle, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/articles/%d", websiteID, locale, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  articles := new(WebsiteHelpdeskLocaleArticlesData)
  resp, err := service.client.Do(req, articles)
  if err != nil {
    return nil, resp, err
  }

  return articles.Data, resp, err
}


// FilterHelpdeskLocaleArticles lists articles for a helpdesk locale in website (filter variant).
func (service *WebsiteService) FilterHelpdeskLocaleArticles(websiteID string, locale string, pageNumber uint, searchTitle string, filterCategoryID string, filterStatusPublished bool, filterStatusDraft bool, filterVisibilityVisible bool, filterVisibilityHidden bool) (*[]WebsiteHelpdeskLocaleArticle, *Response, error) {
  var (
    filterStatusPublishedValue string
    filterStatusDraftValue string
    filterVisibilityVisibleValue string
    filterVisibilityHiddenValue string
  )

  if filterStatusPublished == true {
    filterStatusPublishedValue = "1"
  } else {
    filterStatusPublishedValue = "0"
  }

  if filterStatusDraft == true {
    filterStatusDraftValue = "1"
  } else {
    filterStatusDraftValue = "0"
  }

  if filterVisibilityVisible == true {
    filterVisibilityVisibleValue = "1"
  } else {
    filterVisibilityVisibleValue = "0"
  }

  if filterVisibilityHidden == true {
    filterVisibilityHiddenValue = "1"
  } else {
    filterVisibilityHiddenValue = "0"
  }

  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/articles/%d?search_title=%s&filter_category_id=%s&filter_status_published=%s&filter_status_draft=%s&filter_visibility_visible=%s&filter_visibility_hidden=%s", websiteID, locale, pageNumber, url.QueryEscape(searchTitle), url.QueryEscape(filterCategoryID), url.QueryEscape(filterStatusPublishedValue), url.QueryEscape(filterStatusDraftValue), url.QueryEscape(filterVisibilityVisibleValue), url.QueryEscape(filterVisibilityHiddenValue))
  req, _ := service.client.NewRequest("GET", url, nil)

  articles := new(WebsiteHelpdeskLocaleArticlesData)
  resp, err := service.client.Do(req, articles)
  if err != nil {
    return nil, resp, err
  }

  return articles.Data, resp, err
}


// AddNewHelpdeskLocaleArticle adds a new locale article for helpdesk in website.
func (service *WebsiteService) AddNewHelpdeskLocaleArticle(websiteID string, locale string, articleTitle string) (*WebsiteHelpdeskLocaleArticleNew, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article", websiteID, locale)
  req, _ := service.client.NewRequest("POST", url, WebsiteHelpdeskLocaleArticleItem{Title: &articleTitle})

  articleNew := new(WebsiteHelpdeskLocaleArticleNewData)
  resp, err := service.client.Do(req, articleNew)
  if err != nil {
    return nil, resp, err
  }

  return articleNew.Data, resp, err
}


// CheckIfHelpdeskLocaleArticleExists checks if a locale article exists for helpdesk in website.
func (service *WebsiteService) CheckIfHelpdeskLocaleArticleExists(websiteID string, locale string, articleID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleArticle resolves a locale article for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*WebsiteHelpdeskLocaleArticle, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("GET", url, nil)

  article := new(WebsiteHelpdeskLocaleArticleData)
  resp, err := service.client.Do(req, article)
  if err != nil {
    return nil, resp, err
  }

  return article.Data, resp, err
}


// SaveHelpdeskLocaleArticle saves a locale article for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleArticle(websiteID string, locale string, articleID string, article WebsiteHelpdeskLocaleArticleItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("PUT", url, article)

  return service.client.Do(req, nil)
}


// UpdateHelpdeskLocaleArticle updates a locale article for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleArticle(websiteID string, locale string, articleID string, article WebsiteHelpdeskLocaleArticleItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("PATCH", url, article)

  return service.client.Do(req, nil)
}


// DeleteHelpdeskLocaleArticle deletes a locale article for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleArticleCategory resolves a locale article category for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleArticleCategory(websiteID string, locale string, articleID string) (*WebsiteHelpdeskLocaleArticleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/category", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("GET", url, nil)

  articleCategory := new(WebsiteHelpdeskLocaleArticleCategoryData)
  resp, err := service.client.Do(req, articleCategory)
  if err != nil {
    return nil, resp, err
  }

  return articleCategory.Data, resp, err
}


// UpdateHelpdeskLocaleArticleCategory updates a locale article category for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleArticleCategory(websiteID string, locale string, articleID string, categoryID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/category", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("PATCH", url, WebsiteHelpdeskLocaleArticleCategoryItem{CategoryID: &categoryID})

  return service.client.Do(req, nil)
}


// PublishHelpdeskLocaleArticle publishes a locale article for helpdesk in website.
func (service *WebsiteService) PublishHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*WebsiteHelpdeskLocaleArticlePublish, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/publish", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("POST", url, nil)

  publishData := new(WebsiteHelpdeskLocaleArticlePublishData)
  resp, err := service.client.Do(req, publishData)
  if err != nil {
    return nil, resp, err
  }

  return publishData.Data, resp, err
}


// UnpublishHelpdeskLocaleArticle unpublishes a locale article for helpdesk in website.
func (service *WebsiteService) UnpublishHelpdeskLocaleArticle(websiteID string, locale string, articleID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/article/%s/unpublish", websiteID, locale, articleID)
  req, _ := service.client.NewRequest("POST", url, nil)

  return service.client.Do(req, nil)
}


// ListHelpdeskLocaleCategories lists locale categories for helpdesk in website.
func (service *WebsiteService) ListHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint) (*[]WebsiteHelpdeskLocaleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/categories/%d", websiteID, locale, pageNumber)
  req, _ := service.client.NewRequest("GET", url, nil)

  categories := new(WebsiteHelpdeskLocaleCategoriesData)
  resp, err := service.client.Do(req, categories)
  if err != nil {
    return nil, resp, err
  }

  return categories.Data, resp, err
}


// FilterHelpdeskLocaleCategories lists locale categories for helpdesk in website (filter variant).
func (service *WebsiteService) FilterHelpdeskLocaleCategories(websiteID string, locale string, pageNumber uint, searchName string) (*[]WebsiteHelpdeskLocaleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/categories/%d?search_name=%s", websiteID, locale, pageNumber, url.QueryEscape(searchName))
  req, _ := service.client.NewRequest("GET", url, nil)

  categories := new(WebsiteHelpdeskLocaleCategoriesData)
  resp, err := service.client.Do(req, categories)
  if err != nil {
    return nil, resp, err
  }

  return categories.Data, resp, err
}


// AddHelpdeskLocaleCategory adds a locale category for helpdesk in website.
func (service *WebsiteService) AddHelpdeskLocaleCategory(websiteID string, locale string, categoryName string) (*WebsiteHelpdeskLocaleCategoryNew, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category", websiteID, locale)
  req, _ := service.client.NewRequest("POST", url, WebsiteHelpdeskLocaleCategoryItem{Name: &categoryName})

  categoryNew := new(WebsiteHelpdeskLocaleCategoryNewData)
  resp, err := service.client.Do(req, categoryNew)
  if err != nil {
    return nil, resp, err
  }

  return categoryNew.Data, resp, err
}


// CheckIfHelpdeskLocaleCategoryExists checks if a locale category exists for helpdesk in website.
func (service *WebsiteService) CheckIfHelpdeskLocaleCategoryExists(websiteID string, locale string, categoryID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryID)
  req, _ := service.client.NewRequest("HEAD", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskLocaleCategory resolves a locale category for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskLocaleCategory(websiteID string, locale string, categoryID string) (*WebsiteHelpdeskLocaleCategory, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryID)
  req, _ := service.client.NewRequest("GET", url, nil)

  category := new(WebsiteHelpdeskLocaleCategoryData)
  resp, err := service.client.Do(req, category)
  if err != nil {
    return nil, resp, err
  }

  return category.Data, resp, err
}


// SaveHelpdeskLocaleCategory saves a locale category for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskLocaleCategory(websiteID string, locale string, categoryID string, category WebsiteHelpdeskLocaleCategoryItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryID)
  req, _ := service.client.NewRequest("PUT", url, category)

  return service.client.Do(req, nil)
}


// UpdateHelpdeskLocaleCategory updates a locale category for helpdesk in website.
func (service *WebsiteService) UpdateHelpdeskLocaleCategory(websiteID string, locale string, categoryID string, category WebsiteHelpdeskLocaleCategoryItem) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryID)
  req, _ := service.client.NewRequest("PATCH", url, category)

  return service.client.Do(req, nil)
}


// DeleteHelpdeskLocaleCategory deletes a locale category for helpdesk in website.
func (service *WebsiteService) DeleteHelpdeskLocaleCategory(websiteID string, locale string, categoryID string) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/locale/%s/category/%s", websiteID, locale, categoryID)
  req, _ := service.client.NewRequest("DELETE", url, nil)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskSettings resolves settings for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskSettings(websiteID string) (*WebsiteHelpdeskSettings, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/settings", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  settings := new(WebsiteHelpdeskSettingsData)
  resp, err := service.client.Do(req, settings)
  if err != nil {
    return nil, resp, err
  }

  return settings.Data, resp, err
}


// SaveHelpdeskSettings saves settings for helpdesk in website.
func (service *WebsiteService) SaveHelpdeskSettings(websiteID string, settings WebsiteHelpdeskSettingsUpdate) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/settings", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, settings)

  return service.client.Do(req, nil)
}


// ResolveHelpdeskDomain resolves domain for helpdesk in website.
func (service *WebsiteService) ResolveHelpdeskDomain(websiteID string) (*WebsiteHelpdeskDomain, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/domain", websiteID)
  req, _ := service.client.NewRequest("GET", url, nil)

  domain := new(WebsiteHelpdeskDomainData)
  resp, err := service.client.Do(req, domain)
  if err != nil {
    return nil, resp, err
  }

  return domain.Data, resp, err
}


// RequestHelpdeskDomainChange requests a change in the domain used for helpdesk.
func (service *WebsiteService) RequestHelpdeskDomainChange(websiteID string, domain WebsiteHelpdeskDomainChange) (*Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/domain", websiteID)
  req, _ := service.client.NewRequest("PATCH", url, domain)

  return service.client.Do(req, nil)
}


// GenerateHelpdeskDomainSetupFlow retrieves the domain setup flow for helpdesk.
func (service *WebsiteService) GenerateHelpdeskDomainSetupFlow(websiteID string, custom string) (*WebsiteHelpdeskDomainSetupFlow, *Response, error) {
  url := fmt.Sprintf("website/%s/helpdesk/domain/setup?custom=%s", websiteID, url.QueryEscape(custom))
  req, _ := service.client.NewRequest("GET", url, nil)

  domainSetupFlow := new(WebsiteHelpdeskDomainSetupFlowData)
  resp, err := service.client.Do(req, domainSetupFlow)
  if err != nil {
    return nil, resp, err
  }

  return domainSetupFlow.Data, resp, err
}