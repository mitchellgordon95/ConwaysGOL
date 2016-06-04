package game_manager

type GolManager interface {
	// Manages the game state until the user quits
	Manage()
}
