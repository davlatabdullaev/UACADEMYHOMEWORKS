package country
import "github.com/google/uuid"

type Country struct{
	ID uuid.UUID
	Name string
	Currency string
}