// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// CTFChannel is an object representing the database table.
type CTFChannel struct {
	ID        string      `boil:"id" json:"id" toml:"id" yaml:"id"`
	TopicChan string      `boil:"topic_chan" json:"topic_chan" toml:"topic_chan" yaml:"topic_chan"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	GuildID   string      `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Title     string      `boil:"title" json:"title" toml:"title" yaml:"title"`
	URL       string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	Archived  bool        `boil:"archived" json:"archived" toml:"archived" yaml:"archived"`
	Username  null.String `boil:"username" json:"username,omitempty" toml:"username" yaml:"username,omitempty"`
	Password  null.String `boil:"password" json:"password,omitempty" toml:"password" yaml:"password,omitempty"`
	CtftimeID null.Int    `boil:"ctftime_id" json:"ctftime_id,omitempty" toml:"ctftime_id" yaml:"ctftime_id,omitempty"`

	R *ctfChannelR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ctfChannelL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CTFChannelColumns = struct {
	ID        string
	TopicChan string
	CreatedAt string
	GuildID   string
	Title     string
	URL       string
	Archived  string
	Username  string
	Password  string
	CtftimeID string
}{
	ID:        "id",
	TopicChan: "topic_chan",
	CreatedAt: "created_at",
	GuildID:   "guild_id",
	Title:     "title",
	URL:       "url",
	Archived:  "archived",
	Username:  "username",
	Password:  "password",
	CtftimeID: "ctftime_id",
}

var CTFChannelTableColumns = struct {
	ID        string
	TopicChan string
	CreatedAt string
	GuildID   string
	Title     string
	URL       string
	Archived  string
	Username  string
	Password  string
	CtftimeID string
}{
	ID:        "ctf_channels.id",
	TopicChan: "ctf_channels.topic_chan",
	CreatedAt: "ctf_channels.created_at",
	GuildID:   "ctf_channels.guild_id",
	Title:     "ctf_channels.title",
	URL:       "ctf_channels.url",
	Archived:  "ctf_channels.archived",
	Username:  "ctf_channels.username",
	Password:  "ctf_channels.password",
	CtftimeID: "ctf_channels.ctftime_id",
}

// Generated where

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var CTFChannelWhere = struct {
	ID        whereHelperstring
	TopicChan whereHelperstring
	CreatedAt whereHelpertime_Time
	GuildID   whereHelperstring
	Title     whereHelperstring
	URL       whereHelperstring
	Archived  whereHelperbool
	Username  whereHelpernull_String
	Password  whereHelpernull_String
	CtftimeID whereHelpernull_Int
}{
	ID:        whereHelperstring{field: "\"ctf_channels\".\"id\""},
	TopicChan: whereHelperstring{field: "\"ctf_channels\".\"topic_chan\""},
	CreatedAt: whereHelpertime_Time{field: "\"ctf_channels\".\"created_at\""},
	GuildID:   whereHelperstring{field: "\"ctf_channels\".\"guild_id\""},
	Title:     whereHelperstring{field: "\"ctf_channels\".\"title\""},
	URL:       whereHelperstring{field: "\"ctf_channels\".\"url\""},
	Archived:  whereHelperbool{field: "\"ctf_channels\".\"archived\""},
	Username:  whereHelpernull_String{field: "\"ctf_channels\".\"username\""},
	Password:  whereHelpernull_String{field: "\"ctf_channels\".\"password\""},
	CtftimeID: whereHelpernull_Int{field: "\"ctf_channels\".\"ctftime_id\""},
}

// CTFChannelRels is where relationship names are stored.
var CTFChannelRels = struct {
	Guild               string
	ParentChallChannels string
}{
	Guild:               "Guild",
	ParentChallChannels: "ParentChallChannels",
}

// ctfChannelR is where relationships are stored.
type ctfChannelR struct {
	Guild               *Guild            `boil:"Guild" json:"Guild" toml:"Guild" yaml:"Guild"`
	ParentChallChannels ChallChannelSlice `boil:"ParentChallChannels" json:"ParentChallChannels" toml:"ParentChallChannels" yaml:"ParentChallChannels"`
}

// NewStruct creates a new relationship struct
func (*ctfChannelR) NewStruct() *ctfChannelR {
	return &ctfChannelR{}
}

func (r *ctfChannelR) GetGuild() *Guild {
	if r == nil {
		return nil
	}
	return r.Guild
}

func (r *ctfChannelR) GetParentChallChannels() ChallChannelSlice {
	if r == nil {
		return nil
	}
	return r.ParentChallChannels
}

// ctfChannelL is where Load methods for each relationship are stored.
type ctfChannelL struct{}

var (
	ctfChannelAllColumns            = []string{"id", "topic_chan", "created_at", "guild_id", "title", "url", "archived", "username", "password", "ctftime_id"}
	ctfChannelColumnsWithoutDefault = []string{"id", "topic_chan", "created_at", "guild_id", "title", "url", "archived"}
	ctfChannelColumnsWithDefault    = []string{"username", "password", "ctftime_id"}
	ctfChannelPrimaryKeyColumns     = []string{"id"}
	ctfChannelGeneratedColumns      = []string{}
)

type (
	// CTFChannelSlice is an alias for a slice of pointers to CTFChannel.
	// This should almost always be used instead of []CTFChannel.
	CTFChannelSlice []*CTFChannel
	// CTFChannelHook is the signature for custom CTFChannel hook methods
	CTFChannelHook func(context.Context, boil.ContextExecutor, *CTFChannel) error

	ctfChannelQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ctfChannelType                 = reflect.TypeOf(&CTFChannel{})
	ctfChannelMapping              = queries.MakeStructMapping(ctfChannelType)
	ctfChannelPrimaryKeyMapping, _ = queries.BindMapping(ctfChannelType, ctfChannelMapping, ctfChannelPrimaryKeyColumns)
	ctfChannelInsertCacheMut       sync.RWMutex
	ctfChannelInsertCache          = make(map[string]insertCache)
	ctfChannelUpdateCacheMut       sync.RWMutex
	ctfChannelUpdateCache          = make(map[string]updateCache)
	ctfChannelUpsertCacheMut       sync.RWMutex
	ctfChannelUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var ctfChannelAfterSelectHooks []CTFChannelHook

var ctfChannelBeforeInsertHooks []CTFChannelHook
var ctfChannelAfterInsertHooks []CTFChannelHook

var ctfChannelBeforeUpdateHooks []CTFChannelHook
var ctfChannelAfterUpdateHooks []CTFChannelHook

var ctfChannelBeforeDeleteHooks []CTFChannelHook
var ctfChannelAfterDeleteHooks []CTFChannelHook

var ctfChannelBeforeUpsertHooks []CTFChannelHook
var ctfChannelAfterUpsertHooks []CTFChannelHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CTFChannel) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CTFChannel) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CTFChannel) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CTFChannel) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CTFChannel) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CTFChannel) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CTFChannel) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CTFChannel) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CTFChannel) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range ctfChannelAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCTFChannelHook registers your hook function for all future operations.
func AddCTFChannelHook(hookPoint boil.HookPoint, ctfChannelHook CTFChannelHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		ctfChannelAfterSelectHooks = append(ctfChannelAfterSelectHooks, ctfChannelHook)
	case boil.BeforeInsertHook:
		ctfChannelBeforeInsertHooks = append(ctfChannelBeforeInsertHooks, ctfChannelHook)
	case boil.AfterInsertHook:
		ctfChannelAfterInsertHooks = append(ctfChannelAfterInsertHooks, ctfChannelHook)
	case boil.BeforeUpdateHook:
		ctfChannelBeforeUpdateHooks = append(ctfChannelBeforeUpdateHooks, ctfChannelHook)
	case boil.AfterUpdateHook:
		ctfChannelAfterUpdateHooks = append(ctfChannelAfterUpdateHooks, ctfChannelHook)
	case boil.BeforeDeleteHook:
		ctfChannelBeforeDeleteHooks = append(ctfChannelBeforeDeleteHooks, ctfChannelHook)
	case boil.AfterDeleteHook:
		ctfChannelAfterDeleteHooks = append(ctfChannelAfterDeleteHooks, ctfChannelHook)
	case boil.BeforeUpsertHook:
		ctfChannelBeforeUpsertHooks = append(ctfChannelBeforeUpsertHooks, ctfChannelHook)
	case boil.AfterUpsertHook:
		ctfChannelAfterUpsertHooks = append(ctfChannelAfterUpsertHooks, ctfChannelHook)
	}
}

