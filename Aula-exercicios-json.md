# Aula de exercícios de manipulação de arquivos, json

No terminal criar a pasta bookstore: mkdir bookstore 

cd bookstore  

code .

Criar main

Criar arquivo bookstore.json 

Criar pasta model

Adicionar arquivo bookstore.go na pasta model, package model 

Digite no terminal:
```
go mod init github.com/user/bookstore
go mod tidy
```

## Exercicios

1) Ler o arquivo json e atribuir a struts de bookstore
2) Criar um slice de books
3) Ordenar os books por titulo
4) Ordenar os books pelo menor preco
5) Pegar apenas os books com rating 5
6) Escreva um arquivo com os books com rating 5

## bookstore.json
```json
{
  "location": "Online",
  "sections": [
    {
      "section_name": "Fiction",
      "categories": [
        {
          "category_name": "Fantasy",
          "books": [
            {
              "title": "The Enchanted Forest",
              "author": {
                "name": "Jane Doe",
                "bio": "Jane Doe is a renowned fantasy author.",
                "author_id": 1
              },
              "publication_year": 2021,
              "price": 19.99,
              "ISBN": "978-3-16-148410-0",
              "available_copies": 12,
              "description": "A magical journey into a world of mystery and wonder.",
              "ratings": [
                {
                  "user_id": 101,
                  "rating": 5,
                  "comment": "An amazing adventure!"
                },
                {
                  "user_id": 102,
                  "rating": 4,
                  "comment": "Loved the characters."
                }
              ]
            },
            {
              "title": "A mystic River",
              "author": {
                "name": "John Smith",
                "bio": "John Smith writes thrilling fantasy novels.",
                "author_id": 2
              },
              "publication_year": 2019,
              "price": 15.99,
              "ISBN": "978-1-23-456789-0",
              "available_copies": 8,
              "description": "A tale of magic and mystery along the river.",
              "ratings": [
                {
                  "user_id": 103,
                  "rating": 4,
                  "comment": "A captivating story."
                }
              ]
            }
          ]
        },
        {
          "category_name": "Science Fiction",
          "books": [
            {
              "title": "Galactic Travels",
              "author": {
                "name": "John Smith",
                "bio": "An expert in speculative fiction and astrophysics.",
                "author_id": 2
              },
              "publication_year": 2020,
              "price": 15.99,
              "ISBN": "978-1-23-456789-7",
              "available_copies": 5,
              "description": "Exploring the far reaches of the galaxy.",
              "ratings": []
            }
          ]
        }
      ]
    },
    {
      "section_name": "Non-Fiction",
      "categories": [
        {
          "category_name": "Biography",
          "books": [
            {
              "title": "The Life of Jane Doe",
              "author": {
                "name": "Emily Johnson",
                "bio": "Emily Johnson is a biographer known for her detailed and engaging writing.",
                "author_id": 3
              },
              "publication_year": 2018,
              "price": 12.99,
              "ISBN": "978-0-12-345678-9",
              "available_copies": 10,
              "description": "An in-depth look at the life of Jane Doe.",
              "ratings": [
                {
                  "user_id": 104,
                  "rating": 5,
                  "comment": "A fascinating biography."
                }
              ]
            }
          ]
        }
      ]
    }
  ]
}

```
## bookstore.go
```go
package model

type Bookstore struct {
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Sections []Section `json:"sections"`
}

type Section struct {
	SectionName string     `json:"section_name"`
	Categories  []Category `json:"categories"`
}

type Category struct {
	CategoryName string `json:"category_name"`
	Books        []Book `json:"books"`
}

type Book struct {
	Title           string   `json:"title"`
	Author          Author   `json:"author"`
	PublicationYear int      `json:"publication_year"`
	Price           float64  `json:"price"`
	ISBN            string   `json:"isbn"`
	AvailableCopies int      `json:"available_copies"`
	Description     string   `json:"description"`
	Ratings         []Rating `json:"ratings"`
}

type Author struct {
	Name     string `json:"name"`
	Bio      string `json:"bio"`
	AuthorID int    `json:"author_id"`
}

type Rating struct {
	UserID  int    `json:"user_id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

)
```
