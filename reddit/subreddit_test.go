package reddit

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

var expectedPosts = &Posts{
	Posts: []*Post{
		{
			ID:      "agi5zf",
			FullID:  "t3_agi5zf",
			Created: &Timestamp{time.Date(2019, 1, 16, 5, 57, 51, 0, time.UTC)},
			Edited:  &Timestamp{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},

			Permalink: "/r/test/comments/agi5zf/test/",
			URL:       "https://www.reddit.com/r/test/comments/agi5zf/test/",

			Title: "test",
			Body:  "test",

			Score:            253,
			UpvoteRatio:      0.99,
			NumberOfComments: 1634,

			SubredditName:         "test",
			SubredditNamePrefixed: "r/test",
			SubredditID:           "t5_2qh23",

			Author:   "kmiller0112",
			AuthorID: "t2_30a5ktgt",

			IsSelfPost: true,
			Stickied:   true,
		},
		{
			ID:      "hyhquk",
			FullID:  "t3_hyhquk",
			Created: &Timestamp{time.Date(2020, 7, 27, 0, 5, 10, 0, time.UTC)},
			Edited:  &Timestamp{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},

			Permalink: "/r/test/comments/hyhquk/veggies/",
			URL:       "https://i.imgur.com/LrN2mPw.jpg",

			Title: "Veggies",

			Score:            4,
			UpvoteRatio:      1,
			NumberOfComments: 0,

			SubredditName:         "test",
			SubredditNamePrefixed: "r/test",
			SubredditID:           "t5_2qh23",

			Author:   "MuckleMcDuckle",
			AuthorID: "t2_6fqntbwq",
		},
	},
	After:  "t3_hyhquk",
	Before: "",
}

var expectedSubreddit = &Subreddit{
	ID:      "2rc7j",
	FullID:  "t5_2rc7j",
	Created: &Timestamp{time.Date(2009, 11, 11, 0, 54, 28, 0, time.UTC)},

	URL:          "/r/golang/",
	Name:         "golang",
	NamePrefixed: "r/golang",
	Title:        "The Go Programming Language",
	Description:  "Ask questions and post articles about the Go programming language and related tools, events etc.",
	Type:         "public",

	Subscribers:     116532,
	ActiveUserCount: Int(386),
	NSFW:            false,
	UserIsMod:       false,
	Subscribed:      true,
}

var expectedSubreddits = &Subreddits{
	After:  "t5_2qh0u",
	Before: "",
	Subreddits: []*Subreddit{
		{
			ID:      "2qs0k",
			FullID:  "t5_2qs0k",
			Created: &Timestamp{time.Date(2009, 1, 25, 2, 25, 57, 0, time.UTC)},

			URL:          "/r/Home/",
			Name:         "Home",
			NamePrefixed: "r/Home",
			Title:        "Home",
			Type:         "public",

			Subscribers: 15336,
			NSFW:        false,
			UserIsMod:   false,
			Subscribed:  true,
			Favorite:    false,
		},
		{
			ID:      "2qh1i",
			FullID:  "t5_2qh1i",
			Created: &Timestamp{time.Date(2008, 1, 25, 3, 52, 15, 0, time.UTC)},

			URL:          "/r/AskReddit/",
			Name:         "AskReddit",
			NamePrefixed: "r/AskReddit",
			Title:        "Ask Reddit...",
			Description:  "r/AskReddit is the place to ask and answer thought-provoking questions.",
			Type:         "public",

			Subscribers: 28449174,
			NSFW:        false,
			UserIsMod:   false,
			Subscribed:  true,
			Favorite:    true,
		},
		{
			ID:      "2qh0u",
			FullID:  "t5_2qh0u",
			Created: &Timestamp{time.Date(2008, 1, 25, 0, 31, 9, 0, time.UTC)},

			URL:          "/r/pics/",
			Name:         "pics",
			NamePrefixed: "r/pics",
			Title:        "Reddit Pics",
			Description:  "A place for pictures and photographs.",
			Type:         "public",

			Subscribers: 24987753,
			NSFW:        false,
			UserIsMod:   false,
			Subscribed:  false,
			Favorite:    false,
		},
	},
}

var expectedSubredditNames = []string{
	"golang",
	"golang_infosec",
	"GolangJobOfferings",
	"golanguage",
	"golang2",
	"GolangUnofficial",
	"golang_jobs",
}