// One returns a single ctfChannel record from the query.
func (q ctfChannelQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CTFChannel, error) {
	o := &CTFChannel{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for ctf_channels")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CTFChannel records from the query.
func (q ctfChannelQuery) All(ctx context.Context, exec boil.ContextExecutor) (CTFChannelSlice, error) {
	var o []*CTFChannel

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CTFChannel slice")
	}

	if len(ctfChannelAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CTFChannel records in the query.
func (q ctfChannelQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count ctf_channels rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q ctfChannelQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if ctf_channels exists")
	}

	return count > 0, nil
}

// Guild pointed to by the foreign key.
func (o *CTFChannel) Guild(mods ...qm.QueryMod) guildQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GuildID),
	}

	queryMods = append(queryMods, mods...)

	return Guilds(queryMods...)
}

// ParentChallChannels retrieves all the chall_channel's ChallChannels with an executor via parent_id column.
func (o *CTFChannel) ParentChallChannels(mods ...qm.QueryMod) challChannelQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"chall_channels\".\"parent_id\"=?", o.ID),
	)

	return ChallChannels(queryMods...)
}

// LoadGuild allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (ctfChannelL) LoadGuild(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCTFChannel interface{}, mods queries.Applicator) error {
	var slice []*CTFChannel
	var object *CTFChannel

	if singular {
		var ok bool
		object, ok = maybeCTFChannel.(*CTFChannel)
		if !ok {
			object = new(CTFChannel)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCTFChannel)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCTFChannel))
			}
		}
	} else {
		s, ok := maybeCTFChannel.(*[]*CTFChannel)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCTFChannel)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCTFChannel))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ctfChannelR{}
		}
		args = append(args, object.GuildID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ctfChannelR{}
			}

			for _, a := range args {
				if a == obj.GuildID {
					continue Outer
				}
			}

			args = append(args, obj.GuildID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`guilds`),
		qm.WhereIn(`guilds.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Guild")
	}

	var resultSlice []*Guild
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Guild")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for guilds")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for guilds")
	}

	if len(ctfChannelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Guild = foreign
		if foreign.R == nil {
			foreign.R = &guildR{}
		}
		foreign.R.CTFChannels = append(foreign.R.CTFChannels, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GuildID == foreign.ID {
				local.R.Guild = foreign
				if foreign.R == nil {
					foreign.R = &guildR{}
				}
				foreign.R.CTFChannels = append(foreign.R.CTFChannels, local)
				break
			}
		}
	}

	return nil
}

// LoadParentChallChannels allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (ctfChannelL) LoadParentChallChannels(ctx context.Context, e boil.ContextExecutor, singular bool, maybeCTFChannel interface{}, mods queries.Applicator) error {
	var slice []*CTFChannel
	var object *CTFChannel

	if singular {
		var ok bool
		object, ok = maybeCTFChannel.(*CTFChannel)
		if !ok {
			object = new(CTFChannel)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeCTFChannel)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeCTFChannel))
			}
		}
	} else {
		s, ok := maybeCTFChannel.(*[]*CTFChannel)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeCTFChannel)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeCTFChannel))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &ctfChannelR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &ctfChannelR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`chall_channels`),
		qm.WhereIn(`chall_channels.parent_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load chall_channels")
	}

	var resultSlice []*ChallChannel
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice chall_channels")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on chall_channels")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for chall_channels")
	}

	if len(challChannelAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ParentChallChannels = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &challChannelR{}
			}
			foreign.R.Parent = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ParentID {
				local.R.ParentChallChannels = append(local.R.ParentChallChannels, foreign)
				if foreign.R == nil {
					foreign.R = &challChannelR{}
				}
				foreign.R.Parent = local
				break
			}
		}
	}

	return nil
}

