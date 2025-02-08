package main

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "image/color"
    "log"
)

const (
    screenWidth  = 800
    screenHeight = 600
)

type Game struct {
    ballX, ballY       float64
    ballSpeedX, ballSpeedY float64
    paddle1Y, paddle2Y float64
}

func (g *Game) Update() error {
    // Update ball position
    g.ballX += g.ballSpeedX
    g.ballY += g.ballSpeedY

    // Ball collision with top and bottom
    if g.ballY < 0 || g.ballY > screenHeight-10 {
        g.ballSpeedY = -g.ballSpeedY
    }

    // Ball collision with paddles
    if g.ballX < 20 && g.ballY > g.paddle1Y && g.ballY < g.paddle1Y+100 {
        g.ballSpeedX = -g.ballSpeedX
    }
    if g.ballX > screenWidth-30 && g.ballY > g.paddle2Y && g.ballY < g.paddle2Y+100 {
        g.ballSpeedX = -g.ballSpeedX
    }

    // Ball out of bounds
    if g.ballX < 0 || g.ballX > screenWidth {
        g.ballX, g.ballY = screenWidth/2, screenHeight/2
    }

    // AI paddle movement
    if g.ballSpeedX < 0 { // Left paddle's turn
        if g.ballY > g.paddle1Y+50 && g.paddle1Y < screenHeight-100 {
            g.paddle1Y += 5
        } else if g.ballY < g.paddle1Y+50 && g.paddle1Y > 0 {
            g.paddle1Y -= 5
        }
    } else { // Right paddle's turn
        if g.ballY > g.paddle2Y+50 && g.paddle2Y < screenHeight-100 {
            g.paddle2Y += 5
        } else if g.ballY < g.paddle2Y+50 && g.paddle2Y > 0 {
            g.paddle2Y -= 5
        }
    }

    return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
    // Draw ball
    ebitenutil.DrawRect(screen, g.ballX, g.ballY, 10, 10, color.White)

    // Draw paddles
    ebitenutil.DrawRect(screen, 10, g.paddle1Y, 10, 100, color.White)
    ebitenutil.DrawRect(screen, screenWidth-20, g.paddle2Y, 10, 100, color.White)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
    return screenWidth, screenHeight
}

func main() {
    game := &Game{
        ballX:      screenWidth / 2,
        ballY:      screenHeight / 2,
        ballSpeedX: 4,
        ballSpeedY: 4,
        paddle1Y:   screenHeight / 2,
        paddle2Y:   screenHeight / 2,
    }
    ebiten.SetWindowSize(screenWidth, screenHeight)
    ebiten.SetWindowTitle("Ping Pong")
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}