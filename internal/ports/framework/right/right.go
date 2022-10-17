// right ports/adapters for driven adapters
// example: Database

package rightFrameworkPorts

type DBPort interface {
	CloseDBConnection() error
	AddToHistory(answer int32, operation string) error
}
