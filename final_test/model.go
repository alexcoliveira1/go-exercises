package main

type User struct {
	Email     string     `json:"email"`
	Questions []Question `json:"questions"`
}

type Answer struct {
	ID            string   `json:"id"`
	Body          string   `json:"body"`
	UserEmail     string   `json:"userEmail"`
	Votes         int      `json:"votes"`
	UpvotesList   []string `json:"-"`
	DownvotesList []string `json:"-"`
}

type Question struct {
	ID            string   `json:"id"`
	Content       string   `json:"content"`
	Answer        *Answer  `json:"answer"`
	UserEmail     string   `json:"userEmail"`
	Votes         int      `json:"votes"`
	UpvotesList   []string `json:"-"`
	DownvotesList []string `json:"-"`
}
