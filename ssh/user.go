package ssh

type User struct {
	Username      string
	Uid           int
	Gid           int
	HomeDirectory string
	Shell         string
}
