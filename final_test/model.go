package main

type Answer struct {
	Content       string   `json:"content"`
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
