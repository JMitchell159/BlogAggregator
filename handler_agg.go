package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"errors"
	"fmt"
	"html"
	"io"
	"net/http"
	"time"

	"github.com/JMitchell159/gator/internal/database"
	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("couldn't make the request struct: %v", err)
	}

	req.Header.Set("User-Agent", "gator")
	c := http.Client{
		Timeout: time.Minute,
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, fmt.Errorf("something wrong with the response: %v", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if res.StatusCode > 399 {
		return nil, fmt.Errorf("response failed with status code: %d and \nbody: %s", res.StatusCode, body)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	feed := RSSFeed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling into RSSFeed struct: %v", err)
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i, item := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		feed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &feed, nil
}

func scrapeFeeds(s *state) error {
	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: time.Now(),
		ID:        feed.ID,
	})
	if err != nil {
		return err
	}

	rssFeed, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	fmt.Printf("%s Feed Titles:\n", feed.Name)
	for _, item := range rssFeed.Channel.Item {
		t, err := time.Parse(time.RFC1123, item.PubDate)
		if err != nil {
			t, err = time.Parse(time.RFC1123Z, item.PubDate)
			if err != nil {
				return err
			}
		}

		ok := len(item.Title) > 0
		if _, err := s.db.GetPostByTitle(context.Background(), sql.NullString{
			String: item.Title,
			Valid:  ok,
		}); err == nil {
			continue
		}
		ok1 := len(item.Description) > 0

		post, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: sql.NullString{
				String: item.Title,
				Valid:  ok,
			},
			Url: item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  ok1,
			},
			PublishedAt: t,
			FeedID:      feed.ID,
		})
		if err != nil {
			return err
		}

		fmt.Printf(" > %s\n", post.Title.String)
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	if cmd.args == nil {
		return errors.New("the aggregate handler expects a single argument, a duration string")
	}

	t, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", cmd.args[0])

	ticker := time.NewTicker(t)
	for ; ; <-ticker.C {
		err = scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}