var expectedSearchPosts = &Posts{
	Posts: []*Post{
		{
			ID:      "hybow9",
			FullID:  "t3_hybow9",
			Created: &Timestamp{time.Date(2020, 7, 26, 18, 14, 24, 0, time.UTC)},
			Edited:  &Timestamp{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},

			Permalink: "/r/WatchPeopleDieInside/comments/hybow9/pregnancy_test/",
			URL:       "https://v.redd.it/ra4qnt8bt8d51",

			Title: "Pregnancy test",

			Score:            103829,
			UpvoteRatio:      0.88,
			NumberOfComments: 3748,

			SubredditName:         "WatchPeopleDieInside",
			SubredditNamePrefixed: "r/WatchPeopleDieInside",
			SubredditID:           "t5_3h4zq",

			Author:   "chocolat_ice_cream",
			AuthorID: "t2_3p32m02",
		},
		{
			ID:      "hmwhd7",
			FullID:  "t3_hmwhd7",
			Created: &Timestamp{time.Date(2020, 7, 7, 15, 19, 42, 0, time.UTC)},
			Edited:  &Timestamp{time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)},

			Permalink: "/r/worldnews/comments/hmwhd7/brazilian_president_jair_bolsonaro_tests_positive/",
			URL:       "https://www.theguardian.com/world/2020/jul/07/jair-bolsonaro-coronavirus-positive-test-brazil-president",

			Title: "Brazilian president Jair Bolsonaro tests positive for coronavirus",

			Score:            149238,
			UpvoteRatio:      0.94,
			NumberOfComments: 7415,

			SubredditName:         "worldnews",
			SubredditNamePrefixed: "r/worldnews",
			SubredditID:           "t5_2qh13",

			Author:   "Jeremy_Martin",
			AuthorID: "t2_wgrkg",
		},
	},
	After: "t3_hmwhd7",
}

var expectedRandomSubreddit = &Subreddit{
	FullID:  "t5_2wi4l",
	Created: &Timestamp{time.Date(2013, 3, 1, 4, 4, 18, 0, time.UTC)},

	URL:          "/r/GalaxyS8/",
	Name:         "GalaxyS8",
	NamePrefixed: "r/GalaxyS8",
	Title:        "Samsung Galaxy S8",
	Description:  "The only place for news, discussion, photos, and everything else Samsung Galaxy S8.",
	Type:         "public",

	Subscribers: 52357,
}

var expectedRelationships3 = &Relationships{
	Relationships: []*Relationship{
		{
			ID:      "rel_id1",
			Created: &Timestamp{time.Date(2020, 8, 11, 2, 35, 2, 0, time.UTC)},
			User:    "testuser1",
			UserID:  "t2_user1",
		},
		{
			ID:      "rel_id2",
			Created: &Timestamp{time.Date(2020, 8, 11, 2, 35, 0, 0, time.UTC)},
			User:    "testuser2",
			UserID:  "t2_user2",
		},
	},
	After:  "",
	Before: "",
}

var expectedBans = &Bans{
	Bans: []*Ban{
		{
			Relationship: &Relationship{
				ID:      "rb_123",
				Created: &Timestamp{time.Date(2020, 8, 11, 2, 35, 2, 0, time.UTC)},

				User:   "testuser1",
				UserID: "t2_user1",
			},

			DaysLeft: Int(43),
			Note:     "Spam",
		},
		{
			Relationship: &Relationship{
				ID:      "rb_456",
				Created: &Timestamp{time.Date(2020, 8, 11, 2, 35, 0, 0, time.UTC)},

				User:   "testuser2",
				UserID: "t2_user2",
			},

			DaysLeft: nil,
			Note:     "Spam",
		},
	},
	After:  "",
	Before: "",
}

var expectedModerators = []*Moderator{
	{
		Relationship: &Relationship{
			ID:      "rb_tmatb9",
			User:    "testuser1",
			UserID:  "t2_test1",
			Created: &Timestamp{time.Date(2013, 7, 29, 20, 44, 27, 0, time.UTC)},
		},
		Permissions: []string{"all"},
	},
	{
		Relationship: &Relationship{
			ID:      "rb_5c9s4d",
			User:    "testuser2",
			UserID:  "t2_test2",
			Created: &Timestamp{time.Date(2014, 3, 1, 18, 13, 53, 0, time.UTC)},
		},
		Permissions: []string{"all"},
	},
}

func TestSubredditService_HotPosts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/hot", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.HotPosts(ctx, "test", nil)
	require.NoError(t, err)
	require.Equal(t, expectedPosts, posts)
}

func TestSubredditService_NewPosts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/new", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.NewPosts(ctx, "test", nil)
	require.NoError(t, err)
	require.Equal(t, expectedPosts, posts)
}

