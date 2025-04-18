package rbac

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/mojocn/base64Captcha"
	"image/color"
	"sync"
	"time"
)

// Default configuration constants
const (
	DefaultPoolSize    = 100             // Default size of captcha pool
	DefaultExpiration  = 3 * time.Minute // Default expiration time for captchas
	DefaultRefreshRate = 60              // Threshold to trigger pool refresh when captcha count is lower
)

// Captcha configuration structure
var (
	defaultDriver = &base64Captcha.DriverString{
		Height:          50,
		Width:           200,
		NoiseCount:      0,
		ShowLineOptions: base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "34578abcdefhjkmnprstuvwxy",
		BgColor: &color.RGBA{
			R: 40,
			G: 30,
			B: 89,
			A: 29,
		},
		Fonts: nil,
	}
)

// CodePool Captcha pool structure
type CodePool struct {
	store        base64Captcha.Store         // Captcha storage
	driver       *base64Captcha.DriverString // Captcha driver configuration
	captcha      *base64Captcha.Captcha      // Captcha generator
	pool         map[string]string           // Captcha pool: id -> answer
	mutex        sync.RWMutex                // Read-write lock
	poolSize     int                         // Pool capacity
	refreshRate  int                         // Refresh threshold
	expiration   time.Duration               // Captcha expiration time
	refreshTimer *time.Timer                 // Refresh timer
}

// Global captcha pool instance
var (
	globalPool *CodePool
	once       sync.Once
)

// GetCodePool Get global captcha pool instance
func GetCodePool() *CodePool {
	once.Do(func() {
		globalPool = NewCodePool(DefaultPoolSize, DefaultExpiration, DefaultRefreshRate)
		globalPool.Start()
	})
	return globalPool
}

// NewCodePool Create new captcha pool
func NewCodePool(poolSize int, expiration time.Duration, refreshRate int) *CodePool {
	store := base64Captcha.NewMemoryStore(poolSize*2, expiration)

	pool := &CodePool{
		store:       store,
		driver:      defaultDriver,
		captcha:     base64Captcha.NewCaptcha(defaultDriver, store),
		pool:        make(map[string]string, poolSize),
		poolSize:    poolSize,
		refreshRate: refreshRate,
		expiration:  expiration,
	}

	return pool
}

// ConfigDriver Set captcha driver configuration
func (p *CodePool) ConfigDriver(driver *base64Captcha.DriverString) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if driver != nil {
		p.driver = driver
		p.captcha = base64Captcha.NewCaptcha(p.driver, p.store)
	}
}

// SetExpiration Set captcha expiration time
func (p *CodePool) SetExpiration(expiration time.Duration) {
	if expiration > 0 {
		p.mutex.Lock()
		p.expiration = expiration
		p.mutex.Unlock()

		// Create new storage to apply new expiration time
		p.store = base64Captcha.NewMemoryStore(p.poolSize*2, expiration)
		p.captcha = base64Captcha.NewCaptcha(p.driver, p.store)

		// Refresh pool
		p.RefreshPool(true)
	}
}

// Start Start automatic refresh of captcha pool
func (p *CodePool) Start() {
	go p.autoRefresh()
}

// autoRefresh Automatically refresh captcha pool
func (p *CodePool) autoRefresh() {
	ticker := time.NewTicker(p.expiration / 2) // Check every half of expiration time
	defer ticker.Stop()

	for {
		<-ticker.C
		p.RefreshPool(false)
	}
}

// RefreshPool Refresh captcha pool
func (p *CodePool) RefreshPool(force bool) {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	ctx := context.Background()
	currentSize := len(p.pool)

	// Refresh pool if forced or pool size is below threshold
	if force || currentSize < p.refreshRate {
		neededCodes := p.poolSize - currentSize
		g.Log().Debug(ctx, "Refreshing code pool, current size:", currentSize, "needed:", neededCodes)

		// Generate new captchas to fill the pool
		for i := 0; i < neededCodes; i++ {
			id, _, answer, err := p.captcha.Generate()
			if err != nil {
				g.Log().Error(ctx, "Failed to generate captcha:", err)
				continue
			}
			p.pool[id] = answer
		}

		g.Log().Debug(ctx, "Code pool refreshed, new size:", len(p.pool))
	}
}

// GetCode Get a captcha from the pool
func (p *CodePool) GetCode() (id string, b64s string, err error) {
	p.mutex.Lock()

	// If pool is empty, try to refresh
	if len(p.pool) == 0 {
		p.mutex.Unlock()
		p.RefreshPool(true)
		p.mutex.Lock()

		// If still empty, return error
		if len(p.pool) == 0 {
			p.mutex.Unlock()
			return "", "", errors.New("captcha pool is empty")
		}
	}

	// Take a captcha from the pool
	for id, answer := range p.pool {
		delete(p.pool, id)
		p.mutex.Unlock()

		// Get the captcha image data
		item := p.store.Get(id, false)
		if item == "" {
			// If not found in storage, regenerate the captcha
			p.store.Set(id, answer)
		}

		// Regenerate image data
		itemC, err := p.captcha.Driver.DrawCaptcha(answer)

		if err != nil {
			g.Log().Error(context.Background(), "Failed to draw captcha:", err)
			return "", "", err
		}

		b64s = itemC.EncodeB64string()

		// If pool size is below threshold, asynchronously refresh
		if len(p.pool) < p.refreshRate {
			go p.RefreshPool(false)
		}

		return id, b64s, nil
	}

	p.mutex.Unlock()
	return "", "", errors.New("unable to get captcha")
}

// VerifyCode Verify captcha
func (p *CodePool) VerifyCode(id, answer string) bool {
	return p.store.Verify(id, answer, true)
}

// GetAnswer Get captcha answer (without consuming the captcha)
func (p *CodePool) GetAnswer(id string) string {
	return p.store.Get(id, false)
}

// GenerateCaptcha Generate a single captcha
func GenerateCaptcha() (id string, b64s string, answer string, err error) {
	return GetCodePool().captcha.Generate()
}

// GetCaptcha Get a captcha from the captcha pool
func GetCaptcha() (id string, b64s string, err error) {
	return GetCodePool().GetCode()
}

// VerifyCaptcha Verify captcha
func VerifyCaptcha(id, answer string) bool {
	return GetCodePool().VerifyCode(id, answer)
}

// GetCaptchaAnswer Get captcha answer
func GetCaptchaAnswer(id string) string {
	return GetCodePool().GetAnswer(id)
}
