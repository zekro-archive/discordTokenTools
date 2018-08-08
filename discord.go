package main

import (
	//"fmt"
	"time"
	"github.com/bwmarrin/discordgo"
	"github.com/googollee/go-socket.io"
)


type DUser struct {
	ID,
	Username,
	Discriminator string
}

type DGuild struct {
	ID,
	Name,
	OwnerID string
}

type Data struct {
	Valid    bool
	User     *DUser
	NGuilds  int
}


func DiscordLogin(token string, so socketio.Socket) error {

	dc, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	dc.AddHandlerOnce(func(s *discordgo.Session, e *discordgo.Ready) {
		nguilds := len(e.Guilds)

		data := Data{
			true,
			&DUser{
				e.User.ID,
				e.User.Username,
				e.User.Discriminator,
			},
			nguilds,
		}
		so.Emit("response", data)

		time.Sleep(time.Duration(40 * nguilds) * time.Millisecond)
		guilds := make([]*DGuild, nguilds)
		for i, g := range e.Guilds {
			guilds[i] = &DGuild{
				g.ID,
				g.Name,
				g.OwnerID,
			}
		}
		so.Emit("response_guilds", guilds)

		dc.Close()
	})

	so.Emit("login")

	err = dc.Open()
	if err != nil {
		data := Data{}
		data.Valid = false
		so.Emit("response", data)
		return err
	}

	return nil
}