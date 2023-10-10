package report

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

type OrgInfo struct {
	SystemTableName     string  `protobuf:"bytes,9999,opt,name=system_table_name,json=system_table_name,table_name=org,proto3" json:"system_table_name,omitempty"`
	Id                  int64   `json:"id"`
	Name                string  `json:"name"`
	NickName            string  `json:"nickName"`
	BusinessName        string  `json:"businessName"`
	PhoneNumberOne      string  `json:"phoneNumberOne"`
	PhoneNumberTwo      string  `json:"phoneNumberTwo"`
	IconData            string  `json:"iconData"`
	Address             string  `json:"address"`
	Website             string  `json:"website"`
	Email               string  `json:"email"`
	Locale              string  `json:"locale"`
	LocaleForeign       string  `json:"localeForeign"`
	TimeZone            string  `json:"timeZone"`
	DiffHour            float64 `json:"diffHour"`
	TaxCode             string  `json:"taxCode"`
	AnniversaryDate     int64   `json:"anniversaryDate"`
	AmStart             int64   `json:"amStart"`
	AmEnd               int64   `json:"amEnd"`
	PmStart             int64   `json:"pmStart"`
	PmEnd               int64   `json:"pmEnd"`
	WorkHour            int64   `json:"workHour"`
	BreakTime           int32   `json:"breakTime"`
	Representative      string  `json:"representative"`
	RepresentativeTitle string  `json:"representativeTitle"`
	FontFamily          string  `json:"fontFamily"`
}

type ReportRequest struct {
	TitleKey           string      `json:"titleKey"`
	DbTemplate         int32       `json:"dbTemplate"`
	XlsxTemplate       string      `json:"xlsxTemplate"`
	HtmlHeadTemplate   string      `json:"htmlHeadTemplate"`
	HtmlHeaderTemplate string      `json:"htmlHeaderTemplate"`
	HtmlBodyTemplate   string      `json:"htmlBodyTemplate"`
	HtmlFooterTemplate string      `json:"htmlFooterTemplate"`
	WatermarkValue     string      `json:"watermarkValue"`
	ShowHeaderFooter   int32       `json:"showHeaderFooter"`
	PrintHead          int32       `json:"printHead"`
	Format             string      `json:"format"`
	Landscape          int32       `json:"landscape"`
	MarginLeft         int32       `json:"marginLeft"`
	MarginRight        int32       `json:"marginRight"`
	MarginTop          int32       `json:"marginTop"`
	MarginBottom       int32       `json:"marginBottom"`
	UsedDefaultHead    int32       `json:"usedDefaultHead"`
	UsedDefaultHeader  int32       `json:"usedDefaultHeader"`
	UsedDefaultFooter  int32       `json:"usedDefaultFooter"`
	FontFamily         string      `json:"fontFamily"`
	Params             string      `json:"params"`
	Width              float64     `json:"width"`
	Height             float64     `json:"height"`
	Unit               string      `json:"unit"`
	FontSize           string      `json:"fontSize"`
	CompanyId          int64       `json:"companyId"`
	BranchId           int64       `json:"branchId"`
	Code               string      `json:"code"`
	Name               string      `json:"name"`
	Data               interface{} `json:"data"`
	Background         []byte      `json:"background"`
	BgValid            []byte      `json:"bgValid"`
	BgInvalid          []byte      `json:"bgInvalid"`
	Formula            string      `json:"formula"`
	Filename           string      `json:"filename"`
	Data1              string      `json:"data1"`
	Data2              string      `json:"data2"`
	Data3              string      `json:"data3"`
	Id                 int64       `json:"id"`
	FilesystemId       int64       `json:"filesystemId"`
	FilePath           string      `json:"filePath"`
	FileName           string      `json:"fileName"`
	FileType           string      `json:"fileType"`
	FileMime           string      `json:"fileMime"`
	FileMd5            string      `json:"fileMd5"`
	FileSize           string      `json:"fileSize"`
	FileData           []byte      `json:"fileData"`
	IsPrintHead        int32       `json:"isPrintHead"`
	PrintOpts          interface{} `json:"printOpts"`
	Obj                interface{} `json:"obj"`
}

