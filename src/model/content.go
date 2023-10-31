package model

type Media struct {
	URI        string `json:"uri" dynamodbav:"uri"`
	Dimensions string `json:"dimensions" dynamodbav:"dimensions"`
	Size       string `json:"size" dynamodbav:"size"`
	MimeType   string `json:"mime_type" dynamodbav:"mime_type"`
}

type Content struct {
	CID                  string       `json:"cid" dynamodbav:"cid"`
	Image                string       `json:"image" dynamodbav:"image"`
	Symbol               string       `json:"symbol" dynamodbav:"symbol"`
	Name                 string       `json:"name" dynamodbav:"name"`
	Description          string       `json:"description" dynamodbav:"description"`
	SellerFeeBasisPoints int          `json:"seller_fee_basis_points" dynamodbav:"seller_fee_basis_points"`
	ArtistName           string       `json:"artist_name" dynamodbav:"artist_name"`
	Properties           ContentProps `json:"properties" dynamodbav:"properties"`
	Attributes           []Attribute  `json:"attributes" dynamodbav:"attributes"`
	YearCreated          int          `json:"year_created" dynamodbav:"year_created"`
	CreatedBy            string       `json:"created_by" dynamodbav:"created_by"`
	Artist               string       `json:"artist" dynamodbav:"artist"`
	Edition              int          `json:"edition" dynamodbav:"edition"`
	Media                Media        `json:"media" dynamodbav:"media"`
}

type ContentProps struct {
	Files    []ContentFile `json:"files" dynamodbav:"files"`
	Category string        `json:"category" dynamodbav:"category"`
}

type Attribute struct {
	Value     string `json:"value" dynamodbav:"value"`
	TraitType string `json:"trait_type" dynamodbav:"trait_type"`
}

type ContentFile struct {
	URI  string `json:"uri" dynamodbav:"uri"`
	Type string `json:"type" dynamodbav:"type"`
}
