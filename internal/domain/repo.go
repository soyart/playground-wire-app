package domain

//go:generate mockgen -source=./repo.go -destination=./mock_repo/mock_repo.go -package=mock_repo
type Repo interface {
	Close() error
	Read(key string) ([]byte, error)
}