type ReportTemplate struct {
	Id                     int64   `json:"id"`
	CompanyId              int64   `json:"companyId"`
	BranchId               int64   `json:"branchId"`
	Code                   string  `json:"code"`
	Name                   string  `json:"name"`
	TitleKey               string  `json:"titleKey"`
	Sort                   int64   `json:"sort"`
	XlsxTemplate           string  `json:"xlsxTemplate"`
	HtmlHeadTemplate       string  `json:"htmlHeadTemplate"`
	HtmlHeaderTemplate     string  `json:"htmlHeaderTemplate"`
	HtmlBodyTemplate       string  `json:"htmlBodyTemplate"`
	HtmlFooterTemplate     string  `json:"htmlFooterTemplate"`
	Disabled               int32   `json:"disabled"`
	ShowHeaderFooter       int32   `json:"showHeaderFooter"`
	PrintHead              int32   `json:"printHead"`
	Format                 string  `json:"format"`
	Landscape              int32   `json:"landscape"`
	MarginLeft             int32   `json:"marginLeft"`
	MarginRight            int32   `json:"marginRight"`
	MarginTop              int32   `json:"marginTop"`
	MarginBottom           int32   `json:"marginBottom"`
	SharedHeadTemplateId   int64   `json:"sharedHeadTemplateId"`
	SharedHeaderTemplateId int64   `json:"sharedHeaderTemplateId"`
	SharedFooterTemplateId int64   `json:"sharedFooterTemplateId"`
	ContractId             int64   `json:"contractId"`
	SavePdfFile            int32   `json:"savePdfFile"`
	PdfFileCategory        string  `json:"pdfFileCategory"`
	FontFamily             string  `json:"fontFamily"`
	XlsxTemplateFileName   string  `json:"xlsxTemplateFileName"`
	Params                 string  `json:"params"`
	Width                  float64 `json:"width"`
	Height                 float64 `json:"height"`
	Unit                   string  `json:"unit"`
	FontSize               string  `json:"fontSize"`
	Background             []byte  `json:"background"`
	BgValid                []byte  `json:"bgValid"`
	BgInvalid              []byte  `json:"bgInvalid"`
	Formula                string  `json:"formula"`
	Filename               string  `json:"filename"`
	Data1                  string  `json:"data1"`
	Data2                  string  `json:"data2"`
	Data3                  string  `json:"data3"`
}

type SharedHtmlTemplate struct {
	Id                 int64  `json:"id"`
	CompanyId          int64  `json:"companyId"`
	BranchId           int64  `json:"branchId"`
	Name               string `json:"name"`
	HtmlHeadTemplate   string `json:"htmlHeadTemplate"`
	HtmlHeaderTemplate string `json:"htmlHeaderTemplate"`
	HtmlFooterTemplate string `json:"htmlFooterTemplate"`
	Sort               int64  `json:"sort"`
	Disabled           int32  `json:"disabled"`
}

type XlsxSheet struct {
	SheetName string `json:"sheetName"`

	Data interface{} `json:"data"`
}

func LocalImageToBase64(fullFilePath string) (string, error) {
	dataBytes, err := ioutil.ReadFile(fullFilePath)
	if err != nil {
		return "", err
	}

	return DetectImageType(dataBytes) + base64.StdEncoding.EncodeToString(dataBytes), nil
}

func RemoteImageToBase64(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return DetectImageType(bodyBytes) + base64.StdEncoding.EncodeToString(bodyBytes), nil
}

func MemoryImageToBase64(dataBytes []byte) string {
	return DetectImageType(dataBytes) + base64.StdEncoding.EncodeToString(dataBytes)
}

func DetectImageType(dataBytes []byte) string {
	mimeType := http.DetectContentType(dataBytes)
	switch mimeType {
	case "image/jpg":
		return "data:image/jpg;base64,"
	case "image/jpeg":
		return "data:image/jpeg;base64,"
	case "image/png":
		return "data:image/png;base64,"
	case "image/bmp":
		return "data:image/bmp;base64,"
	case "image/gif":
		return "data:image/gif;base64,"
	case "image/ief":
		return "data:image/ief;base64,"
	case "image/pipeg":
		return "data:image/pipeg;base64,"
	case "image/svg+xml":
		return "data:image/svg+xml;base64,"
	case "image/x-icon":
		return "data:image/x-icon;base64,"
	}
	return "Unkown type"
}
