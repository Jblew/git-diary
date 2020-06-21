package app

// SyncLock locks app synchronisation
func (app *App) SyncLock() {
	app.Mutex.Lock()
}

// SyncUnlock unlocks app synchronisation
func (app *App) SyncUnlock() {
	app.Mutex.Unlock()
}