// SetGuild of the ctfChannel to the related item.
// Sets o.R.Guild to related.
// Adds o to related.R.CTFChannels.
func (o *CTFChannel) SetGuild(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Guild) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"ctf_channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"guild_id"}),
		strmangle.WhereClause("\"", "\"", 2, ctfChannelPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GuildID = related.ID
	if o.R == nil {
		o.R = &ctfChannelR{
			Guild: related,
		}
	} else {
		o.R.Guild = related
	}

	if related.R == nil {
		related.R = &guildR{
			CTFChannels: CTFChannelSlice{o},
		}
	} else {
		related.R.CTFChannels = append(related.R.CTFChannels, o)
	}

	return nil
}

// AddParentChallChannels adds the given related objects to the existing relationships
// of the ctf_channel, optionally inserting them as new records.
// Appends related to o.R.ParentChallChannels.
// Sets related.R.Parent appropriately.
func (o *CTFChannel) AddParentChallChannels(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ChallChannel) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ParentID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"chall_channels\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"parent_id"}),
				strmangle.WhereClause("\"", "\"", 2, challChannelPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ParentID = o.ID
		}
	}

	if o.R == nil {
		o.R = &ctfChannelR{
			ParentChallChannels: related,
		}
	} else {
		o.R.ParentChallChannels = append(o.R.ParentChallChannels, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &challChannelR{
				Parent: o,
			}
		} else {
			rel.R.Parent = o
		}
	}
	return nil
}

