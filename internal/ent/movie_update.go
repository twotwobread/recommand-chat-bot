// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"recommand-chat-bot/internal/ent/movie"
	"recommand-chat-bot/internal/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// MovieUpdate is the builder for updating Movie entities.
type MovieUpdate struct {
	config
	hooks    []Hook
	mutation *MovieMutation
}

// Where appends a list predicates to the MovieUpdate builder.
func (mu *MovieUpdate) Where(ps ...predicate.Movie) *MovieUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MovieUpdate) SetUpdatedAt(t time.Time) *MovieUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MovieUpdate) SetTitle(s string) *MovieUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableTitle(s *string) *MovieUpdate {
	if s != nil {
		mu.SetTitle(*s)
	}
	return mu
}

// SetGenre sets the "genre" field.
func (mu *MovieUpdate) SetGenre(s string) *MovieUpdate {
	mu.mutation.SetGenre(s)
	return mu
}

// SetNillableGenre sets the "genre" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableGenre(s *string) *MovieUpdate {
	if s != nil {
		mu.SetGenre(*s)
	}
	return mu
}

// SetDirector sets the "director" field.
func (mu *MovieUpdate) SetDirector(s string) *MovieUpdate {
	mu.mutation.SetDirector(s)
	return mu
}

// SetNillableDirector sets the "director" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableDirector(s *string) *MovieUpdate {
	if s != nil {
		mu.SetDirector(*s)
	}
	return mu
}

// SetActors sets the "actors" field.
func (mu *MovieUpdate) SetActors(s []string) *MovieUpdate {
	mu.mutation.SetActors(s)
	return mu
}

// AppendActors appends s to the "actors" field.
func (mu *MovieUpdate) AppendActors(s []string) *MovieUpdate {
	mu.mutation.AppendActors(s)
	return mu
}

// SetDescription sets the "description" field.
func (mu *MovieUpdate) SetDescription(s string) *MovieUpdate {
	mu.mutation.SetDescription(s)
	return mu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableDescription(s *string) *MovieUpdate {
	if s != nil {
		mu.SetDescription(*s)
	}
	return mu
}

// SetReleaseDate sets the "release_date" field.
func (mu *MovieUpdate) SetReleaseDate(t time.Time) *MovieUpdate {
	mu.mutation.SetReleaseDate(t)
	return mu
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableReleaseDate(t *time.Time) *MovieUpdate {
	if t != nil {
		mu.SetReleaseDate(*t)
	}
	return mu
}

// Mutation returns the MovieMutation object of the builder.
func (mu *MovieUpdate) Mutation() *MovieMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MovieUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MovieUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MovieUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MovieUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MovieUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := movie.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MovieUpdate) check() error {
	if v, ok := mu.mutation.Title(); ok {
		if err := movie.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Movie.title": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Genre(); ok {
		if err := movie.GenreValidator(v); err != nil {
			return &ValidationError{Name: "genre", err: fmt.Errorf(`ent: validator failed for field "Movie.genre": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Director(); ok {
		if err := movie.DirectorValidator(v); err != nil {
			return &ValidationError{Name: "director", err: fmt.Errorf(`ent: validator failed for field "Movie.director": %w`, err)}
		}
	}
	if v, ok := mu.mutation.Description(); ok {
		if err := movie.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Movie.description": %w`, err)}
		}
	}
	return nil
}

func (mu *MovieUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(movie.Table, movie.Columns, sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt64))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(movie.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.SetField(movie.FieldTitle, field.TypeString, value)
	}
	if value, ok := mu.mutation.Genre(); ok {
		_spec.SetField(movie.FieldGenre, field.TypeString, value)
	}
	if value, ok := mu.mutation.Director(); ok {
		_spec.SetField(movie.FieldDirector, field.TypeString, value)
	}
	if value, ok := mu.mutation.Actors(); ok {
		_spec.SetField(movie.FieldActors, field.TypeJSON, value)
	}
	if value, ok := mu.mutation.AppendedActors(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, movie.FieldActors, value)
		})
	}
	if value, ok := mu.mutation.Description(); ok {
		_spec.SetField(movie.FieldDescription, field.TypeString, value)
	}
	if value, ok := mu.mutation.ReleaseDate(); ok {
		_spec.SetField(movie.FieldReleaseDate, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MovieUpdateOne is the builder for updating a single Movie entity.
type MovieUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MovieUpdateOne) SetUpdatedAt(t time.Time) *MovieUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetTitle sets the "title" field.
func (muo *MovieUpdateOne) SetTitle(s string) *MovieUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableTitle(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetTitle(*s)
	}
	return muo
}

