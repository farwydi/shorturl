package domain

type (
	Mode string

	Config struct {
		Mode Mode

		Endpoint struct {
			Web struct {
				Addr string
			}
		}

		Gateway struct {
			Self struct {
				CacheSize int
			}
		}
	}
)

const (
	ModeDevelopment Mode = "dev"
	ModeRelease     Mode = "release"
)
