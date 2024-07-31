package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/c-bata/go-prompt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	Id    int
	Code  string
	Price uint
}

func handleExit() {
	rawModeOff := exec.Command("/bin/stty", "-raw", "echo")
	rawModeOff.Stdin = os.Stdin
	_ = rawModeOff.Run()
	rawModeOff.Wait()
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "create", Description: "Create a product"},
		{Text: "list", Description: "List products"},
		{Text: "migrate", Description: "Migrate"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	defer handleExit()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	// db.AutoMigrate(&Product{})

	fmt.Println("Please select action.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)

	// Create
	if t == "migrate" {
		db.Create(&Product{Code: "D42", Price: 100})
	}
	// Read

	var product Product
	if t == "create" {
		db.First(&product, 1)                 // find product with integer primary key
		db.First(&product, "code = ?", "D42") // find product with code D42

		// Update - update product's price to 200
		db.Model(&product).Update("Price", 200)
		// Update - update multiple fields
		db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
		db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	}
	if t == "list" {
		db.Debug().First(&product)
		fmt.Println(product)
	}
	// Delete - delete product
	// db.Delete(&product, 1)
}
