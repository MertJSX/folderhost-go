package utils

import (
	"fmt"
	"sync"
	"time"

	"github.com/MertJSX/folder-host-go/database/logs"
	"github.com/MertJSX/folder-host-go/types"
)

var (
	debounceStates = make(map[string]*types.UserDebounceState) // username -> state
	debounceMu     sync.RWMutex
	debounceDelay  = 60 * time.Second // 60 seconds
)

func ScheduleDebouncedLog(username, filePath string) {
	debounceMu.Lock()
	defer debounceMu.Unlock()

	// Eğer kullanıcı için zaten bir timer varsa durdur
	if state, exists := debounceStates[username]; exists {
		state.Mu.Lock()
		if state.Timer != nil {
			state.Timer.Stop()
		}
		state.Mu.Unlock()
	}

	// Yeni state oluştur veya güncelle
	state := &types.UserDebounceState{
		FilePath: filePath,
		Username: username,
	}

	// Yeni timer başlat
	state.Timer = time.AfterFunc(debounceDelay, func() {
		createWriteFileLog(username, filePath)

		// Log oluşturulduktan sonra state'i temizle
		debounceMu.Lock()
		delete(debounceStates, username)
		debounceMu.Unlock()
	})

	debounceStates[username] = state
}

func TriggerPendingLog(username, filePath string) {
	debounceMu.Lock()
	defer debounceMu.Unlock()

	if state, exists := debounceStates[username]; exists {
		state.Mu.Lock()
		if state.Timer != nil {
			state.Timer.Stop()
		}
		state.Mu.Unlock()

		// Hemen log oluştur
		createWriteFileLog(state.Username, state.FilePath)

		// State'i temizle
		delete(debounceStates, username)
	}
}

func createWriteFileLog(username, filePath string) {
	logs.CreateLog(types.AuditLog{
		Username:    username,
		Action:      "Write file",
		Description: fmt.Sprintf("%s wrote the %s file.", username, filePath),
	})
}
