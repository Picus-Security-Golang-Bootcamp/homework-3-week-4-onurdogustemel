package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/onurdogustemel/file_operations/file_operations/csvReader"
	"github.com/onurdogustemel/file_operations/file_operations/dbConnect"
	"github.com/onurdogustemel/file_operations/file_operations/library"
	"gorm.io/gorm"
)

type AuthorRepo struct {
	datb *gorm.DB
}

func RepoAuthorlist(datb *gorm.DB) *AuthorRepo {
	return &AuthorRepo{datb: datb}
}

func (a *AuthorRepo) MigrateAuthors() {
	a.datb.AutoMigrate(&library.Authors{})
}

func (a *AuthorRepo) InsertAuthorDataIntoPostgres() {
	authorlists, err := csvReader.ReadCSVFileForAuthor("authorList.csv")
	if err != nil {
		log.Fatal(err)
	}

	for _, aut := range authorlists {

		a.datb.Where(library.Authors{Author: aut.Author}).Attrs(library.Authors{Author: aut.Author}).FirstOrCreate(&aut)

	}
}

func (a *AuthorRepo) CreateAuthors(autr library.Authors) error {
	result := a.datb.Create(&autr)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

type BookRepo struct {
	db *gorm.DB
}

func RepoBooklist(db *gorm.DB) *BookRepo {
	return &BookRepo{db: db}
}

func (b *BookRepo) Migrate() {
	b.db.AutoMigrate(&library.Books{})
}

func (b *BookRepo) InsertFileDataIntoPostgres() {
	booklists, err := csvReader.ReadCSVFile("booklist.csv")

	if err != nil {
		log.Fatal(err)
	}

	for _, book := range booklists {

		b.db.Where(library.Books{Isbn: book.Isbn}).Attrs(library.Books{Isbn: book.Isbn, Title: book.Title}).FirstOrCreate(&book)

	}
}

func (b *BookRepo) List() []library.Books {
	var newbookSlice []library.Books
	b.db.Find(&newbookSlice)

	return newbookSlice
}

func (b *BookRepo) GetByID(id int) []library.Books {
	var newbookSlice []library.Books
	b.db.Where(`"id" = ?`, id).Find(&newbookSlice)

	return newbookSlice
}

func (b *BookRepo) FindByName(title string) []library.Books {
	var newbookSlice []library.Books
	b.db.Where(`"Title" ILIKE ?`, title).Find(&newbookSlice)

	return newbookSlice
}

func (b *BookRepo) GetBooksWithAuthor(authorname string) []library.Books {
	var newbookSlice []library.Books
	b.db.Where(`"Author" LIKE ?`, "%"+authorname+"%").Find(&newbookSlice)

	return newbookSlice
}

func (b *BookRepo) GetAuthorsWithBooks(booktitle string) []library.Books {

	var newbookSlice []library.Books
	b.db.Where(`"Title" LIKE ?`, "%"+booktitle+"%").Find(&newbookSlice)

	return newbookSlice

}

func (b *BookRepo) CreateBooks(book library.Books) error {
	result := b.db.Create(&book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (b *BookRepo) DeleteByID(id int) error {
	result := b.db.Delete(&library.Books{}, id)

	if result != nil {
		return result.Error
	}
	return nil
}

func main() {

	err1 := godotenv.Load()

	if err1 != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := dbConnect.CreatePostgresConnection()
	if err != nil {
		log.Fatal("Database cannot be initialized")
	}

	log.Println("Postgres connected")

	bookstore := RepoBooklist(db)
	bookstore.Migrate()
	bookstore.InsertFileDataIntoPostgres()

	//bookstore.DeleteByID(8)

	//fmt.Println(bookstore.List())
	//fmt.Println(len(bookstore.List()))
	//fmt.Println(bookstore.GetByID(5))
	//fmt.Println(bookstore.FindByName("the giver"))
	//fmt.Println(bookstore.GetBooksWithAuthor("William"))

	bookstore.CreateBooks(library.Books{Title: "Romeo and Juliet",
		Page:           600,
		Author:         "William Shakespeare",
		NumberOfStocks: 15,
		Price:          10.99,
		StockCode:      "2854",
		Isbn:           "9780320501539"})

	//fmt.Println(bookstore.GetAuthorsWithBooks("The Giver"))

	authorStore := RepoAuthorlist(db)
	authorStore.MigrateAuthors()
	authorStore.InsertAuthorDataIntoPostgres()

	authorStore.CreateAuthors(library.Authors{Author: "William Shakespeare"})
}
