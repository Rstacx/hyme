package music

import (
	"context"
	"os"
	"sync"
	"strconv"
	"github.com/bwmarrin/discordgo"
	"github.com/disgoorg/disgolink/dgolink"
	"github.com/disgoorg/disgolink/lavalink"
)

type Bot struct {
	Link           *dgolink.Link
	PlayerManagers map[string]*PlayerManager
}

type PlayerManager struct {
	lavalink.PlayerEventAdapter
	Player        lavalink.Player
	Queue         []lavalink.AudioTrack
	QueueMu       sync.Mutex
	RepeatingMode int
}

func InitLink(s *discordgo.Session) *dgolink.Link {
	link := dgolink.New(s)
	return link
}

func RegisterNodes(link *dgolink.Link) (lavalink.Node, error) {
	secure, _ := strconv.ParseBool(os.Getenv("LAVALINK_SECURE"))
	node, err := link.AddNode(context.TODO(), lavalink.NodeConfig{
		Name:        "Chisato", // a unique node name
		Host:        os.Getenv("LAVALINK_HOST"),
		Port:        os.Getenv("LAVALINK_PORT"),
		Password:    os.Getenv("LAVALINK_PASSWORD"),
		Secure:      secure,
	})
	return node, err
}