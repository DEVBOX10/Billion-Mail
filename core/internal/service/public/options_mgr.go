package public

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// OptionsMgr Options Manager
type OptionsMgr struct {
	cache      *gcache.Cache // Cache instance
	expiration time.Duration // Cache expiration time
}

// NewOptionsMgr Create options manager
func NewOptionsMgr() *OptionsMgr {
	return &OptionsMgr{
		cache:      gcache.New(),
		expiration: time.Hour * 24, // Default 24 hours cache
	}
}

// OptionsMgrInstance Options manager singleton
var OptionsMgrInstance = NewOptionsMgr()

// SetOption Set option
func (o *OptionsMgr) SetOption(ctx context.Context, key string, value interface{}) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	if value == nil {
		return errors.New("value cannot be nil")
	}

	// Serialize value
	jsonValue, err := o.serialize(value)
	if err != nil {
		return err
	}

	// Save to database
	if err := o.saveToDatabase(ctx, key, jsonValue); err != nil {
		return err
	}

	// Update cache
	cacheKey := o.buildCacheKey(key)
	if err := o.cache.Set(ctx, cacheKey, jsonValue, o.expiration); err != nil {
		return err
	}

	return nil
}

// GetOption Get option and deserialize to specified type
func (o *OptionsMgr) GetOption(ctx context.Context, key string, ptr interface{}) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	if ptr == nil {
		return errors.New("pointer cannot be nil")
	}

	// Try to get from cache
	cacheKey := o.buildCacheKey(key)
	cached, err := o.cache.Get(ctx, cacheKey)
	var jsonValue string

	if err != nil || cached == nil {
		// Cache miss, read from database
		var result struct {
			Value string `json:"value"`
		}

		err := g.DB().Model("bm_options").
			Where("name", key).
			Fields("value").
			Scan(&result)

		if err != nil {
			return errors.New("failed to get option from database: " + err.Error())
		}

		if result.Value == "" {
			return errors.New("option not found")
		}

		jsonValue = result.Value

		// Store in cache
		if err := o.cache.Set(ctx, cacheKey, jsonValue, o.expiration); err != nil {
			g.Log().Warning(ctx, "Failed to set option cache:", err)
		}
	} else {
		// Cache hit
		jsonValue = cached.String()
	}

	// Deserialize
	return o.deserialize(jsonValue, ptr)
}

// GetAllOptions Get all options
func (o *OptionsMgr) GetAllOptions(ctx context.Context) (map[string]string, error) {
	var options []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	}

	err := g.DB().Model("bm_options").Scan(&options)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string, len(options))
	for _, opt := range options {
		result[opt.Name] = opt.Value

		// Update cache
		cacheKey := o.buildCacheKey(opt.Name)
		if err := o.cache.Set(ctx, cacheKey, opt.Value, o.expiration); err != nil {
			g.Log().Warning(ctx, "Failed to set option cache when loading all:", err)
		}
	}

	return result, nil
}

// DeleteOption Delete option
func (o *OptionsMgr) DeleteOption(ctx context.Context, key string) (err error) {
	if key == "" {
		return errors.New("key cannot be empty")
	}

	// Delete from database
	_, err = g.DB().Model("bm_options").Where("name", key).Delete()
	if err != nil {
		return
	}

	// Delete from cache
	cacheKey := o.buildCacheKey(key)
	_, err = o.cache.Remove(ctx, cacheKey)

	return
}

// SetCacheExpiration Set cache expiration time
func (o *OptionsMgr) SetCacheExpiration(expiration time.Duration) {
	o.expiration = expiration
}

// ClearCache Clear all cache
func (o *OptionsMgr) ClearCache(ctx context.Context) error {
	return o.cache.Clear(ctx)
}

// buildCacheKey Build cache key
func (o *OptionsMgr) buildCacheKey(key string) string {
	return "BM_OPTION:" + key
}

// serialize Serialize value to JSON string
func (o *OptionsMgr) serialize(value interface{}) (string, error) {
	// If already a string, return directly
	if str, ok := value.(string); ok {
		return str, nil
	}

	// Serialize other types to JSON
	bytes, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// deserialize Deserialize JSON string to specified type
func (o *OptionsMgr) deserialize(jsonValue string, ptr interface{}) error {
	// If target is string pointer, assign directly
	if strPtr, ok := ptr.(*string); ok {
		*strPtr = jsonValue
		return nil
	}

	// Deserialize from JSON for other types
	return json.Unmarshal([]byte(jsonValue), ptr)
}

// saveToDatabase Save to database
func (o *OptionsMgr) saveToDatabase(ctx context.Context, key string, value string) error {
	_, err := g.DB().Model("bm_options").
		Where("name", key).
		OnConflict("name").
		OnDuplicate("value").
		Data(g.Map{
			"name":  key,
			"value": value,
		}).
		Save()

	return err
}
