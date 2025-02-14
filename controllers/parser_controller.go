package controllers

import (
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"summarize-transactions/dto"
	"time"
)

type CSVType int

const (
	Unknown CSVType = iota
	Nubank
	XP
)

type NubankCsv struct {
	Amount float64 `json:"amount"`
	Title  string  `json:"title"`
	Date   string  `json:"date"`
}

type XpCsv struct {
	Data            string  `json:"data"`
	Estabelecimento string  `json:"estabelecimento"`
	Portador        string  `json:"portador"`
	Parcela         string  `json:"parcela"`
	Valor           float64 `json:"valor"`
}

type StandardCsvType struct {
	Date   string  `json:"date"`
	Title  string  `json:"title"`
	Amount float64 `json:"amount"`
}

type ParserController struct {
	BaseController
}

func NewParserController() *ParserController {
	return &ParserController{
		BaseController: BaseController{},
	}
}

func (controller *ParserController) ParseCsv() gin.HandlerFunc {
	return controller.Run(
		func(c *gin.Context) {
			var payload dto.ParseCsvRequest

			if err := c.ShouldBindJSON(&payload); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
				return
			}

			parsedData, csvType, err := parseCsv(payload.CsvUrl)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			standardData := mapToStandard(parsedData, csvType)

			controller.Ok(c, standardData)
		})
}

func detectCSVType(headers []string) CSVType {
	nubankHeaders := []string{"date", "title", "amount"}
	xpHeaders := []string{"Data", "Estabelecimento", "Portador", "Valor", "Parcela"}

	if equalHeaders(headers, nubankHeaders) {
		return Nubank
	} else if equalHeaders(headers, xpHeaders) {
		return XP
	}
	return Unknown
}

func equalHeaders(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	return slices.Equal(a, b)
}

func cleanCurrency(value string) string {
	re := regexp.MustCompile(`R\$\s?-?([0-9]{1,3}(\.[0-9]{3})*),([0-9]{2})`)
	return re.ReplaceAllStringFunc(value, func(match string) string {
		cleaned := strings.ReplaceAll(match, "R$", "")
		cleaned = strings.TrimSpace(cleaned)
		cleaned = strings.ReplaceAll(cleaned, ".", "")  // Remove thousand separators
		cleaned = strings.ReplaceAll(cleaned, ",", ".") // Convert decimal separator
		return cleaned
	})
}

func parseCsv(url string) (interface{}, CSVType, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, Unknown, fmt.Errorf("failed to download CSV: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, Unknown, fmt.Errorf("failed to fetch CSV, status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, Unknown, fmt.Errorf("failed to read response body: %v", err)
	}

	csvContent := strings.ReplaceAll(string(body), ";", ",")
	csvContent = strings.ReplaceAll(csvContent, "\ufeff", "")
	csvContent = cleanCurrency(csvContent)
	reader := csv.NewReader(strings.NewReader(csvContent))
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, Unknown, fmt.Errorf("failed to read CSV: %v", err)
	}

	if len(rows) < 2 {
		return nil, Unknown, fmt.Errorf("CSV must have at least a header row and one data row")
	}

	//headers := rows[0]
	headers := []string{"Data", "Estabelecimento", "Portador", "Valor", "Parcela"}

	csvType := detectCSVType(headers)

	var jsonData interface{}

	switch csvType {
	case Nubank:
		var data []NubankCsv
		for _, row := range rows[1:] {
			amount, _ := strconv.ParseFloat(row[0], 64)
			data = append(data, NubankCsv{
				Amount: amount,
				Title:  row[1],
				Date:   row[2],
			})
		}
		jsonData = data
	case XP:
		var data []XpCsv
		for _, row := range rows[1:] {
			value, _ := strconv.ParseFloat(row[3], 64)
			data = append(data, XpCsv{
				Data:            row[0],
				Estabelecimento: row[1],
				Portador:        row[2],
				Valor:           value,
				Parcela:         row[4],
			})
		}
		jsonData = data
	default:
		return nil, Unknown, fmt.Errorf("unknown CSV format")
	}

	return jsonData, csvType, nil
}

func mapToStandard(parsedData interface{}, csvType CSVType) []StandardCsvType {
	var standardData []StandardCsvType

	switch csvType {
	case Nubank:
		for _, entry := range parsedData.([]NubankCsv) {
			standardData = append(standardData, StandardCsvType{
				Date:   entry.Date,
				Title:  entry.Title,
				Amount: entry.Amount,
			})
		}
	case XP:
		for _, entry := range parsedData.([]XpCsv) {
			t, err := time.Parse("02/01/2006", entry.Data)
			if err != nil {
				log.Printf("failed to parse date: %v", err)
			}

			formattedDate := t.Format("2006-01-02")

			standardData = append(standardData, StandardCsvType{
				Date:   formattedDate,
				Title:  entry.Estabelecimento,
				Amount: entry.Valor,
			})
		}
	default:
		panic("unhandled default case")
	}
	return standardData
}

var _ IBaseController = (*ParserController)(nil)
