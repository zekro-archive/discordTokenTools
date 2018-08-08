package main

import (
	"time"
	"github.com/bwmarrin/discordgo"
)

// DUser contains ID, Username and Discriminator
// of a Discord user
type DUser struct {
	ID,
	Username,
	Discriminator string
}

// DGuild contains ID, Name and OwnerID of
// a Discord Guild
type DGuild struct {
	ID,
	Name,
	OwnerID string
}

// Data contains the validity of the token 
// (if the bot logged in successfully with the token),
// the User of the Bot (as *DUSer) and the total number
// of Guilds the bot is running on
type Data struct {
	Valid    bool
	User     *DUser
	NGuilds  int
}

// Event is a object passed through a channel to
// fire asyncronous actions
type Event struct {
	Name string
	Data interface{}
}

// GetTokenData get's the token and a channel of *Event's
// passed. Then, a Discord session will be created.
// First, the 'login' event will be fired before the bot
// tries to log in. After, the 'response' event will be
// passed into the channel with the *Data struct instance.
// If the login was successfull, a third event 'response_guilds'
// will be fired with the data (as *DGuild) of all Guilds.
func GetTokenData(token string, c chan *Event) error {
	dc, err := discordgo.New("Bot " + token)
	if err != nil {
		close(c)
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

		c <- &Event{
			"response",
			data,
		}
		
		time.Sleep(time.Duration(40 * nguilds) * time.Millisecond)
		guilds := make([]*DGuild, nguilds)
		for i, g := range e.Guilds {
			guilds[i] = &DGuild{
				g.ID,
				g.Name,
				g.OwnerID,
			}
		}

		c <- &Event{
			"response_guilds",
			guilds,
		}
		
		close(c)
		dc.Close()
	})

	c <- &Event{
		"login",
		nil,
	}
	
	err = dc.Open()
	if err != nil {
		data := Data{}
		data.Valid = false

		c <- &Event{
			"response",
			data,
		}
		close(c)
		return err
	}

	return nil
}

// GetTokenValidity is a simple check function
// for GetTokenData to get just the validity of
// the token without firing events and getting
// token data.
func GetTokenValidity(token string) bool {
	events := make(chan *Event, 2)
	go GetTokenData(token, events)
	<- events // just ignore the login event
	resEvent := <- events
	return resEvent.Data.(Data).Valid
}