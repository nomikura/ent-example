// Code generated by ent, DO NOT EDIT.

package ent

import (
	"entdemo/ent/user"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// LikedTweets holds the value of the liked_tweets edge.
	LikedTweets []*Tweet `json:"liked_tweets,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int

	namedLikedTweets map[string][]*Tweet
	namedLikes       map[string][]*Like
}

// LikedTweetsOrErr returns the LikedTweets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikedTweetsOrErr() ([]*Tweet, error) {
	if e.loadedTypes[0] {
		return e.LikedTweets, nil
	}
	return nil, &NotLoadedError{edge: "liked_tweets"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[1] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		}
	}
	return nil
}

// QueryLikedTweets queries the "liked_tweets" edge of the User entity.
func (u *User) QueryLikedTweets() *TweetQuery {
	return (&UserClient{config: u.config}).QueryLikedTweets(u)
}

// QueryLikes queries the "likes" edge of the User entity.
func (u *User) QueryLikes() *LikeQuery {
	return (&UserClient{config: u.config}).QueryLikes(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// NamedLikedTweets returns the LikedTweets named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedLikedTweets(name string) ([]*Tweet, error) {
	if u.Edges.namedLikedTweets == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedLikedTweets[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedLikedTweets(name string, edges ...*Tweet) {
	if u.Edges.namedLikedTweets == nil {
		u.Edges.namedLikedTweets = make(map[string][]*Tweet)
	}
	if len(edges) == 0 {
		u.Edges.namedLikedTweets[name] = []*Tweet{}
	} else {
		u.Edges.namedLikedTweets[name] = append(u.Edges.namedLikedTweets[name], edges...)
	}
}

// NamedLikes returns the Likes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (u *User) NamedLikes(name string) ([]*Like, error) {
	if u.Edges.namedLikes == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := u.Edges.namedLikes[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (u *User) appendNamedLikes(name string, edges ...*Like) {
	if u.Edges.namedLikes == nil {
		u.Edges.namedLikes = make(map[string][]*Like)
	}
	if len(edges) == 0 {
		u.Edges.namedLikes[name] = []*Like{}
	} else {
		u.Edges.namedLikes[name] = append(u.Edges.namedLikes[name], edges...)
	}
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
