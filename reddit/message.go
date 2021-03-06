package reddit

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/go-querystring/query"
)

// MessageService handles communication with the message
// related methods of the Reddit API.
//
// Reddit API docs: https://www.reddit.com/dev/api/#section_messages
type MessageService struct {
	client *Client
}

// Message is a message.
type Message struct {
	ID      string     `json:"id"`
	FullID  string     `json:"name"`
	Created *Timestamp `json:"created_utc"`

	Subject  string `json:"subject"`
	Text     string `json:"body"`
	ParentID string `json:"parent_id"`

	Author string `json:"author"`
	To     string `json:"dest"`

	IsComment bool `json:"was_comment"`
}

// Messages is a list of messages.
type Messages struct {
	Messages []*Message `json:"messages"`
	After    string     `json:"after"`
	Before   string     `json:"before"`
}

type rootInboxListing struct {
	Kind string       `json:"kind"`
	Data inboxListing `json:"data"`
}

type inboxListing struct {
	Things inboxThings `json:"children"`
	After  string      `json:"after"`
	Before string      `json:"before"`
}

// The returned JSON for comments is a bit different.
// It looks for like the Message struct.
type inboxThings struct {
	Comments []*Message
	Messages []*Message
}

// init initializes or clears the inbox.
func (t *inboxThings) init() {
	t.Comments = make([]*Message, 0)
	t.Messages = make([]*Message, 0)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *inboxThings) UnmarshalJSON(b []byte) error {
	t.init()

	var things []thing
	if err := json.Unmarshal(b, &things); err != nil {
		return err
	}

	for _, thing := range things {
		switch thing.Kind {
		case kindComment:
			v := new(Message)
			if err := json.Unmarshal(thing.Data, v); err == nil {
				t.Comments = append(t.Comments, v)
			}
		case kindMessage:
			v := new(Message)
			if err := json.Unmarshal(thing.Data, v); err == nil {
				t.Messages = append(t.Messages, v)
			}
		}
	}

	return nil
}

func (l *rootInboxListing) getComments() *Messages {
	return &Messages{
		Messages: l.Data.Things.Comments,
		After:    l.Data.After,
		Before:   l.Data.Before,
	}
}

func (l *rootInboxListing) getMessages() *Messages {
	return &Messages{
		Messages: l.Data.Things.Messages,
		After:    l.Data.After,
		Before:   l.Data.Before,
	}
}

// SendMessageRequest represents a request to send a message.
type SendMessageRequest struct {
	// Username, or /r/name for that subreddit's moderators.
	To      string `url:"to"`
	Subject string `url:"subject"`
	Text    string `url:"text"`
	// Optional. If specified, the message will look like it came from the subreddit.
	FromSubreddit string `url:"from_sr,omitempty"`
}

// ReadAll marks all messages/comments as read. It queues up the task on Reddit's end.
// A successful response returns 202 to acknowledge acceptance of the request.
// This endpoint is heavily rate limited.
func (s *MessageService) ReadAll(ctx context.Context) (*Response, error) {
	path := "api/read_all_messages"

	req, err := s.client.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Read marks a message/comment as read via its full ID.
func (s *MessageService) Read(ctx context.Context, ids ...string) (*Response, error) {
	if len(ids) == 0 {
		return nil, errors.New("must provide at least 1 id")
	}

	path := "api/read_message"

	form := url.Values{}
	form.Set("id", strings.Join(ids, ","))

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Unread marks a message/comment as unread via its full ID.
func (s *MessageService) Unread(ctx context.Context, ids ...string) (*Response, error) {
	if len(ids) == 0 {
		return nil, errors.New("must provide at least 1 id")
	}

	path := "api/unread_message"

	form := url.Values{}
	form.Set("id", strings.Join(ids, ","))

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Block blocks the author of a thing via the thing's full ID.
// The thing can be a post, comment or message.
func (s *MessageService) Block(ctx context.Context, id string) (*Response, error) {
	path := "api/block"

	form := url.Values{}
	form.Set("id", id)

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Collapse collapses messages.
func (s *MessageService) Collapse(ctx context.Context, ids ...string) (*Response, error) {
	if len(ids) == 0 {
		return nil, errors.New("must provide at least 1 id")
	}

	path := "api/collapse_message"

	form := url.Values{}
	form.Set("id", strings.Join(ids, ","))

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Uncollapse uncollapses messages.
func (s *MessageService) Uncollapse(ctx context.Context, ids ...string) (*Response, error) {
	if len(ids) == 0 {
		return nil, errors.New("must provide at least 1 id")
	}

	path := "api/uncollapse_message"

	form := url.Values{}
	form.Set("id", strings.Join(ids, ","))

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Delete deletes a message.
func (s *MessageService) Delete(ctx context.Context, id string) (*Response, error) {
	path := "api/del_msg"

	form := url.Values{}
	form.Set("id", id)

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Send sends a message.
func (s *MessageService) Send(ctx context.Context, sendRequest *SendMessageRequest) (*Response, error) {
	if sendRequest == nil {
		return nil, errors.New("sendRequest: cannot be nil")
	}

	path := "api/compose"

	form, err := query.Values(sendRequest)
	if err != nil {
		return nil, err
	}
	form.Set("api_type", "json")

	req, err := s.client.NewRequestWithForm(http.MethodPost, path, form)
	if err != nil {
		return nil, err
	}

	return s.client.Do(ctx, req, nil)
}

// Inbox returns comments and messages that appear in your inbox, respectively.
func (s *MessageService) Inbox(ctx context.Context, opts *ListOptions) (*Messages, *Messages, *Response, error) {
	root, resp, err := s.inbox(ctx, "message/inbox", opts)
	if err != nil {
		return nil, nil, resp, err
	}
	return root.getComments(), root.getMessages(), resp, nil
}

// InboxUnread returns unread comments and messages that appear in your inbox, respectively.
func (s *MessageService) InboxUnread(ctx context.Context, opts *ListOptions) (*Messages, *Messages, *Response, error) {
	root, resp, err := s.inbox(ctx, "message/unread", opts)
	if err != nil {
		return nil, nil, resp, err
	}
	return root.getComments(), root.getMessages(), resp, nil
}

// Sent returns messages that you've sent.
func (s *MessageService) Sent(ctx context.Context, opts *ListOptions) (*Messages, *Response, error) {
	root, resp, err := s.inbox(ctx, "message/sent", opts)
	if err != nil {
		return nil, resp, err
	}
	return root.getMessages(), resp, nil
}

func (s *MessageService) inbox(ctx context.Context, path string, opts *ListOptions) (*rootInboxListing, *Response, error) {
	path, err := addOptions(path, opts)
	if err != nil {
		return nil, nil, err
	}

	req, err := s.client.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(rootInboxListing)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, nil, err
	}

	return root, resp, nil
}
