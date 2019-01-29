// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// ReputationConfig is an object representing the database table.
type ReputationConfig struct {
	GuildID                 int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	PointsName              string           `boil:"points_name" json:"points_name" toml:"points_name" yaml:"points_name"`
	Enabled                 bool             `boil:"enabled" json:"enabled" toml:"enabled" yaml:"enabled"`
	Cooldown                int              `boil:"cooldown" json:"cooldown" toml:"cooldown" yaml:"cooldown"`
	MaxGiveAmount           int64            `boil:"max_give_amount" json:"max_give_amount" toml:"max_give_amount" yaml:"max_give_amount"`
	RequiredGiveRole        null.String      `boil:"required_give_role" json:"required_give_role,omitempty" toml:"required_give_role" yaml:"required_give_role,omitempty"`
	RequiredReceiveRole     null.String      `boil:"required_receive_role" json:"required_receive_role,omitempty" toml:"required_receive_role" yaml:"required_receive_role,omitempty"`
	BlacklistedGiveRole     null.String      `boil:"blacklisted_give_role" json:"blacklisted_give_role,omitempty" toml:"blacklisted_give_role" yaml:"blacklisted_give_role,omitempty"`
	BlacklistedReceiveRole  null.String      `boil:"blacklisted_receive_role" json:"blacklisted_receive_role,omitempty" toml:"blacklisted_receive_role" yaml:"blacklisted_receive_role,omitempty"`
	AdminRole               null.String      `boil:"admin_role" json:"admin_role,omitempty" toml:"admin_role" yaml:"admin_role,omitempty"`
	DisableThanksDetection  bool             `boil:"disable_thanks_detection" json:"disable_thanks_detection" toml:"disable_thanks_detection" yaml:"disable_thanks_detection"`
	MaxRemoveAmount         int64            `boil:"max_remove_amount" json:"max_remove_amount" toml:"max_remove_amount" yaml:"max_remove_amount"`
	AdminRoles              types.Int64Array `boil:"admin_roles" json:"admin_roles,omitempty" toml:"admin_roles" yaml:"admin_roles,omitempty"`
	RequiredGiveRoles       types.Int64Array `boil:"required_give_roles" json:"required_give_roles,omitempty" toml:"required_give_roles" yaml:"required_give_roles,omitempty"`
	RequiredReceiveRoles    types.Int64Array `boil:"required_receive_roles" json:"required_receive_roles,omitempty" toml:"required_receive_roles" yaml:"required_receive_roles,omitempty"`
	BlacklistedGiveRoles    types.Int64Array `boil:"blacklisted_give_roles" json:"blacklisted_give_roles,omitempty" toml:"blacklisted_give_roles" yaml:"blacklisted_give_roles,omitempty"`
	BlacklistedReceiveRoles types.Int64Array `boil:"blacklisted_receive_roles" json:"blacklisted_receive_roles,omitempty" toml:"blacklisted_receive_roles" yaml:"blacklisted_receive_roles,omitempty"`

	R *reputationConfigR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L reputationConfigL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ReputationConfigColumns = struct {
	GuildID                 string
	PointsName              string
	Enabled                 string
	Cooldown                string
	MaxGiveAmount           string
	RequiredGiveRole        string
	RequiredReceiveRole     string
	BlacklistedGiveRole     string
	BlacklistedReceiveRole  string
	AdminRole               string
	DisableThanksDetection  string
	MaxRemoveAmount         string
	AdminRoles              string
	RequiredGiveRoles       string
	RequiredReceiveRoles    string
	BlacklistedGiveRoles    string
	BlacklistedReceiveRoles string
}{
	GuildID:                 "guild_id",
	PointsName:              "points_name",
	Enabled:                 "enabled",
	Cooldown:                "cooldown",
	MaxGiveAmount:           "max_give_amount",
	RequiredGiveRole:        "required_give_role",
	RequiredReceiveRole:     "required_receive_role",
	BlacklistedGiveRole:     "blacklisted_give_role",
	BlacklistedReceiveRole:  "blacklisted_receive_role",
	AdminRole:               "admin_role",
	DisableThanksDetection:  "disable_thanks_detection",
	MaxRemoveAmount:         "max_remove_amount",
	AdminRoles:              "admin_roles",
	RequiredGiveRoles:       "required_give_roles",
	RequiredReceiveRoles:    "required_receive_roles",
	BlacklistedGiveRoles:    "blacklisted_give_roles",
	BlacklistedReceiveRoles: "blacklisted_receive_roles",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelpertypes_Int64Array struct{ field string }

func (w whereHelpertypes_Int64Array) EQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpertypes_Int64Array) NEQ(x types.Int64Array) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpertypes_Int64Array) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpertypes_Int64Array) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpertypes_Int64Array) LT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertypes_Int64Array) LTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertypes_Int64Array) GT(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertypes_Int64Array) GTE(x types.Int64Array) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var ReputationConfigWhere = struct {
	GuildID                 whereHelperint64
	PointsName              whereHelperstring
	Enabled                 whereHelperbool
	Cooldown                whereHelperint
	MaxGiveAmount           whereHelperint64
	RequiredGiveRole        whereHelpernull_String
	RequiredReceiveRole     whereHelpernull_String
	BlacklistedGiveRole     whereHelpernull_String
	BlacklistedReceiveRole  whereHelpernull_String
	AdminRole               whereHelpernull_String
	DisableThanksDetection  whereHelperbool
	MaxRemoveAmount         whereHelperint64
	AdminRoles              whereHelpertypes_Int64Array
	RequiredGiveRoles       whereHelpertypes_Int64Array
	RequiredReceiveRoles    whereHelpertypes_Int64Array
	BlacklistedGiveRoles    whereHelpertypes_Int64Array
	BlacklistedReceiveRoles whereHelpertypes_Int64Array
}{
	GuildID:                 whereHelperint64{field: `guild_id`},
	PointsName:              whereHelperstring{field: `points_name`},
	Enabled:                 whereHelperbool{field: `enabled`},
	Cooldown:                whereHelperint{field: `cooldown`},
	MaxGiveAmount:           whereHelperint64{field: `max_give_amount`},
	RequiredGiveRole:        whereHelpernull_String{field: `required_give_role`},
	RequiredReceiveRole:     whereHelpernull_String{field: `required_receive_role`},
	BlacklistedGiveRole:     whereHelpernull_String{field: `blacklisted_give_role`},
	BlacklistedReceiveRole:  whereHelpernull_String{field: `blacklisted_receive_role`},
	AdminRole:               whereHelpernull_String{field: `admin_role`},
	DisableThanksDetection:  whereHelperbool{field: `disable_thanks_detection`},
	MaxRemoveAmount:         whereHelperint64{field: `max_remove_amount`},
	AdminRoles:              whereHelpertypes_Int64Array{field: `admin_roles`},
	RequiredGiveRoles:       whereHelpertypes_Int64Array{field: `required_give_roles`},
	RequiredReceiveRoles:    whereHelpertypes_Int64Array{field: `required_receive_roles`},
	BlacklistedGiveRoles:    whereHelpertypes_Int64Array{field: `blacklisted_give_roles`},
	BlacklistedReceiveRoles: whereHelpertypes_Int64Array{field: `blacklisted_receive_roles`},
}

