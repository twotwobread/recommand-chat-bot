// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recommand-chat-bot/internal/ent/movie"
	"recommand-chat-bot/internal/ent/predicate"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMovie = "Movie"
)

// MovieMutation represents an operation that mutates the Movie nodes in the graph.
type MovieMutation struct {
	config
	op            Op
	typ           string
	id            *int64
	created_at    *time.Time
	updated_at    *time.Time
	title         *string
	genre         *string
	director      *string
	actors        *[]string
	appendactors  []string
	description   *string
	release_date  *time.Time
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Movie, error)
	predicates    []predicate.Movie
}

var _ ent.Mutation = (*MovieMutation)(nil)

// movieOption allows management of the mutation configuration using functional options.
type movieOption func(*MovieMutation)

// newMovieMutation creates new mutation for the Movie entity.
func newMovieMutation(c config, op Op, opts ...movieOption) *MovieMutation {
	m := &MovieMutation{
		config:        c,
		op:            op,
		typ:           TypeMovie,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMovieID sets the ID field of the mutation.
func withMovieID(id int64) movieOption {
	return func(m *MovieMutation) {
		var (
			err   error
			once  sync.Once
			value *Movie
		)
		m.oldValue = func(ctx context.Context) (*Movie, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Movie.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMovie sets the old Movie of the mutation.
func withMovie(node *Movie) movieOption {
	return func(m *MovieMutation) {
		m.oldValue = func(context.Context) (*Movie, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MovieMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MovieMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Movie entities.
func (m *MovieMutation) SetID(id int64) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MovieMutation) ID() (id int64, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MovieMutation) IDs(ctx context.Context) ([]int64, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int64{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Movie.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *MovieMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MovieMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *MovieMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *MovieMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *MovieMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *MovieMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetTitle sets the "title" field.
func (m *MovieMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *MovieMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *MovieMutation) ResetTitle() {
	m.title = nil
}

// SetGenre sets the "genre" field.
func (m *MovieMutation) SetGenre(s string) {
	m.genre = &s
}

// Genre returns the value of the "genre" field in the mutation.
func (m *MovieMutation) Genre() (r string, exists bool) {
	v := m.genre
	if v == nil {
		return
	}
	return *v, true
}

// OldGenre returns the old "genre" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldGenre(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGenre is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGenre requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGenre: %w", err)
	}
	return oldValue.Genre, nil
}

// ResetGenre resets all changes to the "genre" field.
func (m *MovieMutation) ResetGenre() {
	m.genre = nil
}

// SetDirector sets the "director" field.
func (m *MovieMutation) SetDirector(s string) {
	m.director = &s
}

// Director returns the value of the "director" field in the mutation.
func (m *MovieMutation) Director() (r string, exists bool) {
	v := m.director
	if v == nil {
		return
	}
	return *v, true
}

// OldDirector returns the old "director" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldDirector(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDirector is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDirector requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDirector: %w", err)
	}
	return oldValue.Director, nil
}

// ResetDirector resets all changes to the "director" field.
func (m *MovieMutation) ResetDirector() {
	m.director = nil
}

// SetActors sets the "actors" field.
func (m *MovieMutation) SetActors(s []string) {
	m.actors = &s
	m.appendactors = nil
}

// Actors returns the value of the "actors" field in the mutation.
func (m *MovieMutation) Actors() (r []string, exists bool) {
	v := m.actors
	if v == nil {
		return
	}
	return *v, true
}

// OldActors returns the old "actors" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldActors(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldActors is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldActors requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldActors: %w", err)
	}
	return oldValue.Actors, nil
}

// AppendActors adds s to the "actors" field.
func (m *MovieMutation) AppendActors(s []string) {
	m.appendactors = append(m.appendactors, s...)
}

// AppendedActors returns the list of values that were appended to the "actors" field in this mutation.
func (m *MovieMutation) AppendedActors() ([]string, bool) {
	if len(m.appendactors) == 0 {
		return nil, false
	}
	return m.appendactors, true
}

// ResetActors resets all changes to the "actors" field.
func (m *MovieMutation) ResetActors() {
	m.actors = nil
	m.appendactors = nil
}

// SetDescription sets the "description" field.
func (m *MovieMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *MovieMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *MovieMutation) ResetDescription() {
	m.description = nil
}

// SetReleaseDate sets the "release_date" field.
func (m *MovieMutation) SetReleaseDate(t time.Time) {
	m.release_date = &t
}

// ReleaseDate returns the value of the "release_date" field in the mutation.
func (m *MovieMutation) ReleaseDate() (r time.Time, exists bool) {
	v := m.release_date
	if v == nil {
		return
	}
	return *v, true
}

// OldReleaseDate returns the old "release_date" field's value of the Movie entity.
// If the Movie object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MovieMutation) OldReleaseDate(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldReleaseDate is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldReleaseDate requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReleaseDate: %w", err)
	}
	return oldValue.ReleaseDate, nil
}

// ResetReleaseDate resets all changes to the "release_date" field.
func (m *MovieMutation) ResetReleaseDate() {
	m.release_date = nil
}

// Where appends a list predicates to the MovieMutation builder.
func (m *MovieMutation) Where(ps ...predicate.Movie) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MovieMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MovieMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Movie, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MovieMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MovieMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Movie).
func (m *MovieMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MovieMutation) Fields() []string {
	fields := make([]string, 0, 8)
	if m.created_at != nil {
		fields = append(fields, movie.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, movie.FieldUpdatedAt)
	}
	if m.title != nil {
		fields = append(fields, movie.FieldTitle)
	}
	if m.genre != nil {
		fields = append(fields, movie.FieldGenre)
	}
	if m.director != nil {
		fields = append(fields, movie.FieldDirector)
	}
	if m.actors != nil {
		fields = append(fields, movie.FieldActors)
	}
	if m.description != nil {
		fields = append(fields, movie.FieldDescription)
	}
	if m.release_date != nil {
		fields = append(fields, movie.FieldReleaseDate)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MovieMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case movie.FieldCreatedAt:
		return m.CreatedAt()
	case movie.FieldUpdatedAt:
		return m.UpdatedAt()
	case movie.FieldTitle:
		return m.Title()
	case movie.FieldGenre:
		return m.Genre()
	case movie.FieldDirector:
		return m.Director()
	case movie.FieldActors:
		return m.Actors()
	case movie.FieldDescription:
		return m.Description()
	case movie.FieldReleaseDate:
		return m.ReleaseDate()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MovieMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case movie.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case movie.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case movie.FieldTitle:
		return m.OldTitle(ctx)
	case movie.FieldGenre:
		return m.OldGenre(ctx)
	case movie.FieldDirector:
		return m.OldDirector(ctx)
	case movie.FieldActors:
		return m.OldActors(ctx)
	case movie.FieldDescription:
		return m.OldDescription(ctx)
	case movie.FieldReleaseDate:
		return m.OldReleaseDate(ctx)
	}
	return nil, fmt.Errorf("unknown Movie field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MovieMutation) SetField(name string, value ent.Value) error {
	switch name {
	case movie.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case movie.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case movie.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case movie.FieldGenre:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGenre(v)
		return nil
	case movie.FieldDirector:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDirector(v)
		return nil
	case movie.FieldActors:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetActors(v)
		return nil
	case movie.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case movie.FieldReleaseDate:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReleaseDate(v)
		return nil
	}
	return fmt.Errorf("unknown Movie field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MovieMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MovieMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MovieMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Movie numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MovieMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MovieMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MovieMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Movie nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MovieMutation) ResetField(name string) error {
	switch name {
	case movie.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case movie.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case movie.FieldTitle:
		m.ResetTitle()
		return nil
	case movie.FieldGenre:
		m.ResetGenre()
		return nil
	case movie.FieldDirector:
		m.ResetDirector()
		return nil
	case movie.FieldActors:
		m.ResetActors()
		return nil
	case movie.FieldDescription:
		m.ResetDescription()
		return nil
	case movie.FieldReleaseDate:
		m.ResetReleaseDate()
		return nil
	}
	return fmt.Errorf("unknown Movie field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MovieMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MovieMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MovieMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MovieMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MovieMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MovieMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MovieMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Movie unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MovieMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Movie edge %s", name)
}
