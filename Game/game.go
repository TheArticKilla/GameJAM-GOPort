package game

// Game constants
const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

// Game : an object representing game state
type Game struct {
}

// NewGame : generates a new Game object
func NewGame() (*Game, error) {
	g := &Game{}

	return g, nil
}

// Layout implements the games layout
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
