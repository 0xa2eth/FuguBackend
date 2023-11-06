package secret

// Secret
//
//go:generate gormgen -structs Secret -input .
type Secret struct {
	Id        int32  //
	Authorid  int32  //
	Content   string //
	Timestamp int64  //
	Views     int64  //
}