// ReputationConfigRels is where relationship names are stored.
var ReputationConfigRels = struct {
}{}

// reputationConfigR is where relationships are stored.
type reputationConfigR struct {
}

// NewStruct creates a new relationship struct
func (*reputationConfigR) NewStruct() *reputationConfigR {
	return &reputationConfigR{}
}

// reputationConfigL is where Load methods for each relationship are stored.
type reputationConfigL struct{}

var (
	reputationConfigColumns               = []string{"guild_id", "points_name", "enabled", "cooldown", "max_give_amount", "required_give_role", "required_receive_role", "blacklisted_give_role", "blacklisted_receive_role", "admin_role", "disable_thanks_detection", "max_remove_amount", "admin_roles", "required_give_roles", "required_receive_roles", "blacklisted_give_roles", "blacklisted_receive_roles"}
	reputationConfigColumnsWithoutDefault = []string{"guild_id", "points_name", "enabled", "cooldown", "max_give_amount", "required_give_role", "required_receive_role", "blacklisted_give_role", "blacklisted_receive_role", "admin_role", "admin_roles", "required_give_roles", "required_receive_roles", "blacklisted_give_roles", "blacklisted_receive_roles"}
	reputationConfigColumnsWithDefault    = []string{"disable_thanks_detection", "max_remove_amount"}
	reputationConfigPrimaryKeyColumns     = []string{"guild_id"}
)

