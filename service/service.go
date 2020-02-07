package service

type Service interface {
	List(string, string, string) []string
}
