// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type Bio struct {
	URLCode     string   `json:"url_code"`
	Links       *Links   `json:"links"`
	UserID      string   `json:"user_id"`
	Description string   `json:"description"`
	Skils       []*Skill `json:"skils"`
}

type BioI struct {
	URLCode     string    `json:"url_code"`
	Links       *LinksI   `json:"links"`
	Description string    `json:"description"`
	Skils       []*SkillI `json:"skils"`
}

type Links struct {
	ID        string `json:"ID"`
	Portfolio string `json:"portfolio"`
	Github    string `json:"github"`
	Youtube   string `json:"youtube"`
	Twitter   string `json:"twitter"`
}

type LinksI struct {
	ID        string `json:"ID"`
	Portfolio string `json:"portfolio"`
	Github    string `json:"github"`
	Youtube   string `json:"youtube"`
	Twitter   string `json:"twitter"`
}

type NewUser struct {
	Name string `json:"name"`
	Bio  *BioI  `json:"bio"`
}

type Skill struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type SkillI struct {
	ID   string `json:"ID"`
	Name string `json:"name"`
}

type User struct {
	Name string `json:"name"`
	ID   string `json:"id"`
	Bio  *Bio   `json:"bio"`
}