type (
	// ReputationConfigSlice is an alias for a slice of pointers to ReputationConfig.
	// This should generally be used opposed to []ReputationConfig.
	ReputationConfigSlice []*ReputationConfig

	reputationConfigQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	reputationConfigType                 = reflect.TypeOf(&ReputationConfig{})
	reputationConfigMapping              = queries.MakeStructMapping(reputationConfigType)
	reputationConfigPrimaryKeyMapping, _ = queries.BindMapping(reputationConfigType, reputationConfigMapping, reputationConfigPrimaryKeyColumns)
	reputationConfigInsertCacheMut       sync.RWMutex
	reputationConfigInsertCache          = make(map[string]insertCache)
	reputationConfigUpdateCacheMut       sync.RWMutex
	reputationConfigUpdateCache          = make(map[string]updateCache)
	reputationConfigUpsertCacheMut       sync.RWMutex
	reputationConfigUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single reputationConfig record from the query using the global executor.
func (q reputationConfigQuery) OneG(ctx context.Context) (*ReputationConfig, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single reputationConfig record from the query.
func (q reputationConfigQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ReputationConfig, error) {
	o := &ReputationConfig{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for reputation_configs")
	}

	return o, nil
}

// AllG returns all ReputationConfig records from the query using the global executor.
func (q reputationConfigQuery) AllG(ctx context.Context) (ReputationConfigSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all ReputationConfig records from the query.
func (q reputationConfigQuery) All(ctx context.Context, exec boil.ContextExecutor) (ReputationConfigSlice, error) {
	var o []*ReputationConfig

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ReputationConfig slice")
	}

	return o, nil
}

// CountG returns the count of all ReputationConfig records in the query, and panics on error.
func (q reputationConfigQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all ReputationConfig records in the query.
func (q reputationConfigQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count reputation_configs rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q reputationConfigQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q reputationConfigQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if reputation_configs exists")
	}

	return count > 0, nil
}

// ReputationConfigs retrieves all the records using an executor.
func ReputationConfigs(mods ...qm.QueryMod) reputationConfigQuery {
	mods = append(mods, qm.From("\"reputation_configs\""))
	return reputationConfigQuery{NewQuery(mods...)}
}

// FindReputationConfigG retrieves a single record by ID.
func FindReputationConfigG(ctx context.Context, guildID int64, selectCols ...string) (*ReputationConfig, error) {
	return FindReputationConfig(ctx, boil.GetContextDB(), guildID, selectCols...)
}

// FindReputationConfig retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReputationConfig(ctx context.Context, exec boil.ContextExecutor, guildID int64, selectCols ...string) (*ReputationConfig, error) {
	reputationConfigObj := &ReputationConfig{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"reputation_configs\" where \"guild_id\"=$1", sel,
	)

	q := queries.Raw(query, guildID)

	err := q.Bind(ctx, exec, reputationConfigObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from reputation_configs")
	}

	return reputationConfigObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *ReputationConfig) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ReputationConfig) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reputation_configs provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(reputationConfigColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	reputationConfigInsertCacheMut.RLock()
	cache, cached := reputationConfigInsertCache[key]
	reputationConfigInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			reputationConfigColumns,
			reputationConfigColumnsWithDefault,
			reputationConfigColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"reputation_configs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"reputation_configs\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into reputation_configs")
	}

	if !cached {
		reputationConfigInsertCacheMut.Lock()
		reputationConfigInsertCache[key] = cache
		reputationConfigInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single ReputationConfig record using the global executor.
// See Update for more documentation.
func (o *ReputationConfig) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the ReputationConfig.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ReputationConfig) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	reputationConfigUpdateCacheMut.RLock()
	cache, cached := reputationConfigUpdateCache[key]
	reputationConfigUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			reputationConfigColumns,
			reputationConfigPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update reputation_configs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"reputation_configs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, reputationConfigPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, append(wl, reputationConfigPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update reputation_configs row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for reputation_configs")
	}

	if !cached {
		reputationConfigUpdateCacheMut.Lock()
		reputationConfigUpdateCache[key] = cache
		reputationConfigUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q reputationConfigQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q reputationConfigQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for reputation_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for reputation_configs")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ReputationConfigSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReputationConfigSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"reputation_configs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, reputationConfigPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in reputationConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all reputationConfig")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *ReputationConfig) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ReputationConfig) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no reputation_configs provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(reputationConfigColumnsWithDefault, o)

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

	reputationConfigUpsertCacheMut.RLock()
	cache, cached := reputationConfigUpsertCache[key]
	reputationConfigUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			reputationConfigColumns,
			reputationConfigColumnsWithDefault,
			reputationConfigColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			reputationConfigColumns,
			reputationConfigPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert reputation_configs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(reputationConfigPrimaryKeyColumns))
			copy(conflict, reputationConfigPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"reputation_configs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(reputationConfigType, reputationConfigMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert reputation_configs")
	}

	if !cached {
		reputationConfigUpsertCacheMut.Lock()
		reputationConfigUpsertCache[key] = cache
		reputationConfigUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single ReputationConfig record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *ReputationConfig) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single ReputationConfig record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ReputationConfig) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ReputationConfig provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), reputationConfigPrimaryKeyMapping)
	sql := "DELETE FROM \"reputation_configs\" WHERE \"guild_id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from reputation_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for reputation_configs")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q reputationConfigQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no reputationConfigQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reputation_configs")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reputation_configs")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o ReputationConfigSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReputationConfigSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ReputationConfig slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"reputation_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reputationConfigPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from reputationConfig slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for reputation_configs")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *ReputationConfig) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no ReputationConfig provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *ReputationConfig) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindReputationConfig(ctx, exec, o.GuildID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationConfigSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty ReputationConfigSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReputationConfigSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ReputationConfigSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), reputationConfigPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"reputation_configs\".* FROM \"reputation_configs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, reputationConfigPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReputationConfigSlice")
	}

	*o = slice

	return nil
}

// ReputationConfigExistsG checks if the ReputationConfig row exists.
func ReputationConfigExistsG(ctx context.Context, guildID int64) (bool, error) {
	return ReputationConfigExists(ctx, boil.GetContextDB(), guildID)
}

// ReputationConfigExists checks if the ReputationConfig row exists.
func ReputationConfigExists(ctx context.Context, exec boil.ContextExecutor, guildID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"reputation_configs\" where \"guild_id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if reputation_configs exists")
	}

	return exists, nil
}
