package books_test

import (
  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
  "github.com/mokhan/learn.go/books"
)

var _ = Describe("Book", func() {
  var (
    long_book books.Book
    short_book books.Book
  )

  BeforeEach(func() {
    long_book = books.Book{
      Title: "Les Miserables",
      Author: "Victor Hugo",
      Pages: 1488,
    }

    short_book = books.Book{
      Title: "Fox In Socks",
      Author: "Dr. Seuss",
      Pages: 24,
    }
  });

  Describe("Categorizing book length", func(){
    Context("With more than 300 pages", func(){
      It("should be a novel", func(){
        Expect(long_book.CategoryByLength()).To(Equal("NOVEL"))
      })
    })

    Context("With fewer than 300 pages", func(){
      It("should be a short story", func(){
        Expect(short_book.CategoryByLength()).To(Equal("SHORT STORY"))
      })
    })
  })
})
