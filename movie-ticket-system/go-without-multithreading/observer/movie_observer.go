package observer
import "fmt"

// MovieObserver explicitly isolates Push channels decoupled from Entities matrices
type MovieObserver interface {
	Update(movieTitle string)
}

// UserObserver represents an explicit downstream subscription node
type UserObserver struct {
	UserName string
}

func (u *UserObserver) Update(movieTitle string) {
	fmt.Printf("🔔 Notification for %s: Live mapping of '%s' cleanly pushed to the public!\n", u.UserName, movieTitle)
}
