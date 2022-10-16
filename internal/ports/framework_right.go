// right ports/adapters for driven adpaters
// example: Database

package ports

type DBPort interface {
	CloseDBConnection()
	AddToHistory(answer int32, operation string) error
}
