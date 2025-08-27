package app

// AppHandler responds to an application action.
type AppHandler func(*AppContext) error
