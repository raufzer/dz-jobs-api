package models



type Candidate struct {
	User           User      `db:"user"`
	CandidateID    int       `db:"candidate_id"`
	Resume         string    `db:"resume"`
	Portfolio      string    `db:"portfolio"`
	Skills         string    `db:"skills"`
	TestID         int       `db:"test_id"`
	ProfilePicture string    `db:"profile_picture"`

}