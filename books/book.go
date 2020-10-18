package books

type Book struct {
  Title string
  Author string
  Pages int
};

func (self Book) CategoryByLength() string {
  if self.Pages > 300 {
    return "NOVEL";
  } else {
    return "SHORT STORY";
  }
}
