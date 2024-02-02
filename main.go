package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    welcome, err := UnmarshalWelcome(bytes)
//    bytes, err = welcome.Marshal()

func UnmarshalWelcome(data []byte) (Welcome, error) {
	var r Welcome
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Welcome) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Welcome struct {
	Banner          Banner            `json:"banner"`
	Theme           Theme             `json:"theme"`
	OverrideMessage OverrideMessage   `json:"overrideMessage"`
	Products        []ProductElement  `json:"products"`
	SizeGuides      []SizeGuide       `json:"sizeGuides"`
	Upcoming        []UpcomingElement `json:"upcoming"`
	Blog            []Blog            `json:"blog"`
	Brands          []BrandElement    `json:"brands"`
	Badges          []Badge           `json:"badges"`
	Achievements    []int64           `json:"achievements"`
	Banners         []interface{}     `json:"banners"`
	Campaigns       Campaigns         `json:"campaigns"`
}

type Badge struct {
	Background string    `json:"background"`
	ShareImage string    `json:"shareImage"`
	Image      string    `json:"image"`
	Foreground string    `json:"foreground"`
	StoreName  StoreName `json:"storeName"`
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	SortOrder  int64     `json:"sortOrder"`
}

type Banner struct {
}

type Blog struct {
	Link         string     `json:"link"`
	Categories   []Category `json:"categories"`
	Title        string     `json:"title"`
	Excerpt      string     `json:"excerpt"`
	FeatureMedia string     `json:"feature_media"`
}

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type BrandElement struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type Campaigns struct {
	Campaigns []Campaign    `json:"campaigns"`
	Pages     []PageElement `json:"pages"`
}

