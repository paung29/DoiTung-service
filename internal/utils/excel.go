package utils

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/doitung/DoiTung-service/internal/models"
	"github.com/xuri/excelize/v2"
)

type Sheet struct {
	Name string
	Rows [][]interface{}
}

var invalidSheetNameChars = regexp.MustCompile(`[\[\]\*:/\\?]`)

func BuildWorkBook(sheets []Sheet) ([]byte, error) {
	file := excelize.NewFile()
	defer file.Close()

	usedSheetNames := make(map[string]bool)

	for index, sheet := range sheets {
		sheetName := SafeSheetName(sheet.Name, index+1, usedSheetNames)

		if index == 0 {
			if err := file.SetSheetName("Sheet1", sheetName); err != nil {
				return nil, err
			}
		} else {
			if _, err := file.NewSheet(sheetName); err != nil {
				return nil, err
			}
		}

		for rowIndex, row := range sheet.Rows {
			cell, err := excelize.CoordinatesToCellName(1, rowIndex+1)
			if err != nil {
				return nil, err
			}
			if err := file.SetSheetRow(sheetName, cell, &row); err != nil {
				return nil, err
			}
		}
	}
	buffer, err := file.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil

}

func SafeSheetName(name string, noNameSheetIndex int, usedSheetNames map[string]bool) string {
	name = strings.TrimSpace(invalidSheetNameChars.ReplaceAllString(name, "-"))
	if name == "" {
		name = fmt.Sprintf("Sheet%d", noNameSheetIndex)
	}

	if len([]rune(name)) > 31 {
		name = string([]rune(name)[:31])
	}

	OriginalName := name
	for suffix := 2; usedSheetNames[name]; suffix++ {
		suffixText := fmt.Sprintf("-%d", suffix)
		maxNameLength := 31 - len(suffixText)

		if len([]rune(OriginalName)) > maxNameLength {
			name = string([]rune(OriginalName)[:maxNameLength]) + suffixText
		} else {
			name = OriginalName + suffixText
		}
	}

	usedSheetNames[name] = true
	return name
}

type exportSheet struct {
	Name string
	Rows [][]interface{}
}

type zoneGroup[T any] struct {
	Zone  models.Zone
	Items []T
}

func GroupByZone[T any](items []T, getZone func(T) models.Zone) []zoneGroup[T] {
	groupedItems := make(map[uint][]T)
	zones := make(map[uint]models.Zone)

	for _, item := range items {
		zone := getZone(item)
		groupedItems[zone.ZoneID] = append(
			groupedItems[zone.ZoneID],
			item,
		)
		zones[zone.ZoneID] = zone
	}

	groups := make([]zoneGroup[T], 0, len(groupedItems))

	for zoneID, zoneItems := range groupedItems {
		groups = append(groups, zoneGroup[T]{
			Zone:  zones[zoneID],
			Items: zoneItems,
		})
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i].Zone.ZoneNo < groups[j].Zone.ZoneNo
	})

	return groups
}

func FirstOrZero[T any](items []T) T {
	var zero T
	if len(items) == 0 {
		return zero
	}
	return items[0]
}