func TestSubredditService_RisingPosts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/rising", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.RisingPosts(ctx, "test", nil)
	require.NoError(t, err)
	require.Equal(t, expectedPosts, posts)
}

func TestSubredditService_ControversialPosts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/controversial", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.ControversialPosts(ctx, "test", nil)
	require.NoError(t, err)
	require.Equal(t, expectedPosts, posts)
}

func TestSubredditService_TopPosts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/top", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.TopPosts(ctx, "test", nil)
	require.NoError(t, err)
	require.Equal(t, expectedPosts, posts)
}

func TestSubredditService_Get(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/about.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/golang/about", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	_, _, err = client.Subreddit.Get(ctx, "")
	require.EqualError(t, err, "name: cannot be empty")

	subreddit, _, err := client.Subreddit.Get(ctx, "golang")
	require.NoError(t, err)
	require.Equal(t, expectedSubreddit, subreddit)
}

func TestSubredditService_Popular(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/popular", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Popular(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_New(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/new", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.New(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_Gold(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/gold", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Gold(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_Default(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/default", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Default(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_Subscribed(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/mine/subscriber", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Subscribed(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_Approved(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/mine/contributor", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Approved(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_Moderated(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/mine/moderator", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Moderated(ctx, nil)
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_GetSticky1(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/post/post.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/sticky", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, "1", r.Form.Get("num"))

		fmt.Fprint(w, blob)
	})

	postAndComments, _, err := client.Subreddit.GetSticky1(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, expectedPostAndComments, postAndComments)
}

func TestSubredditService_GetSticky2(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/post/post.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/sticky", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, "2", r.Form.Get("num"))

		fmt.Fprint(w, blob)
	})

	postAndComments, _, err := client.Subreddit.GetSticky2(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, expectedPostAndComments, postAndComments)
}

func TestSubredditService_Subscribe(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		form := url.Values{}
		form.Set("action", "sub")
		form.Set("sr_name", "test,golang,nba")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)
	})

	_, err := client.Subreddit.Subscribe(ctx, "test", "golang", "nba")
	require.NoError(t, err)
}

func TestSubredditService_SubscribeByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		form := url.Values{}
		form.Set("action", "sub")
		form.Set("sr", "t5_test1,t5_test2,t5_test3")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)
	})

	_, err := client.Subreddit.SubscribeByID(ctx, "t5_test1", "t5_test2", "t5_test3")
	require.NoError(t, err)
}

func TestSubredditService_Unsubscribe(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		form := url.Values{}
		form.Set("action", "unsub")
		form.Set("sr_name", "test,golang,nba")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)
	})

	_, err := client.Subreddit.Unsubscribe(ctx, "test", "golang", "nba")
	require.NoError(t, err)
}

func TestSubredditService_UnsubscribeByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/subscribe", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		form := url.Values{}
		form.Set("action", "unsub")
		form.Set("sr", "t5_test1,t5_test2,t5_test3")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)
	})

	_, err := client.Subreddit.UnsubscribeByID(ctx, "t5_test1", "t5_test2", "t5_test3")
	require.NoError(t, err)
}

func TestSubredditService_Favorite(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/favorite", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		form := url.Values{}
		form.Set("sr_name", "testsubreddit")
		form.Set("make_favorite", "true")
		form.Set("api_type", "json")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)
	})

	_, err := client.Subreddit.Favorite(ctx, "testsubreddit")
	require.NoError(t, err)
}

func TestSubredditService_Unfavorite(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/api/favorite", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodPost, r.Method)

		form := url.Values{}
		form.Set("sr_name", "testsubreddit")
		form.Set("make_favorite", "false")
		form.Set("api_type", "json")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)
	})

	_, err := client.Subreddit.Unfavorite(ctx, "testsubreddit")
	require.NoError(t, err)
}