// CTFChannels retrieves all the records using an executor.
func CTFChannels(mods ...qm.QueryMod) ctfChannelQuery {
	mods = append(mods, qm.From("\"ctf_channels\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"ctf_channels\".*"})
	}

	return ctfChannelQuery{q}
}

// FindCTFChannel retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCTFChannel(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*CTFChannel, error) {
	ctfChannelObj := &CTFChannel{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"ctf_channels\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, ctfChannelObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from ctf_channels")
	}

	if err = ctfChannelObj.doAfterSelectHooks(ctx, exec); err != nil {
		return ctfChannelObj, err
	}

	return ctfChannelObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CTFChannel) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ctf_channels provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ctfChannelColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ctfChannelInsertCacheMut.RLock()
	cache, cached := ctfChannelInsertCache[key]
	ctfChannelInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ctfChannelAllColumns,
			ctfChannelColumnsWithDefault,
			ctfChannelColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ctfChannelType, ctfChannelMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ctfChannelType, ctfChannelMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"ctf_channels\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"ctf_channels\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into ctf_channels")
	}

	if !cached {
		ctfChannelInsertCacheMut.Lock()
		ctfChannelInsertCache[key] = cache
		ctfChannelInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CTFChannel.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CTFChannel) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	ctfChannelUpdateCacheMut.RLock()
	cache, cached := ctfChannelUpdateCache[key]
	ctfChannelUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ctfChannelAllColumns,
			ctfChannelPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update ctf_channels, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"ctf_channels\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ctfChannelPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ctfChannelType, ctfChannelMapping, append(wl, ctfChannelPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update ctf_channels row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for ctf_channels")
	}

	if !cached {
		ctfChannelUpdateCacheMut.Lock()
		ctfChannelUpdateCache[key] = cache
		ctfChannelUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q ctfChannelQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for ctf_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for ctf_channels")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CTFChannelSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ctfChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"ctf_channels\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ctfChannelPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in ctfChannel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all ctfChannel")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CTFChannel) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no ctf_channels provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(ctfChannelColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	ctfChannelUpsertCacheMut.RLock()
	cache, cached := ctfChannelUpsertCache[key]
	ctfChannelUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ctfChannelAllColumns,
			ctfChannelColumnsWithDefault,
			ctfChannelColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			ctfChannelAllColumns,
			ctfChannelPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert ctf_channels, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ctfChannelPrimaryKeyColumns))
			copy(conflict, ctfChannelPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"ctf_channels\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ctfChannelType, ctfChannelMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ctfChannelType, ctfChannelMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert ctf_channels")
	}

	if !cached {
		ctfChannelUpsertCacheMut.Lock()
		ctfChannelUpsertCache[key] = cache
		ctfChannelUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CTFChannel record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CTFChannel) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CTFChannel provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ctfChannelPrimaryKeyMapping)
	sql := "DELETE FROM \"ctf_channels\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from ctf_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for ctf_channels")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ctfChannelQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ctfChannelQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ctf_channels")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ctf_channels")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CTFChannelSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(ctfChannelBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ctfChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"ctf_channels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ctfChannelPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from ctfChannel slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for ctf_channels")
	}

	if len(ctfChannelAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CTFChannel) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCTFChannel(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CTFChannelSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CTFChannelSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ctfChannelPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"ctf_channels\".* FROM \"ctf_channels\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ctfChannelPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CTFChannelSlice")
	}

	*o = slice

	return nil
}

// CTFChannelExists checks if the CTFChannel row exists.
func CTFChannelExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"ctf_channels\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if ctf_channels exists")
	}

	return exists, nil
}