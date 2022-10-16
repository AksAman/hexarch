// right ports/adapters for driven adapters
// example: Database

package ports

type DBPort interface {
	CloseDBConnection() error
	AddToHistory(answer int32, operation string) error
}