func TestSubredditService_Search(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/list.json")
	require.NoError(t, err)

	mux.HandleFunc("/subreddits/search", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("q", "golang")
		form.Set("limit", "10")
		form.Set("sort", "activity")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	subreddits, _, err := client.Subreddit.Search(ctx, "golang", &ListSubredditOptions{
		ListOptions: ListOptions{
			Limit: 10,
		},
		Sort: "activity",
	})
	require.NoError(t, err)
	require.Equal(t, expectedSubreddits, subreddits)
}

func TestSubredditService_SearchNames(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/search-names.json")
	require.NoError(t, err)

	mux.HandleFunc("/api/search_reddit_names", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("query", "golang")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	names, _, err := client.Subreddit.SearchNames(ctx, "golang")
	require.NoError(t, err)
	require.Equal(t, expectedSubredditNames, names)
}

func TestSubredditService_SearchPosts(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/search-posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/all/search", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("q", "test")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.SearchPosts(ctx, "test", "", nil)
	require.NoError(t, err)
	require.Equal(t, expectedSearchPosts, posts)
}

func TestSubredditService_SearchPosts_InSubreddit(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/search-posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/search", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("q", "test")
		form.Set("restrict_sr", "true")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.SearchPosts(ctx, "test", "test", nil)
	require.NoError(t, err)
	require.Equal(t, expectedSearchPosts, posts)
}

func TestSubredditService_SearchPosts_InSubreddits(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/search-posts.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test+golang+nba/search", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("q", "test")
		form.Set("restrict_sr", "true")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	posts, _, err := client.Subreddit.SearchPosts(ctx, "test", "test+golang+nba", nil)
	require.NoError(t, err)
	require.Equal(t, expectedSearchPosts, posts)
}

func TestSubredditService_Random(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/random.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/random", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, "true", r.Form.Get("sr_detail"))
		require.Equal(t, "1", r.Form.Get("limit"))

		fmt.Fprint(w, blob)
	})

	subreddit, _, err := client.Subreddit.Random(ctx)
	require.NoError(t, err)
	require.Equal(t, expectedRandomSubreddit, subreddit)
}

func TestSubredditService_RandomNSFW(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/random.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/randnsfw", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, "true", r.Form.Get("sr_detail"))
		require.Equal(t, "1", r.Form.Get("limit"))

		fmt.Fprint(w, blob)
	})

	subreddit, _, err := client.Subreddit.RandomNSFW(ctx)
	require.NoError(t, err)
	require.Equal(t, expectedRandomSubreddit, subreddit)
}

func TestSubredditService_SubmissionText(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/r/test/api/submit_text", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, `{
			"submit_text": "this is a test",
			"submit_text_html": ""
		}`)
	})

	text, _, err := client.Subreddit.SubmissionText(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, "this is a test", text)
}

func TestSubredditService_Banned(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/banned-users.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/banned", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("after", "testafter")
		form.Set("limit", "10")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	bans, _, err := client.Subreddit.Banned(ctx, "test", &ListOptions{After: "testafter", Limit: 10})
	require.NoError(t, err)
	require.Equal(t, expectedBans, bans)
}

func TestSubredditService_Muted(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/relationships.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/muted", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("before", "testbefore")
		form.Set("limit", "50")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	mutes, _, err := client.Subreddit.Muted(ctx, "test", &ListOptions{Before: "testbefore", Limit: 50})
	require.NoError(t, err)
	require.Equal(t, expectedRelationships3, mutes)
}

func TestSubredditService_WikiBanned(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/banned-users.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/wikibanned", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("after", "testafter")
		form.Set("limit", "15")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	bans, _, err := client.Subreddit.WikiBanned(ctx, "test", &ListOptions{After: "testafter", Limit: 15})
	require.NoError(t, err)
	require.Equal(t, expectedBans, bans)
}

func TestSubredditService_Contributors(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/relationships.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/contributors", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("limit", "5")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	contributors, _, err := client.Subreddit.Contributors(ctx, "test", &ListOptions{Limit: 5})
	require.NoError(t, err)
	require.Equal(t, expectedRelationships3, contributors)
}

func TestSubredditService_WikiContributors(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/relationships.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/wikicontributors", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)

		form := url.Values{}
		form.Set("limit", "99")

		err := r.ParseForm()
		require.NoError(t, err)
		require.Equal(t, form, r.Form)

		fmt.Fprint(w, blob)
	})

	contributors, _, err := client.Subreddit.WikiContributors(ctx, "test", &ListOptions{Limit: 99})
	require.NoError(t, err)
	require.Equal(t, expectedRelationships3, contributors)
}

func TestSubredditService_Moderators(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	blob, err := readFileContents("../testdata/subreddit/moderators.json")
	require.NoError(t, err)

	mux.HandleFunc("/r/test/about/moderators", func(w http.ResponseWriter, r *http.Request) {
		require.Equal(t, http.MethodGet, r.Method)
		fmt.Fprint(w, blob)
	})

	moderators, _, err := client.Subreddit.Moderators(ctx, "test")
	require.NoError(t, err)
	require.Equal(t, expectedModerators, moderators)
}