// SetGenre sets the "genre" field.
func (muo *MovieUpdateOne) SetGenre(s string) *MovieUpdateOne {
	muo.mutation.SetGenre(s)
	return muo
}

// SetNillableGenre sets the "genre" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableGenre(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetGenre(*s)
	}
	return muo
}

// SetDirector sets the "director" field.
func (muo *MovieUpdateOne) SetDirector(s string) *MovieUpdateOne {
	muo.mutation.SetDirector(s)
	return muo
}

// SetNillableDirector sets the "director" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableDirector(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetDirector(*s)
	}
	return muo
}

// SetActors sets the "actors" field.
func (muo *MovieUpdateOne) SetActors(s []string) *MovieUpdateOne {
	muo.mutation.SetActors(s)
	return muo
}

// AppendActors appends s to the "actors" field.
func (muo *MovieUpdateOne) AppendActors(s []string) *MovieUpdateOne {
	muo.mutation.AppendActors(s)
	return muo
}

// SetDescription sets the "description" field.
func (muo *MovieUpdateOne) SetDescription(s string) *MovieUpdateOne {
	muo.mutation.SetDescription(s)
	return muo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableDescription(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetDescription(*s)
	}
	return muo
}

// SetReleaseDate sets the "release_date" field.
func (muo *MovieUpdateOne) SetReleaseDate(t time.Time) *MovieUpdateOne {
	muo.mutation.SetReleaseDate(t)
	return muo
}

// SetNillableReleaseDate sets the "release_date" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableReleaseDate(t *time.Time) *MovieUpdateOne {
	if t != nil {
		muo.SetReleaseDate(*t)
	}
	return muo
}

// Mutation returns the MovieMutation object of the builder.
func (muo *MovieUpdateOne) Mutation() *MovieMutation {
	return muo.mutation
}

// Where appends a list predicates to the MovieUpdate builder.
func (muo *MovieUpdateOne) Where(ps ...predicate.Movie) *MovieUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MovieUpdateOne) Select(field string, fields ...string) *MovieUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Movie entity.
func (muo *MovieUpdateOne) Save(ctx context.Context) (*Movie, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MovieUpdateOne) SaveX(ctx context.Context) *Movie {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MovieUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MovieUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MovieUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := movie.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MovieUpdateOne) check() error {
	if v, ok := muo.mutation.Title(); ok {
		if err := movie.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Movie.title": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Genre(); ok {
		if err := movie.GenreValidator(v); err != nil {
			return &ValidationError{Name: "genre", err: fmt.Errorf(`ent: validator failed for field "Movie.genre": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Director(); ok {
		if err := movie.DirectorValidator(v); err != nil {
			return &ValidationError{Name: "director", err: fmt.Errorf(`ent: validator failed for field "Movie.director": %w`, err)}
		}
	}
	if v, ok := muo.mutation.Description(); ok {
		if err := movie.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Movie.description": %w`, err)}
		}
	}
	return nil
}

func (muo *MovieUpdateOne) sqlSave(ctx context.Context) (_node *Movie, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(movie.Table, movie.Columns, sqlgraph.NewFieldSpec(movie.FieldID, field.TypeInt64))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Movie.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, movie.FieldID)
		for _, f := range fields {
			if !movie.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != movie.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(movie.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.SetField(movie.FieldTitle, field.TypeString, value)
	}
	if value, ok := muo.mutation.Genre(); ok {
		_spec.SetField(movie.FieldGenre, field.TypeString, value)
	}
	if value, ok := muo.mutation.Director(); ok {
		_spec.SetField(movie.FieldDirector, field.TypeString, value)
	}
	if value, ok := muo.mutation.Actors(); ok {
		_spec.SetField(movie.FieldActors, field.TypeJSON, value)
	}
	if value, ok := muo.mutation.AppendedActors(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, movie.FieldActors, value)
		})
	}
	if value, ok := muo.mutation.Description(); ok {
		_spec.SetField(movie.FieldDescription, field.TypeString, value)
	}
	if value, ok := muo.mutation.ReleaseDate(); ok {
		_spec.SetField(movie.FieldReleaseDate, field.TypeTime, value)
	}
	_node = &Movie{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}