type Campaign struct {
	DisplayName string    `json:"displayName"`
	Location    string    `json:"location"`
	Image       string    `json:"image"`
	HomePage    string    `json:"homePage"`
	ShowTo      string    `json:"showTo"`
	StoreName   StoreName `json:"storeName"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Invert      bool      `json:"invert"`
	SortOrder   int64     `json:"sortOrder"`
	Type        string    `json:"type"`
}

type PageElement struct {
	Content         []Content `json:"content"`
	IncludeNav      bool      `json:"includeNav"`
	Image           string    `json:"image"`
	Loop            bool      `json:"loop"`
	StoreName       StoreName `json:"storeName"`
	TransparentHead bool      `json:"transparentHead"`
	ID              string    `json:"id"`
	CampaignID      string    `json:"campaignID"`
	Name            string    `json:"name"`
	Type            PageType  `json:"type"`
}

type Content struct {
	Type    ContentType `json:"type"`
	Data    Data        `json:"data"`
	Name    string      `json:"name"`
	Disable *bool       `json:"disable,omitempty"`
}

type Data struct {
	Src           *string  `json:"src,omitempty"`
	Disable       *bool    `json:"disable,omitempty"`
	Effect        *Effect  `json:"effect,omitempty"`
	Name          *string  `json:"name,omitempty"`
	Action        *Action  `json:"action"`
	Tag           *Tag     `json:"tag,omitempty"`
	Fullwidth     *bool    `json:"fullwidth,omitempty"`
	Thumbnail     *string  `json:"thumbnail,omitempty"`
	Scripts       []string `json:"scripts,omitempty"`
	Data          *string  `json:"data,omitempty"`
	HideWithOptin *bool    `json:"hideWithOptin,omitempty"`
}

type Action struct {
	Type  ActionType `json:"type"`
	Value string     `json:"value"`
}

type OverrideMessage struct {
	Active  bool   `json:"active"`
	Message string `json:"message"`
}

type ProductElement struct {
	TermsTemplate                     interface{}        `json:"termsTemplate"`
	LaunchDay                         bool               `json:"launchDay"`
	EntryLimit                        interface{}        `json:"entryLimit"`
	PromoImage                        interface{}        `json:"promoImage"`
	Status                            ProductStatus      `json:"status"`
	SizeGuide                         *string            `json:"sizeGuide"`
	SubTitle                          string             `json:"subTitle"`
	Answers                           []interface{}      `json:"answers"`
	StoreName                         StoreName          `json:"storeName"`
	Name                              string             `json:"name"`
	PremiumEndDate                    interface{}        `json:"premiumEndDate"`
	IsPremium                         bool               `json:"isPremium"`
	Badges                            []string           `json:"badges"`
	MainImage                         MainImageElement   `json:"mainImage"`
	Charity                           bool               `json:"charity"`
	LaunchDayDeliveryMethod           *bool              `json:"launchDayDeliveryMethod"`
	SizeBreakdown                     []interface{}      `json:"sizeBreakdown"`
	Options                           []Option           `json:"options"`
	HideFromLaunchesTable             bool               `json:"hideFromLaunchesTable"`
	ID                                string             `json:"ID"`
	Question                          interface{}        `json:"question"`
	InformationBlocks                 []InformationBlock `json:"informationBlocks"`
	WidgetImage                       *MainImageElement  `json:"widgetImage"`
	Brand                             ProductBrand       `json:"brand"`
	LaunchDayDeliveryMembersExclusive *bool              `json:"launchDayDeliveryMembersExclusive"`
	Filters                           []Filter           `json:"filters"`
	PaymentGateways                   []PaymentGateway   `json:"paymentGateways"`
	LaunchDate                        string             `json:"launchDate"`
	Gender                            Gender             `json:"gender"`
	LaunchDayDelay                    *int64             `json:"launchDayDelay"`
	RelatedProducts                   interface{}        `json:"relatedProducts"`
	BlogCopy                          string             `json:"blogCopy"`
	LaunchStartDate                   string             `json:"launchStartDate"`
	Audio                             interface{}        `json:"audio"`
	HideBeforeRelease                 interface{}        `json:"hideBeforeRelease"`
	LaunchType                        LaunchType         `json:"launchType"`
	Plu                               string             `json:"PLU"`
	Images                            []MainImageElement `json:"images"`
	LoserSKU                          interface{}        `json:"loserSKU"`
	Tandc                             interface{}        `json:"tandc"`
	Link                              interface{}        `json:"link"`
	Description                       string             `json:"description"`
	Price                             Price              `json:"price"`
	Video                             interface{}        `json:"video"`
}

type Filter struct {
	Type string `json:"type"`
}

type MainImageElement struct {
	Original string `json:"original"`
}

type InformationBlock struct {
	Content     string      `json:"content"`
	DisplayName DisplayName `json:"displayName"`
}

type Option struct {
	OptionID string `json:"optionID"`
	Name     string `json:"name"`
	Sku      string `json:"SKU"`
}

type Price struct {
	Currency Currency `json:"currency"`
	Amount   string   `json:"amount"`
}

type SizeGuide struct {
	Value     string    `json:"value"`
	Category  string    `json:"category"`
	Brand     string    `json:"brand"`
	StoreName StoreName `json:"storeName"`
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Gender    Gender    `json:"gender"`
}

type Theme struct {
	CSS           interface{}   `json:"css"`
	Logo          string        `json:"logo"`
	ThemeColour   string        `json:"themeColour"`
	PrimaryButton PrimaryButton `json:"_primaryButton"`
	SplashImage   interface{}   `json:"splashImage"`
}

type PrimaryButton struct {
	BackgroundColor string `json:"backgroundColor"`
	BorderColor     string `json:"borderColor"`
	Color           string `json:"color"`
}

type UpcomingElement struct {
	Name                  string            `json:"name"`
	SubTitle              SubTitle          `json:"subTitle"`
	Price                 Price             `json:"price"`
	MaxPrice              Price             `json:"maxPrice"`
	LaunchDate            string            `json:"launchDate"`
	Plu                   string            `json:"PLU"`
	LaunchesOn            LaunchesOn        `json:"launchesOn"`
	SizeBreakdown         []SizeBreakdown   `json:"sizeBreakdown"`
	Link                  *string           `json:"link"`
	Status                UpcomingStatus    `json:"status"`
	RelatedProducts       []string          `json:"relatedProducts"`
	Brand                 UpcomingBrand     `json:"brand"`
	HideFromLaunchesTable bool              `json:"hideFromLaunchesTable"`
	MainImage             *MainImageElement `json:"mainImage,omitempty"`
}

type StoreName string

const (
	Size StoreName = "size"
)

type ActionType string

const (
	Product ActionType = "product"
)

type Effect string

const (
	Default Effect = "default"
)

type Tag string

const (
	Empty                                     Tag = ""
	ImageCreditedToAdidasArchiveStudioWaldeck Tag = "Image credited to adidas Archive/Studio Waldeck"
)

type ContentType string

const (
	HTML  ContentType = "html"
	Image ContentType = "image"
	OptIn ContentType = "opt-in"
	Page  ContentType = "page"
	Video ContentType = "video"
)

type PageType string

const (
	PagePage   PageType = "page-page"
	PageSlider PageType = "page-slider"
)

type ProductBrand string

const (
	BrandOthers           ProductBrand = "others"
	Jordan                ProductBrand = "jordan"
	PurpleAdidasOriginals ProductBrand = "adidas-originals"
	PurpleNike            ProductBrand = "nike"
)

type Gender string

const (
	GenderOthers Gender = "others"
	Infant       Gender = "infant"
	Junior       Gender = "junior"
	Men          Gender = "men"
	Women        Gender = "women"
)

type DisplayName string

const (
	Delivery             DisplayName = "Delivery"
	InternationalReturns DisplayName = "International Returns"
	UKReturns            DisplayName = "UK Returns"
)

type LaunchType string

const (
	Full   LaunchType = "full"
	Raffle LaunchType = "raffle"
)

type PaymentGateway string

const (
	Adyen  PaymentGateway = "Adyen"
	Klarna PaymentGateway = "Klarna"
)

type Currency string

const (
	Gbp Currency = "GBP"
)

type ProductStatus string

const (
	Archived  ProductStatus = "archived"
	Available ProductStatus = "available"
	Complete  ProductStatus = "complete"
)

type UpcomingBrand string

const (
	FluffyAdidasOriginals UpcomingBrand = "adidas-originals"
	FluffyNike            UpcomingBrand = "nike"
	NewBalance            UpcomingBrand = "new-balance"
	Salomon               UpcomingBrand = "salomon"
)

type LaunchesOn string

const (
	App     LaunchesOn = "app"
	Website LaunchesOn = "website"
)

type SizeBreakdown string

const (
	Mens   SizeBreakdown = "mens"
	Womans SizeBreakdown = "womans"
)

type UpcomingStatus string

const (
	Upcoming UpcomingStatus = "upcoming"
)

type SubTitle string

const (
	MenS       SubTitle = "Men's"
	MenSWomenS SubTitle = "Men's & Women's"
	WomenS     SubTitle = "Women's"
)

func main() {
	url := "https://mosaic-platform.jdmesh.co/stores/size/content?api_key=0ce5f6f477676d95569067180bc4d46d&channel=iphone-mosaic"
	webhookURL := ""

	alreadyPinged := make(map[string]bool)

	for {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}

		var welcome Welcome
		err = json.Unmarshal(body, &welcome)
		if err != nil {
			log.Fatal(err)
		}

		webhookParams := &discordgo.WebhookParams{}

		for _, product := range welcome.Products {
			if product.Status == "available" && (product.SubTitle == "Women's" || product.SubTitle == "Men's") {
				if alreadyPinged[product.ID] {
					continue
				}

				embed := &discordgo.MessageEmbed{
					Title:     fmt.Sprintf("**%s**", product.Name),
					Color:     16646145,
					Fields:    make([]*discordgo.MessageEmbedField, 0),
					Thumbnail: &discordgo.MessageEmbedThumbnail{URL: product.MainImage.Original},
					Footer: &discordgo.MessageEmbedFooter{
						Text:    "FluxyIO Custom",
						IconURL: "https://image.noelshack.com/fichiers/2023/35/6/1693673102-image-3.png",
					},
				}
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: "SubTitle", Value: product.SubTitle, Inline: true})
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: "Status", Value: string(product.Status), Inline: true})
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: "Price", Value: product.Price.Amount + " " + string(product.Price.Currency), Inline: true})

				launchTime, _ := time.Parse(time.RFC3339, product.LaunchDate)
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: "Launch Date", Value: fmt.Sprintf("<t:%d:f>", launchTime.Unix()), Inline: true})

				optionNames := ""
				optionSkus := ""
				for _, option := range product.Options {
					optionNames += fmt.Sprintf("`%s`\n", option.Name)
					optionSkus += fmt.Sprintf("`%s`\n", option.Sku)
				}
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: "Option Names", Value: optionNames, Inline: true})
				embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{Name: "Option SKUs", Value: optionSkus, Inline: true})

				webhookParams.Embeds = []*discordgo.MessageEmbed{embed}

				err := sendWebhook(webhookURL, webhookParams)
				if err != nil {
					fmt.Println("error sending webhook", err)
				}

				alreadyPinged[product.ID] = true
				time.Sleep(time.Second)
			}
		}

		fmt.Printf("checking...")
		time.Sleep(1 * time.Minute)

	}
}

func sendWebhook(url string, params *discordgo.WebhookParams) (err error) {
	var payload []byte
	payload, err = json.Marshal(params)
	if err != nil {
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(resp.Body)
		err = fmt.Errorf("HTTP request failed: %s", string(body))
	}
	return
}
