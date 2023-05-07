package wordDB

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type Word struct {
	gorm.Model
	Literal   string
	Frequency int
}

func CreateWordTable(db *gorm.DB) {
	err := db.AutoMigrate(&Word{})
	if err != nil {
		fmt.Println("Error migrating Word table:", err)
		return
	}

	var count int64
	db.Model(&Word{}).Count(&count)
	if count > 0 {
		return
	}
	file, err := os.Open("./count_1w.txt")
	if err != nil {
		fmt.Println("Error opening count_1w.txt:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		if len(parts) != 2 {
			fmt.Println("Invalid line format:", line)
			continue
		}

		literal := parts[0]
		var frequency, err = strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Error converting frequency to integer:", err)
			continue
		}

		word := &Word{
			Literal:   literal,
			Frequency: frequency,
		}

		result := db.Create(word)
		if result.Error != nil {
			fmt.Printf("Error while inserting wordDB: %v\n", result.Error)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file:", err)
	}
}

func GetFrequency(db *gorm.DB, literal string) (int, error) {
	var word Word
	err := db.Where("literal = ?", literal).First(&word).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, nil
		}
		return 0, err
	}
	return word.Frequency, nil
}
func GetFrequencies(db *gorm.DB, literals []string) ([]int, error) {
	var frequencies []int
	for _, literal := range literals {
		frequency, err := GetFrequency(db, literal)
		if err != nil {
			return nil, err
		}
		frequencies = append(frequencies, frequency)
	}
	return frequencies, nil
}
